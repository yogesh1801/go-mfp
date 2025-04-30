// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Printer type enu,

package ipp

import (
	"fmt"
	"strings"
)

// EnPrinterType is the CUPS extension, bits of the "printer-type"
// printer attribute. Though used as a bitmap, from the IPP point
// of view it is defined as type2 enum
type EnPrinterType int

// EnPrinterType bits:
const (
	// Local printer or class
	EnPrinterLocal EnPrinterType = 0x0000

	// This is printer class, not individual printer
	EnPrinterClass EnPrinterType = 0x0001

	// This is remote printer or class
	EnPrinterRemote EnPrinterType = 0x0002

	// Printer can do B&W printing
	EnPrinterBW EnPrinterType = 0x0004

	// Printer can do color printing
	EnPrinterColor EnPrinterType = 0x0008

	// Printer can do two-sided printing
	EnPrinterDuplex EnPrinterType = 0x0010

	// Printer can staple output
	EnPrinterStaple EnPrinterType = 0x0020

	// Printer can do copies in hardware
	EnPrinterCopies EnPrinterType = 0x0040

	// Printer can quickly collate copies
	EnPrinterCollate EnPrinterType = 0x0080

	// Printer can punch output
	EnPrinterPunch EnPrinterType = 0x0100

	// Printer can cover output
	EnPrinterCover EnPrinterType = 0x0200

	// Printer can bind output
	EnPrinterBind EnPrinterType = 0x0400

	// Printer can sort output
	EnPrinterSort EnPrinterType = 0x0800

	// Letter/Legal/A4-size media
	EnPrinterSmall EnPrinterType = 0x1000

	// Tabloid/B/C/A3/A2-size media
	EnPrinterMedium EnPrinterType = 0x2000

	// D/E/A1/A0-size media
	EnPrinterLarge EnPrinterType = 0x4000

	// Can print on rolls and custom-size media
	EnPrinterVariable EnPrinterType = 0x8000

	// Default printer on network
	EnPrinterDefault EnPrinterType = 0x20000

	// It's a fax queue
	EnPrinterFax EnPrinterType = 0x40000

	// Printer is rejecting jobs
	EnPrinterRejecting EnPrinterType = 0x80000

	// Printer is not shared
	EnPrinterNotShared EnPrinterType = 0x200000

	// Printer requires authentication
	EnPrinterAuthenticated EnPrinterType = 0x400000

	// Printer supports maintenance commands
	EnPrinterCommands EnPrinterType = 0x800000

	// Printer was discovered
	EnPrinterDiscovered EnPrinterType = 0x1000000

	// Scanner-only device
	EnPrinterScanner EnPrinterType = 0x2000000

	// Printer with scanning capabilities
	EnPrinterMfp EnPrinterType = 0x4000000

	// 3D printer
	EnPrinter3D EnPrinterType = 0x8000000
)

// String returns string representation for EnPrinter3D
func (bits EnPrinterType) String() string {
	names := []string{}
	for i := 0; i < 31; i++ {
		bit := EnPrinterType(1 << i)
		if bits&bit != 0 {
			name := enPrinterBitNames[bit]
			if name == "" {
				name = fmt.Sprintf("%#x", int(bit))
			}
			names = append(names, name)
		}
	}

	return strings.Join(names, ",")
}

// enPrinterBitNames contains bit names of EnPrinterType bits
var enPrinterBitNames = map[EnPrinterType]string{
	EnPrinterLocal:         "local",
	EnPrinterClass:         "class",
	EnPrinterRemote:        "remote",
	EnPrinterBW:            "bw",
	EnPrinterColor:         "color",
	EnPrinterDuplex:        "duplex",
	EnPrinterStaple:        "staple",
	EnPrinterCopies:        "copies",
	EnPrinterCollate:       "collate",
	EnPrinterPunch:         "punch",
	EnPrinterCover:         "cover",
	EnPrinterBind:          "bind",
	EnPrinterSort:          "sort",
	EnPrinterSmall:         "small",
	EnPrinterMedium:        "medium",
	EnPrinterLarge:         "large",
	EnPrinterVariable:      "variable",
	EnPrinterDefault:       "default",
	EnPrinterFax:           "fax",
	EnPrinterRejecting:     "rejecting",
	EnPrinterNotShared:     "notsharing",
	EnPrinterAuthenticated: "authenticated",
	EnPrinterCommands:      "commands",
	EnPrinterDiscovered:    "discovered",
	EnPrinterScanner:       "scanner",
	EnPrinterMfp:           "mfp",
	EnPrinter3D:            "3d",
}
