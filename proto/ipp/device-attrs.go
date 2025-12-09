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

// DeviceAttributes represents device attributes, as returned by
// the CUPS-Get-Devices request
type DeviceAttributes struct {
	ObjectRawAttrs
	CUPSDeviceAttributesGroup

	DeviceClass        optional.Val[KwDeviceClass] `ipp:"device-class"`
	DeviceInfo         optional.Val[string]        `ipp:"device-info,text(127)"`
	DeviceMakeAndModel optional.Val[string]        `ipp:"device-make-and-model,text(127)"`
	DeviceURI          optional.Val[string]        `ipp:"device-uri,uri"`
	DeviceID           optional.Val[string]        `ipp:"device-id,text"`
	DeviceLocation     optional.Val[string]        `ipp:"device-location,text(127)"`
}
