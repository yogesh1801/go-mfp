// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP server

package ipp

import (
	"bytes"
	"fmt"
	"mime"
	"net/http"
	"net/http/httputil"

	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/goipp"
)

// Server represents the IPP server.
type Server struct {
	ops map[goipp.Op]*Handler
}

// NewServer returns a new Sever.
func NewServer() *Server {
	s := &Server{
		ops: make(map[goipp.Op]*Handler),
	}
	return s
}

// ServeHTTP handles incoming HTTP request. It implements
// [http.Handler] interface.
//
// Using this interface, the [Server] can work on a top of
// existent [http.Server], [http.ServeMux], [transport.PathMux]
// and so on.
func (s *Server) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	// Setup things
	query := transport.NewServerQuery(w, rq)
	query.SetLogPrefix("IPP")
	ctx := query.RequestContext()

	// Dump request HTTP headers
	dump, _ := httputil.DumpRequest(query.Request(), false)
	log.Debug(ctx, "IPP request received:")
	log.Debug(ctx, "%s", dump)

	// Check HTTP parameters
	if query.RequestMethod() != "POST" {
		s.httpError(query, ErrHTTPMethodNotAllowed)
		return
	}

	ctype := query.RequestContentType()
	mediatype, _, _ := mime.ParseMediaType(ctype)
	if mediatype != goipp.ContentType {
		err := NewErrHTTP(http.StatusUnsupportedMediaType,
			fmt.Sprintf("unsupported media type: %q", ctype))
		s.httpError(query, err)
		return
	}

	// Decode IPP message
	msg := &goipp.Message{}
	err := msg.Decode(query.RequestBody())
	if err != nil {
		s.httpError(query, err)
		return
	}

	// Log the IPP request
	var buf bytes.Buffer
	msg.Print(&buf, true)
	log.Debug(ctx, "IPP request message:")
	log.Debug(ctx, buf.String())

	// Check IPP parameters
	if msg.RequestID == 0 {
		err := NewErrIPP(msg,
			goipp.StatusErrorBadRequest,
			fmt.Sprintf("bad request ID %d", msg.RequestID))

		s.httpError(query, err)
		return
	}

	if msg.Version < goipp.MakeVersion(1, 0) ||
		msg.Version > goipp.DefaultVersion {
		err := NewErrIPP(msg,
			goipp.StatusErrorVersionNotSupported,
			fmt.Sprintf("bad request version %s", msg.Version))

		s.httpError(query, err)
		return
	}

	handler := s.ops[goipp.Op(msg.Code)]
	if handler == nil {
		op := goipp.Op(msg.Code)
		err := NewErrIPP(msg,
			goipp.StatusErrorOperationNotSupported,
			fmt.Sprintf("unsupported operation %s", op))

		s.httpError(query, err)
		return
	}

	// Handle the message
	rsp, err := handler.handle(msg)
	if err != nil {
		s.httpError(query, err)
		return
	}

	// Log the response
	buf.Reset()
	rsp.Print(&buf, false)
	log.Debug(ctx, "IPP response message:")
	log.Debug(ctx, buf.String())

	// Send response
	query.ResponseHeader().Set("Content-Type", "application/ipp")
	query.WriteHeader(http.StatusOK) // At HTTP level everything OK.

	rsp.Encode(query)
}

// RegisterHandler adds the request [Handler].
func (s *Server) RegisterHandler(handler *Handler) {
	s.ops[handler.Op] = handler
}

// httpError finishes HTTP request with an error.
func (s *Server) httpError(query *transport.ServerQuery, err error) {
	switch err := err.(type) {
	case *ErrHTTP:
		query.Reject(err.Status, err)

	case *ErrIPP:
		query.ResponseHeader().Set("Content-Type", "application/ipp")
		query.WriteHeader(http.StatusOK) // At HTTP level everything OK.

		msg := err.Encode()
		msg.Encode(query)

	default:
		query.Reject(http.StatusInternalServerError, err)
	}
}
