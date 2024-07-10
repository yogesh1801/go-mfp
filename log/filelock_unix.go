// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// File locking, the UNIX way

package log

import (
	"os"
	"syscall"
)

// FileLock represents an acquired file lock.
type FileLock struct {
	fd int
}

// FileLockEx acquires an exclusive file lock on file, pointed by [os.File].
//
// If lock already held by another process, this function waits until
// lock is released.
func FileLockEx(file *os.File) (*FileLock, error) {
	fd := int(file.Fd())

	err := syscall.Flock(fd, syscall.LOCK_EX)
	if err != nil {
		return nil, err
	}

	return &FileLock{fd}, nil
}

// Close releases the lock, previously taken by [FileLockEx].
func (fl *FileLock) Close() error {
	return syscall.Flock(fl.fd, syscall.LOCK_EX)
}
