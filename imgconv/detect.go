// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package imgconv

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

// NewDetectReader automatically detects image format and returns
// the appropriate reader.
func NewDetectReader(input io.Reader) (r Reader, err error) {
	var buf [16]byte
	off := 0

	// Prefetch some few bytes for image format detection
	for off < len(buf) {
		var n int
		n, err = input.Read(buf[off:])

		switch {
		case n > 0:
			off += n
		case n == 0 || err == io.EOF:
			err = io.ErrUnexpectedEOF
			fallthrough
		case err != nil:
			return
		}
	}

	input = io.MultiReader(bytes.NewReader(buf[:]), input)

	switch mime := MIMETypeDetect(buf[:]); mime {
	case MIMETypeJPEG:
		r, err = NewJPEGReader(input)
	case MIMETypePNG:
		r, err = NewPNGReader(input)
	case "":
		err = errors.New("image format cannot be detected")
	default:
		err = fmt.Errorf("%s: unsupported format", mime)
	}

	return
}
