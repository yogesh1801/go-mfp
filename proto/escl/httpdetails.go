// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP-level detais

package escl

import (
	"net/http"
	"strings"
)

// HTTPDetails contains HTTP-level details on the [Client]'s operation.
type HTTPDetails struct {
	Status      string         // e.g. "200 OK"
	StatusCode  int            // HTTP status code
	ContentType string         // Response content type
	Header      http.Header    // HTTP response header
	Response    *http.Response // HTTP response
}

// newHTTPDetails fills HTTPDetails from the http.Response
func newHTTPDetails(rsp *http.Response) *HTTPDetails {
	details := &HTTPDetails{
		Status:      rsp.Status,
		StatusCode:  rsp.StatusCode,
		ContentType: strings.ToLower(rsp.Header.Get("Content-Type")),
		Header:      rsp.Header,
		Response:    rsp,
	}

	return details
}
