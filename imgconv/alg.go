// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image transformations algorithms

package imgconv

// algRow represents a single image row in a form, convenient for processing.
// Every element of the slice is the intensity level, in the range [0.0-1.0]:
//   - YYYY      - for the grayscale images
//   - RGBRGBRGB - for the RGB images
type algRow []float32

// algPixelToU8 converts pixel intensity to uint8
func algPixelToU8(c float32) uint8 {
	if c > 1.0 {
		return 0xff
	}

	return uint8(c * 0xff)
}

// algPixelToU16 converts pixel intensity to uint16
func algPixelToU16(c float32) uint16 {
	if c > 1.0 {
		return 0xffff
	}

	return uint16(c * 0xffff)
}

// newAlgRow allocates the new algRow that matches the given Row.
// The Row used only for its color model and size; data is not imported.
func newAlgRow(row Row) algRow {
	sz := row.Width()

	switch row.(type) {
	case RowGray8:
	case RowGray16:
		sz *= 2
	case RowRGBA32:
		sz *= 3
	case RowRGBA64:
		sz *= 6
	}

	return make(algRow, sz)
}

// Import converts Row into algRow
func (r algRow) Import(row Row) {
	switch row := row.(type) {
	case RowGray8:
		for i := range row {
			r[i] = float32(row[i].Y) / 0xff
		}
	case RowGray16:
		for i := range row {
			r[i] = float32(row[i].Y) / 0xffff
		}
	case RowRGBA32:
		for i := range row {
			o := i * 3
			r[o] = float32(row[i].R) / 0xff
			r[o+1] = float32(row[i].G) / 0xff
			r[o+2] = float32(row[i].B) / 0xff
		}
	case RowRGBA64:
		for i := range row {
			o := i * 6
			r[o] = float32(row[i].R) / 0xffff
			r[o+1] = float32(row[i].G) / 0xffff
			r[o+2] = float32(row[i].B) / 0xffff
		}
	}
}

// Export converts algRow into Row
func (r algRow) Export(row Row) {
	switch row := row.(type) {
	case RowGray8:
		for i := range row {
			row[i].Y = algPixelToU8(r[i])
		}
	case RowGray16:
		for i := range row {
			row[i].Y = algPixelToU16(r[i])
		}
	case RowRGBA32:
		for i := range row {
			o := i * 3
			row[i].R = algPixelToU8(r[o])
			row[i].G = algPixelToU8(r[o+1])
			row[i].B = algPixelToU8(r[o+2])
		}
	case RowRGBA64:
		for i := range row {
			o := i * 6
			row[i].R = algPixelToU16(r[o])
			row[i].G = algPixelToU16(r[o+1])
			row[i].B = algPixelToU16(r[o+2])
		}
	}
}

// Combine computes weighted combination of rows r1 and r2 with
// weights w1 and w2 and saves result into its receiver:
//
//	r = r1 * w1 + r2 * w2
//
// All vectors must be of the same size
func (r algRow) Combine(r1, r2 algRow, w1, w2 float32) {
	for i := range r {
		r[i] = r1[i]*w1 + r2[i]*w2
	}
}
