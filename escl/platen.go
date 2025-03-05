// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Platen capabilities

package escl

import (
	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// Platen contains scanner capabilities for the Platen source.
type Platen struct {
	PlatenInputCaps optional.Val[InputSourceCaps] // Platen capabilities
}

// decodePlaten decodes [Platen] from the XML tree.
func decodePlaten(root xmldoc.Element) (
	platen Platen, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	inputcaps := xmldoc.Lookup{Name: NsScan + ":PlatenInputCaps"}
	root.Lookup(&inputcaps)

	// Decode elements
	if inputcaps.Found {
		var caps InputSourceCaps
		caps, err = decodeInputSourceCaps(inputcaps.Elem)

		if err != nil {
			return
		}

		platen.PlatenInputCaps = optional.New(caps)
	}
	return
}

// toXML generates XML tree for the [Platen].
func (platen Platen) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}
	if platen.PlatenInputCaps != nil {
		chld := (*platen.PlatenInputCaps).toXML(
			NsScan + ":PlatenInputCaps")
		elm.Children = append(elm.Children, chld)
	}
	return elm
}
