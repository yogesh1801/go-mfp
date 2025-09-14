// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package abstract

import (
	"io"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// streamfilter applies [FilterOptions] to [io.ReadCloser].
type streamfilter struct {
	filter *Filter      // Underlying Filter
	output DocumentFile // Output image
}

// NewStreamFilter applies [FilterOptions] to io.ReadCloser.
//
// The [Resolution] and format information of the input stream
// must be provided using the 'res' and 'format' parameters.
func NewStreamFilter(in io.ReadCloser,
	res Resolution, format string,
	opt FilterOptions) io.ReadCloser {

	// Bypass filtering if options not set.
	if opt == (FilterOptions{}) {
		return in
	}

	doc := NewDocumentReader(in, res, format)
	filter := NewFilter(doc, opt)
	file, err := filter.Next()
	assert.NoError(err)

	return &streamfilter{filter, file}
}

// Read returns image data after filtering.
func (sf *streamfilter) Read(buf []byte) (int, error) {
	return sf.output.Read(buf)
}

// Close closes the stream filter
func (sf *streamfilter) Close() error {
	return sf.filter.Close()
}
