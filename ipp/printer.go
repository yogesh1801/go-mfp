// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Client-side Printer object

package ipp

// Printer implements Client-side IPP Printer object.
type Printer struct {
	*Client
}

// NewPrinter creates a new Printer object.
// If conf is nil, reasonable defaults are provided automatically
func NewPrinter(printerURL string, conf *ClientConfig) (*Printer, error) {
	// Create a client
	client, err := NewClient(printerURL, conf)
	if err != nil {
		return nil, err
	}

	// Create Printer object
	p := &Printer{
		Client: client,
	}

	return p, nil
}

// GetPrinterAttributes queries Printer attributed
func (p *Printer) GetPrinterAttributes(attrs []string) (
	*PrinterAttributes, error) {
	return nil, nil
}
