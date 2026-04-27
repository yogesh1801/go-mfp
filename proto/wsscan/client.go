// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// WS-Scan client

package wsscan

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Client implements a low-level WS-Scan client.
type Client struct {
	url        *url.URL
	httpClient *transport.Client
}

// NewClient creates a new WS-Scan client.
//
// If tr is nil, [transport.NewTransport] will be used to create
// a new transport.
func NewClient(u *url.URL, tr *transport.Transport) *Client {
	return &Client{
		url:        transport.URLClone(u),
		httpClient: transport.NewClient(tr),
	}
}

// GetScannerElements requests the specified scanner schema elements
// from the WS-Scan server.
func (c *Client) GetScannerElements(
	ctx context.Context,
	elements ...ScannerElemName,
) (*GetScannerElementsResponse, error) {

	req := GetScannerElementsRequest{RequestedElements: elements}
	msg, err := c.sendSOAP(ctx, &req)
	if err != nil {
		return nil, err
	}

	rsp, ok := msg.Body.(*GetScannerElementsResponse)
	if !ok {
		return nil,
			fmt.Errorf("wsscan: unexpected response type %T", msg.Body)
	}

	return rsp, nil
}

// GetJobElements requests the specified job schema elements for the
// job identified by jobID from the WS-Scan server.
func (c *Client) GetJobElements(
	ctx context.Context,
	jobID int,
	elements ...JobElemName,
) (*GetJobElementsResponse, error) {

	req := GetJobElementsRequest{
		JobID:             jobID,
		RequestedElements: elements,
	}
	msg, err := c.sendSOAP(ctx, &req)
	if err != nil {
		return nil, err
	}

	rsp, ok := msg.Body.(*GetJobElementsResponse)
	if !ok {
		return nil,
			fmt.Errorf("wsscan: unexpected response type %T", msg.Body)
	}

	return rsp, nil
}

// sendSOAP wraps body in a SOAP envelope, POSTs it to the server,
// and returns the decoded response [Message].
func (c *Client) sendSOAP(ctx context.Context, body Body) (Message, error) {
	// Build request message
	req := Message{
		Header: Header{
			Action:    body.Action(),
			MessageID: AnyURI(uuid.Random().URN()),
			To:        optional.New(AnyURI(c.url.String())),
			ReplyTo:   optional.New(AnyURI(AddrAnonymous)),
		},
		Body: body,
	}

	// Encode to wire format
	data := req.Encode()

	// POST to server
	httpReq, err := transport.NewRequest(ctx, "POST", c.url,
		bytes.NewReader(data))
	if err != nil {
		return Message{}, err
	}
	httpReq.Header.Set("Content-Type", "application/soap+xml")

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return Message{}, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode/100 != http.StatusOK/100 {
		return Message{}, fmt.Errorf("HTTP %d: %s",
			httpRsp.StatusCode, httpRsp.Status)
	}

	// Read and decode response
	rspData, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return Message{}, err
	}

	root, err := xmldoc.Decode(NsMap, bytes.NewReader(rspData))
	if err != nil {
		return Message{}, err
	}

	return DecodeMessage(root)
}
