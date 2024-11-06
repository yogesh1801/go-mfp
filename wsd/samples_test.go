// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD message samples, for testing

package wsd

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

// Message samples from real devices

const sampleKyoceraECOSYSM2040dnMetadata = `
<?xml version="1.0"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:wsa5="http://www.w3.org/2005/08/addressing" xmlns:pnpx="http://schemas.microsoft.com/windows/pnpx/2005/10" xmlns:df="http://schemas.microsoft.com/windows/2008/09/devicefoundation" xmlns:devprof="http://schemas.xmlsoap.org/ws/2006/02/devprof" xmlns:c14n="http://www.w3.org/2001/10/xml-exc-c14n#" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd" xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:eventsrc="http://schemas.kyoceramita.com/wsf/eventsub/EventingServiceEventSourceSoapBinding" xmlns:eventsub="http://schemas.kyoceramita.com/wsf/eventsub/EventingServiceSubscriptionManagerSoapBinding" xmlns:metadata="http://schemas.xmlsoap.org/ws/2004/09/mex" xmlns:transfer="http://schemas.xmlsoap.org/ws/2004/09/transfer" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsdd="http://docs.oasis-open.org/ws-dd/ns/discovery/2009/01" xmlns:addressing="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:discovery="http://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:eventing="http://schemas.xmlsoap.org/ws/2004/08/eventing" xmlns:xop="http://www.w3.org/2004/08/xop/include" xmlns:scan="http://schemas.microsoft.com/windows/2006/08/wdp/scan" xmlns:print="http://schemas.microsoft.com/windows/2006/08/wdp/print" xmlns:print20="http://schemas.microsoft.com/windows/2014/04/wdp/printV20" xmlns:kmaccmgt="http://www.kyoceramita.com/ws/km-wsdl/setting/account_management" xmlns:kmaddrbook="http://www.kyoceramita.com/ws/km-wsdl/setting/address_book" xmlns:kmauthset="http://www.kyoceramita.com/ws/km-wsdl/setting/authentication_authorization_setting" xmlns:kmboxinfo="http://www.kyoceramita.com/ws/km-wsdl/setting/box_information" xmlns:kmcntinfo="http://www.kyoceramita.com/ws/km-wsdl/log/counter_information" xmlns:kmdevset="http://www.kyoceramita.com/ws/km-wsdl/setting/device_setting" xmlns:kmjobmng="http://www.kyoceramita.com/ws/km-wsdl/job/job_management" xmlns:kmloginfo="http://www.kyoceramita.com/ws/km-wsdl/log/log_information" xmlns:kmpanelset="http://www.kyoceramita.com/ws/km-wsdl/setting/panel_setting" xmlns:kmstored="http://www.kyoceramita.com/ws/km-wsdl/job/stored_data_operation" xmlns:kmscn="http://www.kyoceramita.com/ws/km-wsdl/job/scan_operation" xmlns:kmuserlist="http://www.kyoceramita.com/ws/km-wsdl/setting/user_list" xmlns:kmauth="http://www.kyoceramita.com/ws/km-wsdl/security/authentication_authorization" xmlns:kmdevinfo="http://www.kyoceramita.com/ws/km-wsdl/information/device_information" xmlns:kmdevctrl="http://www.kyoceramita.com/ws/km-wsdl/information/device_control" xmlns:kmfaxset="http://www.kyoceramita.com/ws/km-wsdl/setting/fax_setting" xmlns:kmdevstts="http://www.kyoceramita.com/ws/km-wsdl/status/device_status" xmlns:kmhypasmgt="http://www.kyoceramita.com/ws/km-wsdl/extension/hypas_application_management" xmlns:kmcertmgt="http://www.kyoceramita.com/ws/km-wsdl/setting/certificate_management" xmlns:kmfirmwareupdate="http://www.kyoceramita.com/ws/km-wsdl/extension/firmware_update" xmlns:kmmaint="http://www.kyoceramita.com/ws/km-wsdl/information/maintenance" xmlns:kmwsdl="http://www.kyoceramita.com/ws/km-wsdl/discovery">
  <SOAP-ENV:Header>
    <addressing:MessageID>urn:uuid:206766e0-9c5d-11ef-b13f-a93a87f9617d</addressing:MessageID>
    <addressing:RelatesTo>urn:uuid:cff33f49-2afb-4ac6-b105-a3cb1058cde6</addressing:RelatesTo>
    <addressing:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</addressing:Action>
  </SOAP-ENV:Header>
  <SOAP-ENV:Body xml:space="preserve">
    <metadata:Metadata>
      <metadata:MetadataSection Dialect="http://schemas.xmlsoap.org/ws/2006/02/devprof/ThisDevice">
        <devprof:ThisDevice>
          <devprof:FriendlyName>Kyocera:ECOSYS M2040dn:KM7B6A91</devprof:FriendlyName>
          <devprof:FirmwareVersion>2S0_2000.001.828</devprof:FirmwareVersion>
          <devprof:SerialNumber>VCF9192281</devprof:SerialNumber>
          <df:ContainerId>{4509a320-00a0-008f-00b6-006a7023f0bb}</df:ContainerId>
        </devprof:ThisDevice>
      </metadata:MetadataSection>
      <metadata:MetadataSection Dialect="http://schemas.xmlsoap.org/ws/2006/02/devprof/ThisModel">
        <devprof:ThisModel>
          <devprof:Manufacturer>Kyocera</devprof:Manufacturer>
          <devprof:ManufacturerUrl>http://www.kyoceradocumentsolutions.com</devprof:ManufacturerUrl>
          <devprof:ModelName>ECOSYS M2040dn</devprof:ModelName>
          <devprof:ModelNumber>ECOSYS M2040dn</devprof:ModelNumber>
          <devprof:ModelUrl>http://www.kyoceradocumentsolutions.com</devprof:ModelUrl>
          <devprof:PresentationUrl>http://192.168.1.102</devprof:PresentationUrl>
          <pnpx:DeviceCategory>MFP MobilePrinter</pnpx:DeviceCategory>
        </devprof:ThisModel>
      </metadata:MetadataSection>
      <metadata:MetadataSection Dialect="http://schemas.xmlsoap.org/ws/2006/02/devprof/Relationship">
        <devprof:Relationship Type="http://schemas.xmlsoap.org/ws/2006/02/devprof/host">
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/WSDScanner</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/WSDScanner</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>scan:ScannerServiceType</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/WSDScanner</devprof:ServiceId>
            <pnpx:CompatibleId>http://schemas.microsoft.com/windows/2006/08/wdp/scan/ScannerServiceType</pnpx:CompatibleId>
            <pnpx:HardwareId>VEN_0103&amp;DEV_069D</pnpx:HardwareId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/WSDPrinter</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/WSDPrinter</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>print:PrinterServiceType</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/WSDPrinter</devprof:ServiceId>
            <pnpx:CompatibleId>http://schemas.microsoft.com/windows/2006/08/wdp/print/PrinterServiceType</pnpx:CompatibleId>
            <pnpx:HardwareId>VEN_0103&amp;DEV_069D</pnpx:HardwareId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/account_management</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/account_management</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmaccmgt:account_management</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/AccountManagementService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/address_book</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/address_book</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmaddrbook:address_book</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/AddressBookService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/authentication_authorization_setting</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/authentication_authorization_setting</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmauthset:authentication_authorization_setting</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/AuthenticationAuthorizationSettingService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/box_information</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/box_information</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmboxinfo:box_information</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/BoxInformationService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/log/counter_information</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/log/counter_information</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmcntinfo:counter_information</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/CounterInformationService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/device_setting</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/device_setting</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmdevset:device_setting</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/DeviceSettingService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/job/job_management</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/job/job_management</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmjobmng:job_management</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/JobManagementService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/log/log_information</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/log/log_information</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmloginfo:log_information</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/LogInformationService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/panel_setting</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/panel_setting</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmpanelset:panel_setting</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/PanelSettingService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/job/stored_data_operation</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/job/stored_data_operation</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmstored:stored_data_operation</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/StoredDataOperationService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/job/scan_operation</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/job/scan_operation</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmscn:scan_operation</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/ScanOperationService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/user_list</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/user_list</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmuserlist:user_list</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/UserListService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/security/authentication_authorization</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/security/authentication_authorization</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmauth:authentication_authorization</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/AuthenticationAuthorizationService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/information/device_information</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/information/device_information</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmdevinfo:device_information</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/DeviceInformationService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/information/device_control</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/information/device_control</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmdevctrl:device_control</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/DeviceControlService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/fax_setting</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/fax_setting</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmfaxset:fax_setting</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/FaxSettingService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/status/device_status</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/status/device_status</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmdevstts:device_status</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/DeviceStatusService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/extension/hypas_application_management</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/extension/hypas_application_management</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmhypasmgt:hypas_application_management</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/HypasApplicationManagementService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/setting/certificate_management</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/certificate_management</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmcertmgt:certificate_management</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/CertificateManagementService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/extension/firmware_update</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/extension/firmware_update</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmfirmwareupdate:firmware_update</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/FirmwareUpdateService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/information/maintenance</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/information/maintenance</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmmaint:maaintenance</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/MaintenanceService</devprof:ServiceId>
          </devprof:Hosted>
          <devprof:Hosted>
            <addressing:EndpointReference>
              <addressing:Address>http://192.168.1.102:5358/ws/km-wsdl/discovery</addressing:Address>
            </addressing:EndpointReference>
            <addressing:EndpointReference>
              <addressing:Address>http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/discovery</addressing:Address>
            </addressing:EndpointReference>
            <devprof:Types>kmwsdl:KMWSDL_SERVICE_TYPE</devprof:Types>
            <devprof:ServiceId>uri:4509a320-00a0-008f-00b6-002507510eca/KMWSDLService</devprof:ServiceId>
          </devprof:Hosted>
        </devprof:Relationship>
      </metadata:MetadataSection>
    </metadata:Metadata>
  </SOAP-ENV:Body>
</SOAP-ENV:Envelope>
`
