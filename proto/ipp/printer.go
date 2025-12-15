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

	"github.com/OpenPrinting/go-mfp/proto/ipp/iana"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/goipp"
)

// Printer implements the IPP printer.
type Printer struct {
	server        *Server                        // Underlying IPP server
	attrs         *PrinterAttributes             // Printer attributes
	attrSelection map[string]generic.Set[string] // Attr groups
}

// NewPrinter creates a new [Printer], which facilities and
// behavior is defined by the supplied [PrinterAttributes].
func NewPrinter(attrs *PrinterAttributes, options ServerOptions) *Printer {
	// Create the Printer structure
	server := NewServer(options)
	printer := &Printer{
		server:        server,
		attrs:         attrs,
		attrSelection: make(map[string]generic.Set[string]),
	}

	// Populate Printer.attrSelection
	all := generic.NewSet[string]()
	for name := range iana.PrinterDescription {
		all.Add(name)
	}
	for name := range iana.PrinterStatus {
		all.Add(name)
	}
	all.Del("media-col-database")

	jobTemplate := generic.NewSet[string]()
	all.ForEach(func(name string) {
		name2 := name + "-default"
		if iana.PrinterDescription[name2] != nil {
			jobTemplate.Add(name2)
		}

		name2 = name + "-supported"
		if iana.PrinterDescription[name2] != nil {
			jobTemplate.Add(name2)
		}
	})

	printerDescription := all.Clone()
	jobTemplate.ForEach(func(name string) {
		printerDescription.Del(name)
	})

	printer.attrSelection["all"] = all
	printer.attrSelection["printer-description"] = printerDescription
	printer.attrSelection["job-template"] = jobTemplate

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

	rsp := GetPrinterAttributesResponse{
		ResponseHeader: rq.ResponseHeader(goipp.StatusOk),
		Printer:        printer.attrs,
	}

	// Obtain all attributes.
	//
	// Here we encode GetPrinterAttributesResponse into the goipp.Message
	// with the only purpose to obtain printer attributes.
	attrs := rsp.Encode().Printer
	if printer.server.options.UseRawPrinterAttributes {
		attrs = printer.attrs.RawAttrs().All()
	}

	// Build set of supported attributes.
	supported := generic.NewSet[string]()
	for _, attr := range attrs {
		supported.Add(attr.Name)
	}

	// Prepare filter of returned attributes and build list
	// of unsupported attributes, if any.
	filter := generic.NewSet[string]()

	unsupported := generic.NewSet[string]()
	var unsupportedNames []string

	for _, name := range rq.RequestedAttributes {
		if group, ok := printer.attrSelection[name]; ok {
			filter.Merge(group)
		} else if supported.Contains(name) {
			filter.Add(name)
		} else if unsupported.TestAndAdd(name) {
			unsupportedNames = append(unsupportedNames, name)
		}
	}

	// Now collect actually returned attributes
	var returnedAttrs goipp.Attributes
	for _, attr := range attrs {
		if filter.Contains(attr.Name) {
			returnedAttrs = append(returnedAttrs, attr)
		}
	}

	// Rebuild the response.
	//
	// We don't need printer attributes to be encoded here, because we
	// will replace them directly in the message with the filtered list
	// of attributes. Hence rsp.Printer = nil.
	//
	// FIXME, from the architectural point of view this is really ugly.
	rsp.UnsupportedAttributes = unsupportedNames
	rsp.Printer = nil
	msg := rsp.Encode()

	// Set status code
	msg.Code = goipp.Code(goipp.StatusOk)
	if len(unsupportedNames) > 0 {
		msg.Code = goipp.Code(goipp.StatusOkIgnoredOrSubstituted)
	}

	// Rebuild msg.Groups
	msg.Printer = returnedAttrs
	msg.Groups = nil // Forces Groups to be rebuilt
	msg.Groups = msg.AttrGroups()

	return msg
}
