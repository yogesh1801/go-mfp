// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Default values

package transport

// Default URLs
var (
	// DefaultCupsUNIX is the default CUPS socket URL using
	// UNIX domain sockets
	DefaultCupsUNIX = MustParseURL("unix:/var/run/cups/cups.sock")

	// DefaultCupsLocalhost is the default CUPS socket URL using
	// Localhost TCP connection
	DefaultCupsLocalhost = MustParseURL("ipp://localhost/")
)
