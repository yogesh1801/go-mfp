// MFP - Miulti-Function Printers and scanners toolkit
// Default (typical) configurations
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Default configurations

package defaults

import (
	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/proto/usbip"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/uuid"
)

// ScannerCapabilities the *[abstract.ScannerCapabilities]
// for the typical scanner.
func ScannerCapabilities() *abstract.ScannerCapabilities {
	colorModes := generic.MakeBitset(
		abstract.ColorModeBinary,
		abstract.ColorModeMono,
		abstract.ColorModeColor,
	)

	depths := generic.MakeBitset(
		abstract.ColorDepth8,
	)

	renderings := generic.MakeBitset(
		abstract.BinaryRenderingHalftone,
		abstract.BinaryRenderingThreshold,
	)

	intents := generic.MakeBitset(
		abstract.IntentDocument,
		abstract.IntentTextAndGraphic,
		abstract.IntentPhoto,
		abstract.IntentPreview,
	)

	resolutions := []abstract.Resolution{
		{XResolution: 75, YResolution: 75},
		{XResolution: 150, YResolution: 150},
		{XResolution: 300, YResolution: 300},
		{XResolution: 600, YResolution: 600},
	}

	profile := abstract.SettingsProfile{
		ColorModes:       colorModes,
		Depths:           depths,
		BinaryRenderings: renderings,
		Resolutions:      resolutions,
	}

	inputcaps := &abstract.InputCapabilities{
		MinWidth:   0,
		MaxWidth:   abstract.A4Width,
		MinHeight:  0,
		MaxHeight:  abstract.A4Height,
		MaxXOffset: abstract.A4Width / 2,
		MaxYOffset: abstract.A4Height / 2,
		Intents:    intents,
		Profiles:   []abstract.SettingsProfile{profile},
	}

	caps := &abstract.ScannerCapabilities{
		UUID: uuid.Must(uuid.Parse(
			"169e8d94-9a17-4f14-ae81-52b9176ee9be")),
		MakeAndModel:     "OpenPrinting eSCL scanner",
		SerialNumber:     "OP-0000223321",
		Manufacturer:     "OpenPrinting",
		DocumentFormats:  []string{"image/jpeg", "application/pdf"},
		ADFCapacity:      50,
		CompressionRange: abstract.Range{Min: 2, Normal: 5, Max: 10},
		BrightnessRange:  abstract.Range{Min: -100, Normal: 0, Max: 100},
		ContrastRange:    abstract.Range{Min: -100, Normal: 0, Max: 100},
		Platen:           inputcaps,
		ADFSimplex:       inputcaps,
		ADFDuplex:        inputcaps,
	}
	return caps
}

// USBIPPDescriptor returns the [usbip.USBDeviceDescriptor]
// and endpoints for the typical IPP over USB device.
func USBIPPDescriptor() (
	usbip.USBDeviceDescriptor, []*usbip.Endpoint) {

	endpoints := make([]*usbip.Endpoint, 3)
	for i := range endpoints {
		endpoints[i] = usbip.NewEndpoint(usbip.EndpointInOut,
			usbip.USBXferBulk, 512)
	}

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

	return desc, endpoints
}
