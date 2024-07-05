// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// DNS-SD service description

package dnssd

import "net"

type ServiceCommon struct {
	InstanceName string      // E.g., "My printer"
	Type         string      // E.g., "_ipp._tcp"
	Subtypes     []string    // E.g., "_print._sub._ipp._tcp"
	Addr         net.IP      // Service IP address
	Port         int         // Service IP port, 0 if service is disabled
	TXT          []TxtRecord // TXT records
}
