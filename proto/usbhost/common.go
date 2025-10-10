// MFP - Miulti-Function Printers and scanners toolkit
// USB host API
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common definitions

package usbhost

import "github.com/OpenPrinting/go-mfp/proto/usb"

// Location represents the device location as Bus and Dev numbers
type Location struct {
	Bus int
	Dev int
}

// DeviceInfo contains USB device location and descriptor
type DeviceInfo struct {
	Loc  Location             // Device location
	Desc usb.DeviceDescriptor // Device descriptor
}
