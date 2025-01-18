// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL namespace

package escl

import "github.com/alexpevzner/mfp/xmldoc"

// Namespace prefixes:
const (
	NsScan = "scan"
	NsPWG  = "pwg"
)

// NsMap maps namespace prefixes to URL
var NsMap = xmldoc.Namespace{
	{Prefix: NsScan, URL: "http://schemas.hp.com/imaging/escl/2011/05/03"},
	{Prefix: NsPWG, URL: "http://www.pwg.org/schemas/2010/12/sm"},
}
