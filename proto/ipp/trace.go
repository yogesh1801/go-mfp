// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// trace.Message adapter for the goipp.Message

package ipp

import (
	"bytes"

	"github.com/OpenPrinting/go-mfp/proto/trace"
	"github.com/OpenPrinting/goipp"
)

// goippRequest implements trace.Message interface for the IPP requests
type goippRequest struct {
	msg *goipp.Message
}

var _ = trace.Message(goippRequest{})

// Ext returns file extension for the protocol message files.
// It implements the [trace.Message] interface.
func (goippRequest) Ext() string {
	return "ipp"
}

// Name returns the message name.
// It implements the [trace.Message] interface.
func (rq goippRequest) Name() string {
	return goipp.Op(rq.msg.Code).String()
}

// MarshalLog formats goippRequest for logging.
// It implements the [log.Marshaler] interface.
func (rq goippRequest) MarshalLog() []byte {
	buf := bytes.Buffer{}
	rq.msg.Print(&buf, true)
	return buf.Bytes()
}

// MarshalTrace returns the binary representation of goippRequest.
// It implements the [trace.Message] interface.
func (rq goippRequest) MarshalTrace() []byte {
	data, _ := rq.msg.EncodeBytes()
	return data
}

// goippRequest implements trace.Message interface for the IPP responses
type goippResponse struct {
	msg *goipp.Message
}

var _ = trace.Message(goippResponse{})

// Ext returns file extension for the protocol message files.
// It implements the [trace.Message] interface.
func (goippResponse) Ext() string {
	return "ipp"
}

// Name returns the message name.
// It implements the [trace.Message] interface.
func (rsp goippResponse) Name() string {
	return goipp.Status(rsp.msg.Code).String()
}

// MarshalLog formats goippResponse for logging.
// It implements the [log.Marshaler] interface.
func (rsp goippResponse) MarshalLog() []byte {
	buf := bytes.Buffer{}
	rsp.msg.Print(&buf, false)
	return buf.Bytes()
}

// MarshalTrace returns the binary representation of goippResponse.
// It implements the [trace.Message] interface
func (rsp goippResponse) MarshalTrace() []byte {
	data, _ := rsp.msg.EncodeBytes()
	return data
}
