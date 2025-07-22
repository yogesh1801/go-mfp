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
	AutoExposureSupported             AutoExposureSupported
	BrightnessSupported               BrightnessSupported
	CompressionQualityFactorSupported CompressionQualityFactorSupported
	ContentTypesSupported             ContentTypesSupported
	ContrastSupported                 ContrastSupported
	DocumentSizeAutoDetectSupported   DocumentSizeAutoDetectSupported
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

	autoExposureSupported := xmldoc.Lookup{Name: NsWSCN +
		":AutoExposureSupported", Required: true}
	brightnessSupported := xmldoc.Lookup{Name: NsWSCN +
		":BrightnessSupported", Required: true}
	compressionQualityFactorSupported := xmldoc.Lookup{Name: NsWSCN + ":CompressionQualityFactorSupported", Required: true}
	contentTypesSupported := xmldoc.Lookup{Name: NsWSCN +
		":ContentTypesSupported", Required: true}
	contrastSupported := xmldoc.Lookup{Name: NsWSCN +
		":ContrastSupported", Required: true}
	documentSizeAutoDetectSupported := xmldoc.Lookup{Name: NsWSCN + ":DocumentSizeAutoDetectSupported", Required: true}
	formatsSupported := xmldoc.Lookup{Name: NsWSCN +
		":FormatsSupported", Required: true}
	rotationsSupported := xmldoc.Lookup{Name: NsWSCN +
		":RotationsSupported", Required: true}
	scalingRangeSupported := xmldoc.Lookup{Name: NsWSCN +
		":ScalingRangeSupported", Required: true}

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

	ds.AutoExposureSupported, err = decodeAutoExposureSupported(autoExposureSupported.Elem)
	if err != nil {
		return
	}
	ds.BrightnessSupported, err = decodeBrightnessSupported(brightnessSupported.Elem)
	if err != nil {
		return
	}
	ds.CompressionQualityFactorSupported, err = decodeCompressionQualityFactorSupported(compressionQualityFactorSupported.Elem)
	if err != nil {
		return
	}
	ds.ContentTypesSupported, err = decodeContentTypesSupported(contentTypesSupported.Elem)
	if err != nil {
		return
	}
	ds.ContrastSupported, err = decodeContrastSupported(contrastSupported.Elem)
	if err != nil {
		return
	}
	ds.DocumentSizeAutoDetectSupported, err = decodeDocumentSizeAutoDetectSupported(documentSizeAutoDetectSupported.Elem)
	if err != nil {
		return
	}
	ds.FormatsSupported, err = decodeFormatsSupported(formatsSupported.Elem)
	if err != nil {
		return
	}
	ds.RotationsSupported, err = decodeRotationsSupported(rotationsSupported.Elem)
	if err != nil {
		return
	}
	ds.ScalingRangeSupported, err = decodeScalingRangeSupported(scalingRangeSupported.Elem)
	if err != nil {
		return
	}

	if err = ds.Validate(); err != nil {
		return
	}
	return
}

// Validate checks that all child elements are valid.
func (ds DeviceSettings) Validate() error {
	if _, err := ds.AutoExposureSupported.Bool(); err != nil {
		return fmt.Errorf("AutoExposureSupported: %w", err)
	}
	if _, err := ds.BrightnessSupported.Bool(); err != nil {
		return fmt.Errorf("BrightnessSupported: %w", err)
	}
	if err := ds.CompressionQualityFactorSupported.Validate(); err != nil {
		return fmt.Errorf("CompressionQualityFactorSupported: %w", err)
	}
	if len(ds.ContentTypesSupported.Values) == 0 {
		return fmt.Errorf("ContentTypesSupported requires at least one value")
	}
	if _, err := ds.ContrastSupported.Bool(); err != nil {
		return fmt.Errorf("ContrastSupported: %w", err)
	}
	if _, err := ds.DocumentSizeAutoDetectSupported.Bool(); err != nil {
		return fmt.Errorf("DocumentSizeAutoDetectSupported: %w", err)
	}
	if len(ds.FormatsSupported.Values) == 0 {
		return fmt.Errorf("FormatsSupported requires at least one value")
	}
	if len(ds.RotationsSupported.Values) == 0 {
		return fmt.Errorf("RotationsSupported requires at least one value")
	}
	if err := ds.ScalingRangeSupported.Validate(); err != nil {
		return fmt.Errorf("ScalingRangeSupported: %w", err)
	}
	return nil
}
