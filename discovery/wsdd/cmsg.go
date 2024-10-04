// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UDP RecvFrom control message

package wsdd

import "net/netip"

// cmsg represents a parsed socket control message, as returned
// by RecvFrom by the UDP socket.
type cmsg struct {
	Dst     netip.Addr // Message destination address
	IfIndex int        // Network interface index
}
