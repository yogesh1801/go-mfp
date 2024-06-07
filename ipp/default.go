// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Default values

package ipp

import "github.com/OpenPrinting/goipp"

// Default values for common attributes
var (
	// DefaultCharset is the default value for the "attributes-charset"
	// operation attribute and the "charset-configured" printer attribute
	DefaultCharset = "utf-8"

	// DefaultNaturalLanguage is the default value for
	// "attributes-natural-language" operation attribute.
	DefaultNaturalLanguage = "en-us"

	// DefaultCharsetSupported is the default value for
	// ""charset-supported" printer attribute.
	DefaultCharsetSupported = []string{DefaultCharset}

	// DefaultIppVersionsSupported is the default value for
	// "ipp-features-supported" printer attribute.
	DefaultIppVersionsSupported = []goipp.Version{
		goipp.MakeVersion(2, 0),
		goipp.MakeVersion(1, 0),
		goipp.MakeVersion(1, 1),
	}

	// DefaultRequestHeader is the default value for the
	// RequestHeader structure.
	DefaultRequestHeader = RequestHeader{
		Version:                   goipp.DefaultVersion,
		AttributesCharset:         DefaultCharset,
		AttributesNaturalLanguage: DefaultNaturalLanguage,
	}

	// DefaultResponseHeader is the default value for the
	// RequestHeader structure.
	DefaultResponseHeader = ResponseHeader{
		Version:                   goipp.DefaultVersion,
		Status:                    goipp.StatusOk,
		AttributesCharset:         DefaultCharset,
		AttributesNaturalLanguage: DefaultNaturalLanguage,
		StatusMessage:             "success",
	}
)
