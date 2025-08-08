// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// IPP-USB device implementation

package usbip

import (
	"context"
	"net"
	"net/http"

	"github.com/OpenPrinting/go-mfp/transport"
)

// ippusbAddr returned as local and remote address for the IPP over USB connections.
var ippusbAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 80, Zone: ""}

// NewIPPUSB creates a new IPP over USB server (represented by the
// [http.Server] and its [Endpoint]s
func NewIPPUSB(ctx context.Context,
	numendpoints int,
	handler http.Handler) (*transport.Server, []*Endpoint) {

	endpoints := make([]*Endpoint, numendpoints)
	for i := range endpoints {
		endpoints[i] = NewEndpoint(EndpointInOut, USBXferBulk, 512)
	}

	srv := transport.NewServer(ctx, nil, handler)
	listener := NewEndpointListener(ippusbAddr, ippusbAddr, endpoints)
	go srv.Serve(listener)

	return srv, endpoints
}
