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
// Its parameter is a function with two parameters: Request and Response.
func NewHandler[
	RQT any, RSPT any,
	RQ interface {
		*RQT
		Request
	}, RSP interface {
		*RSPT
		Response
	}](f func(rq RQ, rsp RSP) error) *Handler {

	callback := func(rqMsg *goipp.Message) (*goipp.Message, error) {
		rq := RQ(new(RQT))
		err := rq.Decode(rqMsg)
		if err != nil {
			return nil, err
		}

		rsp := RSP(new(RSPT))
		err = f(rq, rsp)
		if err != nil {
			return nil, err
		}

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
