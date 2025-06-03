// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// PNG Reader and Writer test

package imgconv

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"golang.org/x/image/draw"
)

// TestPNGDecode tests PNG reader
func TestPNGDecode(t *testing.T) {
	type testData struct {
		name  string      // Image name, for logging
		data  []byte      // Image data
		model color.Model // Expected color model
		img   image.Image // Decoded reference images
	}

	tests := []testData{
		{
			name:  "PNG100x75rgb8",
			data:  testutils.Images.PNG100x75rgb8,
			model: color.RGBAModel,
		},
		{
			name:  "PNG100x75rgb8paletted",
			data:  testutils.Images.PNG100x75rgb8paletted,
			model: color.RGBAModel,
		},
		{
			name:  "PNG100x75gray1",
			data:  testutils.Images.PNG100x75gray1,
			model: color.GrayModel,
		},
		{
			name:  "PNG100x75gray8",
			data:  testutils.Images.PNG100x75gray8,
			model: color.GrayModel,
		},
		{
			name:  "PNG100x75gray16",
			data:  testutils.Images.PNG100x75gray16,
			model: color.Gray16Model,
		},
		{
			name:  "PNG100x75rgb16",
			data:  testutils.Images.PNG100x75rgb16,
			model: color.RGBA64Model,
		},
	}

	// Decode reference images. We need to do it only once.
	for i := range tests {
		test := &tests[i]
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}
		test.img = reference
	}

	// Test image decoding
	for _, test := range tests {
		// Create a reader
		in := bytes.NewReader(test.data)
		reader, err := NewPNGReader(in)
		if err != nil {
			t.Errorf("%s: NewPNGReader: %s",
				test.name, err)
			continue
		}

		// Decode test image
		img, err := decodeImage(reader)
		if err != nil {
			t.Errorf("%s: decodeImage: %s",
				test.name, err)
			continue
		}

		// Check ColorModel()
		model := reader.ColorModel()
		if model != test.model {
			t.Errorf("%s: Reader.ColorModel mismatch",
				test.name)
		}

		// Compare images
		diff := imageDiff(test.img, img)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
			reader.Close()
			continue
		}

		reader.Close()
	}

	// Test handling of truncated data
	for _, test := range tests {
		step := 1
		for trunc := 0; trunc < len(test.data)-1; trunc += step {
			// Create a reader
			in := bytes.NewReader(test.data[:trunc])
			reader, err := NewPNGReader(in)
			if err != nil {
				continue // Error is expected
			}

			// Once header is decoded, increment step.
			// Otherwise test longs for too long.
			step = 16

			// Decode test image
			img, err := decodeImage(reader)
			if err != nil {
				reader.Close()
				continue // Error is expected
			}

			// Compare images.
			//
			// If everything is OK so far, images must match
			diff := imageDiff(test.img, img)
			if diff != "" {
				t.Errorf("%s: %s", test.name, diff)
				reader.Close()
				continue
			}

			reader.Close()
		}
	}

	// Test handling of damaged data
	for _, test := range tests {
		damagedHeader := false
		damagedData := false
		for off := 0; off < len(test.data); off++ {
			// Scan until we have seen both damaged header
			// and damaged image data
			if damagedHeader && damagedData {
				break
			}

			// Damage data
			data := generic.CopySlice(test.data)
			data[off] = ^data[off]

			// Create a reader
			in := bytes.NewReader(data)
			reader, err := NewPNGReader(in)
			if err != nil {
				damagedHeader = true
				continue // Error is expected
			}

			// Decode test image
			img, err := decodeImage(reader)
			if err != nil {
				damagedData = true
				continue // Error is expected
			}

			// Compare images.
			//
			// If everything is OK so far, images must match
			diff := imageDiff(test.img, img)
			if diff != "" {
				t.Errorf("%s: %s", test.name, diff)
				reader.Close()
				continue
			}

			reader.Close()
		}

		if !(damagedHeader && damagedData) {
			t.Errorf("%s: damaged data not handled properly",
				test.name)
		}
	}

	// Test handling of I/O errors
	for _, test := range tests {
		expectedErr := errors.New("I/O error, for testing")
		damagedHeader := false
		damagedData := false
		fail := false
		for off := 0; off < len(test.data) && fail; off++ {
			// Scan until we have seen both damaged header
			// and damaged image data
			if damagedHeader && damagedData {
				break
			}

			// Simulate I/O error
			rd := newIoReaderWithError(test.data[:off], expectedErr)
			data := generic.CopySlice(test.data)
			data[off] = ^data[off]

			// Create a reader
			reader, err := NewPNGReader(rd)
			if err != nil {
				damagedHeader = true
				if err != expectedErr {
					fail = true
					t.Errorf("%s:in NewPNGReader\n"+
						"error expected: %s\n"+
						"error present:  %s\n",
						test.name, expectedErr, err)
				}
				continue // Error is expected
			}

			img, err := decodeImage(reader)
			if err != nil {
				damagedData = true
				if err != expectedErr {
					fail = true
					t.Errorf("%s:in decodeImage\n"+
						"error expected: %s\n"+
						"error present:  %s\n",
						test.name, expectedErr, err)
				}
				continue // Error is expected
			}

			// Compare images.
			//
			// If everything is OK so far, images must match
			diff := imageDiff(test.img, img)
			if diff != "" {
				t.Errorf("%s: %s", test.name, diff)
				reader.Close()
				continue
			}

			reader.Close()
		}
	}
}

// TestPNGDecodeErrors tests PNG reader errors, not handled by other tests
func TestPNGDecodeErrors(t *testing.T) {
	// Interlaced images not supported
	in := bytes.NewReader(testutils.Images.PNG100x75rgb8i)
	_, err := NewPNGReader(in)

	if err == nil || !strings.Contains(err.Error(), "interlaced") {
		s := "nil"
		if err != nil {
			s = err.Error()
		}
		t.Errorf("PNG: test for interlaced images failed (err=%s)", s)
	}
}

// TestPNGEncode tests PNG writer
func TestPNGEncode(t *testing.T) {
	type testData struct {
		name string      // Image name, for logging
		data []byte      // Image data
		img  image.Image // Decoded reference images
	}

	tests := []testData{
		{
			name: "PNG100x75rgb8",
			data: testutils.Images.PNG100x75rgb8,
		},
		{
			name: "PNG100x75rgb8paletted",
			data: testutils.Images.PNG100x75rgb8paletted,
		},
		{
			name: "PNG100x75gray1",
			data: testutils.Images.PNG100x75gray1,
		},
		{
			name: "PNG100x75gray8",
			data: testutils.Images.PNG100x75gray8,
		},
		{
			name: "PNG100x75gray16",
			data: testutils.Images.PNG100x75gray16,
		},
		{
			name: "PNG100x75rgb16",
			data: testutils.Images.PNG100x75rgb16,
		},
	}

	// Decode reference images. We need to do it only once.
	for i := range tests {
		test := &tests[i]
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}
		test.img = reference
	}

	buf := &bytes.Buffer{}
	for _, test := range tests {
		// Create image reader
		in := bytes.NewReader(test.data)
		reader, err := NewPNGReader(in)
		if err != nil {
			panic(err)
		}

		// Create image writer
		buf.Reset()
		wid, hei := reader.Size()
		model := reader.ColorModel()

		writer, err := NewPNGWriter(buf, wid, hei, model)
		if err != nil {
			t.Errorf("%s: NewPNGWriter: %s", test.name, err)
			reader.Close()
			continue
		}

		// Test Writer.ColorModel method
		newmodel := writer.ColorModel()
		if newmodel != model {
			t.Errorf("%s: Writer.Model mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.name, newmodel, model)
		}

		// Test Writer.Size method
		newwid, newhei := writer.Size()
		if newwid != wid || newhei != hei {
			t.Errorf("%s: Writer.Size mismatch:\n"+
				"expected: %d x %d\n"+
				"present:  %d x %d\n",
				test.name, wid, hei, newwid, newhei)
		}

		// Recode the image, row by row
		row := reader.NewRow()
		for err == nil {
			_, err = reader.Read(row)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			err = writer.Write(row)
			if err != nil {
				t.Errorf("%s: Writer.Write: %s", test.name, err)
				reader.Close()
				writer.Close()
				continue
			}
		}

		reader.Close()
		writer.Close()

		// Decode just encoded image by reference reader
		img, err := png.Decode(buf)
		if err != nil {
			t.Errorf("%s: error in encoded image: %s",
				test.name, err)
			continue
		}

		diff := imageDiff(test.img, img)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
		}
	}
}

// TestPNGEncodeErrors tests I/O errors handling by the PNG writer
func TestPNGEncodeErrors(t *testing.T) {
	// Decode test image
	reader, err := NewPNGReader(
		bytes.NewReader(testutils.Images.PNG100x75rgb8))
	if err != nil {
		panic(err)
	}

	wid, hei := reader.Size()
	rows := mustDecodeImageRows(reader)
	reader.Close()

	// Test I/O errors handling in writer
	ioerr := errors.New("I/O error")
	headerErr := false
	dataErr := false
	for lim := 0; !(headerErr && dataErr); lim++ {
		w := newIoWriterWithError(io.Discard, lim, ioerr)
		writer, err := NewPNGWriter(w, wid, hei, color.RGBAModel)
		if err != nil {
			headerErr = true
		}

		if writer != nil {
			err = encodeImageRows(writer, rows)
			if err != nil {
				dataErr = true
			}

			writer.Close()
		}

		if err != ioerr {
			t.Errorf("Writer I/O error test:\n"+
				"error expected: %v\n"+
				"error present:  %v\n",
				ioerr, err)
			break
		}
	}
}

// TestPNGStickyEncodeError verifies that once I/O error occurs,
// all subsequent write attempts will return the same error
func TestPNGStickyEncodeError(t *testing.T) {
	// Use large image source, so data will not be fully cached
	// by Writer until the end, and we will hit the error
	reader, err := NewPNGReader(
		bytes.NewReader(testutils.Images.PNG5100x7016))
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	ioerr := errors.New("I/O error")
	w := newIoWriterWithError(io.Discard, 4096, ioerr)

	wid, hei := reader.Size()
	writer, err := NewPNGWriter(w, wid, hei, reader.ColorModel())
	if err != nil {
		panic(err)
	}

	defer writer.Close()

	row := reader.NewRow()
	var y int
	for y = 0; y < hei && err == nil; y++ {
		_, err = reader.Read(row)
		if err != nil {
			panic(err)
		}

		err = writer.Write(row)
	}

	for ; y < hei && err == ioerr; y++ {
		err = writer.Write(row)
	}

	if err != ioerr {
		t.Errorf("Writer sticky error test:\n"+
			"error expected: %v\n"+
			"error present:  %v\n",
			ioerr, err)
	}
}

// TestPNGDecodeEncodeTest decodes PNG sample image, then encodes
// it and compares results.
//
// It does it twice: once directly and the second time wrapping the
// image Row into the structure, which effectively prevents Row to
// be used directly for the encoding and implies conversion of
// pixels.
func TestPNGDecodeEncodeTest(t *testing.T) {
	type testData struct {
		name string
		data []byte
	}

	tests := []testData{
		{
			name: "PNG100x75rgb8",
			data: testutils.Images.PNG100x75rgb8,
		},
		{
			name: "PNG100x75rgb8paletted",
			data: testutils.Images.PNG100x75rgb8paletted,
		},
		{
			name: "PNG100x75gray1",
			data: testutils.Images.PNG100x75gray1,
		},
		{
			name: "PNG100x75gray8",
			data: testutils.Images.PNG100x75gray8,
		},
		{
			name: "PNG100x75gray16",
			data: testutils.Images.PNG100x75gray16,
		},
		{
			name: "PNG100x75rgb16",
			data: testutils.Images.PNG100x75rgb16,
		},
	}

	for _, test := range tests {
		// Decode the image
		reader, err := NewPNGReader(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		wid, hei := reader.Size()
		model := reader.ColorModel()
		rows := mustDecodeImageRows(reader)
		reader.Close()

		// Decode initial image with the stdlin reader, for reference
		ref, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		// Encode image directly, then decode with the stdlib reader
		buf := &bytes.Buffer{}

		writer, err := NewPNGWriter(buf, wid, hei, model)
		mustEncodeImageRows(writer, rows)
		writer.Close()

		img, err := png.Decode(buf)
		if err != nil {
			panic(err)
		}

		diff := imageDiff(ref, img)
		if diff != "" {
			t.Errorf("%s: Direct decoding/encoding:\n%s",
				test.name, diff)
		}

		// Encode image with pixels conversion
		buf.Reset()
		writer, err = NewPNGWriter(buf, wid, hei, model)

		for y := 0; y < hei; y++ {
			type wrap struct{ Row }
			row := wrap{rows[y]}
			err := writer.Write(row)
			if err != nil {
				panic(err)
			}
		}

		writer.Close()

		img, err = png.Decode(buf)
		if err != nil {
			panic(err)
		}

		diff = imageDiff(ref, img)
		if diff != "" {
			t.Errorf("%s: Converted decoding/encoding:\n%s",
				test.name, diff)
		}
	}
}

// TestPNGEncodeExpand tests how the PNG writer expands
// image with the insufficient width or height.
func TestPNGEncodeExpand(t *testing.T) {
	type testData struct {
		name string
		data []byte
	}

	tests := []testData{
		{
			name: "PNG100x75rgb8",
			data: testutils.Images.PNG100x75rgb8,
		},
		{
			name: "PNG100x75rgb8paletted",
			data: testutils.Images.PNG100x75rgb8paletted,
		},
		{
			name: "PNG100x75gray1",
			data: testutils.Images.PNG100x75gray1,
		},
		{
			name: "PNG100x75gray8",
			data: testutils.Images.PNG100x75gray8,
		},
		{
			name: "PNG100x75gray16",
			data: testutils.Images.PNG100x75gray16,
		},
		{
			name: "PNG100x75rgb16",
			data: testutils.Images.PNG100x75rgb16,
		},
	}

	for _, test := range tests {
		// Decode the image
		reader, err := NewPNGReader(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		wid, hei := reader.Size()
		model := reader.ColorModel()

		rows := mustDecodeImageRows(reader)
		reader.Close()

		// Halve the image
		halfwid := wid / 2
		halfhei := hei / 2

		rows = rows[:halfhei]
		for y := range rows {
			rows[y] = rows[y].Slice(0, halfwid)
		}

		// Encode the image on its original size
		buf := &bytes.Buffer{}
		writer, err := NewPNGWriter(buf, wid, hei, model)
		if err != nil {
			panic(err)
		}
		mustEncodeImageRows(writer, rows)
		writer.Close()

		// Prepare reference image for validation
		var expected draw.Image
		bounds := image.Rect(0, 0, wid, hei)
		switch model {
		case color.GrayModel:
			expected = image.NewGray(bounds)
		case color.Gray16Model:
			expected = image.NewGray16(bounds)
		case color.RGBAModel:
			expected = image.NewRGBA(bounds)
		case color.RGBA64Model:
			expected = image.NewRGBA64(bounds)
		}

		draw.Draw(expected, bounds, &image.Uniform{color.White},
			image.ZP, draw.Over)

		rect := image.Rect(0, 0, halfwid, halfhei)

		source, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		draw.Draw(expected, rect, source, image.ZP, draw.Over)

		expanded, err := png.Decode(buf)
		if err != nil {
			panic(err)
		}

		diff := imageDiff(expected, expanded)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
		}
	}
}

// TestPNGCornerCases tests those few things that are not covered
// by other tests.
func TestPNGCornerCases(t *testing.T) {
	// NewPNGWriter should not accept unknown color.Model
	_, err := NewPNGWriter(io.Discard, 100, 75, color.AlphaModel)
	if err == nil {
		t.Errorf("NewPNGWriter accepted invalid color.Model")
	}

	// Excessive rows should not affect encoded image
	reader, err := NewPNGReader(
		bytes.NewReader(testutils.Images.PNG100x75rgb8))
	if err != nil {
		panic(err)
	}

	wid, hei := reader.Size()
	model := reader.ColorModel()

	rows := mustDecodeImageRows(reader)
	reader.Close()

	buf1 := &bytes.Buffer{}
	writer, err := NewPNGWriter(buf1, wid, hei, model)
	if err != nil {
		panic(err)
	}

	mustEncodeImageRows(writer, rows)
	writer.Close()

	buf2 := &bytes.Buffer{}
	writer, err = NewPNGWriter(buf2, wid, hei, model)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		for y := range rows {
			err := writer.Write(rows[y])
			if err != nil {
				panic(err)
			}
		}
	}

	writer.Close()

	if !bytes.Equal(buf1.Bytes(), buf2.Bytes()) {
		t.Errorf("Test for encoding excessive rows failed")
	}
}
