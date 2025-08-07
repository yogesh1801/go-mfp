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
	"net/http"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/proto/usbip"
	"github.com/OpenPrinting/go-mfp/transport"
)

func newUsbipServer(ctx context.Context,
	handler http.Handler) *transport.Server {

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
		Manufacturer:    "OpenPrinting",
		Product:         "Virtual MFP",
		SerialNumber:    "NN-001122334455",
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

	usbipSrv := usbip.NewServer()
	err = usbipSrv.AddDevice(dev)
	assert.NoError(err)

	go usbipSrv.Run(ctx, "localhost", 3240)

	return srv
}
