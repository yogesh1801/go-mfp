// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Protocol parameters

package wsdd

import (
	"net/netip"
)

// WSDD protocol parameters:
var (
	// WSDD IPv4 multicast group address
	wsddMulticastIP4 = netip.MustParseAddrPort("239.255.255.250:3702")

	// WSDD IPv6 multicast group address
	wsddMulticastIP6 = netip.MustParseAddrPort("[ff02::c]:3702")
)
