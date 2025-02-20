// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF capabilities

package escl

import (
	"strconv"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// ADF contains scanner capabilities for the automated document feeder.
type ADF struct {
	ADFSimplexInputCaps optional.Val[InputSourceCaps] // ADF simplex caps
	ADFDuplexInputCaps  optional.Val[InputSourceCaps] // ADF duplex caps
	FeederCapacity      optional.Val[int]             // Feeder capacity
	ADFOptions          []ADFOption                   // ADF options
	Justification       optional.Val[Justification]   // Image justification
}

// decodeADF decodes [ADF] from the XML tree
func decodeADF(root xmldoc.Element) (adf ADF, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	simplex := xmldoc.Lookup{Name: NsScan + ":AdfSimplexInputCaps"}
	duplex := xmldoc.Lookup{Name: NsScan + ":AdfDuplexInputCaps"}
	capacity := xmldoc.Lookup{Name: NsScan + ":FeederCapacity"}
	options := xmldoc.Lookup{Name: NsScan + ":AdfOptions"}
	justification := xmldoc.Lookup{Name: NsScan + ":Justification"}

	root.Lookup(&simplex, &duplex, &capacity, &options, &justification)

	// Decode elements
	if simplex.Found {
		var caps InputSourceCaps
		caps, err = decodeInputSourceCaps(simplex.Elem)
		if err != nil {
			return
		}

		adf.ADFSimplexInputCaps = optional.New(caps)
	}

	if duplex.Found {
		var caps InputSourceCaps
		caps, err = decodeInputSourceCaps(duplex.Elem)
		if err != nil {
			return
		}

		adf.ADFDuplexInputCaps = optional.New(caps)
	}

	if capacity.Found {
		var c int
		c, err = decodeNonNegativeInt(capacity.Elem)
		if err != nil {
			return
		}

		adf.FeederCapacity = optional.New(c)
	}

	if options.Found {
		for _, elem := range options.Elem.Children {
			if elem.Name == NsScan+":AdfOption" {
				var opt ADFOption
				opt, err = decodeADFOption(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						options.Elem, err)
					return
				}

				adf.ADFOptions = append(adf.ADFOptions, opt)
			}
		}
	}

	if justification.Found {
		var jst Justification
		jst, err = decodeJustification(justification.Elem)
		if err != nil {
			return
		}

		adf.Justification = optional.New(jst)
	}

	return
}

// toXML generates XML tree for the [ADF].
func (adf ADF) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	if adf.ADFSimplexInputCaps != nil {
		chld := (*adf.ADFSimplexInputCaps).toXML(
			NsScan + ":AdfSimplexInputCaps")
		elm.Children = append(elm.Children, chld)
	}

	if adf.ADFDuplexInputCaps != nil {
		chld := (*adf.ADFDuplexInputCaps).toXML(
			NsScan + ":AdfDuplexInputCaps")
		elm.Children = append(elm.Children, chld)
	}

	if adf.FeederCapacity != nil {
		chld := xmldoc.Element{
			Name: NsScan + ":FeederCapacity",
			Text: strconv.Itoa(*adf.FeederCapacity),
		}
		elm.Children = append(elm.Children, chld)
	}

	if adf.ADFOptions != nil {
		chld := xmldoc.Element{Name: NsScan + ":AdfOptions"}
		for _, opt := range adf.ADFOptions {
			chld2 := opt.toXML(NsScan + ":AdfOption")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	if adf.Justification != nil {
		chld := (*adf.Justification).toXML(
			NsScan + ":Justification")
		elm.Children = append(elm.Children, chld)
	}

	return elm
}
