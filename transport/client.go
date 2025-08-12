// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP client

package transport

import (
	"net/http"

	"github.com/OpenPrinting/go-mfp/log"
)

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

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(rq *http.Request) (*http.Response, error) {
	// Execute the request
	rsp, err := c.Client.Do(rq)

	// Write log message
	var status string
	if err != nil {
		status = err.Error()
	} else {
		status = rsp.Status
	}

	log.Debug(rq.Context(), "HTTP-CLNT %s %s - %s", rq.Method, rq.URL, status)

	return rsp, err
}
