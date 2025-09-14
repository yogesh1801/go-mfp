// MFP - Miulti-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Name element for DeviceCondition and ConditionHistoryEntry

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ConditionName names the current error condition specified in
// DeviceCondition or ConditionHistoryEntry.
//
// Values are defined by the WS-Scan spec.
type ConditionName string

// Known ConditionName values.
const (
	UnknownConditionName ConditionName = ""
	Calibrating          ConditionName = "Calibrating"
	CoverOpen            ConditionName = "CoverOpen"
	InputTrayEmpty       ConditionName = "InputTrayEmpty"
	InterlockOpen        ConditionName = "InterlockOpen"
	InternalStorageFull  ConditionName = "InternalStorageFull"
	LampError            ConditionName = "LampError"
	LampWarming          ConditionName = "LampWarming"
	MediaJam             ConditionName = "MediaJam"
	MultipleFeedError    ConditionName = "MultipleFeedError"
)

// decodeConditionName decodes [ConditionName] from the XML tree.
func decodeConditionName(root xmldoc.Element) (cn ConditionName, err error) {
	var v string
	v, err = decodeNMTOKEN(root)
	if err != nil {
		err = fmt.Errorf("invalid ConditionName: %q",
			root.Text)
		err = xmldoc.XMLErrWrap(root, err)
		return
	}

	cn = ConditionName(v)
	return
}

// toXML generates XML tree for the [ConditionName].
func (cn ConditionName) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: cn.String(),
	}
}

// String returns a string representation of the [ConditionName].
func (cn ConditionName) String() string {
	return string(cn)
}
