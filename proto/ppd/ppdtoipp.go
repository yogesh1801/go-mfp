// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// PPD->IPP conversion

package ppd

import (
	"errors"
	"io"

	"github.com/OpenPrinting/goipp"
)

// #include "libppd.h"
import "C"

// ToIPP converts content of the PPD file into the IPP printer
// attributes, represented as [goipp.Attributes].
func ToIPP(ppd []byte) (goipp.Attributes, error) {
	// Check for initialization error
	if libppdInitError != nil {
		return nil, libppdInitError
	}

	// Open the temporary file
	fd, err := tmpFDOpen("temporary.ppd")
	if err != nil {
		return nil, err
	}

	defer func() {
		if fd >= 0 {
			fd.Close()
		}
	}()

	// Populate temporary file with the PPD data
	for len(ppd) > 0 {
		var n int
		n, err = fd.Write(ppd)
		if err != nil {
			return nil, err
		}
		ppd = ppd[n:]
	}

	_, err = fd.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Load ppd into the ppd_file_t
	ppdfp := C.libppd_ppdOpenFd(C.int(fd))
	fd = -1 // ppdOpenFd takes file ownership

	if ppdfp == nil {
		return nil, errors.New("ppdOpenFd error")
	}

	defer C.libppd_ppdClose(ppdfp)

	// Convert to ipp_t
	ipp := C.libipp_ppdLoadAttributes(ppdfp)
	if ipp == nil {
		return nil, errors.New("ppdLoadAttributes error")
	}

	defer C.libppd_ippDelete(ipp)

	// Convert to goipp.Attributes.
	return ippImport(ipp)
}
