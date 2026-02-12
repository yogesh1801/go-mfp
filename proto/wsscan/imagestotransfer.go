// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ImagesToTransfer: specifies the number of images to scan for the current job

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ImagesToTransfer specifies the number of images to scan for the current job.
// The text value is an integer in the range from 0 through 2147483648.
//
// Example XML:
//
//	<wscn:ImagesToTransfer wscn:MustHonor="true">10</wscn:ImagesToTransfer>
type ImagesToTransfer ValWithOptions[int]

// decodeImagesToTransfer decodes an ImagesToTransfer from an XML element.
func decodeImagesToTransfer(root xmldoc.Element) (ImagesToTransfer, error) {
	var base ValWithOptions[int]
	decoded, err := base.decodeValWithOptions(root, intValueDecoder)
	if err != nil {
		return ImagesToTransfer{}, err
	}
	return ImagesToTransfer(decoded), nil
}

// toXML converts an ImagesToTransfer to an XML element.
func (itt ImagesToTransfer) toXML(name string) xmldoc.Element {
	return ValWithOptions[int](itt).toXML(name, intValueEncoder)
}
