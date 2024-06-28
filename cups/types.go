// MFP - Miulti-Function Printers and scanners toolkit
// CUPS Client and Server
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common types

package cups

// Default values for common types:
var (
	DefaultGetPrintersSelection = &GetPrintersSelection{}
)

// GetPrintersSelection configures a selection of printers returned
// by Client.CUPSGetPrinter
type GetPrintersSelection struct {
	// Printer name (also, queue name) is the unique name, under
	// which printer is registered in the CUPS system.
	//
	// If this parameter is not empty, it specifies the first printer
	// name to be returned.
	FirstPrinterName string

	// If not zero, specifies maximum number of printers to be returned.
	Limit int

	// Each printer in the CUPS system has its own PrinterID,
	// which is system-unique positive number, assigned to the
	// printer when it is added to the system.
	//
	// If this parameter is not zero, only the printer with specified
	// ID (and matching other criteria) will be returned.
	PrinterID int

	// PrinterLocation is the arbitrary string which can be
	// configured on the printer device by owner or system
	// administrator. Its purpose is to simplify selection of
	// particular devices in a big network. It corresponds to the
	// "printer-location" IPP attribute and may sound like
	// "Printers on 1st floor" or "Printers at reception".
	//
	// If this parameter is not empty, only devices with the specified
	// locatiom will be returned.
	PrinterLocation string

	// TODO
	PrinterType     int
	PrinterTypeMask int

	// If not empty, only printers accessible to that user will
	// be returned. User name is the user's **login** name,
	User string
}
