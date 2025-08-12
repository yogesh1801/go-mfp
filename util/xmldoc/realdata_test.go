// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Test of decoding on real data

package xmldoc

import (
	"os"
	"strings"
	"testing"
)

// Kyocera ECOSYS M2040dn WSDD ProbeMatches message
const KyoceraECOSYSM2040dnProbeMatches = `<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:pnpx="http://schemas.microsoft.com/windows/pnpx/2005/10" xmlns:df="http://schemas.microsoft.com/windows/2008/09/devicefoundation" xmlns:devprof="http://schemas.xmlsoap.org/ws/2006/02/devprof" xmlns:c14n="http://www.w3.org/2001/10/xml-exc-c14n#" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd" xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:eventsub="http://schemas.kyoceramita.com/wsf/eventsub" xmlns:metadata="http://schemas.xmlsoap.org/ws/2004/09/mex" xmlns:addressing="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:discovery="http://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:eventing="http://schemas.xmlsoap.org/ws/2004/08/eventing" xmlns:print="http://schemas.microsoft.com/windows/2006/08/wdp/print" xmlns:scan="http://schemas.microsoft.com/windows/2006/08/wdp/scan"><SOAP-ENV:Header><addressing:MessageID>urn:uuid:bc8b1d56-8d93-11ef-b071-a93a87f9617d</addressing:MessageID><addressing:RelatesTo>urn:uuid:f9f56956-7662-4b32-9459-d0f2823424fe</addressing:RelatesTo><addressing:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</addressing:To><addressing:Action>http://schemas.xmlsoap.org/ws/2005/04/discovery/ProbeMatches</addressing:Action><discovery:AppSequence MessageNumber="525" InstanceId="1727722437"></discovery:AppSequence></SOAP-ENV:Header><SOAP-ENV:Body xml:space="preserve"><discovery:ProbeMatches><discovery:ProbeMatch><addressing:EndpointReference><addressing:Address>uuid:4509a320-00a0-008f-00b6-002507510eca</addressing:Address></addressing:EndpointReference><discovery:Types>devprof:Device scan:ScanDeviceType print:PrintDeviceType</discovery:Types><discovery:XAddrs>http://192.168.1.102:5358/DeviceService/ http://[fe80::217:c8ff:fe7b:6a91]:5358/DeviceService/</discovery:XAddrs><discovery:MetadataVersion>1</discovery:MetadataVersion></discovery:ProbeMatch></discovery:ProbeMatches></SOAP-ENV:Body></SOAP-ENV:Envelope>`

// WSD namespace
var ns = Namespace{
	// SOAP 1.2
	{Prefix: "s", URL: "http://www.w3.org/2003/05/soap-envelope"},

	// SOAP 1.1
	{Prefix: "s", URL: "http://schemas.xmlsoap.org/soap/envelope"},

	// WSD prefixes
	{Prefix: "a", URL: "http://schemas.xmlsoap.org/ws/2004/08/addressing"},
	{Prefix: "d", URL: "http://schemas.xmlsoap.org/ws/2005/04/discovery"},
	{Prefix: "devprof", URL: "http://schemas.xmlsoap.org/ws/2006/02/devprof"},
	{Prefix: "mex", URL: "http://schemas.xmlsoap.org/ws/2004/09/mex"},
	{Prefix: "pnpx", URL: "http://schemas.microsoft.com/windows/pnpx/2005/10"},
	{Prefix: "scan", URL: "http://schemas.microsoft.com/windows/2006/08/wdp/scan"},
	{Prefix: "print", URL: "http://schemas.microsoft.com/windows/2006/08/wdp/print"},
}

// TestRealData tests real data decoding
func TestRealData(t *testing.T) {
	input := strings.NewReader(KyoceraECOSYSM2040dnProbeMatches)
	xml, err := Decode(ns, input)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	s := xml.EncodeIndentString(ns, "  ")
	os.Stdout.Write([]byte(s))
}
