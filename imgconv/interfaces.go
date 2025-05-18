// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common interfaces

package imgconv

import (
	"image/color"
)

// Decoder implements streaming image decoder.
//
// It reads image row by row from the supplied [io.Reader].
type Decoder interface {
	// ColorModel returns the [color.Model] of image being decoded.
	ColorModel() color.Model

	// Size returns the image size.
	Size() (wid, hei int)

	// Read returns the next image [Row].
	Read() (Row, error)

	// Close closes the decoder
	Close()
}

// Encoder implements streaming image encoder.
//
// It writes image row by row into the supplied [io.Writer].
type Encoder interface {
	// Write writes the next image [Row].
	Write(Row) error

	// Close flushes the buffered data and then closes the Encoder
	Close() error
}
