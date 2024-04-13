// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP stuff

package ipp

import "net/http"

var (
	// DefaultHTTPClient is the default HTTP client
	DefaultHTTPClient = &http.Client{}
)
