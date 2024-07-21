// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Hooks for testing

package netstate

import (
	"net"
)

// Overriding these hooks allows tests to simulate
// running on real OS:
var (
	// net.Interfaces
	hookNetInterfaces = net.Interfaces

	// net.Interface.Addrs
	hookNetInterfacesAddrs = (*net.Interface).Addrs
)
