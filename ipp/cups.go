// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS requests and responses

package ipp

import (
	"github.com/OpenPrinting/goipp"
)

// The CUPS-Get-Default operation (0x4001) returns the default printer URI
// and attributes.

type (
	// CUPSGetDefaultRequest operation (0x4001) returns the default printer URI
	// and attributes.
	CUPSGetDefaultRequest struct {
		// IPP version and RequestID.
		Version   goipp.Version
		RequestID uint32

		// Operation attributes
		AttributesCharset         string   `ipp:"!attributes-charset,charset"`
		AttributesNaturalLanguage string   `ipp:"!attributes-natural-language,naturalLanguage"`
		RequestedAttributes       []string `ipp:"requested-attributes,keyword"`
	}

	// CUPSGetDefaultResponse is the CUPS-Get-Default Response.
	CUPSGetDefaultResponse struct {
		// IPP version, RequestID, IPP Status code.
		Version   goipp.Version
		RequestID uint32
		Status    goipp.Status

		// Operation attributes.
		AttributesCharset         string `ipp:"!attributes-charset,charset"`
		AttributesNaturalLanguage string `ipp:"!attributes-natural-language,naturalLanguage"`
		StatusMessage             string `ipp:"?status-message,text"`

		Printers []PrinterAttributes
	}
)

// GetVersion returns IPP version of the Request.
func (rq *CUPSGetDefaultRequest) GetVersion() goipp.Version {
	return rq.Version
}

// GetRequestID returns IPP request ID.
func (rq *CUPSGetDefaultRequest) GetRequestID() uint32 {
	return rq.RequestID
}

// GetOp returns CUPSGetDefaultRequest IPP Operation code.
func (rq *CUPSGetDefaultRequest) GetOp() goipp.Op {
	return goipp.OpCupsGetDefault
}

// Encode encodes CUPSGetDefaultRequest into the goipp.Message.
func (rq *CUPSGetDefaultRequest) Encode() *goipp.Message {
	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: ippEncodeAttrs(rq),
		},
	}

	msg := &goipp.Message{
		Version:   rq.Version,
		Code:      goipp.Code(rq.GetOp()),
		RequestID: rq.RequestID,
		Groups:    groups,
	}

	return msg
}

// Decode decodes CUPSGetDefaultRequest from goipp.Groups.
func (rq *CUPSGetDefaultRequest) Decode(msg *goipp.Message) error {
	return nil
}

// GetVersion returns IPP version of the Response.
func (rsp *CUPSGetDefaultResponse) GetVersion() goipp.Version {
	return rsp.Version
}

// GetRequestID returns IPP request ID.
func (rsp *CUPSGetDefaultResponse) GetRequestID() uint32 {
	return rsp.RequestID
}

// GetStatus returns CUPSGetDefaultResponse IPP Status code.
func (rsp *CUPSGetDefaultResponse) GetStatus() goipp.Status {
	return rsp.Status
}

// Encode encodes CUPSGetDefaultResponse into goipp.Message.
func (rsp *CUPSGetDefaultResponse) Encode() *goipp.Message {
	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: ippEncodeAttrs(rsp),
		},
	}

	for i := range rsp.Printers {
		prn := &rsp.Printers[i]
		groups.Add(goipp.Group{
			Tag:   goipp.TagPrinterGroup,
			Attrs: ippEncodeAttrs(prn),
		})
	}

	msg := &goipp.Message{
		Version:   rsp.Version,
		Code:      goipp.Code(rsp.Status),
		RequestID: rsp.RequestID,
		Groups:    groups,
	}

	return msg
}

// Decode decodes CUPSGetDefaultResponse from goipp.Message.
func (rsp *CUPSGetDefaultResponse) Decode(msg *goipp.Message) error {
	return nil
}

// CUPSGetDefaultRequestOperation is the Operation attributes
// for CUPSGetDefaultRequest.
type CUPSGetDefaultRequestOperation struct {
	OperationAttributes
	RequestedAttributes []string `ipp:"requested-attributes,keyword"`
}
