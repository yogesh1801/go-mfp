// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// Device information

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// devLocation represents the device location as Bus and Dev numbers
type devLocation struct {
	Bus int
	Dev int
}

// devInfo represents the device information
type devInfo struct {
	Location devLocation // Device location
	Device   *Device     // The device
}

// String returns string representation of the device location
func (loc devLocation) String() string {
	return fmt.Sprintf("%d-%d", loc.Bus, loc.Dev)
}

// makeDevBusID makes the new devBusID for the given device location.
func (loc devLocation) BusID() devBusID {
	var id devBusID
	copy(id[:31], loc.String())
	return id
}

// devBusID represents the device BusID, per the USBIP protocol.
type devBusID [32]byte

// String returns string representation of the devBusID
func (busid devBusID) String() string {
	s := busid[:]
	if n := bytes.IndexByte(s, 0); n >= 0 {
		s = s[:n]
	}
	return string(s)
}

// devPath represents the device path, per the USBIP protocol.
// It's role in the USBIP protocol is mostly informative.
type devPath [256]byte

// makeDevPath makes the new devPath or the given bus and
// device numbers
func makeDevPath(loc devLocation) devPath {
	// Obtain the hostname.
	host, err := os.Hostname()

	const maxHostname = 64

	switch {
	case err != nil:
		// If hostname can't be obtained, fallback to some
		// constant sting
		host = "UNKNOWN"

	case len(host) > maxHostname:
		// Hostname is too long. Try to cut at the DNS label boundary
		// with fallback to just cut at some arbitrary position.
		i := strings.IndexByte(host, '.')
		if i > 0 && i <= maxHostname {
			host = host[:i]
		} else {
			host = host[:maxHostname] + "..."
		}
	}

	s := fmt.Sprintf("/sys/devices/%s/usb%d/%d-%d",
		host, loc.Bus, loc.Bus, loc.Dev)

	var path devPath
	copy(path[:], s)

	return path
}
