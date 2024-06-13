// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP server

package ipp

import (
	"fmt"
	"mime"
	"net/http"

	"github.com/OpenPrinting/goipp"
)

// Server represents an IPP server.
type Server struct {
	httpServer *http.Server
	ops        map[goipp.Op]*Handler
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
// existent [http.Server] or [http.ServeMux].
func (s *Server) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	// Check HTTP parameters
	if rq.Method != "POST" {
		s.httpError(w, ErrHTTPMethodNotAllowed)
		return
	}

	ctype := rq.Header.Get("Content-Type")
	mediatype, _, _ := mime.ParseMediaType(ctype)
	if mediatype != goipp.ContentType {
		err := NewErrHTTP(http.StatusUnsupportedMediaType,
			fmt.Sprintf("unsupported media type: %q", ctype))
		s.httpError(w, err)
		return
	}

	// Decode IPP message
	msg := &goipp.Message{}
	err := msg.Decode(rq.Body)
	if err != nil {
		s.httpError(w, err)
		return
	}

	// Check IPP parameters
	if msg.RequestID == 0 {
		err := NewErrIPP(msg,
			goipp.StatusErrorVersionNotSupported,
			fmt.Sprintf("bad request ID %d", msg.RequestID))

		s.httpError(w, err)
		return
	}

	if msg.Version < goipp.MakeVersion(1, 0) ||
		msg.Version > goipp.DefaultVersion {
		err := NewErrIPP(msg,
			goipp.StatusErrorVersionNotSupported,
			fmt.Sprintf("bad request version %s", msg.Version))

		s.httpError(w, err)
		return
	}

	handler := s.ops[goipp.Op(msg.Code)]
	if handler == nil {
		op := goipp.Op(msg.Code)
		err := NewErrIPP(msg,
			goipp.StatusErrorVersionNotSupported,
			fmt.Sprintf("unsupported operation %s", op))

		s.httpError(w, err)
		return
	}

	// Handle the message
	rsp, err := handler.handle(msg)
	if err != nil {
		s.httpError(w, err)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/ipp")
	w.WriteHeader(http.StatusOK) // At HTTP level everything OK.

	rsp.Encode(w)
}

// httpError finishes HTTP request with an error.
func (s *Server) httpError(w http.ResponseWriter, err error) {
AGAIN:
	switch err := err.(type) {
	case *ErrHTTP:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		s.httpNoCache(w)
		w.WriteHeader(err.Status)

		fmt.Fprintf(w, "%3.3d %s\n", err.Status, err.Message)

	case *ErrIPP:
		w.Header().Set("Content-Type", "application/ipp")
		w.WriteHeader(http.StatusOK) // At HTTP level everything OK.

		msg := err.Encode()
		msg.Encode(w)

	default:
		err = NewErrHTTP(http.StatusInternalServerError,
			err.Error())
		goto AGAIN
	}
}

// httpNoCache sets response headers to disable caching.
func (s *Server) httpNoCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
