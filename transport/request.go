// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP requests

package transport

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/OpenPrinting/go-mfp/util/missed"
)

// NewRequest wraps the [http.NewRequestWithContext] with small API
// difference: it uses parsed [url.URL] instead of the URL string.
//
// The convenient method to obtain parsed [url.URL] is to use
// [ParseURL] or [ParseAddr] functions, provided by this package.
//
// See [http.NewRequestWithContext] documentation for details and
// nuances.
func NewRequest(ctx context.Context, method string,
	u *url.URL, body io.Reader) (rq *http.Request, err error) {

	rq, err = http.NewRequestWithContext(ctx, method, u.String(), body)
	if err == nil {
		rq.URL = u
		requestAdjustHost(rq, u)
	}

	return
}

// requestAdjustHost adjust rq.Host according to URL
func requestAdjustHost(rq *http.Request, u *url.URL) {
	// Make sure that port explicitly set in the rq.Host if and
	// only if it doesn't match default for the schema.
	//
	// Take in account, that schema in URL may be ipp/ipps, while
	// rq.Host defaults are relative to http/https, so implicit
	// port number may differ.
	var portCut, portAdd string

	switch rq.URL.Scheme {
	case "ipp":
		portCut, portAdd = ":80", ":631"
	case "ipps":
		portCut, portAdd = ":443", ":631"
	case "http":
		portCut, portAdd = ":80", ""
	case "https":
		portCut, portAdd = ":443", ""
	case "unix":
		rq.Host = "localhost"
		return
	}

	rq.Host, _ = missed.StringsCutSuffix(u.Host, portCut)
	if u.Port() == "" {
		rq.Host += portAdd
	}

	// Remove zone suffix from IPv6 literal
	host := rq.Host
	if strings.HasPrefix(host, "[") {
		beg := strings.IndexByte(host, '%')
		end := strings.IndexByte(host, ']')

		if beg >= 0 && end >= 0 && beg < end {
			rq.Host = host[:beg] + host[end:]
		}
	}
}
