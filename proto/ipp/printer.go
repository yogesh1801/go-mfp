// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP printer implementation.

package ipp

import (
	"context"
	"io"
	"net/http"

	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/proto/ipp/iana"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/goipp"
)

// Printer implements the IPP printer.
type Printer struct {
	server        *Server                        // Underlying IPP server
	attrs         *PrinterAttributes             // Printer attributes
	attrSelection map[string]generic.Set[string] // Attr groups
	q             *queue                         // Job queue
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
		q:             newQueue(),
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
	server.RegisterHandler(NewHandler(printer.handleValidateJob))
	server.RegisterHandler(NewHandler(printer.handleCreateJob))
	server.RegisterHandler(NewHandler(printer.handleSendDocument))

	return printer
}

// ServeHTTP handles incoming HTTP request. It implements
// [http.Handler] interface.
func (printer *Printer) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	printer.server.ServeHTTP(w, rq)
}

// handleGetPrinterAttributes handles Get-Printer-Attributes request.
func (printer *Printer) handleGetPrinterAttributes(
	ctx context.Context,
	rq *GetPrinterAttributesRequest) (*goipp.Message, error) {

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

	return msg, nil
}

// handleValidateJob handles Validate-Job request.
func (printer *Printer) handleValidateJob(
	ctx context.Context,
	rq *ValidateJobRequest) (*goipp.Message, error) {

	rsp := ValidateJobResponse{
		ResponseHeader: rq.ResponseHeader(goipp.StatusOk),
	}

	return rsp.Encode(), nil
}

// handleCreateJob handles Create-Job request.
func (printer *Printer) handleCreateJob(
	ctx context.Context,
	rq *CreateJobRequest) (*goipp.Message, error) {

	// Create new job
	j := newJob(&rq.JobCreateOperation, rq.Job)
	j.Lock()
	defer j.Unlock()

	printer.q.Push(j)

	// Prepare the CreateJobResponse
	rsp := CreateJobResponse{
		ResponseHeader: rq.ResponseHeader(goipp.StatusOk),
		Job: &JobStatus{
			JobID:           j.JobID,
			JobState:        j.JobState,
			JobStateReasons: j.JobStateReasons,
			JobURI:          j.JobURI,
		},
	}

	return rsp.Encode(), nil
}

// handleCreateJob handles Send-Document request.
func (printer *Printer) handleSendDocument(
	ctx context.Context,
	rq *SendDocumentRequest) (*goipp.Message, error) {

	// Lookup the job
	var j *job

	switch {
	case rq.PrinterURI != nil && rq.JobID != nil:
		j = printer.q.JobByID(*rq.JobID)
		if j == nil {
			err := NewErrIPPFromRequest(rq,
				goipp.StatusErrorNotFound,
				"job not found (job-id=%d)", *rq.JobID)
			return nil, err
		}

	case rq.JobURI != nil:
		j = printer.q.JobByURI(*rq.JobURI)
		if j == nil {
			err := NewErrIPPFromRequest(rq,
				goipp.StatusErrorNotFound,
				"job not found (job-uri=%q)", *rq.JobURI)
			return nil, err
		}

	default:
		err := NewErrIPPFromRequest(rq,
			goipp.StatusErrorBadRequest,
			"missed job-id and job-uri attributes")
		return nil, err
	}

	j.Lock()
	defer j.Unlock()

	// Check job state
	if j.JobState != EnJobStatePendingHeld {
		err := NewErrIPPFromRequest(rq,
			goipp.StatusErrorNotPossible,
			"job is not in pending-held state")
		return nil, err
	}

	if j.SendDocumentActive {
		err := NewErrIPPFromRequest(rq,
			goipp.StatusErrorNotPossible,
			"Send-Document already in progress")
		return nil, err
	}

	// Consume the document body
	//
	// FIXME -- this is just stub
	j.SendDocumentActive = true
	j.Unlock()

	n, err := io.Copy(io.Discard, rq.Body)
	if err == nil {
		log.Debug(ctx, "Send-Document: %d bytes received", n)
	} else {
		log.Error(ctx, "Send-Document: %s", err)
	}

	j.Lock()
	j.SendDocumentActive = false

	// Generate response
	rsp := &SendDocumentResponse{
		Job: &JobStatus{
			JobID:           j.JobID,
			JobState:        j.JobState,
			JobStateReasons: j.JobStateReasons,
			JobURI:          j.JobURI,
		},
	}

	return rsp.Encode(), nil
}
