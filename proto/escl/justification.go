// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF image justification

package escl

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// Justification specifies how the ADF justify the document.
type Justification struct {
	XImagePosition ImagePosition // Horizontal image position
	YImagePosition ImagePosition // Vertical image position
}

// decodeJustification decodes [Justification] from the XML tree
func decodeJustification(root xmldoc.Element) (jst Justification, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	xpos := xmldoc.Lookup{Name: NsScan + ":XImagePosition", Required: true}
	ypos := xmldoc.Lookup{Name: NsScan + ":YImagePosition", Required: true}

	missed := root.Lookup(&xpos, &ypos)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	jst.XImagePosition, err = decodeImagePosition(xpos.Elem)
	if err == nil {
		jst.YImagePosition, err = decodeImagePosition(ypos.Elem)
	}

	return
}

// toXML generates XML tree for the [Justification].
func (jst Justification) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			jst.XImagePosition.toXML(NsScan + ":XImagePosition"),
			jst.YImagePosition.toXML(NsScan + ":YImagePosition"),
		},
	}

	return elm
}
