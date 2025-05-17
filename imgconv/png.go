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
// // do_png_read_row wraps png_read_row.
// // The wrapper is required to catch setjmp return as
// // we can't do it from Go
// static inline void
// do_png_read_row(png_struct *png, png_bytep row, png_bytep display_row) {
//     if (setjmp(png_jmpbuf(png))) {
//         return;
//     }
//
//     png_read_row(png, row, display_row);
// }
import "C"

// pngEmptyCString is the empty string, suitable for
// passing as parameter to C functions.
var pngEmptyCString = C.CString("")

// pngDecoder is the common part of the PNG [Decoder].
//
// It is used being wrapped by specialized decoders - [pngDecoderGray8],
// [pngDecoderGray16], [pngDecoderRGB24], or [pngDecoderRGB48] - depending
// on the color mode of the image being decoded. Each wrapper implements the
// appropriate Decoder.At method tailored to its specific image format.
type pngDecoder struct {
	handle   cgo.Handle    // Handle to self
	png      *C.png_struct // Underlying png_structure
	pngInfo  *C.png_info   // libpng png_info structure
	err      error         // Error from the libpng
	input    io.Reader     // Underlying io.Reader
	model    color.Model   // Image color mode
	wid, hei int           // Image size
	rowBytes []C.png_byte  // Row decoding buffer
	row      Row           // Decoded row
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

	C.png_get_IHDR(decoder.png, decoder.pngInfo, &width, &height,
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
	decoder.row = NewRow(decoder.model, decoder.wid)

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

// Next returns the next image [Row].
func (decoder *pngDecoder) Next() (Row, error) {
	// Read the next row
	decoder.readRow()
	if decoder.err != nil {
		return nil, decoder.err
	}

	// Decode the row
	switch row := decoder.row.(type) {
	case RowGray8:
		for x := 0; x < decoder.wid; x++ {
			row[x].Y = uint8(decoder.rowBytes[x])
		}

	case RowGray16:
		for x := 0; x < decoder.wid; x++ {
			off := x * 2
			row[x].Y =
				(uint16(decoder.rowBytes[off]) << 8) |
					uint16(decoder.rowBytes[off+1])
		}

	case RowRGBA32:
		for x := 0; x < decoder.wid; x++ {
			off := x * 3
			s := decoder.rowBytes[off : off+3]

			row[x] = color.RGBA{
				R: uint8(s[0]),
				G: uint8(s[1]),
				B: uint8(s[2]),
				A: 255,
			}
		}

	case RowRGBA64:
		for x := 0; x < decoder.wid; x++ {
			off := x * 6
			s := decoder.rowBytes[off : off+6]
			r := (uint16(s[0]) << 8) | uint16(s[1])
			g := (uint16(s[2]) << 8) | uint16(s[3])
			b := (uint16(s[4]) << 8) | uint16(s[5])
			row[x] = color.RGBA64{R: r, G: g, B: b, A: 65535}
		}
	}

	return decoder.row, nil
}

// readRow reads the next image line.
func (decoder *pngDecoder) readRow() {
	if decoder.err == nil {
		C.do_png_read_row(decoder.png, &decoder.rowBytes[0], nil)
	}
}

// pngErrorCallback is called by the libpng to report a error
//
//export pngErrorCallback
func pngErrorCallback(png *C.png_struct, msg C.png_const_charp) {
	decoder := (*cgo.Handle)(C.png_get_io_ptr(png)).Value().(*pngDecoder)

	if decoder.err == nil {
		decoder.err = fmt.Errorf("PNG: %s", C.GoString(msg))
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
			decoder.err = err
			return 0
		}
	}

	return 1
}
