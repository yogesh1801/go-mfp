// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP client

package transport

import "net/http"

// Client wraps [http.Client]
type Client struct {
	http.Client
}

// NewClient creates a new [Client].
//
// It inherits all methods from the [http.Client].
//
// If [Transport] parameter is nil, new transport will be created
// with the [NewTransport] function.
func NewClient(tr *Transport) *Client {
	if tr == nil {
		tr = NewTransport(nil)
	}

	clnt := &Client{
		Client: http.Client{
			Transport: tr,
		},
	}

	return clnt
}
