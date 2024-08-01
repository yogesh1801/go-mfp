// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network state monitor -- the Linux version

package netstate

import (
	"os"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

// Poll period, if netlink socket is not available
const monitorPollPeriod = 5 * time.Second

// monitorLinus keeps track on a current network state and provides
// notifications when something changes.
//
// It contains Linux implementation of the monitor interface.
type monitorLinux struct {
	lock          sync.Mutex    // Access lock
	snapLast      snapshot      // Last known network state
	errLast       error         // Last error
	errSeq        int64         // Error sequence number
	waitchan      chan struct{} // Channel for clients to wait
	rtnetlinkFile *os.File      // rtnetlink socket as os.File
}

// newMonitor creates a network event monitor.
// Monitor is designed to run as a singleton shared between all users.
// Users should call getMonitor() instead.
func newMonitor() monitor {
	mon := &monitorLinux{
		waitchan: make(chan struct{}),
	}
	go mon.poll()

	return mon
}

// Get returns last known network state and channel to wait for updates.
//
// The returned channel will be closed by monitor when state changes.
func (mon *monitorLinux) Get() (snapshot, <-chan struct{}) {
	mon.lock.Lock()
	defer mon.lock.Unlock()

	return mon.snapLast, mon.waitchan
}

// GetError returns the latest error, if its sequence number
// is greater that supplied by the caller (i.e., caller has
// not seen this error yet). The returned error is wrapped into
// the EventError structure.
//
// If there is no new error, it returns nil.
//
// Additionally it returns a sequence number for the next call.
// The first call should use zero sequence number.
func (mon *monitorLinux) GetError(seq int64) (Event, int64) {
	mon.lock.Lock()
	defer mon.lock.Unlock()

	var evnt Event
	if seq < mon.errSeq {
		evnt = EventError{mon.errLast}
	}

	return evnt, mon.errSeq
}

// awake wakes all sleeping clients.
// It MUST be called under the mon.lock
func (mon *monitorLinux) awake() {
	close(mon.waitchan)
	mon.waitchan = make(chan struct{})
}

// update re-reads network state, updates monitor and awakes
// subscribers when appropriate.
func (mon *monitorLinux) update() {
	snapNext, err := newSnapshot()

	mon.lock.Lock()
	defer mon.lock.Unlock()

	if err != nil {
		mon.setError(err)
	} else if !mon.snapLast.Equal(snapNext) {
		mon.snapLast = snapNext
		mon.awake()
	}
}

// setError saves an error
func (mon *monitorLinux) setError(err error) {
	mon.lock.Lock()
	defer mon.lock.Unlock()

	if mon.errLast == nil || mon.errLast.Error() != err.Error() {
		mon.errLast = err
		mon.errSeq++
		mon.awake()
	}
}

// poll performs polling for network state changes.
func (mon *monitorLinux) poll() {
	for {
		// Try opening rtnetlink socket
		if mon.rtnetlinkFile == nil {
			err := mon.rtnetlinkOpen()
			if err != nil {
				mon.setError(err)
			}
		}

		// Try reading rtnetlink socket
		if mon.rtnetlinkFile != nil {
			err := mon.rtnetlinkRead()
			if err != nil {
				mon.rtnetlinkFile.Close()
				mon.rtnetlinkFile = nil
				mon.setError(err)
			}
		}

		// Fallback to timer-based polling
		if mon.rtnetlinkFile == nil {
			time.Sleep(monitorPollPeriod)
			mon.update()
		}
	}
}

// rtnetlinkRead reads and parses rtnetlink messages.
// If relevant event is received, it calls mon.update()
func (mon *monitorLinux) rtnetlinkRead() error {
	buf := make([]byte, 16384)

	n, err := mon.rtnetlinkFile.Read(buf)
	if err != nil {
		return err
	}

	messages, err := syscall.ParseNetlinkMessage(buf[0:n])
	if err != nil {
		return err
	}

	for _, msg := range messages {
		switch msg.Header.Type {
		case unix.RTM_NEWADDR, unix.RTM_DELADDR:
			mon.update()
			return nil
		}
	}

	return nil
}

// rtnetlinkOpen opens the rtnetlink file
func (mon *monitorLinux) rtnetlinkOpen() error {
	// Open rtnetlink socket
	sock, err := unix.Socket(unix.AF_NETLINK,
		unix.SOCK_RAW|unix.SOCK_CLOEXEC,
		unix.NETLINK_ROUTE)

	if err != nil {
		return err
	}

	// Subscribe to notifications
	var addr unix.SockaddrNetlink
	addr.Family = unix.AF_NETLINK
	addr.Groups = unix.RTMGRP_IPV4_IFADDR | unix.RTMGRP_IPV6_IFADDR

	err = unix.Bind(sock, &addr)
	if err != nil {
		unix.Close(sock)
		return err
	}

	// Wrap socket into os.File and return
	mon.rtnetlinkFile = os.NewFile(uintptr(sock), "rtnetlink")
	return nil
}
