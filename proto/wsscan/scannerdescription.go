// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scanner description

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScannerDescription holds descriptive information about the scanner.
// The ScannerDescription element contains child elements that provide
// descriptive information about the scanner, including its name, location,
// and other details.
//
// XML Usage:
//
//	<wscn:ScannerDescription>
//	  child elements
//	</wscn:ScannerDescription>
//
// Attributes: None
//
// Text value: None
//
// Child elements:
//   - ScannerInfo (optional, repeatable): Administratively assigned
//     descriptive info, one per language
//   - ScannerLocation (optional, repeatable): Administratively assigned
//     location info, one per language
//   - ScannerName (required, repeatable): Administratively assigned
//     user-friendly name, one per language
type ScannerDescription struct {
	ScannerInfo     TextWithLangList
	ScannerLocation TextWithLangList
	ScannerName     TextWithLangList
}

// decodeScannerDescription decodes a [ScannerDescription] from an XML element.
// It extracts child elements ScannerName, ScannerInfo, and ScannerLocation
// from the XML element. ScannerName is required (at least one entry),
// while ScannerInfo and ScannerLocation are optional. Each element may
// appear multiple times, once per language variant.
func decodeScannerDescription(root xmldoc.Element) (
	sd ScannerDescription,
	err error,
) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	nameN := NsWSCN + ":ScannerName"
	infoN := NsWSCN + ":ScannerInfo"
	locationN := NsWSCN + ":ScannerLocation"

	for _, child := range root.Children {
		switch child.Name {
		case nameN:
			var t TextWithLangElement
			t, err = t.decodeTextWithLangElement(child)
			if err != nil {
				return sd, err
			}
			sd.ScannerName = append(sd.ScannerName, t)

		case infoN:
			var t TextWithLangElement
			t, err = t.decodeTextWithLangElement(child)
			if err != nil {
				return sd, err
			}
			sd.ScannerInfo = append(sd.ScannerInfo, t)

		case locationN:
			var t TextWithLangElement
			t, err = t.decodeTextWithLangElement(child)
			if err != nil {
				return sd, err
			}
			sd.ScannerLocation = append(sd.ScannerLocation, t)
		}
	}

	if len(sd.ScannerName) == 0 {
		return sd, xmldoc.XMLErrMissed(nameN)
	}

	return sd, nil
}

// toXML converts a [ScannerDescription] to an XML element.
// It creates an XML element with the given name and adds child elements for
// ScannerName (required), ScannerInfo (optional), and ScannerLocation
// (optional). Each may appear multiple times for different language variants.
func (sd ScannerDescription) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	// Add ScannerName entries (required, at least one)
	for _, sn := range sd.ScannerName {
		elm.Children = append(
			elm.Children,
			sn.toXML(NsWSCN+":ScannerName"),
		)
	}

	// Add ScannerInfo entries (optional)
	for _, si := range sd.ScannerInfo {
		elm.Children = append(
			elm.Children,
			si.toXML(NsWSCN+":ScannerInfo"),
		)
	}

	// Add ScannerLocation entries (optional)
	for _, sl := range sd.ScannerLocation {
		elm.Children = append(
			elm.Children,
			sl.toXML(NsWSCN+":ScannerLocation"),
		)
	}

	return elm
}
