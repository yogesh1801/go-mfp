// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// trace.Writter support

package escl

// traceMessage implements trace.Message interface for eSCL protocol messages.
//
// As eSCL protocol messages are trivial and the only information they
// carry is the message name (everything else is transmitted in the
// HTTP header or attachment), we represent them as string.
type traceMessage string

func (m traceMessage) Protocol() string     { return "eSCL" }
func (m traceMessage) Ext() string          { return "" }
func (m traceMessage) Name() string         { return string(m) }
func (m traceMessage) MarshalLog() []byte   { return nil }
func (m traceMessage) MarshalTrace() []byte { return nil }
