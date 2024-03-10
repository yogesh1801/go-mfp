// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Client-side Printer object

package ippx

import (
	"net/http"
	"net/url"

	"github.com/OpenPrinting/goipp"
)

// Printer implements Client-side IPP Printer object.
type Printer struct {
	// HTTP stuff
	httpURL    *url.URL     // Parsed URL
	httpClient *http.Client // HTTP Client

	// Attributes, common for all requests
	IppVersion          goipp.Version // IPP protocol version
	PrinterURL          string        // Printer URL (ipp://...)
	AttrCharset         string        // "attributes-charset"
	AttrNaturalLanguage string        // "attributes-natural-language"
}

// PrinterConfig represents Printer configuration options.
// Used as parameter to the NewPrinter function.
type PrinterConfig struct {
	HTTPClient *http.Client  // HTTP Client
	IppVersion goipp.Version // IPP protocol version
}

// NewPrinter creates a new Printer object.
// If conf is nil, reasonable defaults are provided automatically
func NewPrinter(printerURL string, conf *PrinterConfig) (*Printer, error) {
	// Parse and validate Printer URL
	httpURL, _, err := urlParse(printerURL)
	if err != nil {
		return nil, err
	}

	// Create Printer object.
	p := &Printer{
		httpURL:             httpURL,
		httpClient:          DefaultHTTPClient,
		IppVersion:          goipp.DefaultVersion,
		PrinterURL:          printerURL,
		AttrCharset:         "utf-8",
		AttrNaturalLanguage: "en-US",
	}

	// Apply PrinterConfig.
	if conf != nil {
		if conf.IppVersion != 0 {
			p.IppVersion = conf.IppVersion
		}

		if conf.HTTPClient != nil {
			p.httpClient = conf.HTTPClient
		}
	}

	return p, nil
}

// GetPrinterAttributes queries Printer attributed
func (p *Printer) GetPrinterAttributes(attrs []string) (
	*PrinterAttributes, error) {
	return nil, nil
}
