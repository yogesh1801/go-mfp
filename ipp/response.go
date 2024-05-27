// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP response

package ipp

import "github.com/OpenPrinting/goipp"

// Response is the IPP response interface.
type Response interface {
	// GetVersion returns IPP version of the Response.
	GetVersion() goipp.Version

	// GetRequestID returns IPP request ID.
	GetRequestID() uint32

	// GetStatus returns Response IPP status code of the Response.
	GetStatus() goipp.Status

	// Encode encodes Response into goipp.Message.
	Encode() *goipp.Message

	// Decode decodes Response from goipp.Message.
	Decode(*goipp.Message) error
}
