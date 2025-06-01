// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversions between Rows and byte slices

package imgconv

import (
	"image/color"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// bytesGray8toRow converts byte slice to Row.
// The byte slice assumed to contain 8-bit grayscale image line.
// It returns the resulting Row length, in pixels.
func bytesGray8toRow(r Row, bytes []byte) int {
	wid := generic.Min(r.Width(), len(bytes))

	switch r := r.(type) {
	case RowGray8:
		for x := 0; x < wid; x++ {
			r[x].Y = bytes[x]
		}
	case RowGrayFP32:
		for x := 0; x < wid; x++ {
			r[x] = float32(bytes[x]) / 0xff
		}
	default:
		for x := 0; x < wid; x++ {
			r.Set(x, color.Gray{bytes[x]})
		}
	}

	return wid
}

// bytesGray16BEtoRow converts byte slice to Row.
// The byte slice assumed to contain 16-bit big endian grayscale image line.
// It returns the resulting Row length, in pixels.
func bytesGray16BEtoRow(r Row, bytes []byte) int {
	wid := generic.Min(r.Width(), len(bytes)/2)

	switch r := r.(type) {
	case RowGray16:
		for x := 0; x < wid; x++ {
			off := x * 2
			v := (uint16(bytes[off]) << 8) | uint16(bytes[off+1])
			r[x].Y = v
		}
	case RowGrayFP32:
		for x := 0; x < wid; x++ {
			off := x * 2
			v := (uint16(bytes[off]) << 8) | uint16(bytes[off+1])
			r[x] = float32(v) / 0xffff
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 2
			v := (uint16(bytes[off]) << 8) | uint16(bytes[off+1])
			r.Set(x, color.Gray16{v})
		}
	}

	return wid
}

// bytesRGB8toRow converts byte slice to Row.
// The byte slice assumed to contain 8-bit R-G-B image line.
// It returns the resulting Row length, in pixels.
func bytesRGB8toRow(r Row, bytes []byte) int {
	wid := generic.Min(r.Width(), len(bytes)/3)

	switch r := r.(type) {
	case RowRGBA32:
		for x := 0; x < wid; x++ {
			off := x * 3
			s := bytes[off : off+3]

			r[x] = color.RGBA{
				R: uint8(s[0]),
				G: uint8(s[1]),
				B: uint8(s[2]),
				A: 255,
			}
		}
	case RowRGBAFP32:
		for x := 0; x < wid; x++ {
			soff := x * 3
			doff := x * 4
			s := bytes[soff : soff+3]
			d := r[doff : doff+4]

			d[0] = float32(s[0]) / 0xff
			d[1] = float32(s[1]) / 0xff
			d[2] = float32(s[2]) / 0xff
			d[3] = 1.0
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 3
			s := bytes[off : off+3]

			c := color.RGBA{
				R: uint8(s[0]),
				G: uint8(s[1]),
				B: uint8(s[2]),
				A: 255,
			}

			r.Set(x, c)
		}
	}

	return wid
}

// bytesRGB16toRow converts byte slice to Row.
// The byte slice assumed to contain 16-bit big endian R-G-B image line.
// It returns the resulting Row length, in pixels.
func bytesRGB16toRow(r Row, bytes []byte) int {
	wid := generic.Min(r.Width(), len(bytes)/6)

	switch r := r.(type) {
	case RowRGBA64:
		for x := 0; x < wid; x++ {
			off := x * 6
			s := bytes[off : off+6]

			R := (uint16(s[0]) << 8) | uint16(s[1])
			G := (uint16(s[2]) << 8) | uint16(s[3])
			B := (uint16(s[4]) << 8) | uint16(s[5])

			r[x] = color.RGBA64{R: R, G: G, B: B, A: 65535}
		}
	case RowRGBAFP32:
		for x := 0; x < wid; x++ {
			soff := x * 6
			doff := x * 4
			s := bytes[soff : soff+6]
			d := r[doff : doff+4]

			d[0] = float32((uint16(s[0])<<8)|uint16(s[1])) / 0xffff
			d[1] = float32((uint16(s[2])<<8)|uint16(s[3])) / 0xffff
			d[2] = float32((uint16(s[4])<<8)|uint16(s[5])) / 0xffff
			d[3] = 1.0
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 6
			s := bytes[off : off+6]

			R := (uint16(s[0]) << 8) | uint16(s[1])
			G := (uint16(s[2]) << 8) | uint16(s[3])
			B := (uint16(s[4]) << 8) | uint16(s[5])

			r.Set(x, color.RGBA64{R: R, G: G, B: B, A: 65535})
		}
	}

	return wid
}

// bytesRGBA8toRow converts byte slice to Row.
// The byte slice assumed to contain 8-bit R-G-B-A image line.
// It returns the resulting Row length, in pixels.
func bytesRGBA8toRow(r Row, bytes []byte) int {
	wid := generic.Min(r.Width(), len(bytes)/4)

	switch r := r.(type) {
	case RowRGBA32:
		for x := 0; x < wid; x++ {
			off := x * 4
			s := bytes[off : off+4]

			r[x] = color.RGBA{
				R: uint8(s[0]),
				G: uint8(s[1]),
				B: uint8(s[2]),
				A: uint8(s[3]),
			}
		}
	case RowRGBAFP32:
		for x := 0; x < wid; x++ {
			off := x * 4
			s := bytes[off : off+4]
			d := r[off : off+4]

			d[0] = float32(s[0]) / 0xff
			d[1] = float32(s[1]) / 0xff
			d[2] = float32(s[2]) / 0xff
			d[3] = float32(s[3]) / 0xff
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 4
			s := bytes[off : off+4]

			c := color.RGBA{
				R: uint8(s[0]),
				G: uint8(s[1]),
				B: uint8(s[2]),
				A: uint8(s[3]),
			}

			r.Set(x, c)
		}
	}

	return wid
}

// bytesRGBA16toRow converts byte slice to Row.
// The byte slice assumed to contain 16-bit big endian R-G-B-A image line.
// It returns the resulting Row length, in pixels.
func bytesRGBA16toRow(r Row, bytes []byte) int {
	wid := generic.Min(r.Width(), len(bytes)/8)

	switch r := r.(type) {
	case RowRGBA64:
		for x := 0; x < wid; x++ {
			off := x * 8
			s := bytes[off : off+8]

			R := (uint16(s[0]) << 8) | uint16(s[1])
			G := (uint16(s[2]) << 8) | uint16(s[3])
			B := (uint16(s[4]) << 8) | uint16(s[5])
			A := (uint16(s[6]) << 8) | uint16(s[7])

			r[x] = color.RGBA64{R: R, G: G, B: B, A: A}
		}
	case RowRGBAFP32:
		for x := 0; x < wid; x++ {
			soff := x * 8
			doff := x * 4
			s := bytes[soff : soff+8]
			d := r[doff : doff+4]

			d[0] = float32((uint16(s[0])<<8)|uint16(s[1])) / 0xffff
			d[1] = float32((uint16(s[2])<<8)|uint16(s[3])) / 0xffff
			d[2] = float32((uint16(s[4])<<8)|uint16(s[5])) / 0xffff
			d[3] = float32((uint16(s[6])<<8)|uint16(s[7])) / 0xffff
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 8
			s := bytes[off : off+8]

			R := (uint16(s[0]) << 8) | uint16(s[1])
			G := (uint16(s[2]) << 8) | uint16(s[3])
			B := (uint16(s[4]) << 8) | uint16(s[5])
			A := (uint16(s[6]) << 8) | uint16(s[7])

			r.Set(x, color.RGBA64{R: R, G: G, B: B, A: A})
		}
	}

	return wid
}

// bytesGray8fromRow converts Row to byte slice.
// The byte slice assumed to contain 8-bit grayscale image line.
// It returns the resulting Row length, in pixels.
func bytesGray8fromRow(bytes []byte, r Row) int {
	wid := generic.Min(r.Width(), len(bytes))

	switch r := r.(type) {
	case RowGray8:
		for x := 0; x < wid; x++ {
			bytes[x] = r[x].Y
		}
	case RowGrayFP32:
		for x := 0; x < wid; x++ {
			bytes[x] = bytesFPtoU8(r[x])
		}
	default:
		for x := 0; x < wid; x++ {
			c := color.GrayModel.Convert(r.At(x)).(color.Gray)
			bytes[x] = c.Y
		}
	}

	return wid
}

// bytesGray16BEfromRow converts Row to byte slice.
// The byte slice assumed to contain 16-bit big endian grayscale image line.
// It returns the resulting Row length, in pixels.
func bytesGray16BEfromRow(bytes []byte, r Row) int {
	wid := generic.Min(r.Width(), len(bytes)/2)

	switch r := r.(type) {
	case RowGray16:
		for x := 0; x < wid; x++ {
			off := x * 2
			bytes[off] = uint8(r[x].Y >> 8)
			bytes[off+1] = uint8(r[x].Y)
		}
	case RowGrayFP32:
		for x := 0; x < wid; x++ {
			off := x * 2
			v := bytesFPtoU16(r[x])
			bytes[off] = uint8(v >> 8)
			bytes[off+1] = uint8(v)
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 2
			c := color.Gray16Model.Convert(r.At(x)).(color.Gray16)
			bytes[off] = uint8(c.Y >> 8)
			bytes[off+1] = uint8(c.Y)
		}
	}

	return wid * 2
}

// bytesRGB8fromRow converts Row to byte slice.
// The byte slice assumed to contain 8-bit R-G-B image line.
// It returns the resulting Row length, in pixels.
func bytesRGB8fromRow(bytes []byte, r Row) int {
	wid := generic.Min(r.Width(), len(bytes)/3)

	switch r := r.(type) {
	case RowRGBA32:
		for x := 0; x < wid; x++ {
			off := x * 3
			d := bytes[off : off+3]
			c := r[x]

			d[0] = c.R
			d[1] = c.G
			d[2] = c.B
		}
	case RowRGBAFP32:
		for x := 0; x < wid; x++ {
			soff := x * 4
			doff := x * 3
			s := r[soff : soff+4]
			d := bytes[doff : doff+3]

			d[0] = bytesFPtoU8(s[0])
			d[1] = bytesFPtoU8(s[1])
			d[2] = bytesFPtoU8(s[2])
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 3
			d := bytes[off : off+3]
			c := color.RGBAModel.Convert(r.At(x)).(color.RGBA)

			d[0] = c.R
			d[1] = c.G
			d[2] = c.B
		}
	}

	return wid * 3
}

// bytesRGB16fromRow converts Row to byte slice.
// The byte slice assumed to contain 16-bit big endian R-G-B image line.
// It returns the resulting Row length, in pixels.
func bytesRGB16fromRow(bytes []byte, r Row) int {
	wid := generic.Min(r.Width(), len(bytes)/6)

	switch r := r.(type) {
	case RowRGBA64:
		for x := 0; x < wid; x++ {
			off := x * 6
			d := bytes[off : off+6]
			c := r[x]

			d[0] = uint8(c.R >> 8)
			d[1] = uint8(c.R)
			d[2] = uint8(c.G >> 8)
			d[3] = uint8(c.G)
			d[4] = uint8(c.B >> 8)
			d[5] = uint8(c.B)
		}
	case RowRGBAFP32:
		for x := 0; x < wid; x++ {
			soff := x * 4
			doff := x * 6
			s := r[soff : soff+4]
			d := bytes[doff : doff+6]

			R := bytesFPtoU16(s[0])
			G := bytesFPtoU16(s[1])
			B := bytesFPtoU16(s[2])

			d[0] = uint8(R >> 8)
			d[1] = uint8(R)
			d[2] = uint8(G >> 8)
			d[3] = uint8(G)
			d[4] = uint8(B >> 8)
			d[5] = uint8(B)
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 6
			d := bytes[off : off+6]

			c := color.RGBA64Model.Convert(r.At(x)).(color.RGBA64)

			d[0] = uint8(c.R >> 8)
			d[1] = uint8(c.R)
			d[2] = uint8(c.G >> 8)
			d[3] = uint8(c.G)
			d[4] = uint8(c.B >> 8)
			d[5] = uint8(c.B)
		}
	}

	return wid * 6
}

// bytesRGBA8fromRow converts Row to byte slice.
// The byte slice assumed to contain 8-bit R-G-B-A image line.
// It returns the resulting Row length, in pixels.
func bytesRGBA8fromRow(bytes []byte, r Row) int {
	wid := generic.Min(r.Width(), len(bytes)/4)

	switch r := r.(type) {
	case RowRGBA32:
		for x := 0; x < wid; x++ {
			off := x * 4
			d := bytes[off : off+4]
			c := r[x]

			d[0] = c.R
			d[1] = c.G
			d[2] = c.B
			d[3] = c.A
		}
	case RowRGBAFP32:
		for x := 0; x < wid; x++ {
			off := x * 4
			s := r[off : off+4]
			d := bytes[off : off+4]

			d[0] = bytesFPtoU8(s[0])
			d[1] = bytesFPtoU8(s[1])
			d[2] = bytesFPtoU8(s[2])
			d[3] = bytesFPtoU8(s[3])
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 4
			d := bytes[off : off+4]
			c := color.RGBAModel.Convert(r.At(x)).(color.RGBA)

			d[0] = c.R
			d[1] = c.G
			d[2] = c.B
			d[3] = c.A
		}
	}

	return wid * 4
}

// bytesRGBA16fromRow converts Row to byte slice.
// The byte slice assumed to contain 16-bit big endian R-G-B-A image line.
// It returns the resulting Row length, in pixels.
func bytesRGBA16fromRow(bytes []byte, r Row) int {
	wid := generic.Min(r.Width(), len(bytes)/8)

	switch r := r.(type) {
	case RowRGBA64:
		for x := 0; x < wid; x++ {
			off := x * 8
			d := bytes[off : off+8]
			c := r[x]

			d[0] = uint8(c.R >> 8)
			d[1] = uint8(c.R)
			d[2] = uint8(c.G >> 8)
			d[3] = uint8(c.G)
			d[4] = uint8(c.B >> 8)
			d[5] = uint8(c.B)
			d[6] = uint8(c.A >> 8)
			d[7] = uint8(c.A)
		}
	case RowRGBAFP32:
		for x := 0; x < wid; x++ {
			soff := x * 4
			doff := x * 8
			s := r[soff : soff+4]
			d := bytes[doff : doff+8]

			R := bytesFPtoU16(s[0])
			G := bytesFPtoU16(s[1])
			B := bytesFPtoU16(s[2])
			A := bytesFPtoU16(s[3])

			d[0] = uint8(R >> 8)
			d[1] = uint8(R)
			d[2] = uint8(G >> 8)
			d[3] = uint8(G)
			d[4] = uint8(B >> 8)
			d[5] = uint8(B)
			d[6] = uint8(A >> 8)
			d[7] = uint8(A)
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 8
			d := bytes[off : off+8]

			c := color.RGBA64Model.Convert(r.At(x)).(color.RGBA64)

			d[0] = uint8(c.R >> 8)
			d[1] = uint8(c.R)
			d[2] = uint8(c.G >> 8)
			d[3] = uint8(c.G)
			d[4] = uint8(c.B >> 8)
			d[5] = uint8(c.B)
			d[6] = uint8(c.A >> 8)
			d[7] = uint8(c.A)
		}
	}

	return wid * 8
}

// bytesFPtoU8 converts float32 valie in range [0.0...1.0] into the
// uint8 value in range [0...0xff]
func bytesFPtoU8(f float32) uint8 {
	return uint8(generic.Min(f, 1.0) * 0xff)
}

// bytesFPtoU16 converts float32 valie in range [0.0...1.0] into the
// uint16 value in range [0...0xffff]
func bytesFPtoU16(f float32) uint16 {
	return uint16(generic.Min(f, 1.0) * 0xffff)
}
