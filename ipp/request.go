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
	// The following methods are implemented by the RequestHeader.
	// Each concrete Request implementation inherits them by
	// embedding this structure:
	//
	//   - GetHeader returns RequestHeader
	//   - GetVersion returns IPP version of the Request.
	//   - GetRequestID returns IPP request ID.
	//   - GetBody returns Request body or nil if body is not set.
	GetHeader() *RequestHeader
	GetVersion() goipp.Version
	GetRequestID() uint32
	GetBody() io.Reader

	// The following methods each concrete Request implementation
	// must define by itself:
	//   - GetOp returns IPP Operation code of the Request.
	//   - Encode encodes Request into the goipp.Message.
	//   - Decode decodes Request from the goipp.Message.
	GetOp() goipp.Op
	Encode() *goipp.Message
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

// GetHeader returns [RequestHeader], which gives uniform
// access to the header of any [Request]
func (rqh *RequestHeader) GetHeader() *RequestHeader {
	return rqh
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
