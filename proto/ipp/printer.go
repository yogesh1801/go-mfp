// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP printer implementation.

package ipp

import (
	"net/http"

	"github.com/OpenPrinting/goipp"
)

// Printer implements the IPP printer.
type Printer struct {
	server *Server            // Underlying IPP server
	attrs  *PrinterAttributes // Printer attributes
}

// NewPrinter creates a new [Printer], which facilities and
// behavior is defined by the supplied [PrinterAttributes].
func NewPrinter(attrs *PrinterAttributes, options ServerOptions) *Printer {
	// Create the Printer structure
	server := NewServer(options)
	printer := &Printer{
		server: server,
		attrs:  attrs,
	}

	// Install request handlers
	server.RegisterHandler(NewHandler(printer.handleGetPrinterAttributes))

	return printer
}

// ServeHTTP handles incoming HTTP request. It implements
// [http.Handler] interface.
func (printer *Printer) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	printer.server.ServeHTTP(w, rq)
}

// handleGetPrinterAttributes handles Get-Printer-Attributes request.
func (printer *Printer) handleGetPrinterAttributes(
	rq *GetPrinterAttributesRequest) *goipp.Message {

	rsp := &GetPrinterAttributesResponse{
		ResponseHeader: rq.ResponseHeader(goipp.StatusOk),
		Printer:        printer.attrs,
	}

	msg := rsp.Encode()
	if printer.server.options.UseRawPrinterAttributes {
		// Replace printer attributes in response
		msg.Printer = printer.attrs.RawAttrs().All()

		// Rebuild msg.Groups
		msg.Groups = nil
		msg.Groups = msg.AttrGroups()
	}

	return msg
}
