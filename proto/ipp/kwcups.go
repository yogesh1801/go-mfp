// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS-specific IPP keywords

package ipp

// KwDeviceClass represents known values of the CUPS' "device-class"
// attribute, used by CUPS-Get-Devices.
//
// See [CUPS Implementation of IPP] for details.
//
// [CUPS Implementation of IPP]: https://www.cups.org/doc/spec-ipp.html
type KwDeviceClass string

const (
	// KwDeviceClassFile is a disk file.
	KwDeviceClassFile KwDeviceClass = "file"

	// KwDeviceClassDirect is a directly connected device, i.e.,
	// parallel or fixed-rate serial port, Centronics, IEEE-1284,
	// and USB printer ports.
	KwDeviceClassDirect KwDeviceClass = "direct"

	// KwDeviceClassSerial is a variable-rate serial port.
	KwDeviceClassSerial KwDeviceClass = "serial"

	// KwDeviceClassNetwork is a network connection, typically via
	// AppSocket, HTTP, IPP, LPD, or SMB/CIFS protocols.
	//
	// As a special exception, IPP over USB also belongs to this
	// class, because this protocol uses a pseudo-network connection.
	KwDeviceClassNetwork KwDeviceClass = "network"
)
