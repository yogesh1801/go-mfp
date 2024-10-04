// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UDP RecvFrom control message -- the Linux version

package wsdd

// cmsgSetSockOpt sets system-specific socket options
// to enable control messages reception on an UDP socket
func cmsgSetSockOpt(fd int, ip6 bool) error {
	return nil
}

// Parse parses the control message from the oob data, returned
// by [net.UDPConn.ReadMsgUDP]
func cmsgParse(oob []byte) (cmsg, error) {
	return cmsg{}, nil
}
