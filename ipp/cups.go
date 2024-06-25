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
		ObjectAttrs
		RequestHeader

		// Operation attributes
		RequestedAttributes []string `ipp:"requested-attributes,keyword"`
	}

	// CUPSGetDefaultResponse is the CUPS-Get-Default Response.
	CUPSGetDefaultResponse struct {
		ObjectAttrs
		ResponseHeader

		// Other attributes.
		Printer *PrinterAttributes
	}

	// CUPSGetPrintersRequest operation (0x4002) returns the printer
	// attributes for every printer known to the system.
	CUPSGetPrintersRequest struct {
		ObjectAttrs
		RequestHeader

		// Operation attributes
		FirstPrinterName    string   `ipp:"?first-printer-name"`
		Limit               int      `ipp:"?limit"`
		PrinterID           int      `ipp:"?printer-id"`
		PrinterLocation     string   `ipp:"?printer-location"`
		PrinterType         int      `ipp:"?printer-type"`
		PrinterTypeMask     int      `ipp:"?printer-type-mask"`
		RequestedUserName   string   `ipp:"?requested-user-name"`
		RequestedAttributes []string `ipp:"?requested-attributes,keyword"`
	}

	// CUPSGetPrintersResponse is the CUPS-Get-Printers Response.
	CUPSGetPrintersResponse struct {
		ObjectAttrs
		ResponseHeader

		// Other attributes.
		Printer []*PrinterAttributes
	}
)

// ----- CUPS-Get-Default methods -----

// GetOp returns CUPSGetDefaultRequest IPP Operation code.
func (rq *CUPSGetDefaultRequest) GetOp() goipp.Op {
	return goipp.OpCupsGetDefault
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetDefaultRequest
func (rq *CUPSGetDefaultRequest) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rq)
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

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetDefaultResponse.
func (rsp *CUPSGetDefaultResponse) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rsp)
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

// ----- CUPS-Get-Printers methods -----

// GetOp returns CUPSGetPrintersRequest IPP Operation code.
func (rq *CUPSGetPrintersRequest) GetOp() goipp.Op {
	return goipp.OpCupsGetPrinters
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetPrintersRequest
func (rq *CUPSGetPrintersRequest) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rq)
}

// Encode encodes CUPSGetPrintersRequest into the goipp.Message.
func (rq *CUPSGetPrintersRequest) Encode() *goipp.Message {
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

// Decode decodes CUPSGetPrintersRequest from goipp.Groups.
func (rq *CUPSGetPrintersRequest) Decode(msg *goipp.Message) error {
	rq.Version = msg.Version
	rq.RequestID = msg.RequestID

	err := ippDecodeAttrs(rq, msg.Operation)
	if err != nil {
		return err
	}

	return nil
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetPrintersResponse.
func (rsp *CUPSGetPrintersResponse) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rsp)
}

// Encode encodes CUPSGetPrintersResponse into goipp.Message.
func (rsp *CUPSGetPrintersResponse) Encode() *goipp.Message {
	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: ippEncodeAttrs(rsp),
		},
	}

	for _, prn := range rsp.Printer {
		groups.Add(goipp.Group{
			Tag:   goipp.TagPrinterGroup,
			Attrs: ippEncodeAttrs(prn),
		})
	}

	msg := goipp.NewMessageWithGroups(rsp.Version, goipp.Code(rsp.Status),
		rsp.RequestID, groups)

	return msg
}

// Decode decodes CUPSGetPrintersResponse from goipp.Message.
func (rsp *CUPSGetPrintersResponse) Decode(msg *goipp.Message) error {
	rsp.Version = msg.Version
	rsp.RequestID = msg.RequestID
	rsp.Status = goipp.Status(msg.Code)

	err := ippDecodeAttrs(rsp, msg.Operation)
	if err != nil {
		return err
	}

	for _, grp := range msg.Groups {
		if grp.Tag == goipp.TagPrinterGroup && len(grp.Attrs) > 0 {
			prn := &PrinterAttributes{}
			err = ippDecodeAttrs(prn, grp.Attrs)
			if err != nil {
				return err
			}

			rsp.Printer = append(rsp.Printer, prn)
		}
	}

	return nil
}
