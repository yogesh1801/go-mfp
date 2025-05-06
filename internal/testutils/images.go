// MFP - Miulti-Function Printers and scanners toolkit
// Utility functions and data BLOBs for testing
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package testutils

import (
	// Import "embed" for its side effects
	_ "embed"
)

// Images contains samples of test images of various formats and sizes
var Images struct {
	BMP100x75  []byte
	JPEG100x75 []byte
	PDF100x75  []byte
	PNG100x75  []byte
	TIFF100x75 []byte
}

//go:embed "data/UEIT-100x75.bmp"
var imageBMP100x75 []byte

//go:embed "data/UEIT-100x75.jpeg"
var imageJPEG100x75 []byte

//go:embed "data/UEIT-100x75.pdf"
var imagePDF100x75 []byte

//go:embed "data/UEIT-100x75.png"
var imagePNG100x75 []byte

//go:embed "data/UEIT-100x75.tiff"
var imageTIFF100x75 []byte

func init() {
	Images.BMP100x75 = imageBMP100x75
	Images.JPEG100x75 = imageJPEG100x75
	Images.PDF100x75 = imagePDF100x75
	Images.PNG100x75 = imagePNG100x75
	Images.TIFF100x75 = imageTIFF100x75
}
