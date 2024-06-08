// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP response

package ipp

import (
	"io"

	"github.com/OpenPrinting/goipp"
)

// Response is the IPP response interface.
type Response interface {
	// GetVersion returns IPP version of the Response.
	GetVersion() goipp.Version

	// GetRequestID returns IPP request ID.
	GetRequestID() uint32

	// GetStatus returns IPP Status code of the Response.
	GetStatus() goipp.Status

	// Encode encodes Response into goipp.Message.
	Encode() *goipp.Message

	// Decode decodes Response from goipp.Message.
	Decode(*goipp.Message) error

	// GetBody returns [Response] Body.
	GetBody() io.ReadCloser

	// SetBody sets [Response] Body.
	SetBody(body io.ReadCloser)
}

// ResponseHeader is the common [Response] header. It contains common
// fields and implements common interfaces.
//
// It should be embedded at the beginning of every structure that
// implements the [Response] interface.
type ResponseHeader struct {
	// IPP version, RequestID, IPP Status code.
	Version   goipp.Version
	RequestID uint32
	Status    goipp.Status

	// Common Operation attributes.
	AttributesCharset         string `ipp:"!attributes-charset,charset"`
	AttributesNaturalLanguage string `ipp:"!attributes-natural-language,naturalLanguage"`
	StatusMessage             string `ipp:"?status-message,text"`

	// Response Body.
	Body io.ReadCloser
}

// GetVersion returns IPP version of the [Response].
func (rsph *ResponseHeader) GetVersion() goipp.Version {
	return rsph.Version
}

// GetRequestID returns IPP request ID of the [Response].
func (rsph *ResponseHeader) GetRequestID() uint32 {
	return rsph.RequestID
}

// GetStatus returns [Response] IPP Status code.
func (rsph *ResponseHeader) GetStatus() goipp.Status {
	return rsph.Status
}

// GetBody returns [Response] Body.
func (rsph *ResponseHeader) GetBody() io.ReadCloser {
	return rsph.Body
}

// SetBody sets [Response] Body.
func (rsph *ResponseHeader) SetBody(body io.ReadCloser) {
	rsph.Body = body
}
