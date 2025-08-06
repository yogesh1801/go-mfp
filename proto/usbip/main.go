// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by Mohammed Imaduddin (mdimad005@gmail.com)
// See LICENSE for license terms and conditions
//
// IPP over USB emulation logic

package main

import (
	"context"
	"fmt"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/log"
)

// Config holds settings used to initialize the virtual IPP-over-USB device.
type Config struct {
	IPPServerURL string `json:"ipp_server_url"`
	DeviceName   string `json:"device_name"`
	VendorID     string `json:"vendor_id"`
	ProductID    string `json:"product_id"`
	Manufacturer string `json:"manufacturer"`
	Product      string `json:"product"`
	Serial       string `json:"serial"`
	ListenIP     string `json:"listen_ip"`
	ListenPort   int    `json:"listen_port"`
	Debug        bool   `json:"debug"`
}

func main() {
	logger := log.NewLogger(log.LevelDebug, log.Console)
	ctx := log.NewContext(context.Background(), logger)

	_, endpoints := NewIPPUSB(2)

	desc := USBDeviceDescriptor{
		BCDUSB:          0x0200,
		Speed:           USBSpeedHigh,
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
		Configurations: []USBConfigurationDescriptor{{
			BMAttributes: USBConfAttrSelfPowered,
			MaxPower:     1,
			Interfaces: []USBInterface{
				{
					[]USBInterfaceDescriptor{
						{
							BInterfaceClass:    7,
							BInterfaceSubClass: 1,
							BInterfaceProtocol: 4,
							Endpoints: []*Endpoint{
								endpoints[0],
							},
						},
					},
				},
				{
					[]USBInterfaceDescriptor{
						{
							BInterfaceClass:    7,
							BInterfaceSubClass: 1,
							BInterfaceProtocol: 4,
							Endpoints: []*Endpoint{
								endpoints[1],
							},
						},
					},
				},
			},
		}},
	}

	dev, err := NewDevice(desc)
	assert.NoError(err)

	srv := NewServer()
	err = srv.AddDevice(dev)
	assert.NoError(err)

	// Get listen settings from config
	listenIP := "localhost"
	listenPort := 3240

	fmt.Printf("Listening on %s:%d\n", listenIP, listenPort)
	fmt.Println("Press Ctrl+C to stop")

	srv.Run(ctx, listenIP, listenPort)
}
