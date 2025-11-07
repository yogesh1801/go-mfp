// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP server hooks

package ipp

import (
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/goipp"
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

	// OnIPPRequest is called when the IPP request is
	// received.
	//
	// The hook can modify the [goipp.Message] request
	// in place or completely replace it by returning
	// the non-nil new value.
	OnIPPRequest func(*transport.ServerQuery,
		*goipp.Message) *goipp.Message

	// OnIPPResponse is called when the IPP response is
	// received.
	//
	// The hook can modify the [goipp.Message] response
	// in place or completely replace it by returning
	// the non-nil new value.
	OnIPPResponse func(*transport.ServerQuery,
		*goipp.Message) *goipp.Message
}
