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
)

// PPDFilter specifies the subset of PPD files and their attributes,
// returned by the CUPS-Get-PPDs request.
type PPDFilter struct {
	ExcludeSchemes      []string             `ipp:"exclude-schemes"`
	IncludeSchemes      []string             `ipp:"include-schemes"`
	Limit               optional.Val[int]    `ipp:"limit"`
	PpdMake             optional.Val[string] `ipp:"ppd-make"`
	PpdMakeAndModel     optional.Val[string] `ipp:"ppd-make-and-model"`
	PpdModelNumber      optional.Val[int]    `ipp:"ppd-model-number"`
	PpdNaturalLanguage  optional.Val[string] `ipp:"ppd-natural-language"`
	PpdProduct          optional.Val[string] `ipp:"ppd-product"`
	PpdPsVersion        optional.Val[string] `ipp:"ppd-psversion"`
	PpdType             optional.Val[string] `ipp:"ppd-type,keyword"`
	RequestedAttributes []string             `ipp:"requested-attributes"`
}

// PPDAttributes represents PPD file attributes, as returned by
// the CUPS-Get-PPDs request
type PPDAttributes struct {
	ObjectRawAttrs
	CUPSPPDAttributesGroup
}
