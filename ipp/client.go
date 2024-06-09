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

	"github.com/alexpevzner/mfp/transport"
)

// Client implements Client-side IPP Printer object.
type Client struct {
	URL        *url.URL     // Destination URL (ipp://...)
	HTTPClient *http.Client // HTTP Client
}

// NewClient creates a new IPP client.
//
// If tr is nil, [transport.NewTransport] will be used to create
// a new transport.
func NewClient(u *url.URL, tr http.RoundTripper) *Client {
	if tr == nil {
		tr = transport.NewTransport(nil)
	}

	c := &Client{
		URL: u,
		HTTPClient: &http.Client{
			Transport: tr,
		},
	}

	return c
}

// Do sends the [Request] and waits for [Response].
//
// It automatically closes Response Body. This is convenient
// for most IPP requests, as body is rarely returned by IPP.
//
// For requests with returned body, use [Client.DoWithBody] instead.
func (c *Client) Do(rq Request, rsp Response) error {
	return c.DoContext(context.Background(), rq, rsp)
}

// DoWithBody Do sends the [Request] and waits for [Response].
//
// On success, caller MUST close Response body after use.
func (c *Client) DoWithBody(rq Request, rsp Response) error {
	return c.DoContextWithBody(context.Background(), rq, rsp)
}

// DoContext sends the Request and waits for Response.
// This is a version of [ipp.Client.Do] with [context.Context].
//
// It automatically closes Response Body. This is convenient
// for most IPP requests, as body is rarely returned by IPP.
//
// For requests with returned body, use [Client.DoContextWithBody] instead.
func (c *Client) DoContext(ctx context.Context,
	rq Request, rsp Response) error {
	err := c.DoContextWithBody(ctx, rq, rsp)
	if err == nil {
		rsp.GetBody().Close()
	}
	return err
}

// DoContextWithBody sends the Request and waits for Response.
// This is a version of [ipp.Client.DoWithBody] with [context.Context].
//
// On success, caller MUST close Response body after use.
func (c *Client) DoContextWithBody(ctx context.Context,
	rq Request, rsp Response) error {

	// Encode IPP message
	buf := &bytes.Buffer{}
	msg := rq.Encode()
	msg.Encode(buf)

	// Attach Request body, if any
	body := rq.GetBody()
	if body == nil {
		body = buf
	} else {
		body = io.MultiReader(buf, body)
	}

	// Create HTTP request
	httpRq, err := transport.NewRequestWithContext(ctx,
		"POST", c.URL.String(), body)
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
	rsp.SetBody(httpRsp.Body)
	return nil

ERROR:
	httpRsp.Body.Close()
	return err
}
