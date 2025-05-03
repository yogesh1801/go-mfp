// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Loopback interface index

package dnssd

import (
	"sync"

	"github.com/OpenPrinting/go-avahi"
)

var (
	// loopback contains index of the loopback interface
	//
	// It is initialized on demand, when backend is created, so all
	// backend code may assume this value is valid.
	loopback avahi.IfIndex = -1

	// loopbackInitLock is the initialization lock
	loopbackInitLock sync.Mutex
)

// loopbackInit initializes loopback
func loopbackInit() error {
	loopbackInitLock.Lock()
	defer loopbackInitLock.Unlock()

	v, err := avahi.Loopback()
	if err == nil {
		loopback = v
	}

	return err
}
