// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image transformation filter for image.Image->draw.Image transformations.

package imgconv

import (
	"image"
	"image/color"
	"image/draw"
	"sync"
	"sync/atomic"
)

// NewTransformer creates a new image filter on a top of existent [Decoder].
//
// This filter performs image transformation by applying external function,
// that takes [image.Image] as input and uses [draw.Image] as output.
//
// The image is read from the underlying Decoder, and transformer
// implements the Decoder interface, from where the transformed image
// can be read.
//
// The transformation function will be executed in the separate
// goroutine.
//
// The source [image.Image] and the target [draw.Image]
// have certain limitations. See [SourceImageAdapter] and
// [TargetImageAdapter] for details.
//
// These limitations allow image transformation to be performed
// in the row-by-row streaming mode, without need to keep the
// entire image in memory.
//
// Despite these limitations, the transformer filter can be used with
// many image transformation algorithms, like image scaling.
//
// The dstWid, dstHei and dstModel arguments define parameters
// of the destination image. Source image parameters are defined
// by the underlying Decoder.
func NewTransformer(input Decoder,
	dstWid, dstHei int, dstModel color.Model,
	transform func(target draw.Image, source image.Image)) Decoder {

	source := NewSourceImageAdapter(input)
	target := NewTargetImageAdapter(dstWid, dstHei, dstModel)

	tr := &transformer{Decoder: target, input: input}

	tr.done.Add(1)
	go func() {
		transform(target, source)
		if err := source.Error(); err != nil {
			tr.err.Store(err)
		}
		target.Flush()
		tr.done.Done()
	}()

	return tr
}

// transformer is the Decoder returned by the NewTransformer
type transformer struct {
	Decoder                // Implements transformer's Decoder interface
	input   Decoder        // Underlying image source
	err     atomic.Value   // Read error from transformer.input
	done    sync.WaitGroup // Wait for goroutine completion
}

// Close closes the transformer
func (tr *transformer) Close() {
	tr.Decoder.Close()
	tr.done.Wait()
	tr.input.Close()
}

// Read returns the next image [Row].
func (tr *transformer) Read(row Row) (int, error) {
	n, err := tr.Decoder.Read(row)
	if err2 := tr.err.Load(); err2 != nil {
		err = err2.(error)
		n = 0
	}
	return n, err
}
