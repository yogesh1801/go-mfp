// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// Convenience functions for network connections

package main

import (
	"io"
	"net"
)

// connReadAll reads the full buffer from the network connection.
// It returns the io.ErrUnexpectedEOF error, if connection was closed
// by peer before all data is received.
func connReadAll(conn net.Conn, buf []byte) error {
	for len(buf) > 0 {
		n, err := conn.Read(buf)
		switch {
		case n > 0:
			// Some data was received. Ignore the error, if any
			buf = buf[n:]

		case err == io.EOF:
			// Connection has closed prematurely
			fallthrough

		case err != nil:
			// Read error
			return err
		}
	}

	return nil
}

// connWriteAll writes the whole buffer into the network connection.
func connWriteAll(conn net.Conn, buf []byte) error {
	for len(buf) > 0 {
		n, err := conn.Write(buf)
		switch {
		case n > 0:
			// Some data was sent. Ignore the error, if any
			buf = buf[n:]

		case err != nil:
			// Write error
			return err
		}
	}

	return nil
}
