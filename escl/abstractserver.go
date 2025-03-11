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
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/alexpevzner/mfp/abstract"
	"github.com/alexpevzner/mfp/transport"
	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// AbstractServerHistorySize specifies how many scan jobs the
// [AbstractServer] keeps on its history.
const AbstractServerHistorySize = 10

// AbstractServer implements eSCL server on a top of [abstract.Scanner].
type AbstractServer struct {
	ctx        context.Context               // Logging context
	base       *url.URL                      // Base URL
	scanner    abstract.Scanner              // Underlying abstract.Scanner
	caps       *abstract.ScannerCapabilities // Scanner capabilities
	esclCaps   ScannerCapabilities           // Caps, translated to eSCL
	esclStatus ScannerStatus                 // Scanner status
	lock       sync.Mutex                    // Access lock
}

// NewAbstractServer returns a new [AbstractServer].
//
// The base URL parameter is requires so server knows how to interpret
// [url.URL.Path] of the incoming requests.
//
// For the standard eSCL server that mimics the behavior of the
// typical hardware eSCL scanner, the URL should be something like
// "http://localhost/eSCL".
func NewAbstractServer(ctx context.Context,
	scanner abstract.Scanner, base *url.URL) *AbstractServer {

	// Canonicalize the base URL
	base = transport.URLClone(base)
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}

	// Create the AbstractServer structure
	srv := &AbstractServer{
		ctx:     ctx,
		scanner: scanner,
		base:    base,
		caps:    scanner.Capabilities(),
	}

	srv.esclStatus = ScannerStatus{
		Version: DefaultVersion,
		State:   ScannerIdle,
	}

	if srv.caps.InputSupported.Contains(abstract.InputADF) {
		srv.esclStatus.ADFState = optional.New(ScannerAdfProcessing)
	}

	return srv
}

// ServeHTTP serves incoming HTTP requests.
// It implements the [http.Handler] interface.
func (srv *AbstractServer) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	// Dispatch the request
	if !strings.HasPrefix(rq.URL.Path, srv.base.Path) {
		srv.httpReject(w, rq, http.StatusNotFound, nil)
		return
	}

	path := rq.URL.Path[len(srv.base.Path):]

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

	// No need to acquire srv.lock here, because srv.esclCaps
	// are immutable.
	xml := srv.esclCaps.ToXML()
	srv.httpSendXML(w, xml)
}

// getScannerStatus handles GET /{root}/ScannerStatus request
func (srv *AbstractServer) getScannerStatus(
	w http.ResponseWriter, rq *http.Request) {

	srv.lock.Lock()
	xml := srv.esclStatus.ToXML()
	srv.lock.Unlock()

	srv.httpSendXML(w, xml)
}

// postScanJobs handles GET /{root}/ScannerStatus request
func (srv *AbstractServer) postScanJobs(
	w http.ResponseWriter, rq *http.Request) {
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

	w.Write([]byte(err.Error()))
	w.Write([]byte("\n"))
}

// httpNoCache set response headers to disable client-side.
// response cacheing.
func (srv *AbstractServer) httpNoCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
