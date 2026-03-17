// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ImageInformation: image data information resulting from a validated scan

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ImageInformation contains information about the resulting image data from a
// scan made with a ScanTicket. MediaBackImageInfo and MediaFrontImageInfo are
// both optional and present only when the corresponding scan side was used.
type ImageInformation struct {
	MediaBackImageInfo  optional.Val[MediaSideImageInfo]
	MediaFrontImageInfo optional.Val[MediaSideImageInfo]
}

// toXML creates an XML element for ImageInformation.
func (ii ImageInformation) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	if ii.MediaBackImageInfo != nil {
		elm.Children = append(elm.Children,
			optional.Get(ii.MediaBackImageInfo).toXML(
				NsWSCN+":MediaBackImageInfo"))
	}
	if ii.MediaFrontImageInfo != nil {
		elm.Children = append(elm.Children,
			optional.Get(ii.MediaFrontImageInfo).toXML(
				NsWSCN+":MediaFrontImageInfo"))
	}

	return elm
}

// decodeImageInformation decodes an ImageInformation from an XML element.
func decodeImageInformation(root xmldoc.Element) (ImageInformation, error) {
	var ii ImageInformation

	mediaBack := xmldoc.Lookup{Name: NsWSCN + ":MediaBackImageInfo"}
	mediaFront := xmldoc.Lookup{Name: NsWSCN + ":MediaFrontImageInfo"}

	root.Lookup(&mediaBack, &mediaFront)

	if mediaBack.Found {
		back, err := decodeMediaSideImageInfo(mediaBack.Elem)
		if err != nil {
			return ii, fmt.Errorf("MediaBackImageInfo: %w", err)
		}
		ii.MediaBackImageInfo = optional.New(back)
	}

	if mediaFront.Found {
		front, err := decodeMediaSideImageInfo(mediaFront.Elem)
		if err != nil {
			return ii, fmt.Errorf("MediaFrontImageInfo: %w", err)
		}
		ii.MediaFrontImageInfo = optional.New(front)
	}

	return ii, nil
}
