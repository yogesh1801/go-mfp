// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Camera capabilities

package escl

import (
	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// Camera contains scanner capabilities for the Camera source.
type Camera struct {
	CameraInputCaps optional.Val[InputSourceCaps] // Camera capabilities
}

// decodeCamera decodes [Camera] from the XML tree.
func decodeCamera(root xmldoc.Element) (
	camera Camera, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	inputcaps := xmldoc.Lookup{Name: NsScan + ":CameraInputCaps"}
	root.Lookup(&inputcaps)

	// Decode elements
	if inputcaps.Found {
		var caps InputSourceCaps
		caps, err = decodeInputSourceCaps(inputcaps.Elem)

		if err != nil {
			return
		}

		camera.CameraInputCaps = optional.New(caps)
	}
	return
}

// toXML generates XML tree for the [Camera].
func (camera Camera) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}
	if camera.CameraInputCaps != nil {
		chld := (*camera.CameraInputCaps).toXML(
			NsScan + ":CameraInputCaps")
		elm.Children = append(elm.Children, chld)
	}
	return elm
}
