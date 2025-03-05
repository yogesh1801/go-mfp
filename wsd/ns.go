// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD namespace

package wsd

import "github.com/alexpevzner/mfp/util/xmldoc"

// Namespace prefixes:
const (
	NsSOAP       = "s"
	NsAddressing = "a"
	NsDiscovery  = "d"
	NsDevprof    = "devprof"
	NsMex        = "mex"
	NsPNPX       = "pnpx"
	NsScan       = "scan"
	NsPrint      = "print"
)

// NsMap maps namespace prefixes to URL
var NsMap = xmldoc.Namespace{
	// SOAP 1.2
	{Prefix: NsSOAP, URL: "http://www.w3.org/2003/05/soap-envelope"},

	// SOAP 1.1
	{Prefix: NsSOAP, URL: "http://schemas.xmlsoap.org/soap/envelope"},

	// WSD prefixes
	{Prefix: NsAddressing, URL: "http://schemas.xmlsoap.org/ws/2004/08/addressing"},
	{Prefix: NsDiscovery, URL: "http://schemas.xmlsoap.org/ws/2005/04/discovery"},
	{Prefix: NsDevprof, URL: "http://schemas.xmlsoap.org/ws/2006/02/devprof"},
	{Prefix: NsMex, URL: "http://schemas.xmlsoap.org/ws/2004/09/mex"},
	{Prefix: NsPNPX, URL: "http://schemas.microsoft.com/windows/pnpx/2005/10"},
	{Prefix: NsScan, URL: "http://schemas.microsoft.com/windows/2006/08/wdp/scan"},
	{Prefix: NsPrint, URL: "http://schemas.microsoft.com/windows/2006/08/wdp/print"},
}
