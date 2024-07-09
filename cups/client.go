// MFP - Miulti-Function Printers and scanners toolkit
// CUPS Client and Server
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS Client

package cups

import (
	"context"
	"net/url"
	"time"

	"github.com/alexpevzner/mfp/ipp"
	"github.com/alexpevzner/mfp/transport"
)

// Client represents the CUPS client.
type Client struct {
	IPPClient *ipp.Client // Underlying IPP client
}

// NewClient creates a new CUPS client.
//
// If tr is nil, [transport.NewTransport] will be used to create
// a new transport.
func NewClient(u *url.URL, tr *transport.Transport) *Client {
	return &Client{
		IPPClient: ipp.NewClient(u, tr),
	}
}

// CUPSGetDefault returns information on default printer.
// The attrs attribute allows to specify list of requested attributes.
func (c *Client) CUPSGetDefault(ctx context.Context,
	attrs []string) (*ipp.PrinterAttributes, error) {

	rq := &ipp.CUPSGetDefaultRequest{
		RequestHeader:       ipp.DefaultRequestHeader,
		RequestedAttributes: attrs,
	}

	rsp := &ipp.CUPSGetDefaultResponse{}

	err := c.IPPClient.Do(ctx, rq, rsp)
	if err != nil {
		return nil, err
	}

	return rsp.Printer, nil
}

// CUPSGetPrinters returns printer attributes for printers known
// to the system.
//
// If [GetPrintersSelection] argument is not nil, it allows to
// specify a subset of printers to be returned.
//
// The attrs attribute allows to specify list of requested attributes.
func (c *Client) CUPSGetPrinters(ctx context.Context,
	sel *GetPrintersSelection, attrs []string) (
	[]*ipp.PrinterAttributes, error) {

	if sel == nil {
		sel = DefaultGetPrintersSelection
	}

	rq := &ipp.CUPSGetPrintersRequest{
		RequestHeader:       ipp.DefaultRequestHeader,
		FirstPrinterName:    sel.FirstPrinterName,
		Limit:               sel.Limit,
		PrinterID:           sel.PrinterID,
		PrinterLocation:     sel.PrinterLocation,
		PrinterType:         sel.PrinterType,
		PrinterTypeMask:     sel.PrinterTypeMask,
		RequestedUserName:   sel.User,
		RequestedAttributes: attrs,
	}

	rsp := &ipp.CUPSGetPrintersResponse{}

	err := c.IPPClient.Do(ctx, rq, rsp)
	if err != nil {
		return nil, err
	}

	return rsp.Printer, nil
}

// CUPSGetDevices performs search for available devices and returns
// found devices.
//
// If [GetDevicesSelection] argument is not nil, it allows to
// specify a subset of devices to be returned.
//
// The attrs attribute allows to specify list of requested attributes.
func (c *Client) CUPSGetDevices(ctx context.Context,
	sel *GetDevicesSelection, attrs []string) (
	[]*ipp.DeviceAttributes, error) {

	if sel == nil {
		sel = DefaultGetDevicesSelection
	}

	tm := 0
	if sel.Timeout != 0 {
		tm = int((sel.Timeout + time.Second - 1) / time.Second)
	}

	rq := &ipp.CUPSGetDevicesRequest{
		RequestHeader:       ipp.DefaultRequestHeader,
		ExcludeSchemes:      sel.ExcludeSchemes,
		IncludeSchemes:      sel.IncludeSchemes,
		Limit:               sel.Limit,
		Timeout:             tm,
		RequestedAttributes: attrs,
	}

	rsp := &ipp.CUPSGetDevicesResponse{}

	err := c.IPPClient.Do(ctx, rq, rsp)
	if err != nil {
		return nil, err
	}

	return rsp.Printer, nil
}
