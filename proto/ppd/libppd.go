// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// libppd binding

package ppd

// #include "libppd.h"
// #cgo LDFLAGS: -l dl
import "C"
import "errors"

// libppdInitError contains the libppd initialization error, if any
var libppdInitError error

// init initializes the libppd binding
func init() {
	s := C.libppd_init()
	if s != nil {
		libppdInitError = errors.New(C.GoString(s))
	}
}
