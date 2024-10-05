// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UDP connection -- Linux-specific stuff

package wsdd

import (
	"fmt"
	"syscall"
)

// sysSetSockOptIP4 sets system-specific socket options for IPv4 connection
func (uc *uconn) sysSetSockOptIP4() error {
	return uc.control(func(fd int) error {
		// Bind socket to interface for outgoing multicasting
		idx := uc.local.Interface().Index()
		err := syscall.SetsockoptInt(fd, syscall.IPPROTO_IP,
			syscall.IP_MULTICAST_IF, idx)
		if err != nil {
			err = fmt.Errorf(
				"setsockopt(IP_MULTICAST_IF,%d):%w",
				idx, err)
			return err
		}

		// Set multicast TTL
		err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP,
			syscall.IP_TTL, 255)
		if err != nil {
			err = fmt.Errorf("setsockopt(IPPROTO_IP):%w", err)
			return err
		}

		return nil
	})
}

// sysSetSockOptIP6 sets system-specific socket options for IPv6 connection
func (uc *uconn) sysSetSockOptIP6() error {
	return uc.control(func(fd int) error {
		// Bind socket to interface for outgoing multicasting
		idx := uc.local.Interface().Index()
		err := syscall.SetsockoptInt(fd, syscall.IPPROTO_IPV6,
			syscall.IPV6_MULTICAST_IF, idx)
		if err != nil {
			err = fmt.Errorf(
				"setsockopt(IPV6_MULTICAST_IF,%d):%w",
				idx, err)
			return err
		}

		// Set multicast TTL
		//
		// FIXME. Looks like IPv6 doesn't have its own option
		// for the multicast TTL, but this fact requires
		// double checking.
		err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP,
			syscall.IP_TTL, 255)
		if err != nil {
			err = fmt.Errorf("setsockopt(IPPROTO_IP):%w", err)
			return err
		}

		return nil
	})
}
