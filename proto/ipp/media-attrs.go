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
	MediaHoleCount    optional.Val[int]                `ipp:"media-hole-count,integer(0:MAX)"`
	MediaInfo         optional.Val[string]             `ipp:"media-info,text(255)"`
	MediaKey          optional.Val[KwMedia]            `ipp:"media-key"`
	MediaOrderCount   optional.Val[int]                `ipp:"media-order-count,integer(1:MAX)"`
	MediaPrePrinted   optional.Val[string]             `ipp:"media-pre-printed,keyword|name"`
	MediaRecycled     optional.Val[string]             `ipp:"media-recycled,keyword|name"`
	MediaSize         optional.Val[MediaSize]          `ipp:"media-size"`
	MediaType         optional.Val[string]             `ipp:"media-type,keyword|name"`
	MediaWeightMetric optional.Val[int]                `ipp:"media-weight-metric,integer(0:MAX)"`

	// ----- PWG5100.7 -----
	MediaBottomMargin optional.Val[int]    `ipp:"media-bottom-margin,integer(0:MAX)"`
	MediaGrain        optional.Val[string] `ipp:"media-grain,keyword|name"`
	MediaLeftMargin   optional.Val[int]    `ipp:"media-left-margin,integer(0:MAX)"`
	MediaRightMargin  optional.Val[int]    `ipp:"media-right-margin,integer(0:MAX)"`
	MediaSizeName     optional.Val[string] `ipp:"media-size-name,keyword|name"`
	MediaSource       optional.Val[string] `ipp:"media-source,keyword|name"`
	MediaThickness    optional.Val[int]    `ipp:"media-thickness,integer(1:MAX)"`
	MediaTooth        optional.Val[string] `ipp:"media-tooth,keyword|name"`
	MediaTopMargin    optional.Val[int]    `ipp:"media-top-margin,integer(0:MAX)"`
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
	MediaSourceFeedDirection   string `ipp:"media-source-feed-direction,keyword"`
	MediaSourceFeedOrientation int    `ipp:"media-source-feed-orientation,enum"`
}

// MediaOverprint represents "media-overprint" collection entry
// in JobAttributes
type MediaOverprint struct {
	MediaOverprintDistance int    `ipp:"media-overprint-distance,integer(0:MAX)"`
	MediaOverprintMethod   string `ipp:"media-overprint-method,keyword"`
}
