// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// media-col-database and friends

package ipp

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/goipp"
)

// MediaColDatabase represents "media-col-database" attribute.
type MediaColDatabase struct {
	MediaColDatabase []MediaColEx `ipp:"media-col-database"`
}

// MediaCol is the "media-col", "media-col-xxx" collection entry.
// It is used in many places.
//
// PWG5100.7
type MediaCol struct {
	// ----- PWG5100.3 -----
	MediaBackCoating  optional.Val[KwMediaBackCoating] `ipp:"media-back-coating"`
	MediaColor        optional.Val[KwColor]            `ipp:"media-color"`
	MediaFrontCoating optional.Val[KwMediaBackCoating] `ipp:"media-front-coating"`
	MediaHoleCount    optional.Val[int]                `ipp:"media-hole-count"`
	MediaInfo         optional.Val[string]             `ipp:"media-info"`
	MediaKey          optional.Val[KwMedia]            `ipp:"media-key"`
	MediaOrderCount   optional.Val[int]                `ipp:"media-order-count"`
	MediaPrePrinted   optional.Val[string]             `ipp:"media-pre-printed"`
	MediaRecycled     optional.Val[string]             `ipp:"media-recycled"`
	MediaSize         optional.Val[MediaSize]          `ipp:"media-size"`
	MediaType         optional.Val[string]             `ipp:"media-type"`
	MediaWeightMetric optional.Val[int]                `ipp:"media-weight-metric"`

	// ----- PWG5100.7 -----
	MediaBottomMargin optional.Val[int]    `ipp:"media-bottom-margin"`
	MediaGrain        optional.Val[string] `ipp:"media-grain"`
	MediaLeftMargin   optional.Val[int]    `ipp:"media-left-margin"`
	MediaRightMargin  optional.Val[int]    `ipp:"media-right-margin"`
	MediaSizeName     optional.Val[string] `ipp:"media-size-name"`
	MediaSource       optional.Val[string] `ipp:"media-source"`
	MediaThickness    optional.Val[int]    `ipp:"media-thickness"`
	MediaTooth        optional.Val[string] `ipp:"media-tooth"`
	MediaTopMargin    optional.Val[int]    `ipp:"media-top-margin"`
}

// MediaColEx is the [MediaCol] with some additional data.
// It is only used for "media-col-database" and "media-col-ready"
// Printer Description attributes.
type MediaColEx struct {
	MediaCol
	MediaSourceProperties optional.Val[MediaSourceProperties] `ipp:"media-source-properties"`
}

// MediaSize represents media size parameter, defined by a pair of
// integer dimensions.
type MediaSize struct {
	XDimension int `ipp:"x-dimension"`
	YDimension int `ipp:"y-dimension"`
}

// MediaSizeRange represents media size parameter, defined by a pair
// if integer or range of integer dimensions.
type MediaSizeRange struct {
	XDimension goipp.IntegerOrRange `ipp:"x-dimension"`
	YDimension goipp.IntegerOrRange `ipp:"y-dimension"`
}

// MediaSourceProperties represents "media-source-properties"
// collectiobn in MediaCol
type MediaSourceProperties struct {
	MediaSourceFeedDirection   string `ipp:"media-source-feed-direction"`
	MediaSourceFeedOrientation int    `ipp:"media-source-feed-orientation"`
}

// MediaOverprint represents "media-overprint" collection entry
// in JobAttributes
type MediaOverprint struct {
	MediaOverprintDistance int    `ipp:"media-overprint-distance"`
	MediaOverprintMethod   string `ipp:"media-overprint-method"`
}
