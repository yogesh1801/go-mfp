// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// USB definitions

package usbip

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// USB limits.
const (
	// USBMaxDevices is the maximum number of devices on bus.
	USBMaxDevices = 127

	// USBMaxConfigurations is the maximum number of configurations
	// per device.
	USBMaxConfigurations = 255

	// USBMaxInterfaces defines the maximum number of interfaces
	// per configuration.
	USBMaxInterfaces = 16

	// USBMaxEndpoints defines the maximum number of endpoints
	// per configuration. Note, this limit includes the reserved
	// zero endpoint.
	USBMaxEndpoints = 16

	// USBMaxStringLength defines the length limit for the strings
	// used in the configuration descriptors.
	USBMaxStringLength = 254 / 2
)

// USBVersion is the BCD-encoded version number.
type USBVersion uint16

// USBBVersion constants.
const (
	// USB 1.0
	USB10 USBVersion = 0x0100

	// USB 1.1
	USB11 USBVersion = 0x0101

	// USB 2.0
	USB20 USBVersion = 0x0200
)

// USBSpeed defines the USB speed codes.
type USBSpeed int

// USBSpeed constants:
const (
	// No speed information
	USBSpeedUnknown USBSpeed = 0

	// USB 1.1 Low Speed (1.5MBit/s).
	USBSpeedLow USBSpeed = 1

	// USB 1.1 Full Speed (1.5MBit/s).  (12MBit/s).
	USBSpeedFull USBSpeed = 2

	// USB 2.0 High Speed (480MBit/s).
	USBSpeedHigh USBSpeed = 3
)

// USBDescriptorType defines the type of the USB descriptor.
type USBDescriptorType uint8

// Known USBDescriptorType values:
const (
	USBDescriptorDevice        USBDescriptorType = 1
	USBDescriptorConfiguration USBDescriptorType = 2
	USBDescriptorString        USBDescriptorType = 3
	USBDescriptorInterface     USBDescriptorType = 4
	USBDescriptorEndpoint      USBDescriptorType = 5
)

// String returns string representation of USBDescriptorType, for logging.
func (t USBDescriptorType) String() string {
	switch t {
	case USBDescriptorDevice:
		return "Device"
	case USBDescriptorConfiguration:
		return "Configuration"
	case USBDescriptorString:
		return "String"
	case USBDescriptorInterface:
		return "Interface"
	case USBDescriptorEndpoint:
		return "Endpoint:"
	}

	return fmt.Sprintf("Unknown(%d)", uint8(t))
}

// USBConfAttributes defines [USBConfigurationDescriptor.BMAttributes] bits.
type USBConfAttributes uint8

// USBConfAttributes assignment:
const (
	USBConfAttrReserved     USBConfAttributes = 1 << 7
	USBConfAttrSelfPowered  USBConfAttributes = 1 << 6
	USBConfAttrRemoteWakeup USBConfAttributes = 1 << 7
)

// USBEndpointAttributes defines [USBEndpointDescriptor.BMAttributes] bits.
type USBEndpointAttributes uint8

// USBEndpointAttributes assignment:
const (
	// Transfer type:
	USBXferControl     USBEndpointAttributes = 0x00
	USBXferIsochronous USBEndpointAttributes = 0x01
	USBXferBulk        USBEndpointAttributes = 0x02
	USBXferInterrupt   USBEndpointAttributes = 0x03
	USBXferMask        USBEndpointAttributes = 0x03

	// Isochronous Synchronization Type:
	USBIsoSyncNone         USBEndpointAttributes = 0x00
	USBIsoSyncAsynchronous USBEndpointAttributes = 0x04
	USBIsoSyncAdaptive     USBEndpointAttributes = 0x08
	USBIsoSyncSynchronous  USBEndpointAttributes = 0x0c
	USBIsoSyncMask         USBEndpointAttributes = 0x0c

	// Isochronous Usage Type:
	USBIsoUsageData     USBEndpointAttributes = 0x00
	USBIsoUsageFeedback USBEndpointAttributes = 0x10
	USBIsoUsageImplicit USBEndpointAttributes = 0x20
	USBIsoUsageMask     USBEndpointAttributes = 0x30
)

// USBDeviceDescriptor represents the USB device descriptor.
//
// This structure and its children structures are very close
// to the USB descriptor structures commonly used in the USB
// documentation, but doesn't match 1:1.
//
// In particular, some fields, like bDescriptorType and bLength
// are omitted and computed automatically. String fields, like
// iManufacturer, are represented by Go strings, not by indices
// of the strings descriptor (indices are assigned automatically
// and appropriate string descriptors are automatically generated
// as well).
//
// Interface numbers, alternate settings numbers and endpoint
// addresses are automatically assigned, based on the device
// configuration layout.
type USBDeviceDescriptor struct {
	BCDUSB          USBVersion                   // USB spec version
	Speed           USBSpeed                     // Device speed
	BDeviceClass    uint8                        // Device class
	BDeviceSubClass uint8                        // Device subclass
	BDeviceProtocol uint8                        // Protocol code
	BMaxPacketSize  uint8                        // Max pkt (8,16,32 or 64)
	IDVendor        uint16                       // Vendor ID
	IDProduct       uint16                       // Product ID
	BCDDevice       USBVersion                   // Device release number
	IManufacturer   string                       // Manufacturer name
	IProduct        string                       // Product name
	ISerialNumber   string                       // Device serial number
	Configurations  []USBConfigurationDescriptor // Device configurations
}

// USBConfigurationDescriptor represents the USB configuration descriptor.
type USBConfigurationDescriptor struct {
	IConfiguration string            // Configuration description
	BMAttributes   USBConfAttributes // Attribute bits
	MaxPower       uint8             // Max power, in 2mA units
	Interfaces     []USBInterface    // Interfaces grouped by alt settings
}

// USBInterface represents collection of [USBInterfaceDescriptor]s
// that belongs to the same interface, ordered by bAlternateSetting.
type USBInterface struct {
	AltSettings []USBInterfaceDescriptor // Ordered by alt setting.
}

// cntEndpoints returns count of endpoints, used by the interface,
// taking all alternate settings into the configuration.
func (iff USBInterface) cntEndpoints() int {
	cnt := 0
	for _, alt := range iff.AltSettings {
		cnt = generic.Max(cnt, alt.cntEndpoints())
	}
	return cnt
}

// USBInterfaceDescriptor represents the USB interface descriptor.
type USBInterfaceDescriptor struct {
	BInterfaceClass    uint8                   // Interface class
	BInterfaceSubClass uint8                   // Interface subclass
	BInterfaceProtocol uint8                   // Interface protocol
	IInterface         string                  // Interface description
	Endpoints          []USBEndpointDescriptor // Interface endpoints
}

// USBEndpointDescriptor represents the USB endpoint descriptor.
type USBEndpointDescriptor struct {
	Type           EndpointType          // Endpoint type
	BMAttributes   USBEndpointAttributes // Endpoint attribute bits
	WMaxPacketSize uint16                // Max packet size, bytes
}

// cntEndpoints returns USBInterfaceDescriptor's count of endpoints.
// Please notice that the [EndpointInOut] endpoints are counted twice.
func (iff USBInterfaceDescriptor) cntEndpoints() int {
	cnt := 0

	for _, ep := range iff.Endpoints {
		switch ep.Type {
		case EndpointIn, EndpointOut:
			cnt++
		case EndpointInOut:
			cnt += 2
		}
	}

	return cnt
}

// USBSetupRequestType is the request type bits. It is used
// in the [USBSetupPacket].
type USBSetupRequestType uint8

// USBSetupRequestType bits:
const (
	// Request direction, input if set
	USBSetupIn USBSetupRequestType = 0x80

	// Request types
	USBRequestTypeStandard USBSetupRequestType = 0x00
	USBRequestTypeClass    USBSetupRequestType = 0x20
	USBRequestTypeVendor   USBSetupRequestType = 0x40
	USBRequestTypeTypeMask USBSetupRequestType = 0x60

	// Recipient
	USBRecipientDevice    USBSetupRequestType = 0x00
	USBRecipientInterface USBSetupRequestType = 0x01
	USBRecipientEndpoint  USBSetupRequestType = 0x02
	USBRecipientOther     USBSetupRequestType = 0x03
	USBRecipientMask      USBSetupRequestType = 0x03
)

// String returns string representation of USBSetupRequest, for logging.
func (t USBSetupRequestType) String() string {
	dir := "->"
	if t&USBSetupIn != 0 {
		dir = "<-"
	}

	ty := "unknown"
	switch t & USBRequestTypeTypeMask {
	case USBRequestTypeStandard:
		ty = "standard"
	case USBRequestTypeClass:
		ty = "class"
	case USBRequestTypeVendor:
		ty = "vendor"
	}

	rec := "unknown"
	switch t & USBRecipientMask {
	case USBRecipientDevice:
		rec = "device"
	case USBRecipientInterface:
		rec = "interface"
	case USBRecipientEndpoint:
		rec = "endpoint"
	case USBRecipientOther:
		rec = "other"
	}

	return fmt.Sprintf("%s%s%s", ty, dir, rec)
}

// USBSetupRequest is the request code
type USBSetupRequest uint8

// USBSetupRequest assigned values:
const (
	/** Request status of the specific recipient */
	USBRequestGetStatus USBSetupRequest = 0x00

	// Clear or disable a specific feature
	USBRequestClearFeature USBSetupRequest = 0x01

	// Set or enable a specific feature
	USBRequestSetFeature USBSetupRequest = 0x03

	// Set device address for all future accesses
	USBRequestSetAddress USBSetupRequest = 0x05

	// Get the specified descriptor
	USBRequestGetDescriptor USBSetupRequest = 0x06

	// Used to update existing descriptors or add new descriptors
	USBRequestSetDescriptor USBSetupRequest = 0x07

	// Get the current device configuration value
	USBRequestGetConfiguration USBSetupRequest = 0x08

	// Set device configuration
	USBRequestSetConfiguration USBSetupRequest = 0x09

	// Return the selected alternate setting for the specified interface
	USBRequestGetInterface USBSetupRequest = 0x0a

	// Select an alternate interface for the specified interface
	USBRequestSetInterface USBSetupRequest = 0x0b
)

// String returns string representation of USBSetupRequest, for logging.
func (r USBSetupRequest) String() string {
	switch r {
	case USBRequestGetStatus:
		return "GET_STATUS"
	case USBRequestClearFeature:
		return "CLEAR_FEATURE"
	case USBRequestSetFeature:
		return "SET_FEATURE"
	case USBRequestSetAddress:
		return "SET_ADDRESS"
	case USBRequestGetDescriptor:
		return "GET_DESCRIPTOR"
	case USBRequestSetDescriptor:
		return "SET_DESCRIPTOR"
	case USBRequestGetConfiguration:
		return "GET_CONFIGURATION"
	case USBRequestSetConfiguration:
		return "SET_CONFIGURATION"
	case USBRequestGetInterface:
		return "SET_INTERFACE"
	case USBRequestSetInterface:
		return "SET_INTERFACE"
	}

	return fmt.Sprintf("Unknown(0x%2.2x)", uint8(r))
}

// USBSetupPacket is the USB Setup Packet
type USBSetupPacket struct {
	RequestType USBSetupRequestType
	Request     USBSetupRequest
	WValue      uint16
	WIndex      uint16
	WLength     uint16
}

// String returns string representation of the USBSetupPacket, for logging
func (p USBSetupPacket) String() string {
	name := fmt.Sprintf("%s (%s)", p.Request, p.RequestType)

	switch p.Request {
	case USBRequestGetDescriptor:
		t := USBDescriptorType(p.WValue >> 8)
		i := p.WValue & 255
		return fmt.Sprintf("%s: %s[%d]", name, t, i)

	case USBRequestSetConfiguration, USBRequestSetInterface:
		return fmt.Sprintf("%s: %d", name, p.WValue)
	}

	return name
}

// Encode returns the binary representation of the setup packet.
func (p USBSetupPacket) Encode() [8]byte {
	enc := newEncoder(8)
	enc.PutU8(uint8(p.RequestType))
	enc.PutU8(uint8(p.Request))
	enc.PutLE16(p.WValue)
	enc.PutLE16(p.WIndex)
	enc.PutLE16(p.WLength)

	var ret [8]byte
	copy(ret[:], enc.Bytes())
	return ret
}

// Decode decodes setup packet from the binary representation.
func (p *USBSetupPacket) Decode(data [8]byte) {
	dec := newDecoder(data[:])
	p.RequestType = USBSetupRequestType(dec.GetU8())
	p.Request = USBSetupRequest(dec.GetU8())
	p.WValue = dec.GetLE16()
	p.WIndex = dec.GetLE16()
	p.WLength = dec.GetLE16()
}
