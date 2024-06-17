// MFP - Miulti-Function Printers and scanners toolkit
// CUPS Client and Server
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS Client

package cups

import (
	"net/url"

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
func (c *Client) CUPSGetDefault(attrs []string) (
	*ipp.PrinterAttributes, error) {

	rq := &ipp.CUPSGetDefaultRequest{
		RequestHeader:       ipp.DefaultRequestHeader,
		RequestedAttributes: attrs,
	}

	rsp := &ipp.CUPSGetDefaultResponse{}

	err := c.IPPClient.Do(rq, rsp)
	if err != nil {
		return nil, err
	}

	return rsp.Printer, nil
}
