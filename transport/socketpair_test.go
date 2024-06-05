// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UNIX socketpair test

package transport

import (
	"errors"
	"net"
	"os"
	"testing"
)

// TestSocketPair tests the socketpair function
func TestSocketPair(t *testing.T) {
	var c1, c2 net.Conn
	var err error

	testErr1 := errors.New("ERROR-1")
	testErr2 := errors.New("ERROR-2")
	testErr3 := errors.New("ERROR-3")

	saveSyscallSocketpair := hookSyscallSocketpair
	saveNetFileConn := hookNetFileConn

	// Cleanup after previous test
	cleanup := func() {
		if c1 != nil {
			c1.Close()
			c1 = nil
		}

		if c2 != nil {
			c2.Close()
			c2 = nil
		}

		err = nil
		hookSyscallSocketpair = saveSyscallSocketpair
		hookNetFileConn = saveNetFileConn
	}

	// socketpair must succeed
	c1, c2, err = socketpair()
	if c1 == nil || c2 == nil || err != nil {
		t.Errorf("socketpair() failed: c1=%v c2=%v err=%v", c1, c2, err)
	}

	cleanup()

	// syscall.Socketpair return error
	hookSyscallSocketpair = func() (fds [2]int, err error) {
		err = testErr1
		return
	}

	c1, c2, err = socketpair()
	if c1 != nil || c2 != nil || err != testErr1 {
		t.Errorf("socketpair():\n"+
			"expected: c1=%v c2=%v err=%v\n"+
			"present:  c1=%v c2=%v err=%v\n",
			nil, nil, testErr1, c1, c2, err)
	}

	cleanup()

	// First call to net.FileConn fails
	hookNetFileConn = func(f *os.File) (net.Conn, error) {
		return nil, testErr2
	}

	c1, c2, err = socketpair()

	if c1 != nil || c2 != nil || err != testErr2 {
		t.Errorf("socketpair():\n"+
			"expected: c1=%v c2=%v err=%v\n"+
			"present:  c1=%v c2=%v err=%v\n",
			nil, nil, testErr2, c1, c2, err)
	}

	cleanup()

	// Second call to net.FileConn fails
	firstCall := true
	hookNetFileConn = func(f *os.File) (net.Conn, error) {
		if firstCall {
			firstCall = false
			return saveNetFileConn(f)
		}
		return nil, testErr3
	}

	c1, c2, err = socketpair()

	if c1 != nil || c2 != nil || err != testErr3 {
		t.Errorf("socketpair():\n"+
			"expected: c1=%v c2=%v err=%v\n"+
			"present:  c1=%v c2=%v err=%v\n",
			nil, nil, testErr3, c1, c2, err)
	}

	cleanup()
}
