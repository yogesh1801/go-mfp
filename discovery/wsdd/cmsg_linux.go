// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UDP RecvFrom control message -- the Linux version

package wsdd

// Parse parses the control message from the oob data, returned
// by [net.UDPConn.ReadMsgUDP]
func (cm *cmsg) Parse(oob []byte) error {
	return nil
}
