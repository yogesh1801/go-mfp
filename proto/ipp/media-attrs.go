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
	MediaColDatabase []MediaCol `ipp:"media-col-database"`
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
	MediaHoleCount    optional.Val[int]                `ipp:"media-hole-count,(0:MAX)"`
	MediaInfo         optional.Val[string]             `ipp:"media-info,text"`
	MediaKey          optional.Val[KwMedia]            `ipp:"media-key"`
	MediaOrderCount   optional.Val[int]                `ipp:"media-order-count,(1:MAX)"`
	MediaPrePrinted   optional.Val[string]             `ipp:"media-pre-printed,keyword"`
	MediaRecycled     optional.Val[string]             `ipp:"media-recycled,keyword"`
	MediaSize         optional.Val[MediaSize]          `ipp:"media-size"`
	MediaType         optional.Val[string]             `ipp:"media-type,keyword"`
	MediaWeightMetric optional.Val[int]                `ipp:"media-weight-metric,(0:MAX)"`

	// ----- PWG5100.7 -----
	MediaBottomMargin     optional.Val[int]                   `ipp:"media-bottom-margin,(0:MAX)"`
	MediaGrain            optional.Val[string]                `ipp:"media-grain,keyword"`
	MediaLeftMargin       optional.Val[int]                   `ipp:"media-left-margin,(0:MAX)"`
	MediaRightMargin      optional.Val[int]                   `ipp:"media-right-margin,(0:MAX)"`
	MediaSizeName         optional.Val[string]                `ipp:"media-size-name,keyword"`
	MediaSourceProperties optional.Val[MediaSourceProperties] `ipp:"media-source-properties"`
	MediaSource           optional.Val[string]                `ipp:"media-source,keyword"`
	MediaThickness        optional.Val[int]                   `ipp:"media-thickness,(1:MAX)"`
	MediaTooth            optional.Val[string]                `ipp:"media-tooth,keyword"`
	MediaTopMargin        optional.Val[int]                   `ipp:"media-top-margin,(0:MAX)"`
}

// MediaSize represents media size parameters (which may be either
// pair of integers or pair of ranges) and used in many places
type MediaSize struct {
	XDimension goipp.IntegerOrRange `ipp:"x-dimension,(1:MAX)"`
	YDimension goipp.IntegerOrRange `ipp:"y-dimension,(1:MAX)"`
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
	MediaOverprintDistance int    `ipp:"media-overprint-distance,(0:MAX)"`
	MediaOverprintMethod   string `ipp:"media-overprint-method,keyword"`
}
