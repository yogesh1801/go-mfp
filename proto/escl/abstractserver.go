// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL server on a top of abstract.Scanner

package escl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/alexpevzner/mfp/abstract"
	"github.com/alexpevzner/mfp/log"
	"github.com/alexpevzner/mfp/transport"
	"github.com/alexpevzner/mfp/util/missed"
	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// AbstractServerHistorySize specifies how many scan jobs the
// [AbstractServer] keeps on its history.
const AbstractServerHistorySize = 10

// AbstractServer implements eSCL server on a top of [abstract.Scanner].
type AbstractServer struct {
	ctx     context.Context               // Logging context
	options AbstractServerOptions         // Server options
	caps    *abstract.ScannerCapabilities // Scanner capabilities
	status  ScannerStatus                 // Scanner status
	lock    sync.Mutex                    // Access lock
}

// AbstractServerOptions represents the [AbstractServerOptions]
// creation options.
type AbstractServerOptions struct {
	Version Version          // eSCL version, DefaultVersion, if not set
	Scanner abstract.Scanner // Underlying abstract.Scanner

	// The BaseURL parameter is required so server knows how to
	// interpret [url.URL.Path] of the incoming requests.
	//
	// For the standard eSCL server that mimics the behavior of the
	// typical hardware eSCL scanner, the URL should be something like
	// "http://localhost/eSCL".
	BaseURL *url.URL
}

// abstractServerQuery maintains an AbstractServer query processing
// context, allowing per-request centralized logging and hooking.
//
// It keeps the reference to the original [http.Request] and wraps
// the corresponding [http.ResponseWriter], passed to the
// AbstractServer.ServeHTTP
type abstractServerQuery struct {
	log                 *log.Record  // Log record for the query
	*http.Request                    // Incoming request
	http.ResponseWriter              // Underlying http.ResponseWriter
	status              atomic.Int32 // HTTP status, 0 if not known yet
}

// newAbstractServerQuery returns the new abstractServerQuery
func newAbstractServerQuery(srv *AbstractServer,
	w http.ResponseWriter, rq *http.Request) *abstractServerQuery {

	query := &abstractServerQuery{
		log:            log.Begin(srv.ctx),
		Request:        rq,
		ResponseWriter: w,
	}

	return query
}

// RequestHeader returns http.Header of the request
func (query *abstractServerQuery) RequestHeader() http.Header {
	return query.Request.Header
}

// Finish must be called when query processing is finished
func (query *abstractServerQuery) Finish() {
	query.log.Commit()
}

// RequestBody returns body of the http.Request
func (query *abstractServerQuery) RequestBody() io.ReadCloser {
	return query.Request.Body
}

// ResponseHeader returns http.Header of the response
func (query *abstractServerQuery) ResponseHeader() http.Header {
	return query.ResponseWriter.Header()
}

// Write writes response body bytes.
func (query *abstractServerQuery) Write(data []byte) (int, error) {
	return query.ResponseWriter.Write(data)
}

// WriteHeader writes HTTP response header.
func (query *abstractServerQuery) WriteHeader(status int) {
	if query.status.CompareAndSwap(0, int32(status)) {
		query.ResponseWriter.WriteHeader(status)
		query.log.Debug("HTTP %s %s -- %d %s",
			query.Method, query.URL,
			status, http.StatusText(status))
		query.log.Flush()
	}
}

// NoCache set response headers to disable client-side response cacheing.
func (query *abstractServerQuery) NoCache() {
	hdr := query.ResponseHeader()
	hdr.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	hdr.Set("Pragma", "no-cache")
	hdr.Set("Expires", "0")
}

// Reject completes request with a error.
func (query *abstractServerQuery) Reject(status int, err error) {
	query.ResponseHeader().Set("Content-Type", "text/plain; charset=utf-8")
	query.NoCache()
	query.WriteHeader(status)

	if err == nil {
		err = errors.New(http.StatusText(status))
	}

	s := fmt.Sprintf("%3.3d %s\n", status, err)
	query.Write([]byte(s))
	query.Write([]byte("\n"))
}

// SendXML sends the XML response.
func (query *abstractServerQuery) SendXML(xml xmldoc.Element) {
	query.ResponseHeader().Set("Content-Type", HTTPContentType)
	query.WriteHeader(http.StatusOK)
	xml.EncodeIndent(query, NsMap, "  ")
}

// NewAbstractServer returns a new [AbstractServer].
func NewAbstractServer(ctx context.Context,
	options AbstractServerOptions) *AbstractServer {

	// Use DefaultVersion, if options.Version is not set
	if options.Version == 0 {
		options.Version = DefaultVersion
	}

	// Canonicalize the base URL
	options.BaseURL = transport.URLClone(options.BaseURL)
	if !strings.HasSuffix(options.BaseURL.Path, "/") {
		options.BaseURL.Path += "/"
	}

	// Create the AbstractServer structure
	srv := &AbstractServer{
		ctx:     ctx,
		options: options,
		caps:    options.Scanner.Capabilities(),
	}

	srv.status = ScannerStatus{
		Version: options.Version,
		State:   ScannerIdle,
	}

	if srv.caps.ADFSimplex != nil || srv.caps.ADFDuplex != nil {
		srv.status.ADFState = optional.New(ScannerAdfProcessing)
	}

	return srv
}

// ServeHTTP serves incoming HTTP requests.
// It implements the [http.Handler] interface.
func (srv *AbstractServer) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	// Create a abstractServerQuery
	query := newAbstractServerQuery(srv, w, rq)
	defer query.Finish()

	// Dispatch the request
	if !strings.HasPrefix(query.URL.Path, srv.options.BaseURL.Path) {
		query.Reject(http.StatusNotFound, nil)
		return
	}

	path, _ := missed.StringsCutPrefix(query.URL.Path,
		srv.options.BaseURL.Path)

	switch path {
	case "ScannerCapabilities":
		if query.Method == "GET" {
			srv.getScannerCapabilities(query)
			return
		}

	case "ScannerStatus":
		if rq.Method == "GET" {
			srv.getScannerStatus(query)
			return
		}

	case "ScanJobs":
		if rq.Method == "POST" {
			srv.postScanJobs(query)
			return
		}
	}

	query.Reject(http.StatusNotFound, nil)
}

// getScannerCapabilities handles GET /{root}/ScannerCapabilities request
func (srv *AbstractServer) getScannerCapabilities(query *abstractServerQuery) {
	ver := srv.status.Version
	xml := fromAbstractScannerCapabilities(ver, srv.caps).ToXML()
	query.SendXML(xml)
}

// getScannerStatus handles GET /{root}/ScannerStatus request
func (srv *AbstractServer) getScannerStatus(query *abstractServerQuery) {
	srv.lock.Lock()
	xml := srv.status.ToXML()
	srv.lock.Unlock()

	query.SendXML(xml)
}

// postScanJobs handles POST /{root}/ScanJobs
func (srv *AbstractServer) postScanJobs(query *abstractServerQuery) {
	// Fetch the XML request body
	xml, err := xmldoc.Decode(NsMap, query.RequestBody())
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

	// Decode ScanSettings request
	ss, err := DecodeScanSettings(xml)
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

	// Convert it into the abstract.ScannerRequest and validate
	absreq := ss.toAbstract()
	err = absreq.Validate(srv.caps)
	if err != nil {
		query.Reject(http.StatusConflict, err)
		return
	}

	query.Reject(http.StatusNotImplemented, nil)
}
