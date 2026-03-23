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
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// AbstractServer implements a WS-Scan server on top of
// [abstract.Scanner].
type AbstractServer struct {
	options           AbstractServerOptions
	caps              *abstract.ScannerCapabilities
	status            ScannerStatus
	defaultScanTicket *ScanTicket
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
	switch msg.Header.Action.Type {
	case ActGetScannerElements:
		srv.getScannerElementsResponse(query, msg)
	default:
		query.Reject(http.StatusBadRequest, nil)
	}
}

// getScannerElements handles GetScannerElements requests.
func (srv *AbstractServer) getScannerElementsResponse(
	query *transport.ServerQuery, msg Message) {

	// Find the GetScannerElements child in the SOAP body
	child, ok := msg.Body.ChildByName(
		NsWSCN + ":GetScannerElementsRequest")
	if !ok {
		query.Reject(http.StatusBadRequest, nil)
		return
	}

	// Decode the request
	req, err := decodeGetScannerElementsRequest(child)
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

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
	srv.sendSOAPResponse(query, msg, ActGetScannerElementsResponse,
		rsp.toXML(NsWSCN+":GetScannerElementsResponse"))
}

// sendSOAPResponse wraps a response body in a SOAP envelope and sends it.
// It reuses the base URL from the request's action so the response
// matches the client's URL scheme.
func (srv *AbstractServer) sendSOAPResponse(
	query *transport.ServerQuery,
	req Message,
	actionType ActionType,
	body xmldoc.Element) {

	rsp := Message{
		Header: Header{
			Action: Action{
				Type:    actionType,
				BaseURL: req.Header.Action.BaseURL,
			},
			MessageID: AnyURI(uuid.Random().URN()),
			To:        optional.New(AnyURI(AddrAnonymous)),
			RelatesTo: optional.New(req.Header.MessageID),
		},
		Body: body,
	}

	query.SendXML(http.StatusOK, NsMap, rsp.toXML())
}
