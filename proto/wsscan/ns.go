// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// WS-Scan namespace

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// Namespace prefixes:
const (
	NsSOAP       = "soap"
	NsAddressing = "wsa"
	NsWSCN       = "wscn"
	NsXML        = "xmlns"
)

// WS-Addressing well-known URIs:
const (
	// AddrAnonymous is the WS-Addressing anonymous role URI,
	// used as the To address in responses.
	AddrAnonymous = "https://schemas.xmlsoap.org/ws/2003/03/addressing/role/anonymous"
)

// NsMap maps namespace prefixes to URLs for XML encoding/decoding.
var NsMap = xmldoc.Namespace{
	// SOAP 1.2
	{Prefix: NsSOAP,
		URL: "http://www.w3.org/2003/05/soap-envelope"},

	// WS-Addressing
	{Prefix: NsAddressing,
		URL: "http://schemas.xmlsoap.org/ws/2004/08/addressing"},

	// WS-Scan
	{Prefix: NsWSCN,
		URL: "http://schemas.microsoft.com/windows/2006/01/wdp/scan"},
}
