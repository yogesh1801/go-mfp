// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS requests and responses

package ipp

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/goipp"
)

type (
	// CUPSGetDefaultRequest operation (0x4001) returns the default printer URI
	// and attributes.
	CUPSGetDefaultRequest struct {
		ObjectRawAttrs
		RequestHeader

		// Operation attributes
		RequestedAttributes []string `ipp:"requested-attributes,keyword"`
	}

	// CUPSGetDefaultResponse is the CUPS-Get-Default Response.
	CUPSGetDefaultResponse struct {
		ObjectRawAttrs
		ResponseHeader

		// Other attributes.
		Printer *PrinterAttributes
	}

	// CUPSGetPrintersRequest operation (0x4002) returns the printer
	// attributes for every printer known to the system.
	CUPSGetPrintersRequest struct {
		ObjectRawAttrs
		RequestHeader

		// Operation attributes
		FirstPrinterName    string   `ipp:"?first-printer-name"`
		Limit               int      `ipp:"?limit"`
		PrinterID           int      `ipp:"?printer-id"`
		PrinterLocation     string   `ipp:"?printer-location,text"`
		PrinterType         int      `ipp:"?printer-type,enum"`
		PrinterTypeMask     int      `ipp:"?printer-type-mask,enum"`
		RequestedUserName   string   `ipp:"?requested-user-name,name"`
		RequestedAttributes []string `ipp:"?requested-attributes,keyword"`
	}

	// CUPSGetPrintersResponse is the CUPS-Get-Printers Response.
	CUPSGetPrintersResponse struct {
		ObjectRawAttrs
		ResponseHeader

		// Other attributes.
		Printer []*PrinterAttributes
	}

	// CUPSGetDevicesRequest operation (0x400b) performs search
	// for available printers and returns all of the supported
	// device-uri's
	CUPSGetDevicesRequest struct {
		ObjectRawAttrs
		RequestHeader

		// Operational attributes
		ExcludeSchemes      []string `ipp:"?exclude-schemes,name"`
		IncludeSchemes      []string `ipp:"?include-schemes,name"`
		Limit               int      `ipp:"?limit,1:MAX"`
		RequestedAttributes []string `ipp:"?requested-attributes,keyword"`
		Timeout             int      `ipp:"?timeout,1:MAX"`
	}

	// CUPSGetDevicesResponse is the CUPS-Get-Devices Response.
	CUPSGetDevicesResponse struct {
		ObjectRawAttrs
		ResponseHeader

		// Other attributes.
		Printer []*DeviceAttributes
	}

	// CUPSGetPPDsRequest operation (0x400c) returns list of available PPDs
	CUPSGetPPDsRequest struct {
		ObjectRawAttrs
		RequestHeader

		// Operational attributes
		ExcludeSchemes      []string          `ipp:"?exclude-schemes,name"`
		IncludeSchemes      []string          `ipp:"?include-schemes,name"`
		Limit               int               `ipp:"?limit,1:MAX"`
		PpdMake             string            `ipp:"?ppd-make,text"`
		PpdMakeAndModel     string            `ipp:"?ppd-make-and-model,text"`
		ModelNumber         optional.Val[int] `ipp:"?ppd-model-number"`
		PpdNaturalLanguage  string            `ipp:"?ppd-natural-language,text"`
		PpdProduct          string            `ipp:"?ppd-product,text"`
		PpdPsversion        string            `ipp:"?ppd-psversion,text"`
		PpdType             string            `ipp:"?ppd-type,keyword"`
		RequestedAttributes []string          `ipp:"?requested-attributes,keyword"`
	}

	// CUPSGetPPDsResponse is the CUPS-Get-PPDs Response.
	CUPSGetPPDsResponse struct {
		ObjectRawAttrs
		ResponseHeader

		// Other attributes.
		PPDs []*PpdAttributes
	}

	// CUPSGetPPDRequest operation (0x400f) returns PPD file from
	// the server.
	CUPSGetPPDRequest struct {
		ObjectRawAttrs
		RequestHeader

		// Operational attributes
		//
		// Use PrinterURI to specify particular print queue
		// or PPDName to request PPD file by its name.
		PrinterURI string `ipp:"?printer-uri,uri"`
		PPDName    string `ipp:"?ppd-name,name"`
	}

	// CUPSGetPPDResponse is the CUPS-Get-PPD Response.
	//
	// If the PPD file is found, goipp.StatusOk is returned with the PPD
	// file represented by the ResponseHeader.Body.
	//
	// If the PPD file cannot be served by the local server because the
	// printer-uri attribute points to an external printer, a
	// goipp.StatusCupsSeeOther is returned and PrinterURI contains
	// the correct URI to use.
	//
	// If the PPD file does not exist, goipp.StatusErrorNotFound is
	// returned.
	CUPSGetPPDResponse struct {
		ObjectRawAttrs
		ResponseHeader

		// Operational attributes
		PrinterURI string `ipp:"?printer-uri,uri"`
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

// Decode decodes CUPSGetDefaultRequest from goipp.Message.
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

// Decode decodes CUPSGetPrintersRequest from goipp.Message.
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

// ----- CUPS-Get-Devices methods -----

// GetOp returns CUPSGetDevicesRequest IPP Operation code.
func (rq *CUPSGetDevicesRequest) GetOp() goipp.Op {
	return goipp.OpCupsGetDevices
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetDevicesRequest
func (rq *CUPSGetDevicesRequest) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rq)
}

// Encode encodes CUPSGetDevicesRequest into the goipp.Message.
func (rq *CUPSGetDevicesRequest) Encode() *goipp.Message {
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

// Decode decodes CUPSGetDevicesRequest from goipp.Message.
func (rq *CUPSGetDevicesRequest) Decode(msg *goipp.Message) error {
	rq.Version = msg.Version
	rq.RequestID = msg.RequestID

	err := ippDecodeAttrs(rq, msg.Operation)
	if err != nil {
		return err
	}

	return nil
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetDevicesResponse.
func (rsp *CUPSGetDevicesResponse) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rsp)
}

// Encode encodes CUPSGetDevicesResponse into goipp.Message.
func (rsp *CUPSGetDevicesResponse) Encode() *goipp.Message {
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

// Decode decodes CUPSGetDevicesResponse from goipp.Message.
func (rsp *CUPSGetDevicesResponse) Decode(msg *goipp.Message) error {
	rsp.Version = msg.Version
	rsp.RequestID = msg.RequestID
	rsp.Status = goipp.Status(msg.Code)

	err := ippDecodeAttrs(rsp, msg.Operation)
	if err != nil {
		return err
	}

	for _, grp := range msg.Groups {
		if grp.Tag == goipp.TagPrinterGroup && len(grp.Attrs) > 0 {
			dev := &DeviceAttributes{}
			err = ippDecodeAttrs(dev, grp.Attrs)
			if err != nil {
				return err
			}

			rsp.Printer = append(rsp.Printer, dev)
		}
	}

	return nil
}

// ----- CUPS-Get-PPDs methods -----

// GetOp returns CUPSGetPPDsRequest IPP Operation code.
func (rq *CUPSGetPPDsRequest) GetOp() goipp.Op {
	return goipp.OpCupsGetPpds
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetPPDsRequest
func (rq *CUPSGetPPDsRequest) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rq)
}

// Encode encodes CUPSGetPPDsRequest into the goipp.Message.
func (rq *CUPSGetPPDsRequest) Encode() *goipp.Message {
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

// Decode decodes CUPSGetPPDsRequest from goipp.Message.
func (rq *CUPSGetPPDsRequest) Decode(msg *goipp.Message) error {
	rq.Version = msg.Version
	rq.RequestID = msg.RequestID

	err := ippDecodeAttrs(rq, msg.Operation)
	if err != nil {
		return err
	}

	return nil
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetPPDsResponse.
func (rsp *CUPSGetPPDsResponse) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rsp)
}

// Encode encodes CUPSGetPPDsResponse into goipp.Message.
func (rsp *CUPSGetPPDsResponse) Encode() *goipp.Message {
	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: ippEncodeAttrs(rsp),
		},
	}

	for _, ppd := range rsp.PPDs {
		groups.Add(goipp.Group{
			Tag:   goipp.TagPrinterGroup,
			Attrs: ippEncodeAttrs(ppd),
		})
	}

	msg := goipp.NewMessageWithGroups(rsp.Version, goipp.Code(rsp.Status),
		rsp.RequestID, groups)

	return msg
}

// Decode decodes CUPSGetPPDsResponse from goipp.Message.
func (rsp *CUPSGetPPDsResponse) Decode(msg *goipp.Message) error {
	rsp.Version = msg.Version
	rsp.RequestID = msg.RequestID
	rsp.Status = goipp.Status(msg.Code)

	err := ippDecodeAttrs(rsp, msg.Operation)
	if err != nil {
		return err
	}

	for _, grp := range msg.Groups {
		if grp.Tag == goipp.TagPrinterGroup && len(grp.Attrs) > 0 {
			ppd := &PpdAttributes{}
			err = ippDecodeAttrs(ppd, grp.Attrs)
			if err != nil {
				return err
			}

			rsp.PPDs = append(rsp.PPDs, ppd)
		}
	}

	return nil
}

// ----- CUPS-Get-PPD methods -----

// GetOp returns CUPSGetPPDRequest IPP Operation code.
func (rq *CUPSGetPPDRequest) GetOp() goipp.Op {
	return goipp.OpCupsGetPpd
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetPPDRequest
func (rq *CUPSGetPPDRequest) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rq)
}

// Encode encodes CUPSGetPPDRequest into the goipp.Message.
func (rq *CUPSGetPPDRequest) Encode() *goipp.Message {
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

// Decode decodes CUPSGetPPDRequest from goipp.Message.
func (rq *CUPSGetPPDRequest) Decode(msg *goipp.Message) error {
	rq.Version = msg.Version
	rq.RequestID = msg.RequestID

	err := ippDecodeAttrs(rq, msg.Operation)
	if err != nil {
		return err
	}

	return nil
}

// KnownAttrs returns information about all known IPP attributes
// of the CUPSGetPPDResponse.
func (rsp *CUPSGetPPDResponse) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(rsp)
}

// Encode encodes CUPSGetPPDResponse into goipp.Message.
func (rsp *CUPSGetPPDResponse) Encode() *goipp.Message {
	groups := goipp.Groups{
		{
			Tag:   goipp.TagOperationGroup,
			Attrs: ippEncodeAttrs(rsp),
		},
	}

	msg := goipp.NewMessageWithGroups(rsp.Version, goipp.Code(rsp.Status),
		rsp.RequestID, groups)

	return msg
}

// Decode decodes CUPSGetPPDResponse from goipp.Message.
func (rsp *CUPSGetPPDResponse) Decode(msg *goipp.Message) error {
	rsp.Version = msg.Version
	rsp.RequestID = msg.RequestID
	rsp.Status = goipp.Status(msg.Code)

	err := ippDecodeAttrs(rsp, msg.Operation)
	if err != nil {
		return err
	}

	return nil
}
