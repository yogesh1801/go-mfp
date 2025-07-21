// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// max value

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// MaxValue defines the maximum value for a scanner configuration element.
type MaxValue int

// decodeMaxValue decodes [MaxValue] from the XML tree.
func decodeMaxValue(root xmldoc.Element) (MaxValue, error) {
	return decodeMaxValueStr(root.Text)
}

// toXML generates XML tree for the [MaxValue].
func (mv MaxValue) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: mv.String(),
	}
}

// String returns a string representation of the [MaxValue].
func (mv MaxValue) String() string {
	return strconv.Itoa(int(mv))
}

// decodeMaxValueStr decodes [MaxValue] from its string representation.
func decodeMaxValueStr(s string) (MaxValue, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid MaxValue: %q", s)
	}
	return MaxValue(v), nil
}
