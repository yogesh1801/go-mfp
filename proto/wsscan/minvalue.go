// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// min value

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// MinValue defines the minimum value for a scanner configuration element.
type MinValue int

// decodeMinValue decodes [MinValue] from the XML tree.
func decodeMinValue(root xmldoc.Element) (MinValue, error) {
	return decodeMinValueStr(root.Text)
}

// toXML generates XML tree for the [MinValue].
func (mv MinValue) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: mv.String(),
	}
}

// String returns a string representation of the [MinValue].
func (mv MinValue) String() string {
	return strconv.Itoa(int(mv))
}

// decodeMinValueStr decodes [MinValue] from its string representation.
func decodeMinValueStr(s string) (MinValue, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid MinValue: %q", s)
	}
	return MinValue(v), nil
}
