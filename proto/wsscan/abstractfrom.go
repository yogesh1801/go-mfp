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
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/optional"
)

// wsscanDPI is used for converting abstract.Dimension to WS-Scan
// dimensions. WS-Scan uses thousandths of an inch (1/1000"),
// so we treat dimensions as dots at 1000 DPI.
const wsscanDPI = 1000

// FromAbstractScannerDescription translates
// [abstract.ScannerCapabilities] into a [ScannerDescription].
func fromAbstractScannerDescription(
	caps *abstract.ScannerCapabilities) ScannerDescription {

	sd := ScannerDescription{
		ScannerName: TextWithLangList{
			{Text: caps.MakeAndModel},
		},
	}

	return sd
}

// FromAbstractScannerConfiguration translates
// [abstract.ScannerCapabilities] into a [ScannerConfiguration].
func fromAbstractScannerConfiguration(
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
// WS-Scan [Dimensions].
func fromAbstractOpticalResolution(
	inp *abstract.InputCapabilities) Dimensions {

	return Dimensions{Width: inp.MaxOpticalXResolution,
		Height: inp.MaxOpticalYResolution}
}

// fromAbstractResolutions extracts unique width and height resolution
// values from all settings profiles into WS-Scan [Resolutions].
func fromAbstractResolutions(
	profiles []abstract.SettingsProfile) Resolutions {

	widthSet := generic.NewSet[int]()
	heightSet := generic.NewSet[int]()

	for _, prof := range profiles {
		for _, res := range prof.Resolutions {
			widthSet.Add(res.XResolution)
			heightSet.Add(res.YResolution)
		}
	}

	var r Resolutions
	widthSet.ForEach(func(w int) {
		r.Widths = append(r.Widths, w)
	})
	heightSet.ForEach(func(h int) {
		r.Heights = append(r.Heights, h)
	})

	sort.Ints(r.Widths)
	sort.Ints(r.Heights)

	return r
}

// fromAbstractScannerRequest converts an [abstract.ScannerRequest]
// into a [ScanTicket]. It is the inverse of [ScanTicket.ToAbstract].
func fromAbstractScannerRequest(req *abstract.ScannerRequest) ScanTicket {
	dp := DocumentParameters{}

	// Input + ADFMode → InputSource
	switch req.Input {
	case abstract.InputPlaten:
		dp.InputSource = optional.New(
			ValWithOptions[InputSourceValue]{Val: InputSourcePlaten})
	case abstract.InputADF:
		switch req.ADFMode {
		case abstract.ADFModeDuplex:
			dp.InputSource = optional.New(
				ValWithOptions[InputSourceValue]{Val: InputSourceADFDuplex})
		default:
			dp.InputSource = optional.New(
				ValWithOptions[InputSourceValue]{Val: InputSourceADF})
		}
	}

	// Build MediaFront (ColorProcessing, Resolution, ScanRegion)
	front := MediaSide{}

	// ColorMode + ColorDepth → ColorProcessing
	ce := abstractColorEntryFrom(req.ColorMode, req.ColorDepth)
	if ce != UnknownColorEntry {
		front.ColorProcessing = optional.New(ValWithOptions[ColorEntry]{Val: ce})
	}

	// Resolution
	if !req.Resolution.IsZero() {
		front.Resolution = optional.New(Resolution{
			Width:  ValWithOptions[int]{Val: req.Resolution.XResolution},
			Height: ValWithOptions[int]{Val: req.Resolution.YResolution},
		})
	}

	// Region → ScanRegion
	if !req.Region.IsZero() {
		sr := ScanRegion{
			ScanRegionWidth: ValWithOptions[int]{
				Val: req.Region.Width.Dots(wsscanDPI)},
			ScanRegionHeight: ValWithOptions[int]{
				Val: req.Region.Height.Dots(wsscanDPI)},
		}
		if req.Region.XOffset != 0 {
			sr.ScanRegionXOffset = optional.New(
				ValWithOptions[int]{Val: req.Region.XOffset.Dots(wsscanDPI)})
		}
		if req.Region.YOffset != 0 {
			sr.ScanRegionYOffset = optional.New(
				ValWithOptions[int]{Val: req.Region.YOffset.Dots(wsscanDPI)})
		}
		front.ScanRegion = optional.New(sr)
	}

	dp.MediaSides = optional.New(MediaSides{MediaFront: front})

	// DocumentFormat (MIME) → Format
	if req.DocumentFormat != "" {
		fv := mimeToFormatValue(req.DocumentFormat)
		if fv != UnknownFormatValue {
			dp.Format = optional.New(ValWithOptions[FormatValue]{Val: fv})
		}
	}

	// Compression
	if req.Compression != nil {
		compression := optional.Get(req.Compression)
		dp.CompressionQualityFactor = optional.New(
			ValWithOptions[int]{Val: compression})
	}

	// Intent → ContentType
	switch req.Intent {
	case abstract.IntentDocument:
		dp.ContentType = optional.New(ValWithOptions[ContentTypeValue]{Val: Text})
	case abstract.IntentPhoto:
		dp.ContentType = optional.New(ValWithOptions[ContentTypeValue]{Val: Photo})
	case abstract.IntentTextAndGraphic:
		dp.ContentType = optional.New(ValWithOptions[ContentTypeValue]{Val: Mixed})
	}

	// Brightness, Contrast, Sharpen → Exposure.ExposureSettings
	if req.Brightness != nil || req.Contrast != nil || req.Sharpen != nil {
		es := ExposureSettings{}
		if req.Brightness != nil {
			es.Brightness = optional.New(ValWithOptions[int]{
				Val: optional.Get(req.Brightness)})
		}
		if req.Contrast != nil {
			es.Contrast = optional.New(ValWithOptions[int]{
				Val: optional.Get(req.Contrast)})
		}
		if req.Sharpen != nil {
			es.Sharpness = optional.New(ValWithOptions[int]{
				Val: optional.Get(req.Sharpen)})
		}
		dp.Exposure = optional.New(Exposure{
			ExposureSettings: optional.New(es),
		})
	}

	return ScanTicket{
		DocumentParameters: optional.New(dp),
	}
}

// abstractColorEntryFrom converts [abstract.ColorMode] and
// [abstract.ColorDepth] into a [ColorEntry].
func abstractColorEntryFrom(
	mode abstract.ColorMode,
	depth abstract.ColorDepth) ColorEntry {

	switch mode {
	case abstract.ColorModeBinary:
		return BlackAndWhite1
	case abstract.ColorModeMono:
		if depth == abstract.ColorDepth16 {
			return Grayscale16
		}
		return Grayscale8
	case abstract.ColorModeColor:
		if depth == abstract.ColorDepth16 {
			return RGB48
		}
		return RGB24
	}
	return UnknownColorEntry
}

// fromAbstractColorEntries extracts unique [ColorEntry] values from
// the color modes and depths defined in settings profiles.
func fromAbstractColorEntries(
	profiles []abstract.SettingsProfile) []ColorEntry {

	seen := generic.NewSet[ColorEntry]()
	var out []ColorEntry

	add := func(e ColorEntry) {
		if seen.TestAndAdd(e) {
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
