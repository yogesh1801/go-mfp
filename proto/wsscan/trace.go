// MFP - Miulti-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by go-mfp authors.
// See LICENSE for license terms and conditions
//
// trace.Writer integration

package wsscan

import "strings"

// traceMessage wraps the Message and implements tracer.Message interface
// on a top of it.
type traceMessage struct {
	msg Message
}

// Trivial methods of traceMessage
func (m traceMessage) Protocol() string     { return "WS-Scan" }
func (m traceMessage) Ext() string          { return "xml" }
func (m traceMessage) MarshalTrace() []byte { return m.MarshalLog() }

// Name returns the message name
func (m traceMessage) Name() string {
	s := m.msg.Body.Action().String()
	if !strings.HasSuffix(s, "Response") {
		s += "Request"
	}
	return s
}

// MarshalLog returns message content as pretty-printed XML
func (m traceMessage) MarshalLog() []byte {
	return []byte(m.msg.Format())
}
