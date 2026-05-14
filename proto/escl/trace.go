// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// trace.Writter support

package escl

import (
	"bytes"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// traceMessage implements trace.Message interface for eSCL protocol messages.
type traceMessage struct {
	name string
	xml  xmldoc.Element
}

// Simple methods of traceMessage
func (m traceMessage) Protocol() string     { return "eSCL" }
func (m traceMessage) Ext() string          { return "xml" }
func (m traceMessage) Name() string         { return m.name }
func (m traceMessage) MarshalTrace() []byte { return m.MarshalLog() }

// MarshalLog returns pretty-printed m.xml, if it is not zero-value
func (m traceMessage) MarshalLog() []byte {
	if !m.xml.IsZero() {
		buf := bytes.Buffer{}
		m.xml.EncodeIndent(&buf, NsMap, "  ")
		return buf.Bytes()
	}
	return nil
}
