// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device attribites, as returned by CUPS-Get-Devices

package ipp

// DeviceAttributes represents device attributes, as returned by
// the CUPS-Get-Devices request
type DeviceAttributes struct {
	ObjectRawAttrs

	DeviceClass        KwDeviceClass `ipp:"?device-class"`
	DeviceInfo         string        `ipp:"?device-info,text"`
	DeviceMakeAndModel string        `ipp:"?device-make-and-model,text"`
	DeviceURI          string        `ipp:"?device-uri,uri"`
	DeviceID           string        `ipp:"device-id,text"`
	DeviceLocation     string        `ipp:"device-location,text"`
}

// KnownAttrs returns information about all known IPP attributes
// of the DeviceAttributes
func (attrs *DeviceAttributes) KnownAttrs() []AttrInfo {
	return ippKnownAttrs(attrs)
}
