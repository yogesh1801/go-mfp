// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common functions for testing

package imgconv

import (
	"fmt"
	"image"
	"image/color"
	"io"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// colorEqual reports if two colors are equal.
// It works by converting both colors to RGB and comparing their components.
func colorEqual(c1, c2 color.Color) bool {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	return r1 == r2 && g1 == g2 && b1 == b2
}

// imageDiff compares two images and reports if they are different.
// If images are equal, it returns an empty string ("").
func imageDiff(img1, img2 image.Image) string {
	if diff := testutils.Diff(img1.Bounds(), img2.Bounds()); diff != "" {
		return fmt.Sprintf("Image.Bounds:\n%s", diff)
	}

	width := img1.Bounds().Dx()
	height := img1.Bounds().Dy()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c1 := img1.At(x, y)
			c2 := img2.At(x, y)
			if !colorEqual(c1, c2) {
				return fmt.Sprintf("Image.At(%d,%d):\n%s",
					x, y, testutils.Diff(c1, c2))
			}
		}
	}

	return ""
}

// readerWithError implements [io.Reader] interface for the byte slice.
// When all data bytes are consumed, it returns the specified error.
type readerWithError struct {
	data []byte
	err  error
}

// newReaderWithError creates a new [io.Reader] that reads from
// the provided data slice. When all data bytes are consumed,
// it returns the specified error instead of the [io.EOF]
func newReaderWithError(data []byte, err error) io.Reader {
	return &readerWithError{data, err}
}

// Read reads from the readerWithError.
// It implements the [io.Reader] interface.
func (r *readerWithError) Read(buf []byte) (int, error) {
	if len(r.data) > 0 {
		n := copy(buf, r.data)
		r.data = r.data[n:]
		return n, nil
	}

	return 0, r.err
}
