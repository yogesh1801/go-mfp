// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Metadata test

package wsd

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

// TestMetagata tests Metadata encoding and decoding
func TestMetagata(t *testing.T) {
	meta := Metadata{
		ThisDevice: ThisDeviceMetadata{
			FriendlyName: LocalizedStringList{
				{String: "I.Fyodorov FP-0001"},
				{String: "И.Фёдоров ПФ-0001", Lang: "ru-RU"},
			},
			FirmwareVersion: "0.0.1",
			SerialNumber:    "FP-8322017",
		},
		ThisModel: ThisModelMetadata{
			Manufacturer: LocalizedStringList{
				{String: "I.Fyodorov"},
				{String: "И.Фёдоров", Lang: "ru-RU"},
			},
			ManufacturerURL: optional.New("http://example.com"),
			ModelName: LocalizedStringList{
				{String: "FP-0001"},
				{String: "ПФ-0001", Lang: "ru-RU"},
			},
			ModelNumber:     "FP-0001",
			ModelURL:        optional.New("http://example.com/FP-0001"),
			PresentationURL: optional.New("http://example.com/FP-0001/pres"),
		},
		Relationship: Relationship{
			Host: &ServiceMetadata{
				EndpointReference: []EndpointReference{
					{"http://127.0.0.1/"},
				},
			},
			Hosted: []ServiceMetadata{
				{
					EndpointReference: []EndpointReference{
						{"http://127.0.0.1/print"},
					},
					Types:     []Type{PrinterServiceType},
					ServiceID: "uri:b827bd97-925c-4502-a7db-4918a0abfc11",
				},
				{
					EndpointReference: []EndpointReference{
						{"http://127.0.0.1/scan"},
					},
					Types:     []Type{ScannerServiceType},
					ServiceID: "uri:6499d366-62a5-4da9-8c18-5af6eea01f22",
				},
			},
		},
	}

	meta2, err := DecodeMetadata(meta.ToXML())
	if err != nil {
		t.Errorf("DecodeMetadata: %s", err)
		return
	}

	if !reflect.DeepEqual(meta, meta2) {
		t.Errorf("encode/decode mismatch\n"+
			"expected: %s\n"+
			"present:  %s\n",
			meta.ToXML().EncodeString(NsMap),
			meta2.ToXML().EncodeString(NsMap),
		)
	}
}

// TestKyoceraECOSYSM2040dnMetadata tests decoding metadate from
// the real device.
func TestKyoceraECOSYSM2040dnMetadata(t *testing.T) {
	msg, err := DecodeMsg([]byte(sampleKyoceraECOSYSM2040dnMetadata))
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	meta, ok := msg.Body.(Metadata)
	if !ok {
		t.Errorf("Message body expected %s, present %s",
			reflect.TypeOf(Metadata{}), msg.Body)
	}

	expected := Metadata{
		ThisDevice: ThisDeviceMetadata{
			FriendlyName:    LocalizedStringList{LocalizedString{String: "Kyocera:ECOSYS M2040dn:KM7B6A91", Lang: ""}},
			FirmwareVersion: "2S0_2000.001.828",
			SerialNumber:    "VCF9192281"},
		ThisModel: ThisModelMetadata{
			Manufacturer:    LocalizedStringList{LocalizedString{String: "Kyocera", Lang: ""}},
			ManufacturerURL: optional.New("http://www.kyoceradocumentsolutions.com"),
			ModelName:       LocalizedStringList{LocalizedString{String: "ECOSYS M2040dn", Lang: ""}},
			ModelNumber:     "ECOSYS M2040dn",
			ModelURL:        optional.New("http://www.kyoceradocumentsolutions.com"),
			PresentationURL: optional.New("http://192.168.1.102")},
		Relationship: Relationship{
			Hosted: []ServiceMetadata{
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/WSDScanner"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/WSDScanner"}},
					Types:     []Type{ScannerServiceType},
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/WSDScanner"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/WSDPrinter"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/WSDPrinter"}},
					Types:     []Type{PrinterServiceType},
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/WSDPrinter"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/account_management"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/account_management"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/AccountManagementService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/address_book"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/address_book"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/AddressBookService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/authentication_authorization_setting"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/authentication_authorization_setting"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/AuthenticationAuthorizationSettingService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/box_information"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/box_information"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/BoxInformationService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/log/counter_information"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/log/counter_information"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/CounterInformationService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/device_setting"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/device_setting"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/DeviceSettingService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/job/job_management"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/job/job_management"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/JobManagementService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/log/log_information"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/log/log_information"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/LogInformationService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/panel_setting"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/panel_setting"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/PanelSettingService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/job/stored_data_operation"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/job/stored_data_operation"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/StoredDataOperationService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/job/scan_operation"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/job/scan_operation"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/ScanOperationService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/user_list"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/user_list"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/UserListService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/security/authentication_authorization"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/security/authentication_authorization"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/AuthenticationAuthorizationService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/information/device_information"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/information/device_information"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/DeviceInformationService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/information/device_control"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/information/device_control"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/DeviceControlService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/fax_setting"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/fax_setting"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/FaxSettingService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/status/device_status"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/status/device_status"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/DeviceStatusService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/extension/hypas_application_management"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/extension/hypas_application_management"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/HypasApplicationManagementService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/setting/certificate_management"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/setting/certificate_management"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/CertificateManagementService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/extension/firmware_update"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/extension/firmware_update"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/FirmwareUpdateService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/information/maintenance"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/information/maintenance"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/MaintenanceService"},
				ServiceMetadata{
					EndpointReference: []EndpointReference{
						EndpointReference{"http://192.168.1.102:5358/ws/km-wsdl/discovery"},
						EndpointReference{"http://[fe80::217:c8ff:fe7b:6a91]:5358/ws/km-wsdl/discovery"}},
					Types:     nil,
					ServiceID: "uri:4509a320-00a0-008f-00b6-002507510eca/KMWSDLService"}}},
	}

	if !reflect.DeepEqual(meta, expected) {
		t.Errorf("metadata content mismatch")
	}
}
