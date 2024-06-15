// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Low-level functions for connections

package transport

import "net"

// connWithSetLinger denotes net.Conn with SetLinger method.
type connWithSetLinger interface {
	SetLinger(sec int) error
}

// connAbort closes connection abortively.
func connAbort(conn net.Conn) {
	if withSetLinger, ok := conn.(connWithSetLinger); ok {
		withSetLinger.SetLinger(0)
	}

	conn.Close()
}
