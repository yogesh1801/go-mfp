// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL client

package escl

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

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
	caps *ScannerCapabilities, details *HTTPDetails, err error) {

	xml, details, err := c.getXML(ctx, "ScannerCapabilities")
	if err == nil {
		caps, err = DecodeScannerCapabilities(xml)
	}

	return
}

// GetScannerStatus requests the [ScannerStatus] from the eSCL scanner.
func (c *Client) GetScannerStatus(ctx context.Context) (
	status *ScannerStatus, details *HTTPDetails, err error) {

	xml, details, err := c.getXML(ctx, "ScannerStatus")
	if err == nil {
		status, err = DecodeScannerStatus(xml)
	}

	return
}

// Scan initializes scanning at the eSCL scanner by sending the
// [ScanSettings] request.
//
// On success it returns the normalized JobUri, that can be
// used for the subsequent call to [Client.NextDocument] and
// [Client.Cancel].
//
// Please notice that this function normalized the JobUri received
// from the server. If you need the raw, unmodified JobUri, use
// [HTTPDetails.Header.Get]("Location") using the provided
// HTTPDetails.
func (c *Client) Scan(ctx context.Context, rq ScanSettings) (
	joburl string, details *HTTPDetails, err error) {

	// Send the request
	details, err = c.post(ctx, "POST", "ScanJobs", rq.ToXML())
	if err != nil {
		return
	}

	// Decode JobUrl
	if vals := details.Header.Values("Location"); len(vals) != 0 {
		// Normalize JobUri
		//
		// Here we strip the JobUri, leaving only the URL Path part.
		// The reasons are:
		//  - for security
		//  - some devices (mostly, Pantum) use DNS-SD host name here
		//    instead of the literal addresses. There is no guarantee,
		//    that these names will properly resolve
		//  - some devices may put malformed IPv6 addresses here, that
		//    doesn't parse at all
		//
		// See also:
		//  https://github.com/alexpevzner/sane-airscan/commit/6e5222c32133793f2e06bf9caa0d36d39f8ef254
		location := vals[0]
		base := path.Join(c.url.Path, "/ScanJobs") + "/"
		i := strings.Index(location, base)
		if i >= 0 {
			joburl = location[i:]
			return
		}
		err = fmt.Errorf("eSCL: ScanJobs response: invalid JobUri: %q", location)
	} else {
		err = errors.New("eSCL: ScanJobs response: missed JobUri")
	}

	return
}

// NextDocument retrieves the next document.
//
// If all scanned documents are consumed, it returns [io.EOF] error,
// but please note that false positives are possible if there were
// no preceding [Client.Scan] request or joburl is invalud.
func (c *Client) NextDocument(ctx context.Context, joburl string) (
	doc io.ReadCloser, details *HTTPDetails, err error) {

	doc, details, err = c.get(ctx, "GET", joburl+"/NextDocument")
	if details != nil && details.StatusCode == http.StatusNotFound {
		err = io.EOF
	}

	return
}

// Cancel cancels the scan operation currently in progress.
// If job is already completed, it may return [io.EOF] or no error.
func (c *Client) Cancel(ctx context.Context, joburl string) (
	details *HTTPDetails, err error) {

	body, details, err := c.get(ctx, "DELETE", joburl)
	if body != nil {
		body.Close()
	}

	if details != nil && details.StatusCode == http.StatusNotFound {
		err = io.EOF
	}

	return
}

// getXML performs GET request, then decodes returned XML.
func (c *Client) getXML(ctx context.Context, subpath string) (
	xml xmldoc.Element, details *HTTPDetails, err error) {

	// Perform GET request
	body, details, err := c.get(ctx, "GET", subpath)
	if err != nil {
		return
	}

	// Decode the body
	xml, err = xmldoc.Decode(NsMap, body)
	body.Close()

	return
}

// dest returns the destination URL for the given subpath (e.g., "ScanJobs").
//
// if subpath doesn't start with "/", it is interpreted relative
// to the Client.url.Path
func (c *Client) dest(subpath string) *url.URL {
	u := transport.URLClone(c.url)
	u.Path = subpath
	if !strings.HasPrefix(subpath, "/") {
		u.Path = path.Join(c.url.Path, subpath)
	}

	return u
}

// get is the common body of the GET-style eSCL client requests.
// The actual method may be "GET" or "DELETE".
func (c *Client) get(ctx context.Context, method, subpath string) (
	body io.ReadCloser, details *HTTPDetails, err error) {

	// Prepare destination URL
	u := c.dest(subpath)

	// Log the request
	log.Debug(ctx, "eSCL request: %s %s", method, u)

	// Perform HTTP request
	httpRq, err := transport.NewRequest(ctx, method, u, nil)
	if err != nil {
		return
	}

	httpRsp, err := c.httpClient.Do(httpRq)
	if err != nil {
		return
	}

	// Decode the response
	details = &HTTPDetails{
		Status:     httpRsp.Status,
		StatusCode: httpRsp.StatusCode,
		Header:     httpRsp.Header,
	}

	if httpRsp.StatusCode/100 != http.StatusOK/100 {
		err = fmt.Errorf("HTTP: %s", httpRsp.Status)
		httpRsp.Body.Close()
		return
	}

	body = httpRsp.Body
	return
}

// post is the common body of the POST-style eSCL client requests.
// The actual method may be "POST" or "PUT".
func (c *Client) post(ctx context.Context, method, subpath string,
	xml xmldoc.Element) (details *HTTPDetails, err error) {

	// Prepare destination URL
	u := c.dest(subpath)

	// Log the request
	log.Debug(ctx, "eSCL request: %s %s", method, u)

	// Perform HTTP request
	var buf bytes.Buffer
	xml.Encode(&buf, NsMap)

	httpRq, err := transport.NewRequest(ctx, method, u, &buf)
	if err != nil {
		return
	}

	httpRq.Header.Set("Content-Type", "text/xml")

	httpRsp, err := c.httpClient.Do(httpRq)
	if err != nil {
		return
	}

	defer httpRsp.Body.Close()

	// Decode the response
	details = &HTTPDetails{
		Status:     httpRsp.Status,
		StatusCode: httpRsp.StatusCode,
		Header:     httpRsp.Header,
	}

	if httpRsp.StatusCode/100 != http.StatusOK/100 {
		err = fmt.Errorf("HTTP: %s", httpRsp.Status)
		return
	}

	return
}
