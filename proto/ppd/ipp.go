// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// goipp.Attributes to/from CUPS ipp_t converter

package ppd

import (
	"errors"
	"io"

	"github.com/OpenPrinting/goipp"
)

// #include "libppd.h"
import "C"

// ippExport converts goipp.Attributes into the CUPS ipp_t.
// The attributes assumed to be printer attributes.
func ippExport(attrs goipp.Attributes) (*C.ipp_t, error) {
	return nil, errors.New("not implemented")
}

// ippImport converts CUPS ipp_t into the goipp.Attributes.
// The attributes assumed to be printer attributes.
func ippImport(ipp *C.ipp_t) (goipp.Attributes, error) {
	// Open the temporary file
	fd, err := tmpFDOpen("temporary.ipp")
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	// Write ipp to the temporary file
	state := C.libppd_ippWriteFile(C.int(fd), ipp)
	if state == C.IPP_STATE_ERROR {
		return nil, errors.New("ippWriteFile error")
	}

	// Reparse as saved data as goipp.Message
	_, err = fd.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	msg := goipp.Message{}
	err = msg.Decode(fd)
	if err != nil {
		return nil, err
	}

	return msg.Printer, nil
}
