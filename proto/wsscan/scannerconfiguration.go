// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ScannerConfiguration: describes the scanner's configurable capabilities

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScannerConfiguration is a collection of elements that describes the
// scanner's configurable capabilities. ADF, Film, and Platen are optional
// and depend on the hardware available on the scanner.
type ScannerConfiguration struct {
	ADF            optional.Val[ADF]
	DeviceSettings DeviceSettings
	Film           optional.Val[Film]
	Platen         optional.Val[Platen]
}

// toXML creates an XML element for ScannerConfiguration.
func (sc ScannerConfiguration) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	if sc.ADF != nil {
		elm.Children = append(elm.Children,
			optional.Get(sc.ADF).toXML(NsWSCN+":ADF"))
	}

	elm.Children = append(elm.Children,
		sc.DeviceSettings.toXML(NsWSCN+":DeviceSettings"))

	if sc.Film != nil {
		elm.Children = append(elm.Children,
			optional.Get(sc.Film).toXML(NsWSCN+":Film"))
	}

	if sc.Platen != nil {
		elm.Children = append(elm.Children,
			optional.Get(sc.Platen).toXML(NsWSCN+":Platen"))
	}

	return elm
}

// decodeScannerConfiguration decodes a ScannerConfiguration from an XML element.
func decodeScannerConfiguration(root xmldoc.Element) (
	ScannerConfiguration, error) {
	var sc ScannerConfiguration

	adf := xmldoc.Lookup{
		Name:     NsWSCN + ":ADF",
		Required: false,
	}
	deviceSettings := xmldoc.Lookup{
		Name:     NsWSCN + ":DeviceSettings",
		Required: true,
	}
	film := xmldoc.Lookup{
		Name:     NsWSCN + ":Film",
		Required: false,
	}
	platen := xmldoc.Lookup{
		Name:     NsWSCN + ":Platen",
		Required: false,
	}

	missed := root.Lookup(&adf, &deviceSettings, &film, &platen)
	if missed != nil {
		return sc, xmldoc.XMLErrMissed(missed.Name)
	}

	if adf.Found {
		a, err := decodeADF(adf.Elem)
		if err != nil {
			return sc, fmt.Errorf("ADF: %w", err)
		}
		sc.ADF = optional.New(a)
	}

	ds, err := decodeDeviceSettings(deviceSettings.Elem)
	if err != nil {
		return sc, fmt.Errorf("DeviceSettings: %w", err)
	}
	sc.DeviceSettings = ds

	if film.Found {
		f, err := decodeFilm(film.Elem)
		if err != nil {
			return sc, fmt.Errorf("Film: %w", err)
		}
		sc.Film = optional.New(f)
	}

	if platen.Found {
		p, err := decodePlaten(platen.Elem)
		if err != nil {
			return sc, fmt.Errorf("Platen: %w", err)
		}
		sc.Platen = optional.New(p)
	}

	return sc, nil
}
