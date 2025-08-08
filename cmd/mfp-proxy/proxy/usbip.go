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

	srv, endpoints := usbip.NewIPPUSB(ctx, 3, handler)

	desc := usbip.USBDeviceDescriptor{
		BCDUSB:          0x0200,
		Speed:           usbip.USBSpeedHigh,
		BDeviceClass:    0,
		BDeviceSubClass: 0,
		BDeviceProtocol: 0,
		BMaxPacketSize:  64,
		IDVendor:        0xdead,
		IDProduct:       0xbeaf,
		BCDDevice:       0x0100,
		IManufacturer:   "OpenPrinting",
		IProduct:        "Virtual MFP",
		ISerialNumber:   "NN-001122334455",
		Configurations: []usbip.USBConfigurationDescriptor{{
			BMAttributes: usbip.USBConfAttrSelfPowered,
			MaxPower:     1,
			Interfaces: []usbip.USBInterface{
				{
					AltSettings: []usbip.USBInterfaceDescriptor{
						{
							BInterfaceClass:    7,
							BInterfaceSubClass: 1,
							BInterfaceProtocol: 4,
							Endpoints: []*usbip.Endpoint{
								endpoints[0],
							},
						},
					},
				},
				{
					AltSettings: []usbip.USBInterfaceDescriptor{
						{
							BInterfaceClass:    7,
							BInterfaceSubClass: 1,
							BInterfaceProtocol: 4,
							Endpoints: []*usbip.Endpoint{
								endpoints[1],
							},
						},
					},
				},
				{
					AltSettings: []usbip.USBInterfaceDescriptor{
						{
							BInterfaceClass:    7,
							BInterfaceSubClass: 1,
							BInterfaceProtocol: 4,
							Endpoints: []*usbip.Endpoint{
								endpoints[2],
							},
						},
					},
				},
			},
		}},
	}

	dev, err := usbip.NewDevice(desc)
	assert.NoError(err)

	usbipSrv := usbip.NewServer(ctx)
	err = usbipSrv.AddDevice(dev)
	assert.NoError(err)

	go usbipSrv.ListenAndServe(addr)

	return srv
}
