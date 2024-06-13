// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP errors

package ipp

import (
	"fmt"

	"github.com/OpenPrinting/goipp"
)

// Common errors, reported as ErrIPP:
var (
// HTTPErrorMethodNotAllowed = NewHTTPError(http.StatusMethodNotAllowed, "")
)

// ErrIPP represents IPP error.
// It consist of the IPP status and optional message text.
// Implements [error] interface.
type ErrIPP struct {
	Version       goipp.Version // IPP version
	RequestID     uint32        // IPP Request ID
	Status        goipp.Status  // IPP status
	StatusMessage string        // Optional error message
}

// NewErrIPP creates a new IPP error.
// If msg is "", [http.StatusText] used instead.
func NewErrIPP(rq *goipp.Message, code goipp.Status, msg string) *ErrIPP {
	ver := rq.Version
	if ver > goipp.DefaultVersion {
		ver = goipp.DefaultVersion
	}

	return &ErrIPP{
		Version:       ver,
		RequestID:     rq.RequestID,
		Status:        code,
		StatusMessage: msg,
	}
}

// Error returns an error string. It implements [error] interface.
func (e *ErrIPP) Error() string {
	msg := e.StatusMessage
	if msg == "" {
		msg = e.Status.String()
	}
	return fmt.Sprintf("IPP %s", msg)
}

// Encode encodes ErrIPP into the goipp.Message.
func (e *ErrIPP) Encode() *goipp.Message {
	msg := &goipp.Message{
		Version:   e.Version,
		RequestID: e.RequestID,
		Code:      goipp.Code(e.Status),
	}

	msg.Operation.Add(goipp.MakeAttribute("attributes-charset",
		goipp.TagCharset, goipp.String("utf-8")))
	msg.Operation.Add(goipp.MakeAttribute("attributes-natural-language",
		goipp.TagLanguage, goipp.String("en-US")))

	if e.StatusMessage != "" {
		msg.Operation.Add(goipp.MakeAttribute("status-message,",
			goipp.TagText, goipp.String(e.StatusMessage)))
	}

	return msg
}
