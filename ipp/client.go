// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP client

package ipp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/OpenPrinting/goipp"
	"github.com/alexpevzner/mfp/transport"
)

// Client implements Client-side IPP Printer object.
type Client struct {
	URL        *url.URL          // Destination URL (ipp://...)
	HTTPClient *transport.Client // HTTP Client
	RequestID  uint32            // RequestID of the next request
}

// NewClient creates a new IPP client.
//
// If tr is nil, [transport.NewTransport] will be used to create
// a new transport.
func NewClient(u *url.URL, tr *transport.Transport) *Client {
	c := &Client{
		URL:        u,
		HTTPClient: transport.NewClient(tr),
	}

	return c
}

// requestid generates a next RequestID
func (c *Client) requestid() uint32 {
	// IPP doesn't allow RequestID to be zero, so roll
	// until first non-zero value
	var id uint32
	for id == 0 {
		id = atomic.AddUint32(&c.RequestID, 1)
	}

	return id
}

// Do sends the [Request] and waits for [Response].
//
// The following Request fields are filled automatically:
//   - Version, if zero, will be set to goipp.DefaultVersion
//   - RequestID will be set to next Client's RequestID in sequence
//
// It automatically closes Response Body. This is convenient
// for most IPP requests, as body is rarely returned by IPP.
//
// For requests with returned body, use [Client.DoWithBody] instead.
func (c *Client) Do(ctx context.Context, rq Request, rsp Response) error {
	err := c.DoWithBody(ctx, rq, rsp)
	if err == nil {
		if body := rsp.Header().Body; body != nil {
			body.Close()
			rsp.Header().Body = nil
		}
	}
	return err
}

// DoWithBody sends the Request and waits for Response.
//
// The following Request fields are filled automatically:
//   - Version, if zero, will be set to goipp.DefaultVersion
//   - RequestID will be set to next Client's RequestID in sequence
//
// On success, caller MUST close Response body after use.
func (c *Client) DoWithBody(ctx context.Context,
	rq Request, rsp Response) error {

	// Encode IPP message
	buf := &bytes.Buffer{}
	msg := rq.Encode()

	if msg.Version == 0 {
		msg.Version = goipp.DefaultVersion
	}

	if msg.RequestID == 0 {
		msg.RequestID = c.requestid()
	}

	msg.Encode(buf)

	// Attach Request body, if any
	body := rq.Header().Body
	if body == nil {
		body = buf
	} else {
		body = io.MultiReader(buf, body)
	}

	// Create HTTP request
	httpRq, err := transport.NewRequest(ctx, "POST", c.URL, body)
	if err != nil {
		return err
	}

	httpRq.Header.Set("Content-Type", "application/ipp")

	// Call server
	httpRsp, err := c.HTTPClient.Do(httpRq)
	if err != nil {
		return err
	}

	if httpRsp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP: %s", httpRsp.Status)
		goto ERROR
	}

	// Decode IPP message
	msg.Reset()
	err = msg.Decode(httpRsp.Body)
	if err != nil {
		goto ERROR
	}

	// Decode Response
	err = rsp.Decode(msg)
	if err != nil {
		goto ERROR
	}

	// Save remainder of body and return
	rsp.Header().Body = httpRsp.Body
	return nil

ERROR:
	httpRsp.Body.Close()
	return err
}
