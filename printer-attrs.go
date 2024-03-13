// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Printer Attributes

package ippx

import (
	"github.com/OpenPrinting/goipp"
)

// Default values for common attributes
var (
	// DefaultCharsetConfigured is the default value for
	// ""charset-configured" printer attribute
	DefaultCharsetConfigured = "utf-8"

	// DefaultCharsetSupported is the default value for
	// ""charset-supported" printer attribute
	DefaultCharsetSupported = []string{DefaultCharsetConfigured}

	// DefaultIppVersionsSupported is the default value for
	// "ipp-features-supported" printer attribute
	DefaultIppVersionsSupported = []goipp.Version{
		goipp.MakeVersion(2, 0),
		goipp.MakeVersion(1, 0),
		goipp.MakeVersion(1, 1),
	}
)

// Standard keywords for CompressionSupported attribute
const (
	CompressionDeflate  = "deflate"  // RFC 1951 ZIP inflate/deflate
	CompressionGzip     = "gzip"     // RFC 1952 GNU zip
	CompressionCompress = "compress" // RFC 1977 UNIX compression
)

// PrinterAttributes represents IPP Printer Attributes
type PrinterAttributes struct {
	// RFC8011, 5.4: Printer Description and Status Attributes
	CharsetConfigured       string                         `ipp:"charset-configured,charset"`
	CharsetSupported        []string                       `ipp:"charset-supported,charset"`
	ColorSupported          bool                           `ipp:"color-supported"`
	CompressionSupported    []string                       `ipp:"compression-supported,keyword"`
	IppFeaturesSupported    []string                       `ipp:"ipp-features-supported,keyword"`
	IppVersionsSupported    []goipp.Version                `ipp:"ipp-versions-supported"`
	MediaSizeSupported      []PrinterMediaSizeSupported    `ipp:"media-size-supported,norange"`
	MediaSizeSupportedRange PrinterMediaSizeSupportedRange `ipp:"media-size-supported,range"`
	OperationsSupported     []goipp.Op                     `ipp:"operations-supported,enum"`
}

// PrinterMediaSizeSupported represents "media-size-supported"
// collection entry, Integer variant
type PrinterMediaSizeSupported struct {
	XDimension int `ipp:"x-dimension"`
	YDimension int `ipp:"y-dimension"`
}

// PrinterMediaSizeSupportedRange represents "media-size-supported"
// collection entry, rangeOfInteger variant
type PrinterMediaSizeSupportedRange struct {
	XDimension goipp.Range `ipp:"x-dimension"`
	YDimension goipp.Range `ipp:"y-dimension"`
}

// IsCharsetSupported tells if charset is supported
func (pa *PrinterAttributes) IsCharsetSupported(cs string) bool {
	for _, supp := range pa.CharsetSupported {
		if cs == supp {
			return true
		}
	}
	return false
}

// IsOperationSupported tells if operation is supported
func (pa *PrinterAttributes) IsOperationSupported(op goipp.Op) bool {
	for _, supp := range pa.OperationsSupported {
		if op == supp {
			return true
		}
	}
	return false
}
