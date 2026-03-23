// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Conversions from abstract.Scanner to WS-Scan data structures

package wsscan

import (
	"sort"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/util/optional"
)

// wsscanDPI is used for converting abstract.Dimension to WS-Scan
// dimensions. WS-Scan uses thousandths of an inch (1/1000"),
// so we treat dimensions as dots at 1000 DPI.
const wsscanDPI = 1000

// FromAbstractScannerConfiguration translates
// [abstract.ScannerCapabilities] into a [ScannerConfiguration].
func FromAbstractScannerConfiguration(
	caps *abstract.ScannerCapabilities) ScannerConfiguration {

	sc := ScannerConfiguration{
		DeviceSettings: fromAbstractDeviceSettings(caps),
	}

	if caps.Platen != nil {
		sc.Platen = optional.New(fromAbstractPlaten(caps.Platen))
	}

	if caps.ADFSimplex != nil || caps.ADFDuplex != nil {
		sc.ADF = optional.New(fromAbstractADF(caps))
	}

	return sc
}

// fromAbstractDeviceSettings builds [DeviceSettings] from abstract
// capabilities.
func fromAbstractDeviceSettings(
	caps *abstract.ScannerCapabilities) DeviceSettings {

	ds := DeviceSettings{
		AutoExposureSupported: BooleanElement("false"),
		BrightnessSupported: fromAbstractBoolSupported(
			caps.BrightnessRange),
		ContrastSupported: fromAbstractBoolSupported(
			caps.ContrastRange),
		DocumentSizeAutoDetectSupported: BooleanElement("false"),
		ContentTypesSupported:           []ContentTypeValue{Auto},
		RotationsSupported:              []RotationValue{Rotation0},
		ScalingRangeSupported: ScalingRangeSupported{
			ScalingWidth:  Range{MinValue: 100, MaxValue: 100},
			ScalingHeight: Range{MinValue: 100, MaxValue: 100},
		},
	}

	if !caps.CompressionRange.IsZero() {
		ds.CompressionQualityFactorSupported = Range{
			MinValue: caps.CompressionRange.Min,
			MaxValue: caps.CompressionRange.Max,
		}
	} else {
		ds.CompressionQualityFactorSupported = Range{
			MinValue: 0,
			MaxValue: 100,
		}
	}

	ds.FormatsSupported = fromAbstractDocumentFormats(caps.DocumentFormats)

	return ds
}

// fromAbstractBoolSupported returns BooleanElement "true" if the
// abstract range is non-zero, "false" otherwise.
func fromAbstractBoolSupported(r abstract.Range) BooleanElement {
	if r.IsZero() {
		return BooleanElement("false")
	}
	return BooleanElement("true")
}

// fromAbstractDocumentFormats maps MIME type strings to WS-Scan
// [FormatValue] entries.
func fromAbstractDocumentFormats(formats []string) []FormatValue {
	var out []FormatValue
	for _, f := range formats {
		v := mimeToFormatValue(f)
		if v != UnknownFormatValue {
			out = append(out, v)
		}
	}
	return out
}

// mimeToFormatValue maps a MIME type string to a WS-Scan [FormatValue].
func mimeToFormatValue(mime string) FormatValue {
	switch mime {
	case "image/bmp", "image/x-ms-bmp":
		return DIB
	case "image/jpeg":
		return JFIF
	case "image/jbig":
		return JBIG
	case "image/jp2":
		return JPEG2K
	case "application/pdf":
		return PDFA
	case "image/png":
		return PNG
	case "image/tiff":
		return TIFFSingleUncompressed
	case "application/vnd.ms-xpsdocument":
		return XPS
	}
	return UnknownFormatValue
}

// fromAbstractPlaten builds a [Platen] from abstract
// [InputCapabilities].
func fromAbstractPlaten(inp *abstract.InputCapabilities) Platen {
	return Platen{
		PlatenColor:             fromAbstractColorEntries(inp.Profiles),
		PlatenMinimumSize:       fromAbstractMinDimensions(inp),
		PlatenMaximumSize:       fromAbstractMaxDimensions(inp),
		PlatenOpticalResolution: fromAbstractOpticalResolution(inp),
		PlatenResolutions:       fromAbstractResolutions(inp.Profiles),
	}
}

// fromAbstractADF builds an [ADF] from abstract scanner capabilities.
func fromAbstractADF(caps *abstract.ScannerCapabilities) ADF {
	adf := ADF{
		ADFSupportsDuplex: BooleanElement("false"),
	}

	if caps.ADFSimplex != nil {
		adf.ADFFront = optional.New(
			fromAbstractADFSide(caps.ADFSimplex))
	}

	if caps.ADFDuplex != nil {
		adf.ADFSupportsDuplex = BooleanElement("true")
		adf.ADFBack = optional.New(
			fromAbstractADFSide(caps.ADFDuplex))
	}

	return adf
}

// fromAbstractADFSide builds an [ADFSide] from abstract
// [InputCapabilities].
func fromAbstractADFSide(inp *abstract.InputCapabilities) ADFSide {
	return ADFSide{
		ADFColor:             fromAbstractColorEntries(inp.Profiles),
		ADFMinimumSize:       fromAbstractMinDimensions(inp),
		ADFMaximumSize:       fromAbstractMaxDimensions(inp),
		ADFOpticalResolution: fromAbstractOpticalResolution(inp),
		ADFResolutions:       fromAbstractResolutions(inp.Profiles),
	}
}

// fromAbstractMinDimensions converts abstract minimum dimensions to
// WS-Scan [Dimensions] (thousandths of an inch).
func fromAbstractMinDimensions(inp *abstract.InputCapabilities) Dimensions {
	return Dimensions{
		Width:  inp.MinWidth.LowerBoundDots(wsscanDPI),
		Height: inp.MinHeight.LowerBoundDots(wsscanDPI),
	}
}

// fromAbstractMaxDimensions converts abstract maximum dimensions to
// WS-Scan [Dimensions] (thousandths of an inch).
func fromAbstractMaxDimensions(inp *abstract.InputCapabilities) Dimensions {
	return Dimensions{
		Width:  inp.MaxWidth.UpperBoundDots(wsscanDPI),
		Height: inp.MaxHeight.UpperBoundDots(wsscanDPI),
	}
}

// fromAbstractOpticalResolution extracts optical resolution as
// WS-Scan [Dimensions]. Defaults to 300 DPI if unset.
func fromAbstractOpticalResolution(
	inp *abstract.InputCapabilities) Dimensions {

	return Dimensions{Width: inp.MaxOpticalXResolution,
		Height: inp.MaxOpticalYResolution}
}

// fromAbstractResolutions extracts unique width and height resolution
// values from all settings profiles into WS-Scan [Resolutions].
func fromAbstractResolutions(
	profiles []abstract.SettingsProfile) Resolutions {

	widthSet := make(map[int]bool)
	heightSet := make(map[int]bool)

	for _, prof := range profiles {
		for _, res := range prof.Resolutions {
			widthSet[res.XResolution] = true
			heightSet[res.YResolution] = true
		}
	}

	var r Resolutions
	for w := range widthSet {
		r.Widths = append(r.Widths, w)
	}
	for h := range heightSet {
		r.Heights = append(r.Heights, h)
	}

	sort.Ints(r.Widths)
	sort.Ints(r.Heights)

	return r
}

// fromAbstractColorEntries extracts unique [ColorEntry] values from
// the color modes and depths defined in settings profiles.
func fromAbstractColorEntries(
	profiles []abstract.SettingsProfile) []ColorEntry {

	seen := make(map[ColorEntry]bool)
	var out []ColorEntry

	add := func(e ColorEntry) {
		if !seen[e] {
			seen[e] = true
			out = append(out, e)
		}
	}

	for _, prof := range profiles {
		if prof.ColorModes.Contains(abstract.ColorModeBinary) {
			add(BlackAndWhite1)
		}

		if prof.ColorModes.Contains(abstract.ColorModeMono) {
			if prof.Depths.Contains(abstract.ColorDepth8) {
				add(Grayscale8)
			}
			if prof.Depths.Contains(abstract.ColorDepth16) {
				add(Grayscale16)
			}
		}

		if prof.ColorModes.Contains(abstract.ColorModeColor) {
			if prof.Depths.Contains(abstract.ColorDepth8) {
				add(RGB24)
			}
			if prof.Depths.Contains(abstract.ColorDepth16) {
				add(RGB48)
			}
		}
	}

	return out
}
