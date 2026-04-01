package wsscan

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
)

// AbstractServer implements a WS-Scan server on top of
// [abstract.Scanner].
type AbstractServer struct {
	options           AbstractServerOptions
	caps              *abstract.ScannerCapabilities
	status            ScannerStatus
	defaultScanTicket *ScanTicket
	document          abstract.Document
	jobID             int
	jobToken          string
	lock              sync.Mutex
}

// AbstractServerOptions allows specifying options that can
// modify the [AbstractServer] behavior.
type AbstractServerOptions struct {
	Scanner abstract.Scanner

	// BasePath is required so the server knows how to
	// interpret incoming request paths.
	BasePath string

	// DefaultScanTicket provides the default scan settings
	// returned in GetScannerElementsResponse. If nil, no
	// DefaultScanTicket element is included.
	DefaultScanTicket *ScanTicket
}

// NewAbstractServer returns a new [AbstractServer].
func NewAbstractServer(options AbstractServerOptions) *AbstractServer {
	srv := &AbstractServer{
		options:           options,
		caps:              options.Scanner.Capabilities(),
		defaultScanTicket: options.DefaultScanTicket,
	}

	srv.status = ScannerStatus{
		ScannerState: Idle,
	}

	return srv
}

// ServeHTTP serves incoming HTTP requests.
// It implements the [http.Handler] interface.
func (srv *AbstractServer) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	query := transport.NewServerQuery(w, rq)
	defer query.Finish()

	// Check path prefix
	if !strings.HasPrefix(query.RequestURL().Path, srv.options.BasePath) {
		query.Reject(http.StatusNotFound, nil)
		return
	}

	// WS-Scan uses POST only
	if query.RequestMethod() != "POST" {
		query.Reject(http.StatusMethodNotAllowed, nil)
		return
	}

	// Read and decode SOAP message
	data, err := io.ReadAll(query.RequestBody())
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

	msg, err := DecodeMessage(data)
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

	// Dispatch by SOAP Action
	switch msg.Header.Action {
	case ActGetScannerElements:
		srv.getScannerElementsResponse(query, msg)
	case ActCreateScanJob:
		srv.getScanJobResponse(query, msg)
	case ActRetrieveImage:
		srv.getRetrieveImageResponse(query, msg)
	default:
		query.Reject(http.StatusBadRequest, nil)
	}
}

// getScannerElementsResponse handles GetScannerElements requests.
func (srv *AbstractServer) getScannerElementsResponse(
	query *transport.ServerQuery, msg Message) {

	req := msg.Body.(GetScannerElementsRequest)

	// Build ElementData for each requested element
	var elements []ElementData

	for _, re := range req.RequestedElements {
		switch re {
		case RequestedElementDescription:
			desc := FromAbstractScannerDescription(srv.caps)
			elements = append(elements, ElementData{
				Name:               ElementDataScannerDescription,
				Valid:              BooleanElement("true"),
				ScannerDescription: optional.New(desc),
			})

		case RequestedElementConfiguration:
			conf := FromAbstractScannerConfiguration(srv.caps)
			elements = append(elements, ElementData{
				Name:                 ElementDataScannerConfiguration,
				Valid:                BooleanElement("true"),
				ScannerConfiguration: optional.New(conf),
			})

		case RequestedElementStatus:
			srv.lock.Lock()
			status := srv.status
			srv.lock.Unlock()
			status.ScannerCurrentTime = time.Now()
			elements = append(elements, ElementData{
				Name:          ElementDataScannerStatus,
				Valid:         BooleanElement("true"),
				ScannerStatus: optional.New(status),
			})
		}
	}

	// Build response
	rsp := GetScannerElementsResponse{
		ScannerElements: elements,
	}

	// Send SOAP response
	srv.sendSOAPResponse(query, msg, rsp)
}

// getScanJobResponse handles CreateScanJob requests.
func (srv *AbstractServer) getScanJobResponse(
	query *transport.ServerQuery, msg Message) {

	req := msg.Body.(CreateScanJobRequest)

	srv.lock.Lock()
	defer srv.lock.Unlock()

	// Check if previous scan is still in progress
	if srv.document != nil {
		query.Reject(http.StatusServiceUnavailable, nil)
		return
	}

	// Convert ScanTicket to abstract.ScannerRequest
	absreq := req.ScanTicket.ToAbstract()

	// Send request to the underlying abstract.Scanner
	ctx := query.RequestContext()
	document, err := srv.options.Scanner.Scan(ctx, absreq)
	if err != nil {
		query.Reject(http.StatusConflict, err)
		return
	}

	// Store document and update status
	srv.document = document
	srv.status.ScannerState = Processing

	// Send SOAP response
	srv.jobID++
	srv.jobToken = uuid.Random().URN()

	rsp := CreateScanJobResponse{
		DocumentFinalParameters: optional.Get(
			req.ScanTicket.DocumentParameters),
		JobID:    srv.jobID,
		JobToken: srv.jobToken,
	}

	srv.sendSOAPResponse(query, msg, rsp)
}

// getRetrieveImageResponse handles RetrieveImage requests.
func (srv *AbstractServer) getRetrieveImageResponse(
	query *transport.ServerQuery, msg Message) {

	req := msg.Body.(RetrieveImageRequest)

	// Validate job credentials
	srv.lock.Lock()

	if srv.document == nil || req.JobID != srv.jobID ||
		req.JobToken != srv.jobToken {
		srv.lock.Unlock()
		query.Reject(http.StatusNotFound, nil)
		return
	}

	// Get next document file
	file, err := srv.document.Next()
	srv.lock.Unlock()

	switch {
	case err == io.EOF:
		srv.finish()
		query.Reject(http.StatusNotFound, nil)
		return
	case err != nil:
		srv.finish()
		query.Reject(http.StatusServiceUnavailable, err)
		return
	}

	// Build response with xop:Include reference
	rsp := RetrieveImageResponse{
		ScanData:    ScanData{ContentID: uuid.Random().String()},
		Image:       io.NopCloser(file),
		ContentType: file.Format(),
	}

	srv.sendMTOMResponse(query, msg, rsp)
}

// finish closes the current document and resets the server state.
func (srv *AbstractServer) finish() {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	srv.document.Close()
	srv.document = nil
	srv.status.ScannerState = Idle
}

// sendMTOMResponse wraps a response body in a SOAP envelope and
// sends it as an MTOM/XOP multipart message with a file attachment.
func (srv *AbstractServer) sendMTOMResponse(
	query *transport.ServerQuery,
	req Message,
	body RetrieveImageResponse) {

	rsp := Message{
		Header: Header{
			Action:    body.Action(),
			MessageID: AnyURI(uuid.Random().URN()),
			To:        optional.New(AnyURI(AddrAnonymous)),
			RelatesTo: optional.New(req.Header.MessageID),
		},
		Body: body,
	}

	// Generate boundary and envelope CID for the multipart message
	boundary := uuid.Random().String()
	envelopeCID := uuid.Random().String()

	// Set headers before writing body
	query.ResponseHeader().Set("Content-Type",
		MTOMContentType(boundary, envelopeCID))
	query.WriteHeader(http.StatusOK)

	// Write the MTOM multipart body
	rsp.WriteMTOM(query, boundary, envelopeCID)
}

// sendSOAPResponse wraps a response body in a SOAP envelope and sends it.
func (srv *AbstractServer) sendSOAPResponse(
	query *transport.ServerQuery,
	req Message,
	body Body) {

	rsp := Message{
		Header: Header{
			Action:    body.Action(),
			MessageID: AnyURI(uuid.Random().URN()),
			To:        optional.New(AnyURI(AddrAnonymous)),
			RelatesTo: optional.New(req.Header.MessageID),
		},
		Body: body,
	}

	query.SendXML(http.StatusOK, NsMap, rsp.toXML())
}
