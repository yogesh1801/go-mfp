// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// PNG Reader and Writer

package imgconv

import (
	"errors"
	"fmt"
	"image/color"
	"io"
	"math"
	"runtime/cgo"
	"unsafe"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// #cgo pkg-config: libpng
//
// #include "pngglue.h"
import "C"

// pngReader implements the [Reader] interface for reading PNG images.
type pngReader struct {
	handle   cgo.Handle    // Handle to self
	png      *C.png_struct // Underlying png_structure
	pngInfo  *C.png_info   // libpng png_info structure
	err      error         // Error from the libpng
	input    io.Reader     // Underlying io.Reader
	model    color.Model   // Image color mode
	wid, hei int           // Image size
	rowBytes []byte        // Row decoding buffer
	y        int           // Current y-coordinate
}

// NewPNGReader creates a new [Decoder] for PNG images
func NewPNGReader(input io.Reader) (Decoder, error) {
	// Create reader structure. Initialize libpng stuff
	reader := &pngReader{input: input}
	reader.handle = cgo.NewHandle(reader)

	reader.png = C.do_png_create_read_struct(
		unsafe.Pointer(&reader.handle))

	reader.pngInfo = C.png_create_info_struct(reader.png)

	// Read image info
	C.do_png_read_info(reader.png, reader.pngInfo)

	var width, height C.png_uint_32
	var depth, colorType, interlace C.int

	C.do_png_get_IHDR(reader.png, reader.pngInfo, &width, &height,
		&depth, &colorType, &interlace, nil, nil)

	if err := reader.err; err != nil {
		reader.Close()
		return nil, err
	}

	if interlace != C.PNG_INTERLACE_NONE {
		reader.Close()
		err := errors.New("PNG: interlaced images not supported")
		return nil, err
	}

	reader.wid, reader.hei = int(width), int(height)

	// Setup input transformations
	var bytesPerPixel int

	gray := (colorType & C.PNG_COLOR_MASK_COLOR) == C.PNG_COLOR_TYPE_GRAY

	if colorType == C.PNG_COLOR_TYPE_PALETTE {
		C.png_set_palette_to_rgb(reader.png)
	}

	if gray && depth < 8 {
		C.png_set_expand_gray_1_2_4_to_8(reader.png)
	}

	if (colorType & C.PNG_COLOR_MASK_ALPHA) != 0 {
		C.png_set_strip_alpha(reader.png)
	}

	if gray {
		reader.model = color.GrayModel
		bytesPerPixel = 1
		if depth == 16 {
			reader.model = color.Gray16Model
			bytesPerPixel = 2
		}
	} else {
		reader.model = color.RGBAModel
		bytesPerPixel = 3
		if depth == 16 {
			reader.model = color.RGBA64Model
			bytesPerPixel = 6
		}
	}

	// Allocate buffers
	reader.rowBytes = make([]byte, bytesPerPixel*int(width))

	return reader, nil
}

// MIMEType returns the MIME type of the image being decoded.
func (*pngReader) MIMEType() string {
	return MIMETypePNG
}

// Close closes the reader.
func (reader *pngReader) Close() {
	C.png_destroy_read_struct(&reader.png, &reader.pngInfo, nil)
	reader.handle.Delete()
}

// ColorModel returns the [color.Model] of image being decoded.
func (reader *pngReader) ColorModel() color.Model {
	return reader.model
}

// Size returns the image size.
func (reader *pngReader) Size() (wid, hei int) {
	return reader.wid, reader.hei
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Reader.Read] function.
func (reader *pngReader) NewRow() Row {
	return NewRow(reader.model, reader.wid)
}

// Read returns the next image [Row].
func (reader *pngReader) Read(row Row) (int, error) {
	// Read the next row
	reader.readRow()
	if reader.err != nil {
		return 0, reader.err
	}

	// Decode the row
	wid := generic.Min(row.Width(), reader.wid)

	switch reader.model {
	case color.GrayModel:
		bytesGray8toRow(row, reader.rowBytes)

	case color.Gray16Model:
		bytesGray16BEtoRow(row, reader.rowBytes)

	case color.RGBAModel:
		bytesRGB8toRow(row, reader.rowBytes)

	case color.RGBA64Model:
		bytesRGB16BEtoRow(row, reader.rowBytes)
	}

	// Update current y
	reader.y++
	if reader.y == reader.hei {
		reader.setError(io.EOF)
	}

	return wid, nil
}

// setError sets the reader.err, if it is not set yet
func (reader *pngReader) setError(err error) {
	if reader.err == nil {
		reader.err = err
	}
}

// readRow reads the next image line.
func (reader *pngReader) readRow() {
	if reader.err == nil {
		C.do_png_read_row(reader.png,
			unsafe.Pointer(&reader.rowBytes[0]), nil)
	}
}

// pngWriter implements the [Writer] interface for writing PNG images
type pngWriter struct {
	handle   cgo.Handle    // Handle to self
	png      *C.png_struct // Underlying png_structure
	pngInfo  *C.png_info   // libpng png_info structure
	err      error         // Error from the libpng
	output   io.Writer     // Underlying io.Writer
	wid, hei int           // Image size
	model    color.Model   // Color model
	rowBytes []byte        // Row encoding buffer
	y        int           // Current y-coordinate
}

// NewPNGWriter creates a new [Writer] for PNG images.
// Supported color models are following:
//   - color.GrayModel
//   - color.Gray16Model
//   - color.RGBAModel
//   - color.RGBA64Model
func NewPNGWriter(output io.Writer,
	wid, hei int, model color.Model) (Encoder, error) {

	// Translate model into libpng terms
	var colorType, depth C.int
	var bytesPerPixel int

	switch model {
	case color.GrayModel:
		colorType = C.PNG_COLOR_TYPE_GRAY
		depth = 8
		bytesPerPixel = 1
	case color.Gray16Model:
		colorType = C.PNG_COLOR_TYPE_GRAY
		depth = 16
		bytesPerPixel = 2
	case color.RGBAModel:
		colorType = C.PNG_COLOR_TYPE_RGB
		depth = 8
		bytesPerPixel = 3
	case color.RGBA64Model:
		colorType = C.PNG_COLOR_TYPE_RGB
		depth = 16
		bytesPerPixel = 6
	default:
		err := errors.New("PNG: unsupported color model")
		return nil, err
	}

	// Create writer structure. Initialize libpng stuff
	writer := &pngWriter{
		output:   output,
		wid:      wid,
		hei:      hei,
		model:    model,
		rowBytes: make([]byte, bytesPerPixel*int(wid)),
	}

	writer.handle = cgo.NewHandle(writer)

	writer.png = C.do_png_create_write_struct(
		unsafe.Pointer(&writer.handle))

	writer.pngInfo = C.png_create_info_struct(writer.png)

	// Create PNG header
	C.do_png_set_IHDR(writer.png, writer.pngInfo,
		C.png_uint_32(wid), C.png_uint_32(hei),
		depth, colorType,
		C.PNG_INTERLACE_NONE,
		C.PNG_COMPRESSION_TYPE_DEFAULT,
		C.PNG_FILTER_TYPE_DEFAULT)

	if writer.err == nil {
		C.png_set_sRGB(writer.png, writer.pngInfo,
			C.PNG_sRGB_INTENT_PERCEPTUAL)

		C.do_png_write_info(writer.png, writer.pngInfo)
	}

	if writer.err != nil {
		writer.Close()
		return nil, writer.err
	}

	return writer, nil
}

// MIMEType returns the MIME type of the image being encoded.
func (*pngWriter) MIMEType() string {
	return MIMETypePNG
}

// Size returns the image size.
func (writer *pngWriter) Size() (wid, hei int) {
	return writer.wid, writer.hei
}

// ColorModel returns the [color.Model] of image being written.
func (writer *pngWriter) ColorModel() color.Model {
	return writer.model
}

// Write writes the next image [Row].
func (writer *pngWriter) Write(row Row) error {
	// Check for pending error
	if writer.err != nil {
		return writer.err
	}

	// Silently ignore excessive rows
	if writer.y == writer.hei {
		return nil
	}

	// Encode the row
	wid := generic.Min(row.Width(), writer.wid)

	var bytesPerPixel int

	switch writer.model {
	case color.GrayModel:
		bytesPerPixel = 1
		bytesGray8fromRow(writer.rowBytes, row)
	case color.Gray16Model:
		bytesPerPixel = 2
		bytesGray16BEfromRow(writer.rowBytes, row)
	case color.RGBAModel:
		bytesPerPixel = 3
		bytesRGB8fromRow(writer.rowBytes, row)
	case color.RGBA64Model:
		bytesPerPixel = 6
		bytesRGB16BEfromRow(writer.rowBytes, row)
	}

	// Fill the tail
	if wid < writer.wid {
		end := writer.wid * bytesPerPixel
		for x := wid * bytesPerPixel; x < end; x++ {
			writer.rowBytes[x] = 0xff
		}
	}

	// Write the row
	C.do_png_write_row(writer.png, unsafe.Pointer(&writer.rowBytes[0]))
	if writer.err == nil {
		writer.y++
	}

	return writer.err
}

// Close flushes the buffered data and then closes the Writer
func (writer *pngWriter) Close() error {
	// Write missed lines
	for writer.err == nil && writer.y < writer.hei {
		writer.Write(RowEmpty{})
	}

	// Finish PNG image
	if writer.err == nil {
		C.do_png_write_end(writer.png, nil)
	}

	// Release allocated resources
	C.png_destroy_write_struct(&writer.png, &writer.pngInfo)
	writer.handle.Delete()

	return writer.err
}

// setError sets the writer.err, if it is not set yet
func (writer *pngWriter) setError(err error) {
	if writer.err == nil {
		writer.err = err
	}
}

// pngErrorCallback is called by the libpng to report a error
// This is the common callback for pngReader and pngWriter
//
//export pngErrorCallback
func pngErrorCallback(png *C.png_struct, msg C.png_const_charp) {
	p := (*cgo.Handle)(C.png_get_io_ptr(png)).Value()
	out := p.(interface{ setError(error) })
	err := fmt.Errorf("PNG: %s", C.GoString(msg))
	out.setError(err)
}

// pngWarningCallback is called by the libpng to report a warning
//
//export pngWarningCallback
func pngWarningCallback(png *C.png_struct, msg C.png_const_charp) {
	// Ignore the warning
}

// pngMallocCallback is called by libpng to allocate a memory
//
//export pngMallocCallback
func pngMallocCallback(png *C.png_struct, size C.size_t) unsafe.Pointer {
	return C.malloc(size)
}

// pngFreeCallback is called by libpng to free the memory
//
//export pngFreeCallback
func pngFreeCallback(png *C.png_struct, p unsafe.Pointer) {
	C.free(p)
}

// pngReadCallback s called by libpng to read from the input stream
//
//export pngReadCallback
func pngReadCallback(png *C.png_struct, data C.png_bytep, size C.size_t) C.int {
	const max = math.MaxInt32

	sz := max
	if C.size_t(sz) > size {
		sz = int(size)
	}

	reader := (*cgo.Handle)(C.png_get_io_ptr(png)).Value().(*pngReader)

	buf := (*[max]byte)(unsafe.Pointer(data))[:sz:sz]
	for len(buf) > 0 {
		n, err := reader.input.Read(buf)
		if n > 0 {
			buf = buf[n:]
		} else if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			reader.setError(err)
			return 0
		}
	}

	return 1
}

// pngWriteCallback s called by libpng to write into the output stream
//
//export pngWriteCallback
func pngWriteCallback(png *C.png_struct, data C.png_bytep, size C.size_t) C.int {
	const max = math.MaxInt32

	sz := max
	if C.size_t(sz) > size {
		sz = int(size)
	}

	writer := (*cgo.Handle)(C.png_get_io_ptr(png)).Value().(*pngWriter)

	buf := (*[max]byte)(unsafe.Pointer(data))[:sz:sz]
	for len(buf) > 0 {
		n, err := writer.output.Write(buf)
		if n > 0 {
			buf = buf[n:]
		} else if err != nil {
			writer.setError(err)
			return 0
		}
	}

	return 1
}
