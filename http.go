// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP stuff

package ippx

import "net/http"

var (
	// DefaultHTTPClient is the default HTTP client
	DefaultHTTPClient = &http.Client{}
)
