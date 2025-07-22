// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// device settings

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// DeviceSettings represents the <wscn:DeviceSettings> element,
// describing the basic capabilities of the scan device.
type DeviceSettings struct {
	AutoExposureSupported             BooleanElement
	BrightnessSupported               BooleanElement
	CompressionQualityFactorSupported CompressionQualityFactorSupported
	ContentTypesSupported             ContentTypesSupported
	ContrastSupported                 BooleanElement
	DocumentSizeAutoDetectSupported   BooleanElement
	FormatsSupported                  FormatsSupported
	RotationsSupported                RotationsSupported
	ScalingRangeSupported             ScalingRangeSupported
}

// toXML generates XML tree for the [DeviceSettings].
func (ds DeviceSettings) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			ds.AutoExposureSupported.toXML(NsWSCN +
				":AutoExposureSupported"),
			ds.BrightnessSupported.toXML(NsWSCN +
				":BrightnessSupported"),
			ds.CompressionQualityFactorSupported.toXML(NsWSCN +
				":CompressionQualityFactorSupported"),
			ds.ContentTypesSupported.toXML(NsWSCN +
				":ContentTypesSupported"),
			ds.ContrastSupported.toXML(NsWSCN +
				":ContrastSupported"),
			ds.DocumentSizeAutoDetectSupported.toXML(NsWSCN +
				":DocumentSizeAutoDetectSupported"),
			ds.FormatsSupported.toXML(NsWSCN +
				":FormatsSupported"),
			ds.RotationsSupported.toXML(NsWSCN +
				":RotationsSupported"),
			ds.ScalingRangeSupported.toXML(NsWSCN +
				":ScalingRangeSupported"),
		},
	}
}

// decodeDeviceSettings decodes [DeviceSettings] from the XML tree.
func decodeDeviceSettings(root xmldoc.Element) (ds DeviceSettings, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup all required XML elements
	autoExposureSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":AutoExposureSupported",
		Required: true,
	}
	brightnessSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":BrightnessSupported",
		Required: true,
	}
	compressionQualityFactorSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":CompressionQualityFactorSupported",
		Required: true,
	}
	contentTypesSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":ContentTypesSupported",
		Required: true,
	}
	contrastSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":ContrastSupported",
		Required: true,
	}
	documentSizeAutoDetectSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":DocumentSizeAutoDetectSupported",
		Required: true,
	}
	formatsSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":FormatsSupported",
		Required: true,
	}
	rotationsSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":RotationsSupported",
		Required: true,
	}
	scalingRangeSupported := xmldoc.Lookup{
		Name:     NsWSCN + ":ScalingRangeSupported",
		Required: true,
	}

	// Perform all lookups at once
	missed := root.Lookup(
		&autoExposureSupported,
		&brightnessSupported,
		&compressionQualityFactorSupported,
		&contentTypesSupported,
		&contrastSupported,
		&documentSizeAutoDetectSupported,
		&formatsSupported,
		&rotationsSupported,
		&scalingRangeSupported,
	)

	if missed != nil {
		return ds, xmldoc.XMLErrMissed(missed.Name)
	}

	// Decode all required elements
	if ds.AutoExposureSupported, err = decodeBooleanElement(
		autoExposureSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("AutoExposureSupported: %w", err)
	}

	if ds.BrightnessSupported, err = decodeBooleanElement(
		brightnessSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("BrightnessSupported: %w", err)
	}

	if ds.CompressionQualityFactorSupported, err = decodeCompressionQualityFactorSupported(
		compressionQualityFactorSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("CompressionQualityFactorSupported: %w", err)
	}

	if ds.ContentTypesSupported, err = decodeContentTypesSupported(
		contentTypesSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("ContentTypesSupported: %w", err)
	}

	if ds.ContrastSupported, err = decodeBooleanElement(
		contrastSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("ContrastSupported: %w", err)
	}

	if ds.DocumentSizeAutoDetectSupported, err = decodeBooleanElement(
		documentSizeAutoDetectSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("DocumentSizeAutoDetectSupported: %w", err)
	}

	if ds.FormatsSupported, err = decodeFormatsSupported(
		formatsSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("FormatsSupported: %w", err)
	}

	if ds.RotationsSupported, err = decodeRotationsSupported(
		rotationsSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("RotationsSupported: %w", err)
	}

	if ds.ScalingRangeSupported, err = decodeScalingRangeSupported(
		scalingRangeSupported.Elem,
	); err != nil {
		return ds, fmt.Errorf("ScalingRangeSupported: %w", err)
	}

	return ds, nil
}
