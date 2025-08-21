// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package abstract

import (
	"bytes"
	"image"

	"github.com/OpenPrinting/go-mfp/imgconv"
)

// Filter runs on a top of existent [Document] and performs various
// transformations of the images, containing in the Document, such
// as changing output format (say, PNG->JPEG), image scaling and
// resizing, brightness and contrast adjustment and so on.
//
// Filter implements the [Document] interface.
type Filter struct {
	input   Document            // Input document
	format  string              // MIME type of the output format
	res     Resolution          // Resolution after filtering
	reg     Region              // Image region after filtering
	curfile *filterDocumentFile // Current DocumentFile, nil if none
}

// NewFilter creates a new [Filter] on a top of existent [Document].
func NewFilter(input Document) *Filter {
	return &Filter{
		input:  input,
		format: imgconv.MIMETypePNG, // FIXME
	}
}

// SetOutputFormat sets the output format (a MIME type)
func (filter *Filter) SetOutputFormat(format string) error {
	filter.format = format
	return nil
}

// SetResolution adds resolution-change resampling filter.
func (filter *Filter) SetResolution(res Resolution) {
	filter.res = res
}

// SetRegion adds a filter that extracts image [Region].
//
// Region position is interpreted relative to the top-left
// Document corner. Some parts of the Region or the entire
// Region may fall outside the document boundaries.
//
// It works by cropping document or adding additional space.
func (filter *Filter) SetRegion(reg Region) {
	filter.reg = reg
}

// Resolution returns the document's rendering resolution in DPI
// (dots per inch).
func (filter *Filter) Resolution() Resolution {
	if !filter.res.IsZero() {
		return filter.res
	}
	return filter.input.Resolution()
}

// Next returns the next [DocumentFile].
func (filter *Filter) Next() (DocumentFile, error) {
	// Close current DocumentFile, if any
	if filter.curfile != nil {
		filter.curfile.close()
		filter.curfile = nil
	}

	// Obtain next DocumentFile from the underlying source Document
	input, err := filter.input.Next()
	if err != nil {
		return nil, err
	}

	// Create filtering pipeline
	_, pipeline, err := imgconv.NewDetectReader(input)
	if err != nil {
		return nil, err
	}

	// Resample to resolution
	res := filter.input.Resolution()
	if !filter.res.IsZero() && filter.res != res {
		wid, hei := pipeline.Size()
		newwid := wid * filter.res.XResolution / res.XResolution
		newhei := hei * filter.res.YResolution / res.YResolution

		pipeline = imgconv.NewScaler(pipeline, newwid, newhei)
		res = filter.res
	}

	// Resize image
	if !filter.reg.IsZero() {
		rect := image.Rect(
			0, 0,
			filter.reg.Width.Dots(res.XResolution),
			filter.reg.Height.Dots(res.XResolution),
		)

		rect = rect.Add(image.Point{
			X: filter.reg.XOffset.Dots(res.XResolution),
			Y: filter.reg.YOffset.Dots(res.XResolution),
		})

		pipeline = imgconv.NewResizer(pipeline, rect)
	}

	// Create filterDocumentFile
	file := &filterDocumentFile{
		filter:   filter,
		input:    input,
		pipeline: pipeline,
		row:      pipeline.NewRow(),
		output:   &bytes.Buffer{},
	}

	// Create encoder
	wid, hei := pipeline.Size()
	model := pipeline.ColorModel()

	file.encoder, err = imgconv.NewPNGWriter(file.output, wid, hei, model)
	if err != nil {
		pipeline.Close()
		return nil, err
	}

	return file, nil
}

// Close closes the Document. It implicitly closes the current
// image being read.
func (filter *Filter) Close() error {
	if filter.curfile != nil {
		filter.curfile.close()
		filter.curfile = nil
	}
	filter.input.Close()
	return nil
}

// filterDocumentFile represents the [DocumentFile] of the
// filtered [Document].
type filterDocumentFile struct {
	filter   *Filter        // Back link to the Filter
	input    DocumentFile   // Underlying DocumentFile
	pipeline imgconv.Reader // Image data decoding/filtering pipeline
	row      imgconv.Row    // Temporary Row for encoding
	output   *bytes.Buffer  // Output stream buffer
	encoder  imgconv.Writer // Image encoder
}

// Format returns the MIME type of the image format used by
// the document file.
func (file *filterDocumentFile) Format() string {
	if file.filter.format != "" {
		return file.filter.format
	}
	return file.input.Format()

}

// Read reads the document file content as a sequence of bytes.
// It implements the [io.Reader] interface.
func (file *filterDocumentFile) Read(buf []byte) (int, error) {
	// Run filtering pipeline until we have some output data
	for file.output.Len() == 0 {
		// Don't let buffer to grow indefinitely
		file.output.Reset()

		// Read and encode the next image Row
		_, err := file.pipeline.Read(file.row)
		if err != nil {
			return 0, err
		}

		err = file.encoder.Write(file.row)
		if err != nil {
			return 0, err
		}
	}

	// Return buffered data
	return file.output.Read(buf)
}

// close closes the filterDocumentFile.
func (file *filterDocumentFile) close() {
	file.pipeline.Close()
}
