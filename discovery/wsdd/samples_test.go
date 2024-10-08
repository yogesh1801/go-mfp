// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSDD message samples, for testing

package wsdd

// Message samples, from the Microsoft "Web Services on Devices"
// specification.
//
// https://learn.microsoft.com/pdf?url=https%3A%2F%2Flearn.microsoft.com%2Fen-us%2Fwindows%2Fwin32%2Fwsdapi%2Ftoc.json#D17-

const sampleHello = `
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="https://www.w3.org/2003/05/soap-envelope" xmlns:wsa="https://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsd="https://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:wsdp="https://schemas.xmlsoap.org/ws/2006/02/devprof">
<soap:Header>
    <wsa:To>
	urn:schemas-xmlsoap-org:ws:2005:04:discovery
    </wsa:To>
    <wsa:Action>
	https://schemas.xmlsoap.org/ws/2005/04/discovery/Hello
    </wsa:Action>
    <wsa:MessageID>
	urn:uuid:0f5d604c-81ac-4abc-8010-51dbffad55f2
    </wsa:MessageID>
    <wsd:AppSequence InstanceId="2" SequenceId="urn:uuid:369a7d7b-5f87-48a4-aa9a-189edf2a8772" MessageNumber="14">
    </wsd:AppSequence>
</soap:Header>
<soap:Body>
    <wsd:Hello>
	<wsa:EndpointReference>
	    <wsa:Address>
		urn:uuid:37f86d35-e6ac-4241-964f-1d9ae46fb366
	    </wsa:Address>
	</wsa:EndpointReference>
	<wsd:Types>wsdp:Device</wsd:Types>
	<wsd:MetadataVersion>2</wsd:MetadataVersion>
    </wsd:Hello>
</soap:Body>
</soap:Envelope>
`
const sampleBye = `
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="https://www.w3.org/2003/05/soap-envelope" xmlns:wsa="https://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsd="https://schemas.xmlsoap.org/ws/2005/04/discovery">
<soap:Header>
    <wsa:To>
	urn:schemas-xmlsoap-org:ws:2005:04:discovery
    </wsa:To>
    <wsa:Action>
	https://schemas.xmlsoap.org/ws/2005/04/discovery/Bye
    </wsa:Action>
    <wsa:MessageID>
	urn:uuid:193ccfa0-347d-41a1-9285-f500b6b96a15
    </wsa:MessageID>
    <wsd:AppSequence InstanceId="2" SequenceId="urn:uuid:369a7d7b-5f87-48a4-aa9a-189edf2a8772" MessageNumber="21">
    </wsd:AppSequence>
</soap:Header>
<soap:Body>
    <wsd:Bye>
	<wsa:EndpointReference>
	    <wsa:Address>
		urn:uuid:37f86d35-e6ac-4241-964f-1d9ae46fb366
	    </wsa:Address>
	</wsa:EndpointReference>
    </wsd:Bye>
</soap:Body>
</soap:Envelope>
`
