// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// JPEG Reader and Writer

package imgconv

import (
	"errors"
	"image/color"
	"io"
	"runtime/cgo"
	"unsafe"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// #cgo pkg-config: libjpeg
//
// #include "jpegglue.h"
import "C"

// jpegReaderimplements the [Reader] interface for reading JPEG images.
type jpegReader struct {
	handle     cgo.Handle                       // Handle to self
	jpeg       *C.struct_jpeg_decompress_struct // JPEG decoder
	jpegErrMgr *C.struct_jpeg_error_mgr         // JPEG error manager
	jpegSrcMgr *C.struct_jpeg_source_mgr        // JPEG source manager
	buf        [65536]byte                      // Input buffer
	err        error                            // Error from the libjpeg
	input      io.Reader                        // Underlying io.Reader
	model      color.Model                      // Image color mode
	wid, hei   int                              // Image size
	rowBytes   []byte                           // Row decoding buffer
	y          int                              // Current y-coordinate
}

// jpegPanic used to distinguish between panics triggered by the
// jpegErrorCallback and other reasons
type jpegPanic struct{}

// NewJPEGReader creates a new [Reader] for JPEG images.
func NewJPEGReader(input io.Reader) (r Reader, err error) {
	// Create reader structure.
	reader := &jpegReader{
		input: input,
	}

	p := C.calloc(C.size_t(unsafe.Sizeof(*reader.jpeg)), 1)
	reader.jpeg = (*C.struct_jpeg_decompress_struct)(p)

	p = C.calloc(C.size_t(unsafe.Sizeof(*reader.jpegErrMgr)), 1)
	reader.jpegErrMgr = (*C.struct_jpeg_error_mgr)(p)

	p = C.calloc(C.size_t(unsafe.Sizeof(*reader.jpegSrcMgr)), 1)
	reader.jpegSrcMgr = (*C.struct_jpeg_source_mgr)(p)

	reader.handle = cgo.NewHandle(reader)

	defer func() {
		if err != nil {
			reader.Close()
		}
	}()

	// Initialize libjpeg stuff
	defer func() {
		p := recover()
		if _, ok := p.(jpegPanic); p != nil && !ok {
			panic(p)
		}

		err = reader.err
	}()

	C.do_jpeg_init_decompress(reader.jpeg,
		reader.jpegErrMgr, reader.jpegSrcMgr, C.uintptr_t(reader.handle))

	rc := C.jpeg_read_header(reader.jpeg, 1)
	if rc != C.JPEG_HEADER_OK {
		err := errors.New("JPEG: invalid header")
		return nil, err
	}

	ok := C.jpeg_start_decompress(reader.jpeg)
	if ok == 0 {
		err := errors.New("JPEG: invalid image")
		return nil, err
	}

	// Obtain image parameters
	reader.wid = int(reader.jpeg.image_width)
	reader.hei = int(reader.jpeg.image_height)

	if reader.jpeg.num_components == 1 {
		reader.model = color.GrayModel
		reader.rowBytes = make([]byte, reader.wid)
	} else {
		reader.model = color.RGBAModel
		reader.rowBytes = make([]byte, reader.wid*3)
	}

	return reader, nil
}

// Close closes the reader.
func (reader *jpegReader) Close() {
	C.jpeg_destroy_decompress(reader.jpeg)

	C.free(unsafe.Pointer(reader.jpeg))
	C.free(unsafe.Pointer(reader.jpegErrMgr))
	C.free(unsafe.Pointer(reader.jpegSrcMgr))

	reader.jpeg = nil
	reader.jpegErrMgr = nil
	reader.jpegSrcMgr = nil

	reader.handle.Delete()
}

// ColorModel returns the [color.Model] of image being decoded.
func (reader *jpegReader) ColorModel() color.Model {
	return reader.model
}

// Size returns the image size.
func (reader *jpegReader) Size() (wid, hei int) {
	return reader.wid, reader.hei
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Reader.Read] function.
func (reader *jpegReader) NewRow() Row {
	return NewRow(reader.model, reader.wid)
}

// Read returns the next image [Row].
func (reader *jpegReader) Read(row Row) (int, error) {
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

	case color.RGBAModel:
		bytesRGB8toRow(row, reader.rowBytes)
	}

	// Update current y
	reader.y++
	if reader.y == reader.hei {
		reader.setError(io.EOF)
	}

	return wid, nil
}

// readRow reads the next image line.
func (reader *jpegReader) readRow() {
	if reader.err != nil {
		return
	}

	defer func() {
		p := recover()
		if _, ok := p.(jpegPanic); p != nil && !ok {
			panic(p)
		}
	}()

	C.do_jpeg_read_scanline(reader.jpeg,
		unsafe.Pointer(&reader.rowBytes[0]))
}

// setError sets the reader.err, if it is not set yet
func (reader *jpegReader) setError(err error) {
	if reader.err == nil {
		reader.err = err
	}
}

// jpegInitSourceCallback called by libjpeg to initialize the source.
//
//export jpegInitSourceCallback
func jpegInitSourceCallback(jpeg C.j_decompress_ptr) {
}

// jpegFillInputBufferCallback called by libjpeg to fill input buffer.
//
//export jpegFillInputBufferCallback
func jpegFillInputBufferCallback(jpeg C.j_decompress_ptr) C.boolean {
	p := (cgo.Handle)(unsafe.Pointer(jpeg.client_data)).Value()
	reader := p.(*jpegReader)

	n, err := reader.input.Read(reader.buf[:])
	if err != nil {
		reader.setError(err)
		return 0
	}

	reader.jpeg.src.bytes_in_buffer = C.size_t(n)
	next := (*C.JOCTET)(unsafe.Pointer(&reader.buf[0]))
	reader.jpeg.src.next_input_byte = next

	return 1
}

// jpegSkipInputDataCallback called by libjpeg to skip input data.
//
//export jpegSkipInputDataCallback
func jpegSkipInputDataCallback(jpeg C.j_decompress_ptr, n C.long) {
	p := (cgo.Handle)(unsafe.Pointer(jpeg.client_data)).Value()
	reader := p.(*jpegReader)

	skip := C.size_t(n)

	// If we have enough bytes in buffer, just update the
	// buffer and we are done.
	if skip <= reader.jpeg.src.bytes_in_buffer {
		reader.jpeg.src.bytes_in_buffer -= skip

		next := unsafe.Pointer(reader.jpeg.src.next_input_byte)
		reader.jpeg.src.next_input_byte = (*C.JOCTET)(
			unsafe.Pointer(uintptr(next) + uintptr(skip)))

		return
	}

	// Otherwise, we need to drain some from the input source
	skip -= reader.jpeg.src.bytes_in_buffer
	reader.jpeg.src.bytes_in_buffer = 0

	lim := io.LimitedReader{R: reader.input, N: int64(skip)}
	_, err := io.Copy(io.Discard, &lim)

	reader.setError(err)
}

// jpegTermSourceCallback called by libjpeg to cleanup the source.
//
//export jpegTermSourceCallback
func jpegTermSourceCallback(jpeg C.j_decompress_ptr) {
}

// jpegWriter implements the [Writer] interface for writing JPEG images
type jpegWriter struct {
	handle     cgo.Handle                     // Handle to self
	jpeg       *C.struct_jpeg_compress_struct // JPEG encoder
	jpegErrMgr *C.struct_jpeg_error_mgr       // JPEG error manager
	jpegDstMgr *C.struct_jpeg_destination_mgr // JPEG destination manager
	buf        [65536]byte                    // Input buffer
	err        error                          // Error from the libjpeg
	output     io.Writer                      // Underlying io.Writer
	wid, hei   int                            // Image size
	model      color.Model                    // Color model
	rowBytes   []byte                         // Row encoding buffer
	y          int                            // Current y-coordinate
}

// NewJPEGWriter creates a new [Writer] for JPEG images.
// Supported color models are following:
//   - color.GrayModel
//   - color.RGBAModel
//
// The quality is the [0...100] integer that defines the tradeoff between
// level of compression and image quality. 0 is the best compression, lowest
// quality, 100 is the best quality, lowest compression.
func NewJPEGWriter(output io.Writer,
	wid, hei int, model color.Model, quality int) (w Writer, err error) {

	// Check model
	if model != color.GrayModel && model != color.RGBAModel {
		err := errors.New("JPEG: unsupported color model")
		return nil, err
	}

	// Create writer structure.
	writer := &jpegWriter{
		output: output,
		wid:    wid,
		hei:    hei,
		model:  model,
	}

	p := C.calloc(C.size_t(unsafe.Sizeof(*writer.jpeg)), 1)
	writer.jpeg = (*C.struct_jpeg_compress_struct)(p)

	p = C.calloc(C.size_t(unsafe.Sizeof(*writer.jpegErrMgr)), 1)
	writer.jpegErrMgr = (*C.struct_jpeg_error_mgr)(p)

	p = C.calloc(C.size_t(unsafe.Sizeof(*writer.jpegDstMgr)), 1)
	writer.jpegDstMgr = (*C.struct_jpeg_destination_mgr)(p)

	writer.handle = cgo.NewHandle(writer)

	defer func() {
		if err != nil {
			writer.Close()
		}
	}()

	// Initialize libjpeg stuff
	defer func() {
		p := recover()
		if _, ok := p.(jpegPanic); p != nil && !ok {
			panic(p)
		}

		err = writer.err
	}()

	C.do_jpeg_init_compress(writer.jpeg,
		writer.jpegErrMgr, writer.jpegDstMgr, C.uintptr_t(writer.handle))

	writer.jpeg.image_width = C.JDIMENSION(wid)
	writer.jpeg.image_height = C.JDIMENSION(hei)
	writer.jpeg.data_precision = 8
	if model == color.RGBAModel {
		writer.jpeg.input_components = 3
		writer.jpeg.in_color_space = C.JCS_RGB
		writer.rowBytes = make([]byte, writer.wid*3)
	} else {
		writer.jpeg.input_components = 1
		writer.jpeg.in_color_space = C.JCS_GRAYSCALE
		writer.rowBytes = make([]byte, writer.wid)
	}

	C.jpeg_set_defaults(writer.jpeg)
	C.jpeg_set_quality(writer.jpeg, C.int(quality), C.TRUE)

	// Use 4:4:4 subsampling
	writer.jpeg.comp_info.h_samp_factor = 1
	writer.jpeg.comp_info.v_samp_factor = 1

	C.jpeg_start_compress(writer.jpeg, 1)

	return writer, nil
}

// Size returns the image size.
func (writer *jpegWriter) Size() (wid, hei int) {
	return writer.wid, writer.hei
}

// ColorModel returns the [color.Model] of image being written.
func (writer *jpegWriter) ColorModel() color.Model {
	return writer.model
}

// Write writes the next image [Row].
func (writer *jpegWriter) Write(row Row) error {
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
	case color.RGBAModel:
		bytesPerPixel = 3
		bytesRGB8fromRow(writer.rowBytes, row)
	}

	// Fill the tail
	if wid < writer.wid {
		end := writer.wid * bytesPerPixel
		for x := wid * bytesPerPixel; x < end; x++ {
			writer.rowBytes[x] = 0xff
		}
	}

	// Write the row
	defer func() {
		p := recover()
		if _, ok := p.(jpegPanic); p != nil && !ok {
			panic(p)
		}
	}()

	C.do_jpeg_write_scanline(writer.jpeg,
		unsafe.Pointer(&writer.rowBytes[0]))

	return writer.err
}

// Close flushes the buffered data and then closes the Writer
func (writer *jpegWriter) Close() error {
	writer.finish()

	C.free(unsafe.Pointer(writer.jpeg))
	C.free(unsafe.Pointer(writer.jpegErrMgr))
	C.free(unsafe.Pointer(writer.jpegDstMgr))

	writer.jpeg = nil
	writer.jpegErrMgr = nil
	writer.jpegDstMgr = nil

	writer.handle.Delete()

	return writer.err
}

// finish ends the compression and flushes buffered output data.
func (writer *jpegWriter) finish() {
	defer func() {
		p := recover()
		if _, ok := p.(jpegPanic); p != nil && !ok {
			panic(p)
		}
	}()

	C.jpeg_finish_compress(writer.jpeg)
}

// setError sets the writer.err, if it is not set yet
func (writer *jpegWriter) setError(err error) {
	if writer.err == nil {
		writer.err = err
	}
}

// flush writes out all data from the output buffer
func (writer *jpegWriter) flush() {
	sz := len(writer.buf) - int(writer.jpeg.dest.free_in_buffer)
	data := writer.buf[:sz]

	// Write all data in the buffer
	for len(data) > 0 && writer.err == nil {
		n, err := writer.output.Write(data)
		data = data[n:]
		writer.setError(err)
	}

	// Reset the buffer
	writer.jpeg.dest.free_in_buffer = C.size_t(len(writer.buf))
	next := (*C.JOCTET)(unsafe.Pointer(&writer.buf[0]))
	writer.jpeg.dest.next_output_byte = next
}

// jpegInitDestination called by libjpeg to initialize the destination.
//
//export jpegInitDestination
func jpegInitDestination(jpeg C.j_compress_ptr) {
	p := (cgo.Handle)(unsafe.Pointer(jpeg.client_data)).Value()
	writer := p.(*jpegWriter)

	writer.jpeg.dest.free_in_buffer = C.size_t(len(writer.buf))
	next := (*C.JOCTET)(unsafe.Pointer(&writer.buf[0]))
	writer.jpeg.dest.next_output_byte = next
}

// jpegEmptyOutputBuffer called by libjpeg to flush out the output buffer.
//
//export jpegEmptyOutputBuffer
func jpegEmptyOutputBuffer(jpeg C.j_compress_ptr) C.boolean {
	p := (cgo.Handle)(unsafe.Pointer(jpeg.client_data)).Value()
	writer := p.(*jpegWriter)

	writer.flush()

	if writer.err != nil {
		return C.TRUE
	}

	return C.FALSE
}

// jpegTermDestination called by libjpeg to finish with the destination.
//
//export jpegTermDestination
func jpegTermDestination(jpeg C.j_compress_ptr) {
	p := (cgo.Handle)(unsafe.Pointer(jpeg.client_data)).Value()
	writer := p.(*jpegWriter)

	writer.flush()
}

// jpegErrorCallback is the error callback.
// It must not return.
//
//export jpegErrorCallback
func jpegErrorCallback(common C.j_common_ptr) {
	jpegOutputMessageCallback(common)
	panic(jpegPanic{})
}

// jpegOutputMessageCallback is the error message output callback.
//
//export jpegOutputMessageCallback
func jpegOutputMessageCallback(common C.j_common_ptr) {
	var buf [C.JMSG_LENGTH_MAX]C.char
	C.do_jpeg_format_message(common, &buf[0])

	p := (cgo.Handle)(unsafe.Pointer(common.client_data)).Value()
	out := p.(interface{ setError(error) })
	err := errors.New(C.GoString(&buf[0]))
	out.setError(err)
}
