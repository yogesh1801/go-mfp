// MFP - Multi-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// IPP Scanner keyword types

package ipp

// KwInputSource represents "input-source" keyword values.
//
// See PWG5100.15.
type KwInputSource string

const (
	// KwInputSourcePlaten means scan from the scanner glass or platen.
	KwInputSourcePlaten KwInputSource = "platen"

	// KwInputSourceADF means scan from the auto-document feeder.
	KwInputSourceADF KwInputSource = "adf"

	// KwInputSourceFilmReader means scan from a microfilm reader.
	KwInputSourceFilmReader KwInputSource = "film-reader"
)

// KwInputContentType represents "input-content-type" keyword values.
//
// See PWG5100.15, 7.1.1.6.
type KwInputContentType string

const (
	// KwInputContentTypeAuto means automatically determine the type of document.
	KwInputContentTypeAuto KwInputContentType = "auto"

	// KwInputContentTypeHalftone means the document contains halftoned images.
	KwInputContentTypeHalftone KwInputContentType = "halftone"

	// KwInputContentTypeLineArt means the document contains line art.
	KwInputContentTypeLineArt KwInputContentType = "line-art"

	// KwInputContentTypeMagazine means the document is a magazine.
	KwInputContentTypeMagazine KwInputContentType = "magazine"

	// KwInputContentTypePhoto means the document is a photograph.
	KwInputContentTypePhoto KwInputContentType = "photo"

	// KwInputContentTypeText means the document only contains text.
	KwInputContentTypeText KwInputContentType = "text"

	// KwInputContentTypeTextAndPhoto means the document contains a combination of text and photographs.
	KwInputContentTypeTextAndPhoto KwInputContentType = "text-and-photo"
)

// KwInputFilmScanMode represents "input-film-scan-mode" keyword values.
//
// See PWG5100.15, 7.1.1.8.
type KwInputFilmScanMode string

const (
	// KwInputFilmScanModeBlackAndWhiteNegativeFilm means the film is black-and-white negatives.
	KwInputFilmScanModeBlackAndWhiteNegativeFilm KwInputFilmScanMode = "black-and-white-negative-film"

	// KwInputFilmScanModeColorNegativeFilm means the film is color negatives.
	KwInputFilmScanModeColorNegativeFilm KwInputFilmScanMode = "color-negative-film"

	// KwInputFilmScanModeColorSlideFilm means the film is color slides (positives).
	KwInputFilmScanModeColorSlideFilm KwInputFilmScanMode = "color-slide-film"

	// KwInputFilmScanModeNotApplicable means the type of film is not applicable to the usage.
	KwInputFilmScanModeNotApplicable KwInputFilmScanMode = "not-applicable"
)

// KwInputColorMode represents "input-color-mode" keyword values.
//
// See PWG5100.15.
type KwInputColorMode string

const (
	// KwInputColorModeAuto means the scanner chooses the color mode.
	KwInputColorModeAuto KwInputColorMode = "auto"

	// KwInputColorModeBiLevel means black and white (1-bit) scanning.
	KwInputColorModeBiLevel KwInputColorMode = "bi-level"

	// KwInputColorModeColor means full color scanning.
	KwInputColorModeColor KwInputColorMode = "color"

	// KwInputColorModeMonochrome means grayscale scanning.
	KwInputColorModeMonochrome KwInputColorMode = "monochrome"
)
