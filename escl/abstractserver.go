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
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/alexpevzner/mfp/abstract"
	"github.com/alexpevzner/mfp/log"
	"github.com/alexpevzner/mfp/transport"
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
		Version: DefaultVersion,
		State:   ScannerIdle,
	}

	if srv.caps.ADFSimplex != nil || srv.caps.ADFDuplex != nil {
		srv.status.ADFState = optional.New(ScannerAdfProcessing)
	}

	return srv
}

// GetVersion returns eSCL [Version], implemented by the server.
func (srv *AbstractServer) GetVersion() Version {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	return srv.status.Version
}

// SetVersion sets eSCL [Version], implemented by the server.
//
// As version affects some aspects of the server behavior, this
// is not recommended to change the eSCL on the running server.
// Do it at the initialization time only.
func (srv *AbstractServer) SetVersion(ver Version) {
	srv.lock.Lock()
	srv.status.Version = ver
	srv.lock.Unlock()
}

// ServeHTTP serves incoming HTTP requests.
// It implements the [http.Handler] interface.
func (srv *AbstractServer) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	// Log the request
	log.Debug(srv.ctx, "HTTP %s %s", rq.Method, rq.URL)

	// Dispatch the request
	if !strings.HasPrefix(rq.URL.Path, srv.options.BaseURL.Path) {
		srv.httpReject(w, rq, http.StatusNotFound, nil)
		return
	}

	path := rq.URL.Path[len(srv.options.BaseURL.Path):]

	switch path {
	case "ScannerCapabilities":
		if rq.Method == "GET" {
			srv.getScannerCapabilities(w, rq)
			return
		}

	case "ScannerStatus":
		if rq.Method == "GET" {
			srv.getScannerStatus(w, rq)
			return
		}

	case "ScanJobs":
		if rq.Method == "POST" {
			srv.postScanJobs(w, rq)
			return
		}
	}

	srv.httpReject(w, rq, http.StatusNotFound, nil)
}

// getScannerCapabilities handles GET /{root}/ScannerCapabilities request
func (srv *AbstractServer) getScannerCapabilities(
	w http.ResponseWriter, rq *http.Request) {

	srv.lock.Lock()
	ver := srv.status.Version
	xml := fromAbstractScannerCapabilities(ver, srv.caps).ToXML()
	srv.lock.Unlock()

	srv.httpSendXML(w, xml)
}

// getScannerStatus handles GET /{root}/ScannerStatus request
func (srv *AbstractServer) getScannerStatus(
	w http.ResponseWriter, rq *http.Request) {

	srv.lock.Lock()
	xml := srv.status.ToXML()
	srv.lock.Unlock()

	srv.httpSendXML(w, xml)
}

// postScanJobs handles POST /{root}/ScanJobs
func (srv *AbstractServer) postScanJobs(
	w http.ResponseWriter, rq *http.Request) {

	// Fetch the XML request body
	xml, err := xmldoc.Decode(NsMap, rq.Body)
	if err != nil {
		srv.httpReject(w, rq, http.StatusBadRequest, err)
		return
	}

	// Decode ScanSettings request
	ss, err := DecodeScanSettings(xml)
	if err != nil {
		srv.httpReject(w, rq, http.StatusBadRequest, err)
		return
	}

	_ = ss

	srv.httpReject(w, rq, http.StatusNotImplemented, nil)
}

// httpSendXML sends XML response
func (srv *AbstractServer) httpSendXML(w http.ResponseWriter,
	xml xmldoc.Element) {

	w.Header().Set("Content-Type", HTTPContentType)
	w.WriteHeader(http.StatusOK)
	xml.EncodeIndent(w, NsMap, "  ")
}

// httpReject completes request with a error.
func (srv *AbstractServer) httpReject(w http.ResponseWriter, rq *http.Request,
	status int, err error) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	srv.httpNoCache(w)
	w.WriteHeader(status)

	if err == nil {
		err = errors.New(http.StatusText(status))
	}

	s := fmt.Sprintf("%3.3d %s\n", status, err)
	w.Write([]byte(s))
	w.Write([]byte("\n"))
}

// httpNoCache set response headers to disable client-side.
// response cacheing.
func (srv *AbstractServer) httpNoCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
