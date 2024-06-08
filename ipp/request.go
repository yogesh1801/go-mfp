// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP request

package ipp

import (
	"io"

	"github.com/OpenPrinting/goipp"
)

// Request is the IPP request interface.
type Request interface {
	// GetVersion returns IPP version of the Request.
	GetVersion() goipp.Version

	// GetRequestID returns IPP request ID.
	GetRequestID() uint32

	// GetBody returns Request body or nil if body is not set.
	GetBody() io.Reader

	// GetOp returns IPP Operation code of the Request.
	GetOp() goipp.Op

	// Encode encodes Request into the goipp.Message.
	Encode() *goipp.Message

	// Decode decodes Request from the goipp.Message.
	Decode(*goipp.Message) error
}

// RequestHeader is the common [Request] header. It contains common
// fields and implements common interfaces.
//
// It should be embedded at the beginning of every structure that
// implements the [Request] interface.
type RequestHeader struct {
	// IPP version and RequestID.
	Version   goipp.Version
	RequestID uint32

	// Common Operation attributes
	AttributesCharset         string `ipp:"!attributes-charset,charset"`
	AttributesNaturalLanguage string `ipp:"!attributes-natural-language,naturalLanguage"`

	// Request body. Sent after IPP message. May be nil
	Body io.Reader
}

// GetVersion returns IPP version of the Request.
func (rqh *RequestHeader) GetVersion() goipp.Version {
	return rqh.Version
}

// GetRequestID returns IPP request ID.
func (rqh *RequestHeader) GetRequestID() uint32 {
	return rqh.RequestID
}

// GetBody request body of request or nil if Request has no body.
// Request body (PDF document, for example) transmitted after
// Request IPP message.
func (rqh *RequestHeader) GetBody() io.Reader {
	return rqh.Body
}
