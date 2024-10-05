// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UDP RecvFrom control message -- the Linux version

package wsdd

import (
	"syscall"
	"unsafe"
)

// cmsgSetSockOpt sets system-specific socket options on IPv4
// UDP socket to enable control messages reception on an UDP socket
func cmsgSetSockOptIP4(fd int) error {
	err := syscall.SetsockoptInt(fd, syscall.IPPROTO_IP,
		syscall.IP_PKTINFO, 1)
	return err
}

// cmsgSetSockOpt sets system-specific socket options on IPv6
// UDP socket to enable control messages reception on an UDP socket
func cmsgSetSockOptIP6(fd int) error {
	err := syscall.SetsockoptInt(fd, syscall.IPPROTO_IPV6,
		syscall.IPV6_RECVPKTINFO, 1)

	return err
}

// Parse parses the control message from the oob data, returned
// by [net.UDPConn.ReadMsgUDP] for the IPv4 UDP socket
//
// It returns cmsg{}, if control message is missed or cannot
// be decoded.
func cmsgParseIP4(msgs []syscall.SocketControlMessage) cmsg {
	for _, msg := range msgs {
		hdr := msg.Header
		if hdr.Level == syscall.IPPROTO_IP &&
			hdr.Type == syscall.IP_PKTINFO &&
			len(msg.Data) >= syscall.SizeofInet4Pktinfo {

			var pktinfo syscall.Inet4Pktinfo
			p := (*[syscall.SizeofInet4Pktinfo]byte)(
				unsafe.Pointer(&pktinfo))[:]

			copy(p, msg.Data)

			cm := cmsg{
				IfIndex: int(pktinfo.Ifindex),
			}

			return cm
		}
	}
	return cmsg{}
}

// Parse parses the control message from the oob data, returned
// by [net.UDPConn.ReadMsgUDP] for the IPv6 UDP socket
//
// It returns cmsg{}, if control message is missed or cannot
// be decoded.
func cmsgParseIP6(msgs []syscall.SocketControlMessage) cmsg {
	for _, msg := range msgs {
		hdr := msg.Header
		if hdr.Level == syscall.IPPROTO_IPV6 &&
			hdr.Type == syscall.IPV6_PKTINFO &&
			len(msg.Data) >= syscall.SizeofInet6Pktinfo {

			var pktinfo syscall.Inet6Pktinfo
			p := (*[syscall.SizeofInet6Pktinfo]byte)(
				unsafe.Pointer(&pktinfo))[:]

			copy(p, msg.Data)

			cm := cmsg{
				IfIndex: int(pktinfo.Ifindex),
			}

			return cm
		}
	}
	return cmsg{}
}
