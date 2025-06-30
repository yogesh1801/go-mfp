// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL client

package escl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Client implements a low-level eSCL client.
type Client struct {
	url        *url.URL          // Destination URL (http://...)
	httpClient *transport.Client // HTTP Client
}

// NewClient creates a new eSCL client.
//
// If tr is nil, [transport.NewTransport] will be used to create
// a new transport.
func NewClient(u *url.URL, tr *transport.Transport) *Client {
	c := &Client{
		url:        transport.URLClone(u),
		httpClient: transport.NewClient(tr),
	}

	return c
}

// GetScannerCapabilities requests the [ScannerCapabilities] from
// the eSCL scanner.
func (c *Client) GetScannerCapabilities(ctx context.Context) (
	caps ScannerCapabilities, details *HTTPDetails, err error) {

	xml, details, err := c.get(ctx, "ScannerCapabilities")
	if err == nil {
		caps, err = DecodeScannerCapabilities(xml)
	}

	return
}

// GetScannerStatus requests the [ScannerStatus] from the eSCL scanner.
func (c *Client) GetScannerStatus(ctx context.Context) (
	status ScannerStatus, details *HTTPDetails, err error) {

	xml, details, err := c.get(ctx, "ScannerStatus")
	if err == nil {
		status, err = DecodeScannerStatus(xml)
	}

	return
}

// Scan initializes scanning at the eSCL scanner by sending the
// [ScanSettings] request.
func (c *Client) Scan(ctx context.Context, rq ScanSettings) (
	joburl string, details *HTTPDetails, err error) {
	return "", nil, errors.New("not implemented")
}

// NextDocument retrieves the next document.
func (c *Client) NextDocument(ctx context.Context, joburl string) (
	doc io.ReadCloser, details *HTTPDetails, err error) {
	return nil, nil, errors.New("not implemented")
}

// Cancel cancels the scan operation currently in progress.
func (c *Client) Cancel(ctx context.Context, joburl string) (
	details *HTTPDetails, err error) {
	return nil, errors.New("not implemented")
}

// get is the common body of the GET-style eSCL client requests.
func (c *Client) get(ctx context.Context, subpath string) (
	xml xmldoc.Element, details *HTTPDetails, err error) {

	// Prepare destination URL
	u := transport.URLClone(c.url)
	u.Path = path.Join(u.Path, subpath)

	// Log the request
	log.Debug(ctx, "eSCL request: GET %s", u)

	// Create HTTP request
	httpRq, err := transport.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return
	}

	httpRsp, err := c.httpClient.Do(httpRq)
	if err != nil {
		return
	}

	defer httpRsp.Body.Close()

	details = &HTTPDetails{
		Status:     httpRsp.Status,
		StatusCode: httpRsp.StatusCode,
		Header:     httpRsp.Header,
	}

	// Decode the response
	if httpRsp.StatusCode/100 != http.StatusOK/100 {
		err = fmt.Errorf("HTTP: %w", err)
		return
	}

	xml, err = xmldoc.Decode(NsMap, httpRsp.Body)
	return
}
