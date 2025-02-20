// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner capabilities

package escl

import (
	"fmt"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/xmldoc"
)

// ScannerCapabilities defines the scanner capabilities.
//
// eSCL Technical Specification, 8.1.4.
type ScannerCapabilities struct {
	// General options
	Version         Version                 // eSCL protocol version
	MakeAndModel    optional.Val[string]    // Device make and model
	SerialNumber    optional.Val[string]    // Device-unique serial number
	Manufacturer    optional.Val[string]    // Device manufacturer
	UUID            optional.Val[uuid.UUID] // Device UUID
	AdminURI        optional.Val[string]    // Configuration mage URL
	IconURI         optional.Val[string]    // Device icon URL
	SettingProfiles []SettingProfile        // Common settings profs

	// Inputs capabilities
	Platen optional.Val[Platen] // Platen capabilities
	Camera optional.Val[Camera] // Camera capabilities
	ADF    optional.Val[ADF]    // ADF capabilities

	// Image transform ranges
	BrightnessSupport        optional.Val[Range] // Brightness
	CompressionFactorSupport optional.Val[Range] // Lower num, better image
	ContrastSupport          optional.Val[Range] // Contrast
	GammaSupport             optional.Val[Range] // Gamma (y = x^(1/g))
	HighlightSupport         optional.Val[Range] // Image Highlight
	NoiseRemovalSupport      optional.Val[Range] // Noise removal level
	ShadowSupport            optional.Val[Range] // The lower, the darger
	SharpenSupport           optional.Val[Range] // Image sharpen
	ThresholdSupport         optional.Val[Range] // For BlackAndWhite1

	// Automatic detection and removal of the blank pages
	BlankPageDetection           optional.Val[bool] // Detection supported
	BlankPageDetectionAndRemoval optional.Val[bool] // Auto-remove supported
}

// DecodeScannerCapabilities decodes [ScannerCapabilities] from the
// XML tree.
func DecodeScannerCapabilities(root xmldoc.Element) (
	scancaps ScannerCapabilities, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	ver := xmldoc.Lookup{Name: NsPWG + ":Version", Required: true}
	mdl := xmldoc.Lookup{Name: NsPWG + ":MakeAndModel"}
	ser := xmldoc.Lookup{Name: NsPWG + ":SerialNumber"}
	mfg := xmldoc.Lookup{Name: NsPWG + ":Manufacturer"}
	uu := xmldoc.Lookup{Name: NsScan + ":UUID"}
	admin := xmldoc.Lookup{Name: NsScan + ":AdminURI"}
	icon := xmldoc.Lookup{Name: NsScan + ":IconURI"}
	profiles := xmldoc.Lookup{Name: NsScan + ":SettingProfiles"}
	platen := xmldoc.Lookup{Name: NsScan + ":Platen"}
	camera := xmldoc.Lookup{Name: NsScan + ":Camera"}
	adf := xmldoc.Lookup{Name: NsScan + ":Adf"}
	brightness := xmldoc.Lookup{Name: NsScan + ":BrightnessSupport"}
	compression := xmldoc.Lookup{Name: NsScan + ":CompressionFactorSupport"}
	contrast := xmldoc.Lookup{Name: NsScan + ":ContrastSupport"}
	gamma := xmldoc.Lookup{Name: NsScan + ":GammaSupport"}
	highlight := xmldoc.Lookup{Name: NsScan + ":HighlightSupport"}
	noiseRemoval := xmldoc.Lookup{Name: NsScan + ":NoiseRemovalSupport"}
	shadow := xmldoc.Lookup{Name: NsScan + ":ShadowSupport"}
	sharpen := xmldoc.Lookup{Name: NsScan + ":SharpenSupport"}
	threshold := xmldoc.Lookup{Name: NsScan + ":ThresholdSupport"}
	blankDetection := xmldoc.Lookup{Name: NsScan + ":BlankPageDetection"}
	blankRemoval := xmldoc.Lookup{
		Name: NsScan + ":BlankPageDetectionAndRemoval"}

	missed := root.Lookup(&ver, &mdl, &ser, &mfg, &uu, &admin, &icon,
		&profiles, &platen, &camera, &adf,
		&brightness, &compression, &contrast, &gamma, &highlight,
		&noiseRemoval, &shadow, &sharpen, &threshold,
		&blankDetection, &blankRemoval)

	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	scancaps.Version, err = decodeVersion(ver.Elem)
	if err != nil {
		return
	}

	if mdl.Found {
		scancaps.MakeAndModel = optional.New(mdl.Elem.Text)
	}

	if ser.Found {
		scancaps.SerialNumber = optional.New(ser.Elem.Text)
	}

	if mfg.Found {
		scancaps.Manufacturer = optional.New(mfg.Elem.Text)
	}

	if uu.Found {
		var u uuid.UUID
		u, err = uuid.Parse(uu.Elem.Text)
		if err != nil {
			err = fmt.Errorf("invalid UUID: %q", uu.Elem.Text)
			return
		}

		scancaps.UUID = optional.New(u)
	}

	if profiles.Found {
		for _, elem := range profiles.Elem.Children {
			if elem.Name == NsScan+":SettingProfile" {
				var prof SettingProfile
				prof, err = decodeSettingProfile(elem)

				if err != nil {
					err = xmldoc.XMLErrWrap(
						profiles.Elem, err)
					return
				}

				scancaps.SettingProfiles = append(
					scancaps.SettingProfiles, prof)
			}
		}
	}

	if platen.Found {
		var v Platen
		v, err = decodePlaten(platen.Elem)
		if err != nil {
			return
		}

		scancaps.Platen = optional.New(v)
	}

	if camera.Found {
		var v Camera
		v, err = decodeCamera(camera.Elem)
		if err != nil {
			return
		}

		scancaps.Camera = optional.New(v)
	}

	if adf.Found {
		var v ADF
		v, err = decodeADF(adf.Elem)
		if err != nil {
			return
		}

		scancaps.ADF = optional.New(v)
	}

	if brightness.Found {
		var r Range
		r, err = decodeRange(brightness.Elem)
		if err != nil {
			return
		}

		scancaps.BrightnessSupport = optional.New(r)
	}

	if compression.Found {
		var r Range
		r, err = decodeRange(compression.Elem)
		if err != nil {
			return
		}

		scancaps.CompressionFactorSupport = optional.New(r)
	}

	if contrast.Found {
		var r Range
		r, err = decodeRange(contrast.Elem)
		if err != nil {
			return
		}

		scancaps.ContrastSupport = optional.New(r)
	}

	if gamma.Found {
		var r Range
		r, err = decodeRange(gamma.Elem)
		if err != nil {
			return
		}

		scancaps.GammaSupport = optional.New(r)
	}

	if highlight.Found {
		var r Range
		r, err = decodeRange(highlight.Elem)
		if err != nil {
			return
		}

		scancaps.HighlightSupport = optional.New(r)
	}

	if noiseRemoval.Found {
		var r Range
		r, err = decodeRange(noiseRemoval.Elem)
		if err != nil {
			return
		}

		scancaps.NoiseRemovalSupport = optional.New(r)
	}

	if shadow.Found {
		var r Range
		r, err = decodeRange(shadow.Elem)
		if err != nil {
			return
		}

		scancaps.ShadowSupport = optional.New(r)
	}

	if sharpen.Found {
		var r Range
		r, err = decodeRange(sharpen.Elem)
		if err != nil {
			return
		}

		scancaps.SharpenSupport = optional.New(r)
	}

	if threshold.Found {
		var r Range
		r, err = decodeRange(threshold.Elem)
		if err != nil {
			return
		}

		scancaps.ThresholdSupport = optional.New(r)
	}

	if blankDetection.Found {
		var flg bool
		flg, err = decodeBool(blankDetection.Elem)
		if err != nil {
			return
		}

		scancaps.BlankPageDetection = optional.New(flg)
	}

	if blankRemoval.Found {
		var flg bool
		flg, err = decodeBool(blankRemoval.Elem)
		if err != nil {
			return
		}

		scancaps.BlankPageDetectionAndRemoval = optional.New(flg)
	}

	return
}
