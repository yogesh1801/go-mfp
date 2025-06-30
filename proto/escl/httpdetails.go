// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP-level detais

package escl

import "net/http"

// HTTPDetails contains HTTP-level details on the [Client]'s operation.
type HTTPDetails struct {
	Status     string      // e.g. "200 OK"
	StatusCode int         // HTTP status code
	Header     http.Header // HTTP response header
}
