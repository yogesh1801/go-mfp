// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// device condition

package wsscan

import (
	"fmt"
	"time"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// DeviceCondition represents the <wscn:DeviceCondition> element,
// providing details about one of the scanner's currently active conditions.
type DeviceCondition struct {
	Component Component
	Name      ConditionName
	Severity  Severity
	Time      time.Time
}

// toXML generates XML tree for the [DeviceCondition].
func (dc DeviceCondition) toXML(name string) xmldoc.Element {
	children := []xmldoc.Element{
		dc.Component.toXML(NsWSCN + ":Component"),
		dc.Name.toXML(NsWSCN + ":Name"),
		dc.Severity.toXML(NsWSCN + ":Severity"),
		xmldoc.Element{
			Name: NsWSCN + ":Time",
			Text: dc.Time.Format(time.RFC3339),
		},
	}
	return xmldoc.Element{
		Name:     name,
		Children: children,
	}
}

// decodeDeviceCondition decodes [DeviceCondition] from the XML tree.
func decodeDeviceCondition(root xmldoc.Element) (
	dc DeviceCondition, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	component := xmldoc.Lookup{
		Name:     NsWSCN + ":Component",
		Required: true,
	}
	name := xmldoc.Lookup{
		Name:     NsWSCN + ":Name",
		Required: true,
	}
	severity := xmldoc.Lookup{
		Name:     NsWSCN + ":Severity",
		Required: true,
	}
	time := xmldoc.Lookup{
		Name:     NsWSCN + ":Time",
		Required: true,
	}

	missed := root.Lookup(
		&component,
		&name,
		&severity,
		&time,
	)
	if missed != nil {
		return dc, xmldoc.XMLErrMissed(missed.Name)
	}

	if dc.Component, err = decodeComponent(component.Elem); err != nil {
		return dc, fmt.Errorf("component: %w", err)
	}
	if dc.Name, err = decodeConditionName(name.Elem); err != nil {
		return dc, fmt.Errorf("name: %w", err)
	}
	if dc.Severity, err = decodeSeverity(severity.Elem); err != nil {
		return dc, fmt.Errorf("severity: %w", err)
	}
	if dc.Time, err = decodeTime(time.Elem); err != nil {
		return dc, fmt.Errorf("time: %w", err)
	}

	return dc, nil
}
