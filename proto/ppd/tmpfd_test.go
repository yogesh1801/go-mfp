// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// Temporary files test

package ppd

import (
	"bytes"
	"io"
	"syscall"
	"testing"
)

// Make sure tmpFD implements all expected interfaces
var (
	_ io.Closer = tmpFD(0)
	_ io.Reader = tmpFD(0)
	_ io.Writer = tmpFD(0)
	_ io.Seeker = tmpFD(0)
)

// TestTmpFD tests the tmpFD implementation
func TestTmpFD(t *testing.T) {
	// Open the file
	fd, err := tmpFDOpen("test.file")
	if err != nil {
		t.Errorf("tmpFDOpen: %s", err)
		return
	}

	defer func() {
		if fd >= 0 {
			fd.Close()
		}
	}()

	// Write data
	data := []byte("hello, world")
	n, err := fd.Write(data)
	if err != nil {
		t.Errorf("tmpFD.Write: %s", err)
		return
	}

	if n != len(data) {
		t.Errorf("tmpFD.Write: lengths mismatch:\n"+
			"expected: %d\n"+
			"present:  %d",
			len(data), n)
		return
	}

	// Seek here and there
	off, err := fd.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Errorf("tmpFD.Seek: %s", err)
		return
	}

	if off != int64(len(data)) {
		t.Errorf("tmpFD.Seek(io.SeekCurrent): offset mismatch:\n"+
			"expected: %d\n"+
			"present:  %d",
			len(data), off)
		return
	}

	off, err = fd.Seek(-1, io.SeekEnd)
	if err != nil {
		t.Errorf("tmpFD.Seek: %s", err)
		return
	}

	if off != int64(len(data)-1) {
		t.Errorf("tmpFD.Seek(io.SeekEnd): offset mismatch:\n"+
			"expected: %d\n"+
			"present:  %d",
			len(data), off)
		return
	}

	off, err = fd.Seek(0, io.SeekStart)
	if err != nil {
		t.Errorf("tmpFD.Seek: %s", err)
		return
	}

	if off != 0 {
		t.Errorf("tmpFD.Seek(io.SeekStart): offset mismatch:\n"+
			"expected: %d\n"+
			"present:  %d",
			0, off)
		return
	}

	_, err = fd.Seek(0, 12345)
	if err != syscall.EINVAL {
		t.Errorf("tmpFD.Seek(invalid whence): error mismatch:\n"+
			"expected: %v\n"+
			"present:  %v",
			syscall.EINVAL, err)
		return
	}

	// Read data
	buf := make([]byte, 1024)
	n, err = fd.Read(buf)

	if err != nil {
		t.Errorf("tmpFD.Read: %s", err)
		return
	}

	if n != len(data) {
		t.Errorf("tmpFD.Write: lengths mismatch:\n"+
			"expected: %d\n"+
			"present:  %d",
			len(data), n)
		return
	}

	buf = buf[:n]
	if !bytes.Equal(data, buf) {
		t.Errorf("tmpFD.Write: data mismatch:\n"+
			"expected: %q\n"+
			"present:  %q",
			data, buf)
		return
	}

	// Close file
	err = fd.Close()
	fd = -1 // Prevent double close

	if err != nil {
		t.Errorf("tmpFD.Close: %s", err)
	}
}
