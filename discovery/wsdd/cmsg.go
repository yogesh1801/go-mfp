// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UDP RecvFrom control message

package wsdd

// cmsg represents a parsed socket control message, as returned
// by RecvFrom by the UDP socket.
type cmsg struct {
	IfIndex int // Network interface index
}
