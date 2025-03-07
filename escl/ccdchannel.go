// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan CCD color channel

package escl

import "github.com/alexpevzner/mfp/util/xmldoc"

// CCDChannel specifies which CCD color channel to use for grayscale
// and monochrome scannig.
type CCDChannel int

// Known CCD Channels.
const (
	UnknownCCDChannel CCDChannel = iota // Unknown CCD
	Red                                 // Use the RED channel
	Green                               // Use the Green channel
	Blue                                // Use the Blue channel
	NTSC                                // NTSC-standard mix
	GrayCcd                             // Dedicated hardware Gray CCD
	GrayCcdEmulated                     // Emulated Gray CCD (1/3 RGB)
)

// decodeCCDChannel decodes [CCDChannel] from the XML tree.
func decodeCCDChannel(root xmldoc.Element) (ccd CCDChannel, err error) {
	return decodeEnum(root, DecodeCCDChannel)
}

// toXML generates XML tree for the [CCDChannel].
func (ccd CCDChannel) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: ccd.String(),
	}
}

// String returns a string representation of the [CCDChannel]
func (ccd CCDChannel) String() string {
	switch ccd {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case NTSC:
		return "NTSC"
	case GrayCcd:
		return "GrayCcd"
	case GrayCcdEmulated:
		return "GrayCcdEmulated"
	}

	return "Unknown"
}

// DecodeCCDChannel decodes [CCDChannel] out of its XML string representation.
func DecodeCCDChannel(s string) CCDChannel {
	switch s {
	case "Red":
		return Red
	case "Green":
		return Green
	case "Blue":
		return Blue
	case "NTSC":
		return NTSC
	case "GrayCcd":
		return GrayCcd
	case "GrayCcdEmulated":
		return GrayCcdEmulated
	}

	return UnknownCCDChannel
}
