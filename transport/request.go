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
		}
	}

	return
}
