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
)

// NewRequest is similar to [http.NewRequest], but it uses [ParseURL]
// for URL string handling.
func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return NewRequestWithContext(context.Background(), method, url, body)
}

// NewRequestWithContext is similar to [http.NewRequestWithContext],
// but it uses [ParseURL] for URL string handling.
func NewRequestWithContext(ctx context.Context,
	method, url string, body io.Reader) (rq *http.Request, err error) {

	// We must Parse URL by ourselves, and then replace rq.URL,
	// parsed by the http.NewRequestWithContext.
	//
	// It guarantees that URL error handling and post-processing are
	// ours in a cost of double URL parsing.
	u, err := ParseURL(url)
	if err == nil {
		rq, err = http.NewRequestWithContext(ctx, method, url, body)
		if err == nil {
			rq.URL = u
			requestAdjustHost(rq, u)
		}
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

	rq.Host, _ = strings.CutSuffix(u.Host, portCut)
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
