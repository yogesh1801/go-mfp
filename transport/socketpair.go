// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UNIX socketpair

package transport

import (
	"net"
	"os"
	"syscall"
)

// This is very hard to make these functions to fail in a real life,
// so wrap them into hooks, for testing.
var (
	hookSyscallSocketpair = func() ([2]int, error) {
		return syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	}

	hookNetFileConn = func(f *os.File) (net.Conn, error) {
		return net.FileConn(f)
	}
)

// socketpair creates a pair of interconnected UNIX sockets.
func socketpair() (c1, c2 net.Conn, err error) {
	// Call socketpair(2) syscall
	fds, err := hookSyscallSocketpair()
	if err != nil {
		return
	}

	// Wrap into os.File
	//
	// Note, net.FileConn() will duplicate underlying sockets,
	// so closing fd1 and fd2 is our responsibility.
	fd1 := os.NewFile(uintptr(fds[0]), "socketpair-0")
	fd2 := os.NewFile(uintptr(fds[0]), "socketpair-1")

	defer fd1.Close()
	defer fd2.Close()

	// Now wrap into net.Conn
	c1, err = hookNetFileConn(fd1)
	if err != nil {
		return
	}

	c2, err = hookNetFileConn(fd2)
	if err != nil {
		c1.Close()
		c1 = nil
	}

	return
}
