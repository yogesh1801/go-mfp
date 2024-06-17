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
		RequestHeader

		// Operation attributes
		RequestedAttributes []string `ipp:"requested-attributes,keyword"`
	}

	// CUPSGetDefaultResponse is the CUPS-Get-Default Response.
	CUPSGetDefaultResponse struct {
		ResponseHeader

		// Other attributes.
		Printer *PrinterAttributes
	}
)

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

	msg := goipp.NewMessageWithGroups(rq.Version, goipp.Code(rq.GetOp()),
		rq.RequestID, groups)

	return msg
}

// Decode decodes CUPSGetDefaultRequest from goipp.Groups.
func (rq *CUPSGetDefaultRequest) Decode(msg *goipp.Message) error {
	rq.Version = msg.Version
	rq.RequestID = msg.RequestID

	err := ippDecodeAttrs(rq, msg.Operation)
	if err != nil {
		return err
	}

	return nil
}

// Encode encodes CUPSGetDefaultResponse into goipp.Message.
func (rsp *CUPSGetDefaultResponse) Encode() *goipp.Message {
	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: ippEncodeAttrs(rsp),
		},
	}

	if rsp.Printer != nil {
		groups.Add(goipp.Group{
			Tag:   goipp.TagPrinterGroup,
			Attrs: ippEncodeAttrs(rsp.Printer),
		})
	}

	msg := goipp.NewMessageWithGroups(rsp.Version, goipp.Code(rsp.Status),
		rsp.RequestID, groups)

	return msg
}

// Decode decodes CUPSGetDefaultResponse from goipp.Message.
func (rsp *CUPSGetDefaultResponse) Decode(msg *goipp.Message) error {
	rsp.Version = msg.Version
	rsp.RequestID = msg.RequestID
	rsp.Status = goipp.Status(msg.Code)

	err := ippDecodeAttrs(rsp, msg.Operation)
	if err != nil {
		return err
	}

	if len(msg.Printer) != 0 {
		rsp.Printer = &PrinterAttributes{}
		err = ippDecodeAttrs(rsp.Printer, msg.Printer)
		if err != nil {
			return err
		}
	}

	return nil
}
