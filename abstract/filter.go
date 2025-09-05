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
	"image/color"
	"io"

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
	opt     FilterOptions       // Filter options
	curfile *filterDocumentFile // Current DocumentFile, nil if none
}

// FilterOptions define image transformations, performed
// by the [Filter].
type FilterOptions struct {
	// OutputFormat specified the MIME type of the output
	// image. If set to "", the output format will be choosen
	// automatically.
	OutputFormat string

	// Res requests image resampling into the specified resolution.
	// Use zero value of [Resolution] to skip this step.
	Res Resolution

	// Reg requests image clipping to the specified region.
	// Use zero value of [Region] to skip this step.
	Reg Region

	// Mode requests image conversion into the particular
	// [ColorMode].
	// Use [ColorModeUnset] to bypass this step.
	Mode ColorMode

	// Depth requests image conversion into the different
	// color depth/
	// Use [ColorDepthUnset] to bypass this step.
	Depth ColorDepth
}

// NewFilter creates a new [Filter] on a top of existent [Document].
func NewFilter(input Document, opt FilterOptions) *Filter {
	filter := &Filter{
		input: input,
		opt:   opt,
	}

	if filter.opt.OutputFormat == "" {
		filter.opt.OutputFormat = imgconv.MIMETypePNG //FIXME
	}

	return filter
}

// Resolution returns the document's rendering resolution in DPI
// (dots per inch).
func (filter *Filter) Resolution() Resolution {
	if !filter.opt.Res.IsZero() {
		return filter.opt.Res
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
	pipeline, err := imgconv.NewDetectReader(input)
	if err != nil {
		return nil, err
	}

	// Resample to resolution
	res := filter.input.Resolution()
	if !filter.opt.Res.IsZero() && filter.opt.Res != res {
		wid, hei := pipeline.Size()
		newwid := wid * filter.opt.Res.XResolution / res.XResolution
		newhei := hei * filter.opt.Res.YResolution / res.YResolution

		pipeline = imgconv.NewScaler(pipeline, newwid, newhei)
		res = filter.opt.Res
	}

	// Resize image
	if !filter.opt.Reg.IsZero() {
		rect := image.Rect(
			0, 0,
			filter.opt.Reg.Width.Dots(res.XResolution),
			filter.opt.Reg.Height.Dots(res.XResolution),
		)

		rect = rect.Add(image.Point{
			X: filter.opt.Reg.XOffset.Dots(res.XResolution),
			Y: filter.opt.Reg.YOffset.Dots(res.XResolution),
		})

		pipeline = imgconv.NewResizer(pipeline, rect)
	}

	// Honor color mode conversion options
	model := pipeline.ColorModel()

	switch filter.opt.Mode {
	case ColorModeMono:
		switch model {
		case color.RGBAModel:
			model = color.GrayModel
		case color.RGBA64Model:
			model = color.Gray16Model
		}
	case ColorModeColor:
		switch model {
		case color.GrayModel:
			model = color.RGBAModel
		case color.Gray16Model:
			model = color.RGBA64Model
		}
	}

	switch filter.opt.Depth {
	case ColorDepth8:
		switch model {
		case color.Gray16Model:
			model = color.GrayModel
		case color.RGBA64Model:
			model = color.RGBAModel
		}
	case ColorDepth16:
		switch model {
		case color.GrayModel:
			model = color.Gray16Model
		case color.RGBAModel:
			model = color.RGBA64Model
		}
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

	switch filter.opt.OutputFormat {
	default:
		// FIXME
		//
		// For now, while not all image formats are implemented,
		// we fallback to JPEG, as it is safe default. Later it will be fixed.
		fallthrough

	case imgconv.MIMETypeJPEG:
		file.encoder, err = imgconv.NewJPEGWriter(file.output, wid, hei, model, 100)
	case imgconv.MIMETypePNG:
		file.encoder, err = imgconv.NewPNGWriter(file.output, wid, hei, model)
	}

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
	encoder  imgconv.Writer // Image encoder; nil if closed
	err      error          // Sticky error
}

// Format returns the MIME type of the image format used by
// the document file.
func (file *filterDocumentFile) Format() string {
	if file.filter.opt.OutputFormat != "" {
		return file.filter.opt.OutputFormat
	}
	return file.input.Format()
}

// Read reads the document file content as a sequence of bytes.
// It implements the [io.Reader] interface.
func (file *filterDocumentFile) Read(buf []byte) (int, error) {
	// Run filtering pipeline until we have some output data
	for file.output.Len() == 0 && file.err == nil {
		// Don't let buffer to grow indefinitely
		file.output.Reset()

		// Read the next image Row
		_, file.err = file.pipeline.Read(file.row)
		if file.err != nil {
			if file.err == io.EOF {
				err := file.encoder.Close()
				file.encoder = nil

				if err != nil {
					file.err = err
				}
			}
			break
		}

		file.err = file.encoder.Write(file.row)
	}

	// Return buffered data
	if file.output.Len() != 0 {
		return file.output.Read(buf)
	}

	return 0, file.err
}

// close closes the filterDocumentFile.
func (file *filterDocumentFile) close() {
	file.pipeline.Close()
	if file.encoder != nil {
		file.encoder.Close()
	}
}
