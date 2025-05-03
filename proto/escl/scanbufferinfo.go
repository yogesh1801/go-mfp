// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan Buffer Info

package escl

import (
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScanBufferInfo is the scanner response, that represents the negotiated
// (adjusted by the scanner) scanning parameters ([ScanSettings]) and image
// parameters estimation.
//
// eSCL Technical Specification, 10.
type ScanBufferInfo struct {
	ScanSettings ScanSettings // Returned scanning parameters
	ImageWidth   int          // Estimated image width
	ImageHeight  int          // Estimated image height
	BytesPerLine int          // Bytes per Line size of uncompressed image
}

// DecodeScanBufferInfo decodes [ScanBufferInfo] from the XML tree.
func DecodeScanBufferInfo(root xmldoc.Element) (
	info ScanBufferInfo, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	ss := xmldoc.Lookup{Name: NsScan + ":ScanSettings", Required: true}
	wid := xmldoc.Lookup{Name: NsScan + ":ImageWidth", Required: true}
	hei := xmldoc.Lookup{Name: NsScan + ":ImageHeight", Required: true}
	bpl := xmldoc.Lookup{Name: NsScan + ":BytesPerLine", Required: true}

	missed := root.Lookup(&ss, &wid, &hei, &bpl)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	info.ScanSettings, err = DecodeScanSettings(ss.Elem)
	if err == nil {
		info.ImageWidth, err = decodeNonNegativeInt(wid.Elem)
	}
	if err == nil {
		info.ImageHeight, err = decodeNonNegativeInt(hei.Elem)
	}
	if err == nil {
		info.BytesPerLine, err = decodeNonNegativeInt(bpl.Elem)
	}

	return
}

// ToXML generates XML tree for the [ScanBufferInfo].
func (info ScanBufferInfo) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsScan + ":ScanBufferInfo",
		Children: []xmldoc.Element{
			info.ScanSettings.ToXML(),
			xmldoc.WithText(NsScan+":ImageWidth",
				strconv.FormatUint(
					uint64(info.ImageWidth), 10)),
			xmldoc.WithText(NsScan+":ImageHeight",
				strconv.FormatUint(
					uint64(info.ImageHeight), 10)),
			xmldoc.WithText(NsScan+":BytesPerLine",
				strconv.FormatUint(
					uint64(info.BytesPerLine), 10)),
		},
	}

	return elm
}
