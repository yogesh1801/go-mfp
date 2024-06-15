// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP request handler

package ipp

import (
	"github.com/OpenPrinting/goipp"
)

// Handler is the IPP request handler. It implements http.Handler interface.
type Handler struct {
	Op       goipp.Op
	callback func(*goipp.Message) (*goipp.Message, error)
}

// NewHandler creates a new IPP handler.
//
// Its parameter is a function with a single parameter, a pointer
// to structure that implements [Request] interface, and return value
// is of the [Response] type:
//
//	func DoCUPSGetDefaultRequest(rq *CUPSGetDefaultRequest) Response {
//	. . .
//	}
//
//	handler := NewHandler(DoCUPSGetDefaultRequest)
func NewHandler[RQT any, RQ interface {
	*RQT
	Request
}](f func(rq RQ) Response) *Handler {

	callback := func(rqMsg *goipp.Message) (*goipp.Message, error) {
		rq := RQ(new(RQT))
		err := rq.Decode(rqMsg)
		if err != nil {
			return nil, err
		}

		rsp := f(rq)
		msg := rsp.Encode()

		return msg, nil
	}

	return &Handler{
		Op:       RQ.GetOp(nil),
		callback: callback,
	}
}

// handle handles the received request.
func (h *Handler) handle(rq *goipp.Message) (*goipp.Message, error) {
	return h.callback(rq)
}
