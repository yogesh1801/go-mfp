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
	"time"
)

// WSDD protocol parameters:
var (
	// WSDD IPv4 multicast group address
	wsddMulticastIP4 = netip.MustParseAddrPort("239.255.255.250:3702")

	// WSDD IPv6 multicast group address
	wsddMulticastIP6 = netip.MustParseAddrPort("[ff02::c]:3702")

	// Timeout for the metadata Get request (performed via HTTP)
	wsddMetadataGetTimeout = 5 * time.Second

	// Response size limit for the metadata Get request (to mitigate
	// possible DOS attack)
	wsddMetadataGetMaxResponse = 64536
)
