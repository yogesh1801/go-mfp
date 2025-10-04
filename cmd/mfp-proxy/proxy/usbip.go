// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by Mohammed Imaduddin (mdimad005@gmail.com)
// See LICENSE for license terms and conditions
//
// IPP over USB emulation logic

package proxy

import (
	"context"
	"net"
	"net/http"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/modeling/defaults"
	"github.com/OpenPrinting/go-mfp/proto/usbip"
	"github.com/OpenPrinting/go-mfp/transport"
)

// newUsbipServer creates the new USBIP server representing
// an IPP over USB MFP device.
//
// The server accepts incoming USBIP connection on a provided
// address and forwards incoming IPP over USB requests (which
// are essentially the HTTP requests) to the provided http.Handler.
func newUsbipServer(ctx context.Context,
	addr net.Addr, handler http.Handler) *transport.Server {

	// Obtain device descriptor and its endpoints
	desc := defaults.USBIPPDescriptor()

	// Create USB device.
	dev := usbip.MustNewDevice(desc)

	// IPP over USB is the HTTP-based protocol.
	// Create HTTP server on a top of the USB device endpoints.
	srv := transport.NewServer(ctx, nil, handler)

	endpoints := dev.EndpointsByClass(7, 1, 4)
	listener := usbip.NewEndpointListener(addr, addr, endpoints)

	go srv.Serve(listener)

	// Create USBIP server
	usbipSrv := usbip.NewServer(ctx)
	err := usbipSrv.AddDevice(dev)
	assert.NoError(err)

	go usbipSrv.ListenAndServe(addr)

	return srv
}
