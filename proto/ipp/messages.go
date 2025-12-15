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
		OperationGroup

		// Operation attributes
		PrinterURI          string               `ipp:"printer-uri"`
		RequestedAttributes []string             `ipp:"requested-attributes"`
		DocumentFormat      optional.Val[string] `ipp:"document-format"`
	}

	// GetPrinterAttributesResponse is the CUPS-Get-Default Response.
	GetPrinterAttributesResponse struct {
		ObjectRawAttrs
		ResponseHeader
		OperationGroup

		// Names of unsupported attributes
		UnsupportedAttributes []string

		// Returned printer attributes
		Printer *PrinterAttributes
	}

	// ValidateJobRequest operation (0x0004) performs the print job
	// validation as if the document was actually printed, bypassing
	// the actual fetching and printing of the document data.
	ValidateJobRequest struct {
		ObjectRawAttrs
		RequestHeader

		OperationGroup

		// Operation attributes
		PrinterURI         string `ipp:"printer-uri"`
		RequestingUserName string `ipp:"requesting-user-name"`

		Job *JobAttributes
	}

	// ValidateJobResponse is the Validate-Job response.
	ValidateJobResponse struct {
		ObjectRawAttrs
		ResponseHeader
		OperationGroup

		// Unsupported attributes, if any
		UnsupportedAttributes goipp.Attributes
	}
)

// ----- Get-Printer-Attributes methods -----

// GetOp returns GetPrinterAttributesRequest IPP Operation code.
func (rq *GetPrinterAttributesRequest) GetOp() goipp.Op {
	return goipp.OpGetPrinterAttributes
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

// Encode encodes GetPrinterAttributesResponse into goipp.Message.
func (rsp *GetPrinterAttributesResponse) Encode() *goipp.Message {
	enc := ippEncoder{}

	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: enc.Encode(rsp),
		},
	}

	if len(rsp.UnsupportedAttributes) > 0 {
		names := make(goipp.Values, 0, len(rsp.UnsupportedAttributes))
		for _, name := range rsp.UnsupportedAttributes {
			names.Add(goipp.TagKeyword, goipp.String(name))
		}

		attr := goipp.Attribute{
			Name:   "requested-attributes",
			Values: names,
		}

		groups.Add(goipp.Group{
			Tag:   goipp.TagUnsupportedGroup,
			Attrs: goipp.Attributes{attr},
		})
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

// ----- Validate-Job methods -----

// GetOp returns ValidateJobRequest IPP Operation code.
func (rq *ValidateJobRequest) GetOp() goipp.Op {
	return goipp.OpValidateJob
}

// Encode encodes ValidateJobRequest into the goipp.Message.
func (rq *ValidateJobRequest) Encode() *goipp.Message {
	enc := ippEncoder{}

	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: enc.Encode(rq),
		},

		{
			Tag:   goipp.TagJobGroup,
			Attrs: enc.Encode(rq.Job),
		},
	}

	msg := goipp.NewMessageWithGroups(rq.Version, goipp.Code(rq.GetOp()),
		rq.RequestID, groups)

	return msg
}

// Decode decodes ValidateJobRequest from goipp.Message.
func (rq *ValidateJobRequest) Decode(
	msg *goipp.Message, opt DecodeOptions) error {

	rq.Version = msg.Version
	rq.RequestID = msg.RequestID

	dec := ippDecoder{opt: opt}

	err := dec.Decode(rq, msg.Operation)
	if err != nil {
		return err
	}

	rq.Job, err = DecodeJobAttributes(msg.Printer, opt)
	if err != nil {
		return err
	}

	return nil
}

// Encode encodes ValidateJobResponse into goipp.Message.
func (rsp *ValidateJobResponse) Encode() *goipp.Message {
	enc := ippEncoder{}

	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: enc.Encode(rsp),
		},
	}

	if len(rsp.UnsupportedAttributes) > 0 {
		groups = append(groups, goipp.Group{
			Tag:   goipp.TagUnsupportedGroup,
			Attrs: rsp.UnsupportedAttributes,
		})
	}

	msg := goipp.NewMessageWithGroups(rsp.Version, goipp.Code(rsp.Status),
		rsp.RequestID, groups)

	return msg
}

// Decode decodes ValidateJobResponse from goipp.Message.
func (rsp *ValidateJobResponse) Decode(
	msg *goipp.Message, opt DecodeOptions) error {

	rsp.Version = msg.Version
	rsp.RequestID = msg.RequestID
	rsp.Status = goipp.Status(msg.Code)
	rsp.UnsupportedAttributes = msg.Unsupported

	return nil
}
