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
	"io"

	"github.com/OpenPrinting/go-mfp/imgconv"
)

// Filter defines certain transformations of the [Document]s, such
// as recoding into the different format (say, PNG->JPEG), image
// scaling and resizing, brightness and contrast adjustment and so on.
//
// Filter can be applied to the existent Document, returning a
// transformed Document.
type Filter struct {
	format string     // MIME type of the output format
	res    Resolution // Resolution after filtering
	reg    Region     // Image region after filtering
}

// NewFilter creates a new [Filter].
func NewFilter() *Filter {
	return &Filter{
		format: DocumentFormatPNG, // FIXME
	}
}

// SetOutputFormat sets the output format (a MIME type)
func (f *Filter) SetOutputFormat(format string) error {
	f.format = format
	return nil
}

// ChangeResolution adds resolution-change resampling filter.
func (f *Filter) ChangeResolution(x, y int) {
	f.res = Resolution{XResolution: x, YResolution: y}
}

// ChangeRegion adds a filter that extracts image [Region].
//
// Region position is interpreted relative to the top-left
// Document corner. Some parts of the Region or the entire
// Region may fall outside the document boundaries.
//
// It works by cropping document or adding additional space.
func (f *Filter) ChangeRegion(reg Region) {
	f.reg = reg
}

// Apply applies filter to the [Document]. It returns a
// new, filtered Document.
func (f *Filter) Apply(input Document) Document {
	doc := &filterDocument{
		filter: f,
		input:  input,
	}

	return doc
}

// filterDocument represents the filtered [Document].
type filterDocument struct {
	filter  *Filter             // Back link to the filter
	input   Document            // Input document
	curfile *filterDocumentFile // Current DocumentFile, nil if none
}

// Resolution returns the document's rendering resolution in DPI
// (dots per inch).
func (doc *filterDocument) Resolution() Resolution {
	return doc.filter.res
}

// Next returns the next [DocumentFile].
func (doc *filterDocument) Next() (DocumentFile, error) {
	// Close current DocumentFile, if any
	if doc.curfile != nil {
		doc.curfile.close()
		doc.curfile = nil
	}

	// Obtain next DocumentFile from the underlying source Document
	input, err := doc.input.Next()
	if err != nil {
		return nil, err
	}

	// Create filtering pipeline
	pipeline, err := imgconv.NewPNGDecoder(input)
	if err != nil {
		return nil, err
	}

	// Resample to resolution
	filter := doc.filter
	res := doc.input.Resolution()
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
		doc:      doc,
		pipeline: pipeline,
		row:      pipeline.NewRow(),
		output:   &bytes.Buffer{},
	}

	// Create encoder
	wid, hei := pipeline.Size()
	model := pipeline.ColorModel()

	file.encoder, err = imgconv.NewPNGEncoder(file.output, wid, hei, model)
	if err != nil {
		pipeline.Close()
		return nil, err
	}

	return file, io.EOF
}

// Close closes the Document. It implicitly closes the current
// image being read.
func (doc *filterDocument) Close() error {
	if doc.curfile != nil {
		doc.curfile.close()
		doc.curfile = nil
	}
	return nil
}

// filterDocumentFile represents the [DocumentFile] of the
// filtered [Document].
type filterDocumentFile struct {
	doc      *filterDocument // Back link to the filterDocument
	pipeline imgconv.Decoder // Image data decoding/filtering pipeline
	row      imgconv.Row     // Temporary Row for encoding
	output   *bytes.Buffer   // Output stream buffer
	encoder  imgconv.Encoder // Image encoder
}

// Format returns the MIME type of the image format used by
// the document file.
func (file *filterDocumentFile) Format() string {
	return file.doc.filter.format
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
