// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UDP multicasting

package wsdd

import (
	"fmt"
	"net"
	"net/netip"
	"sync/atomic"
	"syscall"

	"github.com/alexpevzner/mfp/discovery/netstate"
)

// mconn wraps net.UDPConn and prepares it to be used for
// the UDP multicasts reception.
type mconn struct {
	*net.UDPConn                // Underlying UDP connection
	group        netip.AddrPort // Multicast group
	closed       atomic.Bool    // Connection is closed
}

// newMconn creates a new multicast connection
func newMconn(group netip.AddrPort) (*mconn, error) {
	// Address must be multicast
	if !group.Addr().IsMulticast() {
		err := fmt.Errorf("%s not multicast", group.Addr())
		return nil, err
	}

	// Prepare net.UDPAddr structure
	addr := &net.UDPAddr{
		IP:   net.IP(group.Addr().AsSlice()),
		Port: int(group.Port()),
		Zone: group.Addr().Zone(),
	}

	// Open UDP connection.
	//
	// Note, with the multicast address being given,
	// net.ListenUDP creates UDP socket bound to the
	// 0.0.0.0:port (or [::0]:port) address with
	// SO_REUSEADDR option being set.
	//
	// This socket can be joined multiple multicast
	// groups and suitable for the multicast reception.
	network := "udp4"
	if group.Addr().Is6() {
		network = "udp6"
	}

	conn, err := net.ListenUDP(network, addr)
	if err != nil {
		return nil, err
	}

	// Fill the mconn structure
	mc := &mconn{
		UDPConn: conn,
		group:   group,
	}

	// Do system-specific setup
	if mc.Is4() {
		err = mc.control(cmsgSetSockOptIP4)
	} else {
		err = mc.control(cmsgSetSockOptIP6)
	}

	if err != nil {
		mc.Close()
		return nil, err
	}

	return mc, nil
}

// Close closes the connection
func (mc *mconn) Close() {
	mc.closed.Store(true)
	mc.UDPConn.Close()
}

// LocalAddrPort returns connection's local address and port
func (mc *mconn) LocalAddrPort() netip.AddrPort {
	return mc.group
}

// IsClosed reports if connection is closed
func (mc *mconn) IsClosed() bool {
	return mc.closed.Load()
}

// Is4 reports if connection uses IPv4 address family
func (mc *mconn) Is4() bool {
	return mc.group.Addr().Is4()
}

// Is6 reports if connection uses IPv6 address family
func (mc *mconn) Is6() bool {
	return mc.group.Addr().Is6()
}

// Join joins the multicast group, specified during mcast
// creation, on a network interface, specified by the local
// parameter.
func (mc *mconn) Join(local netstate.Addr) error {
	if mc.Is6() {
		return mc.joinIP6(local)
	}
	return mc.joinIP4(local)
}

// Leave leaves the multicast group, specified during mcast
// creation, on a network interface, specified by the local
// parameter.
func (mc *mconn) Leave(local netstate.Addr) error {
	if mc.Is6() {
		return mc.leaveIP6(local)
	}
	return mc.leaveIP4(local)
}

// RecvFrom receives a packet from the UDP connection
func (mc *mconn) RecvFrom(b []byte) (n int, from netip.AddrPort,
	cmsg cmsg, err error) {

	var oob [8192]byte

	n, ooblen, _, from, err := mc.UDPConn.ReadMsgUDPAddrPort(b, oob[:])
	if err != nil {
		return
	}

	from = netip.AddrPortFrom(from.Addr().Unmap(), from.Port())

	msgs, err := syscall.ParseSocketControlMessage(oob[:ooblen])
	if err != nil {
		return
	}

	if mc.Is4() {
		cmsg = cmsgParseIP4(msgs)
	} else {
		cmsg = cmsgParseIP6(msgs)
	}

	return
}

// joinIP4 is the mcast.Join for IP4 connections
func (mc *mconn) joinIP4(local netstate.Addr) error {
	if !mc.Is4() {
		err := fmt.Errorf("Can't join IP4 group on IP6 connection")
		return err
	}

	mreq := syscall.IPMreqn{
		Multiaddr: mc.group.Addr().As4(),
		Address:   local.Addr().As4(),
		Ifindex:   int32(local.Interface().Index()),
	}

	err := mc.control(func(fd int) error {
		return syscall.SetsockoptIPMreqn(fd, syscall.IPPROTO_IP,
			syscall.IP_ADD_MEMBERSHIP, &mreq)
	})

	return err
}

// joinIP6 is the mcast.Join for IP6 connections
func (mc *mconn) joinIP6(local netstate.Addr) error {
	if !mc.Is6() {
		err := fmt.Errorf("Can't join IP4 group on IP6 connection")
		return err
	}

	mreq := syscall.IPv6Mreq{
		Multiaddr: mc.group.Addr().As16(),
		Interface: uint32(local.Interface().Index()),
	}

	err := mc.control(func(fd int) error {
		return syscall.SetsockoptIPv6Mreq(fd, syscall.IPPROTO_IPV6,
			syscall.IPV6_JOIN_GROUP, &mreq)
	})

	return err
}

// leaveIP4 is the mcast.Leave for IP4 connections
func (mc *mconn) leaveIP4(local netstate.Addr) error {
	if !mc.Is4() {
		err := fmt.Errorf("Can't leave IP4 group on IP6 connection")
		return err
	}

	mreq := syscall.IPMreqn{
		Multiaddr: mc.group.Addr().As4(),
		Address:   local.Addr().As4(),
		Ifindex:   int32(local.Interface().Index()),
	}

	err := mc.control(func(fd int) error {
		return syscall.SetsockoptIPMreqn(fd, syscall.IPPROTO_IP,
			syscall.IP_DROP_MEMBERSHIP, &mreq)
	})

	return err
}

// leaveIP6 is the mcast.Leave for IP6 connections
func (mc *mconn) leaveIP6(local netstate.Addr) error {
	if !mc.Is6() {
		err := fmt.Errorf("Can't leave IP4 group on IP6 connection")
		return err
	}

	mreq := syscall.IPv6Mreq{
		Multiaddr: mc.group.Addr().As16(),
		Interface: uint32(local.Interface().Index()),
	}

	err := mc.control(func(fd int) error {
		return syscall.SetsockoptIPv6Mreq(fd, syscall.IPPROTO_IPV6,
			syscall.IPV6_LEAVE_GROUP, &mreq)
	})

	return err
}

// control invokes f on the underlying connection's
// file descriptor.
func (mc *mconn) control(f func(fd int) error) error {
	rawconn, err := mc.SyscallConn()
	if err != nil {
		return err
	}

	var err2 error
	err = rawconn.Control(func(fd uintptr) {
		err2 = f(int(fd))
	})

	if err != nil {
		return err
	}

	return err2
}
