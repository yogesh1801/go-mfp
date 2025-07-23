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
	lock     sync.Mutex                    // Access lock
}

// AbstractServerOptions represents the [AbstractServerOptions]
// creation options.
type AbstractServerOptions struct {
	Version Version          // eSCL version, DefaultVersion, if not set
	Scanner abstract.Scanner // Underlying abstract.Scanner

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
	query := transport.NewServerQuery(srv.ctx, w, rq)
	defer query.Finish()

	// Dispatch the request
	if !strings.HasPrefix(query.RequestURL().Path, srv.options.BasePath) {
		query.Reject(http.StatusNotFound, nil)
		return
	}

	path, _ := missed.StringsCutPrefix(query.RequestURL().Path,
		srv.options.BasePath)

	// Handle {root}-relative requests
	var action func(*transport.ServerQuery)

	srv.lock.Lock()

	switch path {
	case "ScannerCapabilities":
		if query.RequestMethod() == "GET" {
			action = srv.getScannerCapabilities
		}

	case "ScannerStatus":
		if query.RequestMethod() == "GET" {
			action = srv.getScannerStatus
		}

	case "ScanJobs":
		if query.RequestMethod() == "POST" {
			action = srv.postScanJobs
		}
	}

	// Handle {JobUri}-relative requests
	if action == nil && srv.document != nil {
		joburi := srv.status.Jobs[0].JobURI

		switch query.RequestMethod() {
		case "GET":
			switch query.RequestURL().Path {
			case joburi + "/NextDocument":
				action = srv.getJobURINextDocument
			case joburi + "/ScanImageInfo":
				action = srv.getJobURIScanImageInfo
			}

		case "DELETE":
			if query.RequestURL().Path == joburi {
				action = srv.deleteJobURI
			}
		}
	}

	srv.lock.Unlock()

	if action != nil {
		action(query)
	} else {
		query.Reject(http.StatusNotFound, nil)
	}
}

// getScannerCapabilities handles GET /{root}/ScannerCapabilities request
func (srv *AbstractServer) getScannerCapabilities(query *transport.ServerQuery) {
	ver := srv.status.Version
	xml := fromAbstractScannerCapabilities(ver, srv.caps).ToXML()
	query.SendXML(http.StatusOK, NsMap, xml)
}

// getScannerStatus handles GET /{root}/ScannerStatus request
func (srv *AbstractServer) getScannerStatus(query *transport.ServerQuery) {
	srv.lock.Lock()
	xml := srv.status.ToXML()
	srv.lock.Unlock()

	query.SendXML(http.StatusOK, NsMap, xml)
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

	// Decode ScanSettings request
	ss, err := DecodeScanSettings(xml)
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
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

	srv.status.PushJobInfo(info, AbstractServerHistorySize)

	// Complete the request
	query.Created(joburi)
}

// getJobURINextDocument handles GET /{JobUri}/NextDocument
func (srv *AbstractServer) getJobURINextDocument(query *transport.ServerQuery) {
	srv.lock.Lock()
	file, err := srv.document.Next()
	srv.lock.Unlock()

	switch {
	case err == io.EOF:
		srv.finish(JobCompleted, JobCompletedSuccessfully)
		query.Reject(http.StatusNotFound, nil)

	case err != nil:
		srv.finish(JobCanceled, AbortedBySystem)
		query.Reject(http.StatusServiceUnavailable, err)

	default:
		query.SendData(http.StatusOK, file.Format(), file)
	}
}

// getJobURIScanImageInfo handles GET /{JobUri}/ScanImageInfo
func (srv *AbstractServer) getJobURIScanImageInfo(query *transport.ServerQuery) {
	query.Reject(http.StatusNotImplemented, nil)
}

// deleteJobURI handles DELETE /{JobUri}
func (srv *AbstractServer) deleteJobURI(query *transport.ServerQuery) {
	srv.finish(JobCanceled, JobCanceledByUser)
	query.WriteHeader(http.StatusOK)
}

// finish finishes the current job and updates server state
func (srv *AbstractServer) finish(state JobState, reason JobStateReason) {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	srv.document.Close()
	srv.document = nil
	srv.status.State = ScannerIdle
	srv.status.Jobs[0].JobState = state
	if reason != UnknownJobStateReason {
		srv.status.Jobs[0].JobStateReasons = []JobStateReason{reason}
	}
}
