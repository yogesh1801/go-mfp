// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// List of keywords

package modeling

import "strings"

// keywordMap contains str.ToLower(GoName)->ProtocolName mappings
// for keywords.
var keywordMap = map[string]string{}

// init populates the keywordMap
func init() {
	for _, kw := range keywordList {
		keywordMap[strings.ToLower(kw)] = kw
	}
}

// keywordNormalize returns a keyword with the normalized spelling.
func keywordNormalize(kw string) string {
	if norm, ok := keywordMap[strings.ToLower(kw)]; ok {
		return norm
	}
	return kw
}

// keywordList defines proper spelling of the keywords used in the MFP
// models.
//
// These names are based on a Go field names in the protocol structures,
// and in the most cases 1:1 corresponds to the protocol names.
//
// But sometimes, at the Go side, golint dictates us uppercase spelling.
// For example, Uuid named "UUID" at the Go side, but in the model/protocols
// it is named "Uuid'. So we need to implement the proper mapping
var keywordList = []string{
	"ActualHeight",
	"ActualWidth",
	"Adf",
	"AdfDuplexInputCaps",
	"AdfOptions",
	"AdfSimplexInputCaps",
	"AdfState",
	"AdminUri",
	"Age",
	"BaseURL",
	"BinaryRendering",
	"BinaryRenderings",
	"BlankPageDetected",
	"BlankPageDetection",
	"BlankPageDetectionAndRemoval",
	"Brightness",
	"BrightnessSupport",
	"BytesPerLine",
	"Camera",
	"CameraInputCaps",
	"CCDChannel",
	"CCDChannels",
	"ColorMode",
	"ColorModes",
	"ColorSpace",
	"ColorSpaces",
	"CompressionFactor",
	"CompressionFactorSupport",
	"ContentRegionUnits",
	"ContentType",
	"ContentTypes",
	"Contrast",
	"ContrastSupport",
	"DiscreteResolutions",
	"DocumentFormat",
	"DocumentFormatExt",
	"DocumentFormats",
	"DocumentFormatsExt",
	"Duplex",
	"EdgeAutoDetection",
	"FeedDirection",
	"FeedDirections",
	"FeederCapacity",
	"Gamma",
	"GammaSupport",
	"Height",
	"Highlight",
	"HighlightSupport",
	"IconUri",
	"ImageHeight",
	"ImagesCompleted",
	"ImagesToTransfer",
	"ImageWidth",
	"InputSource",
	"Intent",
	"Jobs",
	"JobState",
	"JobStateReasons",
	"JobUri",
	"JobUuid",
	"Justification",
	"MakeAndModel",
	"Manufacturer",
	"Max",
	"MaxHeight",
	"MaxOpticalXResolution",
	"MaxOpticalYResolution",
	"MaxPhysicalHeight",
	"MaxPhysicalWidth",
	"MaxScanRegions",
	"MaxWidth",
	"MaxXOffset",
	"MaxYOffset",
	"Min",
	"MinHeight",
	"MinWidth",
	"NoiseRemoval",
	"NoiseRemovalSupport",
	"Normal",
	"Platen",
	"PlatenInputCaps",
	"ResolutionRange",
	"RiskyBottomMargins",
	"RiskyLeftMargins",
	"RiskyRightMargins",
	"RiskyTopMargins",
	"Scanner",
	"ScanRegions",
	"ScanSettings",
	"SerialNumber",
	"SettingProfiles",
	"Shadow",
	"ShadowSupport",
	"Sharpen",
	"SharpenSupport",
	"State",
	"Step",
	"SupportedIntents",
	"SupportedResolutions",
	"Threshold",
	"ThresholdSupport",
	"TransferRetryCount",
	"Uuid",
	"Version",
	"Width",
	"XImagePosition",
	"XOffset",
	"XResolution",
	"XResolutionRange",
	"YImagePosition",
	"YOffset",
	"YResolution",
	"YResolutionRange",
}
