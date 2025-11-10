// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// Temporary files

package ppd

import (
	"io"
	"syscall"
	"unsafe"
)

// #define _GNU_SOURCE
// #include <stdlib.h>
// #include <unistd.h>
// #include <sys/mman.h>
import "C"

// tmpFD represents the temporary file.
//
// tmpFD is the valid OS file handle (implemented via memfd_create).
//
// In the same time, tmpFD implements the [io.Closer], [io.Reader],
// [io.Writer] and [io.Seeker] interfaces.
//
// Once tmpFD is closed, its content is removed.
type tmpFD int

// tmpFDOpen creates the new temporary file.
//
// The name argument is passed to the memfd_create system call
// and barely used for the debugging purposes.
func tmpFDOpen(name string) (tmpFD, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	fd, err := C.memfd_create(cname, C.MFD_CLOEXEC)
	return tmpFD(fd), err
}

// Close implements the [io.Closer] interface for the tmpFD.
func (fd tmpFD) Close() error {
	return syscall.Close(int(fd))
}

// Read implements the [io.Reader] interface for the tmpFD.
func (fd tmpFD) Read(buf []byte) (int, error) {
	return syscall.Read(int(fd), buf)
}

// Write implements the [io.Writer] interface for the tmpFD.
func (fd tmpFD) Write(buf []byte) (int, error) {
	return syscall.Write(int(fd), buf)
}

// Seek implements the [io.Seeker] interface for the tmpFD.
func (fd tmpFD) Seek(off int64, whence int) (int64, error) {
	// Map whence
	switch whence {
	case io.SeekStart:
		whence = int(C.SEEK_SET)
	case io.SeekCurrent:
		whence = int(C.SEEK_CUR)
	case io.SeekEnd:
		whence = int(C.SEEK_END)
	default:
		return -1, syscall.EINVAL
	}

	// Call syscall.Seek
	return syscall.Seek(int(fd), off, whence)
}
