// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Protocol messages

package trace

import (
	"github.com/OpenPrinting/go-mfp/log"
)

// Message represents a single protocol message, which can
// be either request or response.
type Message interface {
	// Ext returns file extension for the protocol message files.
	Ext() string

	// Name returns the message name, which will be something
	// like "Get-Printer-Attributes" for the IPP request or
	// "successful-ok" for the IPP response, and something similar
	// for other protocols.
	Name() string

	// Message must be able to write itself to log
	log.Marshaler

	// Message must be able to write itself as a slice of
	// bytes, for saving in trace files.
	MarshalTrace() []byte
}
