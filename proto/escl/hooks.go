// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL server hooks

package escl

import (
	"io"

	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// HookAction hints low-level [ServerHooks] callbacks on action
// they were called for.
type HookAction int

// Known HookAction values:
const (
	HookUnknownAction HookAction = iota
	HookScannerCapabilities
	HookScannerStatus
	HookScanJobs
	HookNextDocument
	HookDelete
)

// ServerHooks allows to specify set of hooks (callbacks) that
// will be called during the request processing and can modify
// the request handling.
//
// Every hook is optional and can be set to nil.
//
// If hook calls the [transport.ServerQuery.WriteHeader] function,
// the query considered completed and further processing is
// not performed.
type ServerHooks struct {
	// OnHTTPRequest is called when the HTTP request is just
	// received.
	OnHTTPRequest func(*transport.ServerQuery)

	// OnScannerCapabilitiesRequest is called when the eSCL
	// ScannerCapabilities request is received.
	OnScannerCapabilitiesRequest func(*transport.ServerQuery)

	// OnScannerCapabilitiesResponse is called when the eSCL
	// ScannerCapabilities response is generated.
	//
	// The hook can modify the [ScannerCapabilities] in place or
	// completely replace it by returning the non-nil new value.
	OnScannerCapabilitiesResponse func(*transport.ServerQuery,
		*ScannerCapabilities) *ScannerCapabilities

	// OnScannerStatusRequest is called when the eSCL
	// ScannerStatus request is received.
	OnScannerStatusRequest func(*transport.ServerQuery)

	// OnScannerStatusResponse is called when the eSCL
	// ScannerStatus response is generated.
	//
	// The hook can modify the [ScannerStatus] in place or
	// completely replace it by returning the non-nil new value.
	OnScannerStatusResponse func(*transport.ServerQuery,
		*ScannerStatus) *ScannerStatus

	// OnScanJobsRequest is called when the eSCL ScanJobs request
	// is received.
	//
	// The hook can modify the [ScanSettings] in place or
	// completely replace it by returning the non-nil new value.
	OnScanJobsRequest func(*transport.ServerQuery,
		*ScanSettings) *ScanSettings

	// OnScanJobsResponse is called when the eSCL ScanJobs response
	// is generated.
	//
	// The hook can modify the resulting JobURI by returning
	// the non-empty new value.
	//
	// The ScanSettings parameters is provided just for information.
	OnScanJobsResponse func(*transport.ServerQuery,
		*ScanSettings) (joburi string)

	// OnNextDocumentRequest is called when the eSCL NextDocument
	// request is received.
	//
	// The hook can replace the effective JobURI by returning
	// the non-empty new value.
	OnNextDocumentRequest func(query *transport.ServerQuery,
		joburi string) string

	// OnNextDocumentResponse is called when the eSCL NextDocument
	// response is generated.
	//
	// The hook can replace the resulting [io.ReadCloser]
	// it by returning the non-nil new value.
	OnNextDocumentResponse func(*transport.ServerQuery,
		io.ReadCloser) io.ReadCloser

	// OnDeleteRequest is called when the eSCL DELETE
	// request is received
	//
	// The hook can replace the effective JobURI by returning
	// the non-empty new value.
	OnDeleteRequest func(query *transport.ServerQuery,
		joburi string) string

	// OnXMLRequest is called when XML request is received.
	//
	// The hook can replace the resulting [xmldoc.Element]
	// it by returning the non-zero new value.
	OnXMLRequest func(*transport.ServerQuery,
		HookAction, xmldoc.Element) xmldoc.Element

	// OnXMLResponse is called when XML response is generated.
	//
	// The hook can replace the resulting [xmldoc.Element]
	// it by returning the non-zero new value.
	OnXMLResponse func(*transport.ServerQuery,
		HookAction, xmldoc.Element) xmldoc.Element
}
