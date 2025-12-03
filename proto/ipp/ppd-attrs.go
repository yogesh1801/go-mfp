// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device attribites, as returned by CUPS-Get-Devices

package ipp

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/goipp"
)

// PPDFilter specifies the subset of PPD files and their attributes,
// returned by the CUPS-Get-PPDs request.
type PPDFilter struct {
	ExcludeSchemes      []string             `ipp:"exclude-schemes,name"`
	IncludeSchemes      []string             `ipp:"include-schemes,name"`
	Limit               optional.Val[int]    `ipp:"limit,(1:MAX)"`
	PpdMake             optional.Val[string] `ipp:"ppd-make,text"`
	PpdMakeAndModel     optional.Val[string] `ipp:"ppd-make-and-model,text"`
	ModelNumber         optional.Val[int]    `ipp:"ppd-model-number"`
	PpdNaturalLanguage  optional.Val[string] `ipp:"ppd-natural-language,text"`
	PpdProduct          optional.Val[string] `ipp:"ppd-product,text"`
	PpdPsversion        optional.Val[string] `ipp:"ppd-psversion,text"`
	PpdType             optional.Val[string] `ipp:"ppd-type,keyword"`
	RequestedAttributes []string             `ipp:"requested-attributes,keyword"`
}

// PPDAttributes represents PPD file attributes, as returned by
// the CUPS-Get-PPDs request
type PPDAttributes struct {
	ObjectRawAttrs
}

// Set sets [goipp.Attibute]. It updates the appropriate structure
// field and Object's raw attributes.
func (attrs *PPDAttributes) Set(attr goipp.Attribute) error {
	return attrs.set(attr, attrs)
}
