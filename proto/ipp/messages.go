// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP protocol messages

package ipp

import (
	"errors"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/goipp"
)

// Standard attribute groups for Get-Printer-Attributes.
const (
	// GetPrinterAttributesAll requests all printer attributes,
	// except the media-col-database.
	GetPrinterAttributesAll = "all"

	// GetPrinterAttributesJobTemplate requests the Job Template
	// Attributes.
	GetPrinterAttributesJobTemplate = "job-template"

	// GetPrinterAttributesPrinterDescription requests the
	// Printer Description Attributes.
	GetPrinterAttributesPrinterDescription = "printer-description"

	// GetPrinterAttributesMediaColDatabase requests the collection
	// of supported media types.
	//
	// Note, the "media-col-database" is not returned by the
	// printer unless explicitly requested, even if "all" attributes
	// are requested.
	GetPrinterAttributesMediaColDatabase = "media-col-database"
)

type (
	// GetPrinterAttributesRequest operation (0x000b) returns
	// the requested printer attributes.
	GetPrinterAttributesRequest struct {
		ObjectRawAttrs
		RequestHeader

		// Operation attributes
		PrinterURI          string               `ipp:"printer-uri,uri"`
		RequestedAttributes []string             `ipp:"requested-attributes,keyword"`
		DocumentFormat      optional.Val[string] `ipp:"document-format,mimeMediaType"`
	}

	// GetPrinterAttributesResponse is the CUPS-Get-Default Response.
	GetPrinterAttributesResponse struct {
		ObjectRawAttrs
		ResponseHeader

		// Other attributes.
		Printer *PrinterAttributes
	}
)

// ----- Get-Printer-Attributes methods -----

// GetOp returns GetPrinterAttributesRequest IPP Operation code.
func (rq *GetPrinterAttributesRequest) GetOp() goipp.Op {
	return goipp.OpGetPrinterAttributes
}

// KnownAttrs returns information about all known IPP attributes
// of the GetPrinterAttributesRequest
func (rq *GetPrinterAttributesRequest) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rq)
}

// Set sets [goipp.Attibute]. It updates the appropriate structure
// field and Object's raw attributes.
func (rq *GetPrinterAttributesRequest) Set(attr goipp.Attribute) error {
	return rq.set(attr, rq)
}

// Encode encodes GetPrinterAttributesRequest into the goipp.Message.
func (rq *GetPrinterAttributesRequest) Encode() *goipp.Message {
	enc := ippEncoder{}

	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: enc.Encode(rq),
		},
	}

	msg := goipp.NewMessageWithGroups(rq.Version, goipp.Code(rq.GetOp()),
		rq.RequestID, groups)

	return msg
}

// Decode decodes GetPrinterAttributesRequest from goipp.Message.
func (rq *GetPrinterAttributesRequest) Decode(
	msg *goipp.Message, opt DecodeOptions) error {

	rq.Version = msg.Version
	rq.RequestID = msg.RequestID

	dec := ippDecoder{opt: opt}
	err := dec.Decode(rq, msg.Operation)
	if err != nil {
		return err
	}

	return nil
}

// KnownAttrs returns information about all known IPP attributes
// of the GetPrinterAttributesResponse.
func (rsp *GetPrinterAttributesResponse) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rsp)
}

// Set sets [goipp.Attibute]. It updates the appropriate structure
// field and Object's raw attributes.
func (rsp *GetPrinterAttributesResponse) Set(attr goipp.Attribute) error {
	return rsp.set(attr, rsp)
}

// Encode encodes GetPrinterAttributesResponse into goipp.Message.
func (rsp *GetPrinterAttributesResponse) Encode() *goipp.Message {
	enc := ippEncoder{}

	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: enc.Encode(rsp),
		},
	}

	if rsp.Printer != nil {
		groups.Add(goipp.Group{
			Tag:   goipp.TagPrinterGroup,
			Attrs: enc.Encode(rsp.Printer),
		})
	}

	msg := goipp.NewMessageWithGroups(rsp.Version, goipp.Code(rsp.Status),
		rsp.RequestID, groups)

	return msg
}

// Decode decodes GetPrinterAttributesResponse from goipp.Message.
func (rsp *GetPrinterAttributesResponse) Decode(
	msg *goipp.Message, opt DecodeOptions) error {

	rsp.Version = msg.Version
	rsp.RequestID = msg.RequestID
	rsp.Status = goipp.Status(msg.Code)

	dec := ippDecoder{opt: opt}
	err := dec.Decode(rsp, msg.Operation)
	if err != nil {
		return err
	}

	if len(msg.Printer) == 0 {
		err = errors.New("missed printer attributes in response")
		return err
	}

	rsp.Printer, err = DecodePrinterAttributes(msg.Printer, opt)
	if err != nil {
		return err
	}

	return nil
}
