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
	ExcludeSchemes      []string             `ipp:"exclude-schemes,1setOf name"`
	IncludeSchemes      []string             `ipp:"include-schemes,1setOf name"`
	Limit               optional.Val[int]    `ipp:"limit,integer(1:MAX)"`
	PpdMake             optional.Val[string] `ipp:"ppd-make,text(127)"`
	PpdMakeAndModel     optional.Val[string] `ipp:"ppd-make-and-model,text(127)"`
	PpdModelNumber      optional.Val[int]    `ipp:"ppd-model-number"`
	PpdNaturalLanguage  optional.Val[string] `ipp:"ppd-natural-language,naturallanguage"`
	PpdProduct          optional.Val[string] `ipp:"ppd-product,text(127)"`
	PpdPsVersion        optional.Val[string] `ipp:"ppd-psversion,text(127)"`
	PpdType             optional.Val[string] `ipp:"ppd-type,keyword"`
	RequestedAttributes []string             `ipp:"requested-attributes,1setOf keyword"`
}

// PPDAttributes represents PPD file attributes, as returned by
// the CUPS-Get-PPDs request
type PPDAttributes struct {
	ObjectRawAttrs
	CUPSPPDAttributesGroup
}
