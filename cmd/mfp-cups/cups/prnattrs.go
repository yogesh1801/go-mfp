// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Printer information pretty-printer

package cups

import (
	"fmt"
	"io"

	"github.com/OpenPrinting/go-mfp/proto/ipp"
	"github.com/OpenPrinting/goipp"
)

// prnAttrsRequested lists attributes that provide a general printer
// information, hence they are always requested by commands like
// "get-default", "get-printers" and similar.
var prnAttrsRequested = []string{
	"device-uri",
	"printer-id",
	"printer-is-shared",
	"printer-is-temporary",
	"printer-name",
	"printer-type",
	"printer-uri-supported",
}

// prnAttrsFormat pretty-prints [ipp.PrinterAttributes]
func prnAttrsFormat(w io.Writer, prn *ipp.PrinterAttributes) {
	fmt.Fprintf(w, "%s:\n", prn.PrinterName)

	fmt.Fprintf(w, "  General information:\n")
	fmt.Fprintf(w, "    URL:          %s\n", prn.PrinterURISupported)
	fmt.Fprintf(w, "    Device URI:   %s\n", prn.DeviceURI)
	fmt.Fprintf(w, "    ID:           %d\n", prn.PrinterID)
	fmt.Fprintf(w, "    Shared:       %v\n", prn.PrinterIsShared)
	fmt.Fprintf(w, "    Temporary:    %v\n", prn.PrinterIsTemporary)
	fmt.Fprintf(w, "    Printer Type: 0x%x\n", int(prn.PrinterType))
	fmt.Fprintf(w, "    Decoded Type: %s\n", prn.PrinterType)
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "  Printer attributes:\n")

	f := goipp.NewFormatter()
	f.SetIndent(4)
	f.FmtAttributes(prn.RawAttrs().All().Clone())

	f.WriteTo(w)
}
