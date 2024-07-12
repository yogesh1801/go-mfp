// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network evens notifications -- the Linux version

package netstate

import (
	"context"
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/alexpevzner/mfp/log"
	"golang.org/x/sys/unix"
)

// Poll period, if netlink socket is not available
const notifierPollPeriod = 5 * time.Second

// Notifier provides notifications on network state changes.
type Notifier struct {
	ctx              context.Context // The Context
	interfaces       []net.Interface // Network interfaces
	rtnetlinkFile    *os.File        // rtnetlink socket as os.File
	rtnetkinkLastErr error           // Last rtnetlink error
	doneWaitLock     sync.Mutex      // Notifier.doneWait lock
}

// NewNotifier creates a new network event notifier.
func NewNotifier(ctx context.Context) *Notifier {
	not := &Notifier{ctx: ctx}
	go not.poll()
	return not
}

// update re-reads network state, updates Notifier and sends events
// to subscribers.
func (not *Notifier) update() {
}

// poll performs polling for network state changes.
func (not *Notifier) poll() {
	log.Debug(not.ctx, "netevent: started network montoting")

	timer := time.NewTimer(time.Hour)
	timer.Stop()

	for not.ctx.Err() == nil {
		// Try opening rtnetlink socket
		if not.rtnetlinkFile == nil {
			err := not.rtnetlinkOpen()
			if err != nil {
				not.rtnetlinkError(err)
			}
		}

		// Try reading rtnetlink socket
		if not.rtnetlinkFile != nil {
			err := not.rtnetlinkRead()
			if err != nil {
				not.rtnetlinkError(err)
			}
		}

		// Fallback to timer-based polling
		if not.rtnetlinkFile == nil {
			timer.Reset(notifierPollPeriod)

			select {
			case <-timer.C:
				not.update()
			case <-not.ctx.Done():
				// Stop the timer, so it will be immediately
				// eligible for garbage collection
				timer.Stop()
			}
		}
	}
}

// doneWait waits for [context.Context], used when [Notifier] was
// created, to be canceled, then terminates Notifier.poll
func (not *Notifier) doneWait() {
	<-not.ctx.Done()

	not.doneWaitLock.Lock()
	not.rtnetlinkFile.Close()
	not.doneWaitLock.Unlock()
}

// rtnetlinkError logs rtnetlink error and performs appropriate actions
func (not *Notifier) rtnetlinkError(err error) {
	if err != not.rtnetkinkLastErr {
		if not.rtnetlinkFile != nil {
			not.rtnetlinkFile.Close()
			not.rtnetlinkFile = nil
		}

		not.rtnetkinkLastErr = err
		log.Warning(not.ctx, "netevent: rtnetlink: %s", err)
	}
}

// rtnetlinkRead reads and parses rtnetlink messages.
// If relevant event is received, it calls not.update()
func (not *Notifier) rtnetlinkRead() error {
	buf := make([]byte, 16384)

	n, err := not.rtnetlinkFile.Read(buf)
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
			not.update()
			return nil
		}
	}

	return nil
}

// rtnetlinkOpen opens the rtnetlink file
func (not *Notifier) rtnetlinkOpen() error {
	// Synchronize with Notifier.doneWait
	not.doneWaitLock.Lock()
	defer not.doneWaitLock.Unlock()

	err := not.ctx.Err()
	if err != nil {
		return err
	}

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
	not.rtnetlinkFile = os.NewFile(uintptr(sock), "rtnetlink")
	return nil
}
