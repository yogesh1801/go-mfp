// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// PNG Decoder and Encoder

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
// #include <png.h>
// #include <setjmp.h>
// #include <stdlib.h>
//
// void pngErrorCallback(png_struct *png, png_const_charp msg);
// void pngWarningCallback(png_struct *png, png_const_charp msg);
// void *pngMallocCallback(png_struct *png, size_t size);
// void pngFreeCallback(png_struct *png, void *p);
// int  pngReadCallback(png_struct *png, png_bytep data, size_t size);
// int  pngWriteCallback(png_struct *png, png_bytep data, size_t size);
//
// // do_pngErrorCallback wraps pngErrorCallback.
// // The wrapper is required, because we cannot call png_longjmp from Go.
// static inline void
// do_pngErrorCallback(png_struct *png, const char *message) {
//     pngErrorCallback(png, message);
//     png_longjmp(png, 1);
// }
//
// // do_pngReadCallback wraps pngReadCallback.
// // It calls png_error() in a case of an error, as we can't do it from Go.
// static inline void
// do_pngReadCallback(png_struct *png, png_bytep data, size_t size) {
//     if (!pngReadCallback(png, data, size)) {
//         png_error(png, "");
//     }
// }
//
// // do_pngWriteCallback wraps pngWriteCallback.
// // It calls png_error() in a case of an error, as we can't do it from Go.
// static inline void
// do_pngWriteCallback(png_struct *png, png_bytep data, size_t size) {
//     if (!pngWriteCallback(png, data, size)) {
//         png_error(png, "");
//     }
// }
//
// // do_png_create_read_struct wraps png_create_read_struct_2.
// // This is the convenience wrapper.
// static inline png_struct*
// do_png_create_read_struct(void *p) {
//     png_struct *png;
//
//     png = png_create_read_struct_2(PNG_LIBPNG_VER_STRING,
//         p, do_pngErrorCallback, pngWarningCallback,
//         p, pngMallocCallback, pngFreeCallback);
//
//     png_set_read_fn(png, p, do_pngReadCallback);
//
//     return png;
// }
//
// // do_png_create_write_struct wraps png_create_write_struct_2.
// // This is the convenience wrapper.
// static inline png_struct*
// do_png_create_write_struct(void *p) {
//     png_struct *png;
//
//     png = png_create_write_struct_2(PNG_LIBPNG_VER_STRING,
//         p, do_pngErrorCallback, pngWarningCallback,
//         p, pngMallocCallback, pngFreeCallback);
//
//     png_set_write_fn(png, p, do_pngWriteCallback, NULL);
//
//     return png;
// }
//
// // do_png_read_info wraps png_read_info.
// // The wrapper is required to catch setjmp return as
// // we can't do it from Go
// static inline void
// do_png_read_info(png_struct *png, png_info *info_ptr) {
//     if (setjmp(png_jmpbuf(png))) {
//         return;
//     }
//
//     png_read_info(png, info_ptr);
// }
//
// // do_png_write_info wraps png_write_info.
// // The wrapper is required to catch setjmp return as
// // we can't do it from Go
// static inline void
// do_png_write_info(png_struct *png, png_info *info_ptr) {
//     if (setjmp(png_jmpbuf(png))) {
//         return;
//     }
//
//     png_write_info(png, info_ptr);
// }
//
// // do_png_get_IHDR wraps png_get_IHDR.
// // The wrapper is required to catch setjmp return as
// // we can't do it from Go
// static inline png_uint_32
// do_png_get_IHDR(png_struct *png, png_info *info_ptr,
//                 png_uint_32 *width, png_uint_32 *height, int *bit_depth,
//                 int *color_type, int *interlace_type, int *compression_type,
//                 int *filter_type) {
//
//     if (setjmp(png_jmpbuf(png))) {
//         return 0;
//     }
//
//     return png_get_IHDR(png, info_ptr, width, height, bit_depth,
//                  color_type, interlace_type, compression_type,
//                  filter_type);
// }
//
// // do_png_read_row wraps png_read_row.
// // The wrapper is required to catch setjmp return as we can't do it from Go
// static inline void
// do_png_read_row(png_struct *png, png_bytep row, png_bytep display_row) {
//     if (setjmp(png_jmpbuf(png))) {
//         return;
//     }
//
//     png_read_row(png, row, display_row);
// }
//
// // do_png_write_row wraps png_write_row.
// // The wrapper is required to catch setjmp return as we can't do it from Go
// static inline void
// do_png_write_row(png_struct *png, png_bytep row) {
//     if (setjmp(png_jmpbuf(png))) {
//         return;
//     }
//
//     png_write_row(png, row);
// }
//
// // do_png_write_end wraps png_write_end.
// // The wrapper is required to catch setjmp return as
// // we can't do it from Go
// static inline void
// do_png_write_end(png_struct *png, png_info *info_ptr) {
//     if (setjmp(png_jmpbuf(png))) {
//         return;
//     }
//
//     png_write_end(png, info_ptr);
// }
import "C"

// ErrPNGUnexpectedEOF is returned when the Decoder encounters an io.EOF
// error while more data is still required from the input stream.
var ErrPNGUnexpectedEOF = errors.New("PNG: unexpected EOF")

// pngDecoder implements the [Decoder] interface for reading PNG images.
type pngDecoder struct {
	handle   cgo.Handle    // Handle to self
	png      *C.png_struct // Underlying png_structure
	pngInfo  *C.png_info   // libpng png_info structure
	err      error         // Error from the libpng
	input    io.Reader     // Underlying io.Reader
	model    color.Model   // Image color mode
	wid, hei int           // Image size
	rowBytes []C.png_byte  // Row decoding buffer
	y        int           // Current y-coordinate
}

// NewPNGDecoder creates a new [Decoder] for PNG images
func NewPNGDecoder(input io.Reader) (Decoder, error) {
	// Create decoder structure. Initialize libpng stuff
	decoder := &pngDecoder{input: input}
	decoder.handle = cgo.NewHandle(decoder)

	decoder.png = C.do_png_create_read_struct(
		unsafe.Pointer(&decoder.handle))

	decoder.pngInfo = C.png_create_info_struct(decoder.png)

	// Read image info
	C.do_png_read_info(decoder.png, decoder.pngInfo)

	var width, height C.png_uint_32
	var depth, colorType, interlace C.int

	C.do_png_get_IHDR(decoder.png, decoder.pngInfo, &width, &height,
		&depth, &colorType, &interlace, nil, nil)

	if err := decoder.err; err != nil {
		decoder.Close()
		return nil, err
	}

	if interlace != C.PNG_INTERLACE_NONE {
		decoder.Close()
		err := errors.New("PNG: interlaced images not supported")
		return nil, err
	}

	decoder.wid, decoder.hei = int(width), int(height)

	// Setup input transformations
	var bytesPerPixel int

	gray := (colorType & C.PNG_COLOR_MASK_COLOR) == C.PNG_COLOR_TYPE_GRAY

	if colorType == C.PNG_COLOR_TYPE_PALETTE {
		C.png_set_palette_to_rgb(decoder.png)
	}

	if gray && depth < 8 {
		C.png_set_expand_gray_1_2_4_to_8(decoder.png)
	}

	if (colorType & C.PNG_COLOR_MASK_ALPHA) != 0 {
		C.png_set_strip_alpha(decoder.png)
	}

	if gray {
		decoder.model = color.GrayModel
		bytesPerPixel = 1
		if depth == 16 {
			decoder.model = color.Gray16Model
			bytesPerPixel = 2
		}
	} else {
		decoder.model = color.RGBAModel
		bytesPerPixel = 3
		if depth == 16 {
			decoder.model = color.RGBA64Model
			bytesPerPixel = 6
		}
	}

	// Allocate buffers
	decoder.rowBytes = make([]C.png_byte, bytesPerPixel*int(width))

	return decoder, nil
}

// Close closes the decoder.
func (decoder *pngDecoder) Close() {
	C.png_destroy_read_struct(&decoder.png, &decoder.pngInfo, nil)
	decoder.handle.Delete()
}

// ColorModel returns the [color.Model] of image being decoded.
func (decoder *pngDecoder) ColorModel() color.Model {
	return decoder.model
}

// Size returns the image size.
func (decoder *pngDecoder) Size() (wid, hei int) {
	return decoder.wid, decoder.hei
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Decoder.Read] function.
func (decoder *pngDecoder) NewRow() Row {
	return NewRow(decoder.model, decoder.wid)
}

// Read returns the next image [Row].
func (decoder *pngDecoder) Read(row Row) (int, error) {
	// Read the next row
	decoder.readRow()
	if decoder.err != nil {
		return 0, decoder.err
	}

	// Decode the row
	wid := generic.Min(row.Width(), decoder.wid)

	switch decoder.model {
	case color.GrayModel:
		row := row.(RowGray8)
		for x := 0; x < wid; x++ {
			row[x].Y = uint8(decoder.rowBytes[x])
		}

	case color.Gray16Model:
		row := row.(RowGray16)
		for x := 0; x < wid; x++ {
			off := x * 2
			row[x].Y =
				(uint16(decoder.rowBytes[off]) << 8) |
					uint16(decoder.rowBytes[off+1])
		}

	case color.RGBAModel:
		row := row.(RowRGBA32)
		for x := 0; x < wid; x++ {
			off := x * 3
			s := decoder.rowBytes[off : off+3]

			row[x] = color.RGBA{
				R: uint8(s[0]),
				G: uint8(s[1]),
				B: uint8(s[2]),
				A: 255,
			}
		}

	case color.RGBA64Model:
		row := row.(RowRGBA64)
		for x := 0; x < wid; x++ {
			off := x * 6
			s := decoder.rowBytes[off : off+6]
			r := (uint16(s[0]) << 8) | uint16(s[1])
			g := (uint16(s[2]) << 8) | uint16(s[3])
			b := (uint16(s[4]) << 8) | uint16(s[5])
			row[x] = color.RGBA64{R: r, G: g, B: b, A: 65535}
		}
	}

	// Update current y
	decoder.y++
	if decoder.y == decoder.hei {
		decoder.err = io.EOF
	}

	return wid, nil
}

// readRow reads the next image line.
func (decoder *pngDecoder) readRow() {
	if decoder.err == nil {
		C.do_png_read_row(decoder.png, &decoder.rowBytes[0], nil)
	}
}

// pngEncoder implements the [Encoder] interface for writing PNG images
type pngEncoder struct {
	handle   cgo.Handle    // Handle to self
	png      *C.png_struct // Underlying png_structure
	pngInfo  *C.png_info   // libpng png_info structure
	err      error         // Error from the libpng
	output   io.Writer     // Underlying io.Writer
	wid, hei int           // Image size
	model    color.Model   // Color model
	rowBytes []C.png_byte  // Row encoding buffer
	y        int           // Current y-coordinate
}

// NewPNGEncoder creates a new [Encoder] for PNG images.
// Supported color models are following:
//   - color.GrayModel
//   - color.Gray16Model
//   - color.RGBAModel
//   - color.RGBA64Model
func NewPNGEncoder(output io.Writer,
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

	// Create encoder structure. Initialize libpng stuff
	encoder := &pngEncoder{
		output:   output,
		wid:      wid,
		hei:      hei,
		model:    model,
		rowBytes: make([]C.png_byte, bytesPerPixel*int(wid)),
	}

	encoder.handle = cgo.NewHandle(encoder)

	encoder.png = C.do_png_create_write_struct(
		unsafe.Pointer(&encoder.handle))

	encoder.pngInfo = C.png_create_info_struct(encoder.png)

	// Create PNG header
	C.png_set_IHDR(encoder.png, encoder.pngInfo,
		C.png_uint_32(wid), C.png_uint_32(hei),
		depth, colorType,
		C.PNG_INTERLACE_NONE,
		C.PNG_COMPRESSION_TYPE_DEFAULT,
		C.PNG_FILTER_TYPE_DEFAULT)

	C.png_set_sRGB(encoder.png, encoder.pngInfo,
		C.PNG_sRGB_INTENT_PERCEPTUAL)

	C.do_png_write_info(encoder.png, encoder.pngInfo)

	if encoder.err != nil {
		encoder.Close()
		return nil, encoder.err
	}

	return encoder, nil
}

// Write writes the next image [Row].
func (encoder *pngEncoder) Write(row Row) error {
	// Check for pending error
	if encoder.err != nil {
		return encoder.err
	}

	// Silently ignore excessive rows
	if encoder.y == encoder.hei {
		return nil
	}

	// Encode the row
	wid := generic.Min(row.Width(), encoder.wid)

	var bytesPerPixel int

	switch encoder.model {
	case color.GrayModel:
		bytesPerPixel = 1
		encoder.encodeGray8(row, wid)
	case color.Gray16Model:
		bytesPerPixel = 2
		encoder.encodeGray16(row, wid)
	case color.RGBAModel:
		bytesPerPixel = 3
		encoder.encodeRGBA32(row, wid)
	case color.RGBA64Model:
		bytesPerPixel = 6
		encoder.encodeRGBA64(row, wid)
	}

	// Fill the tail
	if wid < encoder.wid {
		end := encoder.wid * bytesPerPixel
		for x := wid * bytesPerPixel; x < end; x++ {
			encoder.rowBytes[x] = 0xff
		}
	}

	// Write the row
	C.do_png_write_row(encoder.png, &encoder.rowBytes[0])
	if encoder.err == nil {
		encoder.y++
	}

	return encoder.err
}

// Close flushes the buffered data and then closes the Encoder
func (encoder *pngEncoder) Close() error {
	// Write missed lines
	for encoder.err == nil && encoder.y < encoder.hei {
		encoder.Write(RowEmpty{})
	}

	// Finish PNG image
	if encoder.err == nil {
		C.do_png_write_end(encoder.png, nil)
	}

	// Release allocated resources
	C.png_destroy_write_struct(&encoder.png, &encoder.pngInfo)
	encoder.handle.Delete()

	return encoder.err
}

// encodeGray8 encodes a row of the 8-bit grayscale image
func (encoder *pngEncoder) encodeGray8(row Row, wid int) {
	if row, ok := row.(RowGray8); ok {
		for x := 0; x < wid; x++ {
			encoder.rowBytes[x] = C.png_byte(row[x].Y)
		}
		return
	}

	for x := 0; x < wid; x++ {
		c := color.GrayModel.Convert(row.At(x)).(color.Gray)
		encoder.rowBytes[x] = C.png_byte(c.Y)
	}
}

// encodeGray16 encodes a row of the 16-bit grayscale image
func (encoder *pngEncoder) encodeGray16(row Row, wid int) {
	if row, ok := row.(RowGray16); ok {
		for x := 0; x < wid; x++ {
			off := x * 2
			encoder.rowBytes[off] = C.png_byte(row[x].Y >> 8)
			encoder.rowBytes[off+1] = C.png_byte(row[x].Y)
		}
		return
	}

	for x := 0; x < wid; x++ {
		off := x * 2
		c := color.Gray16Model.Convert(row.At(x)).(color.Gray16)
		encoder.rowBytes[off] = C.png_byte(c.Y >> 8)
		encoder.rowBytes[off+1] = C.png_byte(c.Y)
	}
}

// encodeRGBA32 encodes a row of the 32-bit RGBA image
func (encoder *pngEncoder) encodeRGBA32(row Row, wid int) {
	if row, ok := row.(RowRGBA32); ok {
		for x := 0; x < wid; x++ {
			off := x * 3
			s := encoder.rowBytes[off : off+3]

			s[0] = C.png_byte(row[x].R)
			s[1] = C.png_byte(row[x].G)
			s[2] = C.png_byte(row[x].B)
		}
		return
	}

	for x := 0; x < wid; x++ {
		off := x * 3
		s := encoder.rowBytes[off : off+3]

		c := color.RGBAModel.Convert(row.At(x)).(color.RGBA)
		s[0] = C.png_byte(c.R)
		s[1] = C.png_byte(c.G)
		s[2] = C.png_byte(c.B)
	}
}

// encodeRGBA64 encodes a row of the 64-bit RGBA image
func (encoder *pngEncoder) encodeRGBA64(row Row, wid int) {
	if row, ok := row.(RowRGBA64); ok {
		for x := 0; x < wid; x++ {
			off := x * 6
			s := encoder.rowBytes[off : off+6]

			s[0] = C.png_byte(row[x].R >> 8)
			s[1] = C.png_byte(row[x].R)
			s[2] = C.png_byte(row[x].G >> 8)
			s[3] = C.png_byte(row[x].G)
			s[4] = C.png_byte(row[x].B >> 8)
			s[5] = C.png_byte(row[x].B)
		}
		return
	}

	for x := 0; x < wid; x++ {
		off := x * 6
		s := encoder.rowBytes[off : off+6]

		c := color.RGBAModel.Convert(row.At(x)).(color.RGBA64)
		s[0] = C.png_byte(c.R >> 8)
		s[1] = C.png_byte(c.R)
		s[2] = C.png_byte(c.G >> 8)
		s[3] = C.png_byte(c.G)
		s[4] = C.png_byte(c.B >> 8)
		s[5] = C.png_byte(c.B)
	}
}

// pngErrorCallback is called by the libpng to report a error
// This is the common callback for pngDecoder and pngEncoder
//
//export pngErrorCallback
func pngErrorCallback(png *C.png_struct, msg C.png_const_charp) {
	p := (*cgo.Handle)(C.png_get_io_ptr(png)).Value()
	switch p := p.(type) {
	case (*pngDecoder):
		if p.err == nil {
			p.err = fmt.Errorf("PNG: %s", C.GoString(msg))
		}
	case (*pngEncoder):
		if p.err == nil {
			p.err = fmt.Errorf("PNG: %s", C.GoString(msg))
		}
	}
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

	decoder := (*cgo.Handle)(C.png_get_io_ptr(png)).Value().(*pngDecoder)

	buf := (*[max]byte)(unsafe.Pointer(data))[:sz:sz]
	for len(buf) > 0 {
		n, err := decoder.input.Read(buf)
		if n > 0 {
			buf = buf[n:]
		} else if err != nil {
			if err == io.EOF {
				err = ErrPNGUnexpectedEOF
			}
			decoder.err = err
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

	encoder := (*cgo.Handle)(C.png_get_io_ptr(png)).Value().(*pngEncoder)

	buf := (*[max]byte)(unsafe.Pointer(data))[:sz:sz]
	for len(buf) > 0 {
		n, err := encoder.output.Write(buf)
		if n > 0 {
			buf = buf[n:]
		} else if err != nil {
			encoder.err = err
			return 0
		}
	}

	return 1
}
