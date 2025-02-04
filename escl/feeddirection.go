// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner feed directions

package escl

import "github.com/alexpevzner/mfp/xmldoc"

// FeedDirection specifies the feed direction of the input media
// (affects the resulting image orientation).
type FeedDirection int

// Known feed directions.
const (
	UnknownFeedDirection FeedDirection = iota // Unknown CCD
	LongEdgeFeed                              // Longest edge scanned first
	ShortEdgeFeed                             // Shortest edge scanned 1st
)

// decodeFeedDirection decodes [FeedDirection] from the XML tree.
func decodeFeedDirection(root xmldoc.Element) (feed FeedDirection, err error) {
	return decodeEnum(root, DecodeFeedDirection, NsScan)
}

// toXML generates XML tree for the [FeedDirection].
func (feed FeedDirection) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: NsScan + ":" + feed.String(),
	}
}

// String returns a string representation of the [FeedDirection]
func (feed FeedDirection) String() string {
	switch feed {
	case LongEdgeFeed:
		return "LongEdgeFeed"
	case ShortEdgeFeed:
		return "ShortEdgeFeed"
	}

	return "Unknown"
}

// DecodeFeedDirection decodes [FeedDirection] out of its XML
// string representation.
func DecodeFeedDirection(s string) FeedDirection {
	switch s {
	case "LongEdgeFeed":
		return LongEdgeFeed
	case "ShortEdgeFeed":
		return ShortEdgeFeed
	}

	return UnknownFeedDirection
}
