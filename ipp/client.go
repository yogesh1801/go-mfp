// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP client

package ipp

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/OpenPrinting/goipp"
)

// Client implements Client-side IPP Printer object.
type Client struct {
	URL    string       // Destination URL (ipp://...)
	Config ClientConfig // Client configuration with all defaults resolved

	// HTTP stuff
	httpURL *url.URL // Parsed URL
}

// ClientConfig contains Client configuration options.
// Used as parameter to the NewClient function.
type ClientConfig struct {
	// HTTPClient specifies a HTTP client to use.
	//
	// Please notice that in Go stdlib http.Client is the connection
	// manager. So all Clients that share the same HTTP client will
	// share pool of active connections.
	//
	// If set to nil, DefaultHTTPClient will be used.
	HTTPClient *http.Client

	// IppVersion specifies IPP protocol version to use.
	//
	// If set to 0, goipp.DefaultVersion will be used.
	IppVersion goipp.Version

	// AttrCharset contains "attributes-charset" value
	// for all requests.
	//
	// If set to "", DefaultCharset will be used.
	AttrCharset string

	// AttrNaturalLanguage contains "attributes-natural-language"
	// value for all requests.
	//
	// If set to "", DefaultNaturalLanguage will be used.
	AttrNaturalLanguage string
}

// NewClient creates a new IPP client.
// If conf is nil, reasonable defaults are provided automatically
func NewClient(strURL string, conf *ClientConfig) (*Client, error) {
	// Parse and validate Printer URL
	httpURL, _, err := urlParse(strURL)
	if err != nil {
		return nil, err
	}

	// Create Client object.
	client := &Client{
		URL:     strURL,
		httpURL: httpURL,
	}

	// Resolve all defaults in the configuration
	if conf != nil {
		client.Config = *conf
	}

	if client.Config.HTTPClient == nil {
		client.Config.HTTPClient = DefaultHTTPClient
	}

	if client.Config.IppVersion == 0 {
		client.Config.IppVersion = goipp.DefaultVersion
	}

	if client.Config.AttrCharset == "" {
		client.Config.AttrCharset = DefaultCharset
	}

	if client.Config.AttrNaturalLanguage == "" {
		client.Config.AttrNaturalLanguage = DefaultNaturalLanguage
	}

	return client, nil
}

// Do sends the Request and waits for Response.
func (c *Client) Do(rq Request, rsp Response) error {
	return c.DoContext(context.Background(), rq, rsp)
}

// DoContext sends the Request and waits for Response.
// This is a version of [ipp.Client.Do] with [context.Context].
func (c *Client) DoContext(context.Context, Request, Response) error {
	return errors.New("not implemented")
}
