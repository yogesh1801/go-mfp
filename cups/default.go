// MFP - Miulti-Function Printers and scanners toolkit
// CUPS Client and Server
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Default values

package cups

import "github.com/OpenPrinting/go-mfp/transport"

// Default URLs
var (
	// DefaultUNIXURL is the default CUPS socket URL using
	// UNIX domain sockets
	DefaultUNIXURL = transport.MustParseURL("unix:/var/run/cups/cups.sock")

	// DefaultLocalhostURL is the default CUPS socket URL using
	// Localhost TCP connection
	DefaultLocalhostURL = transport.MustParseURL("ipp://localhost/")
)
