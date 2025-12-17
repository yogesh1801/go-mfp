// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP request handler

package ipp

import (
	"io"

	"github.com/OpenPrinting/goipp"
)

// Handler is the IPP request handler. It implements http.Handler interface.
type Handler struct {
	Op       goipp.Op
	callback func(*goipp.Message, io.Reader) (*goipp.Message, error)
}

// NewHandler creates a new IPP handler from the function that
// consumes [Request] and returns the [goipp.Message] response:
//
//	func DoCUPSGetDefaultRequest(rq *CUPSGetDefaultRequest) *goipp.Message {
//	. . .
//	}
//
//	handler := NewHandler(DoCUPSGetDefaultRequest)
func NewHandler[RQT any,
	RQ interface {
		*RQT
		Request
	}](f func(rq RQ) *goipp.Message) *Handler {

	callback := func(rqMsg *goipp.Message, body io.Reader) (
		*goipp.Message, error) {

		rq := RQ(new(RQT))
		rq.Header().setBody(body)

		err := rq.Decode(rqMsg, DecodeOptions{})
		if err != nil {
			return nil, err
		}

		msg := f(rq)

		return msg, nil
	}

	return &Handler{
		Op:       RQ.GetOp(nil),
		callback: callback,
	}
}

// handle handles the received request.
func (h *Handler) handle(rq *goipp.Message, body io.Reader) (
	*goipp.Message, error) {
	return h.callback(rq, body)
}
