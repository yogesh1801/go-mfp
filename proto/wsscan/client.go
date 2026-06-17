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
	"mime"
	"mime/multipart"
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

// CreateScanJob asks the WS-Scan server to create a scan job from the
// given ticket. req must be non-nil and have ScanTicket populated;
// DestinationToken and ScanIdentifier are optional and used only for
// device-initiated scans.
func (c *Client) CreateScanJob(
	ctx context.Context,
	req *CreateScanJobRequest,
) (*CreateScanJobResponse, error) {

	msg, err := c.sendSOAP(ctx, req)
	if err != nil {
		return nil, err
	}

	rsp, ok := msg.Body.(*CreateScanJobResponse)
	if !ok {
		return nil,
			fmt.Errorf("wsscan: unexpected response type %T", msg.Body)
	}

	return rsp, nil
}

// CancelJob cancels the scan job identified by jobID.
func (c *Client) CancelJob(
	ctx context.Context,
	jobID int,
) (*CancelJobResponse, error) {

	req := CancelJobRequest{
		JobID: jobID,
	}
	msg, err := c.sendSOAP(ctx, &req)
	if err != nil {
		return nil, err
	}

	rsp, ok := msg.Body.(*CancelJobResponse)
	if !ok {
		return nil,
			fmt.Errorf("wsscan: unexpected response type %T", msg.Body)
	}

	return rsp, nil
}

// GetActiveJobs returns a summary of all currently-active scan jobs.
func (c *Client) GetActiveJobs(
	ctx context.Context,
) (*GetActiveJobsResponse, error) {

	req := GetActiveJobsRequest{}
	msg, err := c.sendSOAP(ctx, &req)
	if err != nil {
		return nil, err
	}

	rsp, ok := msg.Body.(*GetActiveJobsResponse)
	if !ok {
		return nil,
			fmt.Errorf("wsscan: unexpected response type %T", msg.Body)
	}

	return rsp, nil
}

// GetJobHistory returns a summary of completed scan jobs.
func (c *Client) GetJobHistory(
	ctx context.Context,
) (*GetJobHistoryResponse, error) {

	req := GetJobHistoryRequest{}
	msg, err := c.sendSOAP(ctx, &req)
	if err != nil {
		return nil, err
	}

	rsp, ok := msg.Body.(*GetJobHistoryResponse)
	if !ok {
		return nil,
			fmt.Errorf("wsscan: unexpected response type %T", msg.Body)
	}

	return rsp, nil
}

// RetrieveImage retrieves a scanned image from the device. The returned
// [RetrieveImageResponse.Image] must be closed by the caller when done
// to release the underlying HTTP connection.
func (c *Client) RetrieveImage(
	ctx context.Context,
	req *RetrieveImageRequest,
) (*RetrieveImageResponse, error) {

	httpRsp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	mediaType, params, err := mime.ParseMediaType(
		httpRsp.Header.Get("Content-Type"))
	if err != nil {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: invalid Content-Type: %w", err)
	}
	if mediaType != "multipart/related" {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: expected multipart/related, got %q",
			mediaType)
	}
	boundary := params["boundary"]
	if boundary == "" {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: missing multipart boundary")
	}

	mr := multipart.NewReader(httpRsp.Body, boundary)

	// Part 1: SOAP envelope
	soapPart, err := mr.NextPart()
	if err != nil {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: reading SOAP part: %w", err)
	}
	soapData, err := io.ReadAll(soapPart)
	if err != nil {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: reading SOAP part: %w", err)
	}
	root, err := xmldoc.Decode(NsMap, bytes.NewReader(soapData))
	if err != nil {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: decoding SOAP envelope: %w", err)
	}
	msg, err := DecodeMessage(root)
	if err != nil {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: decoding SOAP message: %w", err)
	}
	rsp, ok := msg.Body.(*RetrieveImageResponse)
	if !ok {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: unexpected response type %T", msg.Body)
	}

	// Part 2: image data — streamed; closing Image closes httpRsp.Body
	imagePart, err := mr.NextPart()
	if err != nil {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("wsscan: reading image part: %w", err)
	}
	rsp.ContentType = imagePart.Header.Get("Content-Type")
	rsp.Image = struct {
		io.Reader
		io.Closer
	}{imagePart, httpRsp.Body}

	return rsp, nil
}

// doRequest wraps body in a SOAP envelope, POSTs it to the server,
// and returns the raw HTTP response. The caller is responsible for
// reading and closing the response body.
func (c *Client) doRequest(
	ctx context.Context, body Body) (*http.Response, error) {

	req := Message{
		Header: Header{
			Action:    body.Action(),
			MessageID: AnyURI(uuid.Random().URN()),
			To:        optional.New(AnyURI(c.url.String())),
			ReplyTo:   optional.New(AnyURI(AddrAnonymous)),
		},
		Body: body,
	}

	data := req.Encode()

	httpReq, err := transport.NewRequest(ctx, "POST", c.url,
		bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/soap+xml")

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if httpRsp.StatusCode/100 != http.StatusOK/100 {
		httpRsp.Body.Close()
		return nil, fmt.Errorf("HTTP %d: %s",
			httpRsp.StatusCode, httpRsp.Status)
	}

	return httpRsp, nil
}

// sendSOAP wraps body in a SOAP envelope, POSTs it to the server,
// and returns the decoded response [Message].
func (c *Client) sendSOAP(ctx context.Context, body Body) (Message, error) {
	httpRsp, err := c.doRequest(ctx, body)
	if err != nil {
		return Message{}, err
	}
	defer httpRsp.Body.Close()

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
