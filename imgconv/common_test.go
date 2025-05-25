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
	"github.com/OpenPrinting/go-mfp/util/generic"
	"golang.org/x/image/draw"
)

// decodeImage reads the entire image out of the decoder
// and returns it as image.Image
func decodeImage(decoder Decoder) (image.Image, error) {
	wid, hei := decoder.Size()
	bounds := image.Rect(0, 0, wid, hei)

	var img draw.Image

	switch decoder.ColorModel() {
	case color.GrayModel:
		img = image.NewGray(bounds)
	case color.Gray16Model:
		img = image.NewGray16(bounds)
	case color.RGBAModel:
		img = image.NewRGBA(bounds)
	case color.RGBA64Model:
		img = image.NewRGBA64(bounds)
	default:
		panic("internal error")
	}

	row := decoder.NewRow()
	for y := 0; y < hei; y++ {
		_, err := decoder.Read(row)
		if err != nil {
			return nil, err
		}

		// Use row.Width() instead of the image width, returned
		// by Decoder.Size, so it also will be test-covered.
		wid := row.Width()
		for x := 0; x < wid; x++ {
			img.Set(x, y, row.At(x))
		}
	}

	return img, nil
}

// decodeImage reads the entire image out of the decoder
// and returns it as slice of rows.
func decodeImageRows(decoder Decoder) ([]Row, error) {
	_, hei := decoder.Size()
	rows := make([]Row, hei)
	for y := 0; y < hei; y++ {
		rows[y] = decoder.NewRow()
		_, err := decoder.Read(rows[y])
		if err != nil {
			return nil, err
		}
	}

	return rows, nil
}

// mustDecodeImageRows reads the entire image out of the decoder
// and returns it as slice of rows.
//
// It panics in a case of error.
func mustDecodeImageRows(decoder Decoder) []Row {
	rows, err := decodeImageRows(decoder)
	if err != nil {
		panic(err)
	}
	return rows
}

// encodeImageRows encodes the entire image, represented as a slice of rows.
func encodeImageRows(encoder Encoder, rows []Row) error {
	for y := range rows {
		err := encoder.Write(rows[y])
		if err != nil {
			return err
		}
	}
	return nil
}

// mustEncodeImageRows encodes the entire image, represented as a slice
// of rows.
//
// It panics in a case of error.
func mustEncodeImageRows(encoder Encoder, rows []Row) {
	err := encodeImageRows(encoder, rows)
	if err != nil {
		panic(err)
	}
}

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

// writerWithError wraps [io.Writer]. When the specified amount
// of bytes are written, it returns the specified error.
type writerWithError struct {
	dst io.Writer // Destination io.Writer
	lim int       // Write limit
	err error     // Error to be returned
}

// newWriterWithError returns [io.Writer], that bypasses up to the
// lim amount of bytes into the dst, then returns err.
func newWriterWithError(dst io.Writer, lim int, err error) io.Writer {
	return &writerWithError{dst: dst, lim: lim, err: err}
}

// Write implements io.Writer interface for the writerWithError.
func (w *writerWithError) Write(data []byte) (int, error) {
	if lim := w.lim; lim > 0 {
		lim = generic.Min(lim, len(data))
		data = data[:lim]
		n, err := w.dst.Write(data)
		if n > 0 {
			w.lim -= n
		}
		return n, err
	}

	return 0, w.err
}

// decoderWithError implements [Decoder] interface. It returns
// specified amount of image lines (quite meaningless lines),
// then returns the specified error.
type decoderWithError struct {
	model    color.Model // Image color model
	wid, hei int         // Image size
	lim      int         // Limit of successfully returned rows
	y        int         // Current y-coordinate
	err      error       // Error to be returned
}

// newDecoderWithError creates a new decoderWithError
func newDecoderWithError(model color.Model,
	wid, hei, lim int, err error) Decoder {
	return &decoderWithError{
		model: model,
		wid:   wid,
		hei:   hei,
		lim:   lim,
		err:   err,
	}
}

// ColorModel returns the [color.Model] of image being decoded.
func (decoder *decoderWithError) ColorModel() color.Model {
	return decoder.model
}

// Size returns the image size.
func (decoder *decoderWithError) Size() (wid, hei int) {
	return decoder.wid, decoder.hei
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Decoder.Read] function.
func (decoder *decoderWithError) NewRow() Row {
	return NewRow(decoder.model, decoder.wid)
}

// Read returns the next image [Row].
// The Row type must match the [Decoder]'s [color.Model].
//
// It returns the resulting row length, in pixels, or an error.
func (decoder *decoderWithError) Read(row Row) (int, error) {
	switch decoder.y {
	case decoder.lim:
		return 0, decoder.err
	case decoder.hei:
		return 0, io.EOF
	}

	wid := generic.Min(decoder.wid, row.Width())
	row.Slice(0, wid).Fill(color.White)
	decoder.y++

	return decoder.wid, nil
}

// Close closes the decoder.
func (decoder *decoderWithError) Close() {
}

// encoderWithError implements [Encoder] interface.
type encoderWithError struct {
	model    color.Model // Image color model
	wid, hei int         // Image size
	lim      int         // Limit of successfully returned rows
	y        int         // Current y-coordinate
	err      error       // Error to be returned
}

// newEncoderWithError creates a new encoderWithError
func newEncoderWithError(model color.Model,
	wid, hei, lim int, err error) Encoder {
	return &encoderWithError{
		model: model,
		wid:   wid,
		hei:   hei,
		lim:   lim,
		err:   err,
	}
}

// ColorModel returns the [color.Model] of image being decoded.
// It consumes the specified amount of image lines, then returns
// the specified error.
func (encoder *encoderWithError) ColorModel() color.Model {
	return encoder.model
}

// Size returns the image size.
func (encoder *encoderWithError) Size() (wid, hei int) {
	return encoder.wid, encoder.hei
}

// Write writes the next image [Row].
func (encoder *encoderWithError) Write(Row) error {
	switch encoder.y {
	case encoder.lim:
		return encoder.err
	case encoder.hei:
		return nil
	}

	encoder.y++
	return nil
}

// Close closes the encoder.
func (encoder *encoderWithError) Close() error {
	return nil
}
