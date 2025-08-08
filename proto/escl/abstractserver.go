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
	"io"
	"net/http"
	"path"
	"strings"
	"sync"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/missed"
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// AbstractServerHistorySize specifies how many scan jobs the
// [AbstractServer] keeps on its history.
const AbstractServerHistorySize = 10

// AbstractServer implements eSCL server on a top of [abstract.Scanner].
type AbstractServer struct {
	ctx      context.Context               // Logging context
	options  AbstractServerOptions         // Server options
	caps     *abstract.ScannerCapabilities // Scanner capabilities
	status   ScannerStatus                 // Scanner status
	document abstract.Document             // Document being server
	joburi   string                        // Current JobURI, "" if none
	lock     sync.Mutex                    // Access lock
}

// AbstractServerOptions represents the [AbstractServerOptions]
// creation options.
type AbstractServerOptions struct {
	Version Version          // eSCL version, DefaultVersion, if not set
	Scanner abstract.Scanner // Underlying abstract.Scanner
	Hooks   ServerHooks      // eSCL server hooks

	// The BasePath parameter is required so server knows how to
	// interpret [url.URL.Path] of the incoming requests.
	//
	// For the standard eSCL server that mimics the behavior of the
	// typical hardware eSCL scanner, the URL should be something like
	// "/eSCL".
	BasePath string
}

// NewAbstractServer returns a new [AbstractServer].
func NewAbstractServer(ctx context.Context,
	options AbstractServerOptions) *AbstractServer {

	// Use DefaultVersion, if options.Version is not set
	if options.Version == 0 {
		options.Version = DefaultVersion
	}

	// Canonicalize the base path
	options.BasePath = transport.CleanURLPath(options.BasePath + "/")

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
	// Create a transport.ServerQuery
	query := transport.NewServerQuery(w, rq)
	defer query.Finish()

	// Call the OnHTTPRequest hook
	if srv.options.Hooks.OnHTTPRequest != nil {
		srv.options.Hooks.OnHTTPRequest(query)
		if query.IsStatusSet() {
			return
		}
	}

	// Dispatch the request
	if !strings.HasPrefix(query.RequestURL().Path, srv.options.BasePath) {
		query.Reject(http.StatusNotFound, nil)
		return
	}

	path := query.RequestURL().Path
	subpath, _ := missed.StringsCutPrefix(path, srv.options.BasePath)
	method := query.RequestMethod()

	// Dispatch the request
	var action func(*transport.ServerQuery)

	const NextDocument = "/NextDocument"
	const ScanImageInfo = "/ScanImageInfo"

	switch {
	// Handle {root}-relative requests
	case method == "GET" && subpath == "ScannerCapabilities":
		action = srv.getScannerCapabilities

	case method == "GET" && subpath == "ScannerStatus":
		action = srv.getScannerStatus

	case method == "POST" && subpath == "ScanJobs":
		action = srv.postScanJobs

	// Handle {JobUri}-relative requests
	case method == "GET" && strings.HasSuffix(path, NextDocument):
		joburi := path[:len(path)-len(NextDocument)]
		action = func(*transport.ServerQuery) {
			srv.getJobURINextDocument(query, joburi)
		}

	case method == "GET" && strings.HasSuffix(path, ScanImageInfo):
		joburi := path[:len(path)-len(ScanImageInfo)]
		action = func(*transport.ServerQuery) {
			srv.getJobURIScanImageInfo(query, joburi)
		}

	case method == "DELETE":
		action = func(*transport.ServerQuery) {
			srv.deleteJobURI(query, path)
		}
	}

	if action != nil {
		action(query)
	} else {
		query.Reject(http.StatusNotFound, nil)
	}
}

// getScannerCapabilities handles GET /{root}/ScannerCapabilities request
func (srv *AbstractServer) getScannerCapabilities(query *transport.ServerQuery) {
	// Call OnScannerCapabilitiesRequest hook
	if srv.options.Hooks.OnScannerCapabilitiesRequest != nil {
		srv.options.Hooks.OnScannerCapabilitiesRequest(query)
		if query.IsStatusSet() {
			return
		}
	}

	// Generate eSCL ScannerCapabilities
	ver := srv.status.Version
	caps := FromAbstractScannerCapabilities(ver, srv.caps)

	// Call OnScannerCapabilitiesResponse hook
	if srv.options.Hooks.OnScannerCapabilitiesResponse != nil {
		caps2 := srv.options.Hooks.OnScannerCapabilitiesResponse(
			query, caps)
		if query.IsStatusSet() {
			return
		}

		if caps2 != nil {
			caps = caps2
		}
	}

	// Generate and send XML response
	srv.sendXML(query, HookScannerCapabilities, caps)
}

// getScannerStatus handles GET /{root}/ScannerStatus request
func (srv *AbstractServer) getScannerStatus(query *transport.ServerQuery) {
	if srv.options.Hooks.OnScannerStatusRequest != nil {
		srv.options.Hooks.OnScannerStatusRequest(query)
		if query.IsStatusSet() {
			return
		}
	}

	srv.lock.Lock()
	status := srv.status
	srv.lock.Unlock()

	if srv.options.Hooks.OnScannerStatusResponse != nil {
		status2 := srv.options.Hooks.OnScannerStatusResponse(
			query, &status)
		if query.IsStatusSet() {
			return
		}

		if status2 != nil {
			status = *status2
		}
	}

	srv.sendXML(query, HookScannerStatus, &status)
}

// postScanJobs handles POST /{root}/ScanJobs
func (srv *AbstractServer) postScanJobs(query *transport.ServerQuery) {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	// Fetch the XML request body
	xml, err := xmldoc.Decode(NsMap, query.RequestBody())
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

	// Call OnXMLRequest hook
	if srv.options.Hooks.OnXMLRequest != nil {
		xml2 := srv.options.Hooks.OnXMLRequest(
			query, HookScanJobs, xml)
		if query.IsStatusSet() {
			return
		}

		if !xml2.IsZero() {
			xml = xml2
		}
	}

	// Decode ScanSettings request
	ss, err := DecodeScanSettings(xml)
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

	// Call OnScanJobsRequest hook
	if srv.options.Hooks.OnScanJobsRequest != nil {
		ss2 := srv.options.Hooks.OnScanJobsRequest(query, ss)
		if query.IsStatusSet() {
			return
		}

		if ss2 != nil {
			ss = ss2
		}
	}

	// Check if previous request already in progress
	if srv.document != nil {
		err := errors.New("Device is busy with the previous request")
		query.Reject(http.StatusServiceUnavailable, err)
		return
	}

	// Convert it into the abstract.ScannerRequest and validate
	absreq := ss.ToAbstract()

	// Generate a new Job UUID. Do it now, because in theory
	// it can fail (though very unlikely), so do it before
	// the job is created
	uu, err := uuid.Random()
	if err != nil {
		query.Reject(http.StatusServiceUnavailable, err)
		return
	}

	// Send request to the underlying abstract.Scanner
	document, err := srv.options.Scanner.Scan(srv.ctx, absreq)
	if err != nil {
		query.Reject(http.StatusConflict, err)
		return
	}

	// Update server status
	srv.document = document
	srv.status.State = ScannerProcessing

	jobuuid := uu.URN()
	joburi := path.Join(srv.options.BasePath, "ScanJobs", jobuuid)

	info := JobInfo{
		JobURI:   joburi,
		JobUUID:  optional.New(jobuuid),
		JobState: JobProcessing,
	}

	srv.joburi = joburi
	srv.status.PushJobInfo(info, AbstractServerHistorySize)

	// Call OnScanJobsResponse hook
	if srv.options.Hooks.OnScanJobsResponse != nil {
		joburi2 := srv.options.Hooks.OnScanJobsResponse(query, ss, info)
		if query.IsStatusSet() {
			return
		}

		if joburi2 != "" {
			joburi = joburi2
		}
	}

	// Complete the request
	query.Created(joburi)
}

// getJobURINextDocument handles GET /{JobUri}/NextDocument
func (srv *AbstractServer) getJobURINextDocument(
	query *transport.ServerQuery, joburi string) {

	// Call OnNextDocumentRequest hook
	if srv.options.Hooks.OnNextDocumentRequest != nil {
		joburi2 := srv.options.Hooks.OnNextDocumentRequest(
			query, joburi)
		if query.IsStatusSet() {
			return
		}

		if joburi2 != "" {
			joburi = joburi2
		}
	}

	// Fetch the next document file
	srv.lock.Lock()

	var file abstract.DocumentFile
	var err error
	var info JobInfo

	if srv.document != nil && srv.joburi == joburi {
		file, err = srv.document.Next()
		info = srv.status.Jobs[0]
	}

	srv.lock.Unlock()

	// Handle possible error conditions
	switch {
	case err == nil && file == nil:
		query.Reject(http.StatusNotFound, nil)
		return

	case err == io.EOF:
		srv.finish(JobCompleted, JobCompletedSuccessfully)
		query.Reject(http.StatusNotFound, nil)
		return

	case err != nil:
		srv.finish(JobCanceled, AbortedBySystem)
		query.Reject(http.StatusServiceUnavailable, err)
		return
	}

	// Call OnNextDocumentResponse hook
	if srv.options.Hooks.OnNextDocumentResponse != nil {
		file2 := srv.options.Hooks.OnNextDocumentResponse(
			query, info, file)
		if query.IsStatusSet() {
			return
		}

		if file2 != nil {
			file = file2
		}
	}

	// Send the response
	query.SendData(http.StatusOK, file.Format(), file)
}

// getJobURIScanImageInfo handles GET /{JobUri}/ScanImageInfo
func (srv *AbstractServer) getJobURIScanImageInfo(
	query *transport.ServerQuery, joburi string) {
	query.Reject(http.StatusNotImplemented, nil)
}

// deleteJobURI handles DELETE /{JobUri}
func (srv *AbstractServer) deleteJobURI(
	query *transport.ServerQuery, joburi string) {

	// Call OnDeleteRequest hook
	if srv.options.Hooks.OnDeleteRequest != nil {
		joburi2 := srv.options.Hooks.OnDeleteRequest(query, joburi)
		if query.IsStatusSet() {
			return
		}

		if joburi2 != "" {
			joburi = joburi2
		}
	}

	// Check the joburi
	srv.lock.Lock()
	jobOK := srv.document != nil && srv.joburi == joburi
	srv.lock.Unlock()

	if !jobOK {
		query.Reject(http.StatusNotFound, nil)
		return
	}

	// Finish the job
	srv.finish(JobCanceled, JobCanceledByUser)
	query.WriteHeader(http.StatusOK)
}

// finish finishes the current job and updates server state
func (srv *AbstractServer) finish(state JobState, reason JobStateReason) {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	srv.document.Close()
	srv.document = nil
	srv.joburi = ""
	srv.status.State = ScannerIdle
	srv.status.Jobs[0].JobState = state
	if reason != UnknownJobStateReason {
		srv.status.Jobs[0].JobStateReasons = []JobStateReason{reason}
	}
}

// sendXML generates and sends the XML response to the query.
func (srv *AbstractServer) sendXML(query *transport.ServerQuery,
	action HookAction, rsp interface{ ToXML() xmldoc.Element }) {

	xml := rsp.ToXML()
	if srv.options.Hooks.OnXMLResponse != nil {
		xml2 := srv.options.Hooks.OnXMLResponse(query, action, xml)

		if query.IsStatusSet() {
			return
		}

		if !xml2.IsZero() {
			xml = xml2
		}
	}

	query.SendXML(http.StatusOK, NsMap, xml)
}
