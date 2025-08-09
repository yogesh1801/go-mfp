// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP protocol sniffer

package ipp

import (
	"io"
	"net/http"

	"github.com/OpenPrinting/goipp"
)

// Sniffer contains a set of hooks which are called for the
// IPP protocol sniffing purposes.
type Sniffer struct {
	// Request, if not nil, is called when IPP request is
	// being sent to the destination.
	//
	// The sequence number is incremented for each new
	// request.
	Request func(seqnum uint64,
		rq *http.Request, msg *goipp.Message, body io.Reader)

	// Response, if not nil, is called when IPP response has been
	// being received from the destination.
	//
	// The sequence number of the response matches the sequence
	// number of the request.
	Response func(seqnum uint64,
		rsp *http.Response, msg *goipp.Message, body io.Reader)
}
