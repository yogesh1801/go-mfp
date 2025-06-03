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
	"image/png"
	"io"
	"math"
	"os"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"golang.org/x/image/draw"
)

// saveImage saves image as a PNG file.
func saveImage(name string, img image.Image) {
	flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	fp, err := os.OpenFile(name, flags, 0644)
	if err != nil {
		panic(err)
	}

	err = png.Encode(fp, img)
	if err != nil {
		panic(err)
	}

	fp.Close()
}

// decodeImage reads the entire image out of the Reader
// and returns it as image.Image
func decodeImage(reader Reader) (image.Image, error) {
	wid, hei := reader.Size()
	bounds := image.Rect(0, 0, wid, hei)

	var img draw.Image

	switch reader.ColorModel() {
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

	row := reader.NewRow()
	for y := 0; y < hei; y++ {
		_, err := reader.Read(row)
		if err != nil {
			return nil, err
		}

		// Use row.Width() instead of the image width, returned
		// by Reader.Size, so it also will be test-covered.
		wid := row.Width()
		for x := 0; x < wid; x++ {
			img.Set(x, y, row.At(x))
		}
	}

	return img, nil
}

// decodeImage reads the entire image out of the Reader
// and returns it as slice of rows.
func decodeImageRows(reader Reader) ([]Row, error) {
	_, hei := reader.Size()
	rows := make([]Row, hei)
	for y := 0; y < hei; y++ {
		rows[y] = reader.NewRow()
		_, err := reader.Read(rows[y])
		if err != nil {
			return nil, err
		}
	}

	return rows, nil
}

// mustDecodeImageRows reads the entire image out of the Reader
// and returns it as slice of rows.
//
// It panics in a case of error.
func mustDecodeImageRows(reader Reader) []Row {
	rows, err := decodeImageRows(reader)
	if err != nil {
		panic(err)
	}
	return rows
}

// encodeImageRows encodes the entire image, represented as a slice of rows.
func encodeImageRows(writer Writer, rows []Row) error {
	for y := range rows {
		err := writer.Write(rows[y])
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
func mustEncodeImageRows(writer Writer, rows []Row) {
	err := encodeImageRows(writer, rows)
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

// imageEuclideanDistance compares two images for similarity by computing
// the summary euclidean distance between image points, relative to
// the max possible euclidean distance (sqtr(3)) multiplied by the
// total number of points.
//
// Image sizes must be equal.
func imageEuclideanDistance(img1, img2 image.Image) float64 {
	assert.Must(img1.Bounds() == img1.Bounds())

	wid := img1.Bounds().Dx()
	hei := img1.Bounds().Dy()

	sum := 0.0
	for y := 0; y < hei; y++ {
		for x := 0; x < wid; x++ {
			c1 := img1.At(x, y)
			c2 := img2.At(x, y)
			r1, g1, b1, _ := c1.RGBA()
			r2, g2, b2, _ := c2.RGBA()

			r1fp := float64(r1) / 0xffff
			g1fp := float64(g1) / 0xffff
			b1fp := float64(b1) / 0xffff

			r2fp := float64(r2) / 0xffff
			g2fp := float64(g2) / 0xffff
			b2fp := float64(b2) / 0xffff

			d := math.Sqrt(
				(r1fp-r2fp)*(r1fp-r2fp) +
					(g1fp-g2fp)*(g1fp-g2fp) +
					(b1fp-b2fp)*(b1fp-b2fp))
			sum += d
		}
	}

	sum /= float64(wid) * float64(hei) * math.Sqrt(3)
	return sum
}

// ioReaderWithError implements [io.Reader] interface for the byte slice.
// When all data bytes are consumed, it returns the specified error.
type ioReaderWithError struct {
	data []byte
	err  error
}

// newReaderWithError creates a new [io.Reader] that reads from
// the provided data slice. When all data bytes are consumed,
// it returns the specified error instead of the [io.EOF]
func newIoReaderWithError(data []byte, err error) io.Reader {
	return &ioReaderWithError{data, err}
}

// Read reads from the ioReaderWithError.
// It implements the [io.Reader] interface.
func (r *ioReaderWithError) Read(buf []byte) (int, error) {
	if len(r.data) > 0 {
		n := copy(buf, r.data)
		r.data = r.data[n:]
		return n, nil
	}

	return 0, r.err
}

// ioWriterWithError wraps [io.Writer]. When the specified amount
// of bytes are written, it returns the specified error.
type ioWriterWithError struct {
	dst io.Writer // Destination io.Writer
	lim int       // Write limit
	err error     // Error to be returned
}

// newWriterWithError returns [io.Writer], that bypasses up to the
// lim amount of bytes into the dst, then returns err.
func newIoWriterWithError(dst io.Writer, lim int, err error) io.Writer {
	return &ioWriterWithError{dst: dst, lim: lim, err: err}
}

// Write implements io.Writer interface for the ioWriterWithError.
func (w *ioWriterWithError) Write(data []byte) (int, error) {
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

// readerWithError implements [Reader] interface. It returns
// specified amount of image lines (quite meaningless lines),
// then returns the specified error.
type readerWithError struct {
	model    color.Model // Image color model
	wid, hei int         // Image size
	lim      int         // Limit of successfully returned rows
	y        int         // Current y-coordinate
	err      error       // Error to be returned
}

// newReaderWithError creates a new readerWithError
func newReaderWithError(model color.Model,
	wid, hei, lim int, err error) Reader {
	return &readerWithError{
		model: model,
		wid:   wid,
		hei:   hei,
		lim:   lim,
		err:   err,
	}
}

// ColorModel returns the [color.Model] of image being decoded.
func (reader *readerWithError) ColorModel() color.Model {
	return reader.model
}

// Size returns the image size.
func (reader *readerWithError) Size() (wid, hei int) {
	return reader.wid, reader.hei
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Reader.Read] function.
func (reader *readerWithError) NewRow() Row {
	return NewRow(reader.model, reader.wid)
}

// Read returns the next image [Row].
// It returns the resulting row length, in pixels, or an error.
func (reader *readerWithError) Read(row Row) (int, error) {
	switch reader.y {
	case reader.lim:
		return 0, reader.err
	case reader.hei:
		return 0, io.EOF
	}

	wid := generic.Min(reader.wid, row.Width())
	row.Slice(0, wid).Fill(color.White)
	reader.y++

	return reader.wid, nil
}

// Close closes the reader.
func (reader *readerWithError) Close() {
}

// writerWithError implements [Writer] interface. It consumes
// specified amount of image lines then returns the specified
// error.
type writerWithError struct {
	model    color.Model // Image color model
	wid, hei int         // Image size
	lim      int         // Limit of successfully returned rows
	y        int         // Current y-coordinate
	err      error       // Error to be returned
}

// newWriterWithError creates a new writerWithError
func newWriterWithError(model color.Model,
	wid, hei, lim int, err error) Writer {
	return &writerWithError{
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
func (writer *writerWithError) ColorModel() color.Model {
	return writer.model
}

// Size returns the image size.
func (writer *writerWithError) Size() (wid, hei int) {
	return writer.wid, writer.hei
}

// Write writes the next image [Row].
func (writer *writerWithError) Write(Row) error {
	switch writer.y {
	case writer.lim:
		return writer.err
	case writer.hei:
		return nil
	}

	writer.y++
	return nil
}

// Close closes the writer.
func (writer *writerWithError) Close() error {
	return nil
}
