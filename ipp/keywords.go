// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP keywords

package ipp

import "reflect"

// KwCompression represents standard keyword values for compression.
// See RFC8011, 5.4.32
type KwCompression string

const (
	// KwCompressionNone is no compression
	KwCompressionNone KwCompression = "none"

	// KwCompressionDeflate is RFC 1951 ZIP inflate/deflate
	KwCompressionDeflate KwCompression = "deflate"

	// KwCompressionGzip is RFC 1952 GNU zip
	KwCompressionGzip KwCompression = "gzip"

	// KwCompressionCompress is RFC 1977 UNIX compression
	KwCompressionCompress KwCompression = "compress"
)

// kwRegisteredTypes lists all registered keyword types for IPP codec.
var kwRegisteredTypes = map[reflect.Type]struct{}{
	reflect.TypeOf(KwCompression("")): struct{}{},
}
