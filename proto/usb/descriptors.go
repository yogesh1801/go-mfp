// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// USB descriptors

package usb

import (
	"encoding/binary"
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// USB limits.
const (
	// MaxDevices is the maximum number of devices on bus.
	MaxDevices = 127

	// MaxConfigurations is the maximum number of configurations
	// per device.
	MaxConfigurations = 255

	// MaxInterfaces defines the maximum number of interfaces
	// per configuration.
	MaxInterfaces = 16

	// MaxEndpoints defines the maximum number of endpoints
	// per configuration. Note, this limit includes the reserved
	// zero endpoint.
	MaxEndpoints = 16

	// MaxStringLength defines the length limit for the strings
	// used in the configuration descriptors.
	MaxStringLength = 254 / 2
)

// Version is the BCD-encoded version number.
type Version uint16

// BVersion constants.
const (
	// USB 1.0
	USB10 Version = 0x0100

	// USB 1.1
	USB11 Version = 0x0101

	// USB 2.0
	USB20 Version = 0x0200
)

// Speed defines the USB speed codes.
type Speed int

// Speed constants:
const (
	// No speed information
	SpeedUnknown Speed = 0

	// USB 1.1 Low Speed (1.5MBit/s).
	SpeedLow Speed = 1

	// USB 1.1 Full Speed (1.5MBit/s).  (12MBit/s).
	SpeedFull Speed = 2

	// USB 2.0 High Speed (480MBit/s).
	SpeedHigh Speed = 3
)

// DescriptorType defines the type of the USB descriptor.
type DescriptorType uint8

// Known DescriptorType values:
const (
	DescriptorDevice        DescriptorType = 1
	DescriptorConfiguration DescriptorType = 2
	DescriptorString        DescriptorType = 3
	DescriptorInterface     DescriptorType = 4
	DescriptorEndpoint      DescriptorType = 5
)

// String returns string representation of DescriptorType, for logging.
func (t DescriptorType) String() string {
	switch t {
	case DescriptorDevice:
		return "Device"
	case DescriptorConfiguration:
		return "Configuration"
	case DescriptorString:
		return "String"
	case DescriptorInterface:
		return "Interface"
	case DescriptorEndpoint:
		return "Endpoint:"
	}

	return fmt.Sprintf("Unknown(%d)", uint8(t))
}

// ConfAttributes defines [ConfigurationDescriptor.BMAttributes] bits.
type ConfAttributes uint8

// ConfAttributes assignment:
const (
	ConfAttrReserved     ConfAttributes = 1 << 7
	ConfAttrSelfPowered  ConfAttributes = 1 << 6
	ConfAttrRemoteWakeup ConfAttributes = 1 << 7
)

// EndpointType represents the endpoint type (in/out/bidir).
//
// Note, the hardware USB doesn't have such a thing that bidirectional
// endpoint, but it is much more convenient from the software point of
// view that having a pair of endpoints, one per direction, to implement
// a logically bidirectional endpoint.
//
// So at the USB side the bidirectional endpoint are represented by
// the pair of uni-directional endpoints.
type EndpointType int

// Endpoint types:
const (
	EndpointIn    EndpointType = iota // Input (host->device)
	EndpointOut                       // Output (device->host)
	EndpointInOut                     // Input/Output (bidirectional)
)

// EndpointAttributes defines [EndpointDescriptor.BMAttributes] bits.
type EndpointAttributes uint8

// EndpointAttributes assignment:
const (
	// Transfer type:
	XferControl     EndpointAttributes = 0x00
	XferIsochronous EndpointAttributes = 0x01
	XferBulk        EndpointAttributes = 0x02
	XferInterrupt   EndpointAttributes = 0x03
	XferMask        EndpointAttributes = 0x03

	// Isochronous Synchronization Type:
	IsoSyncNone         EndpointAttributes = 0x00
	IsoSyncAsynchronous EndpointAttributes = 0x04
	IsoSyncAdaptive     EndpointAttributes = 0x08
	IsoSyncSynchronous  EndpointAttributes = 0x0c
	IsoSyncMask         EndpointAttributes = 0x0c

	// Isochronous Usage Type:
	IsoUsageData     EndpointAttributes = 0x00
	IsoUsageFeedback EndpointAttributes = 0x10
	IsoUsageImplicit EndpointAttributes = 0x20
	IsoUsageMask     EndpointAttributes = 0x30
)

// DeviceDescriptor represents the USB device descriptor.
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
type DeviceDescriptor struct {
	BCDUSB          Version                   // USB spec version
	Speed           Speed                     // Device speed
	BDeviceClass    uint8                     // Device class
	BDeviceSubClass uint8                     // Device subclass
	BDeviceProtocol uint8                     // Protocol code
	BMaxPacketSize  uint8                     // Max pkt (8,16,32 or 64)
	IDVendor        uint16                    // Vendor ID
	IDProduct       uint16                    // Product ID
	BCDDevice       Version                   // Device release number
	IManufacturer   string                    // Manufacturer name
	IProduct        string                    // Product name
	ISerialNumber   string                    // Device serial number
	Configurations  []ConfigurationDescriptor // Device configurations
}

// ConfigurationDescriptor represents the USB configuration descriptor.
type ConfigurationDescriptor struct {
	IConfiguration string         // Configuration description
	BMAttributes   ConfAttributes // Attribute bits
	MaxPower       uint8          // Max power, in 2mA units
	Interfaces     []Interface    // Interfaces grouped by alt settings
}

// Interface represents collection of [InterfaceDescriptor]s
// that belongs to the same interface, ordered by bAlternateSetting.
type Interface struct {
	AltSettings []InterfaceDescriptor // Ordered by alt setting.
}

// CntEndpoints returns count of endpoints, used by the interface,
// taking all alternate settings into the configuration.
func (iff Interface) CntEndpoints() int {
	cnt := 0
	for _, alt := range iff.AltSettings {
		cnt = generic.Max(cnt, alt.CntEndpoints())
	}
	return cnt
}

// InterfaceDescriptor represents the USB interface descriptor.
type InterfaceDescriptor struct {
	BInterfaceClass    uint8                // Interface class
	BInterfaceSubClass uint8                // Interface subclass
	BInterfaceProtocol uint8                // Interface protocol
	IInterface         string               // Interface description
	Endpoints          []EndpointDescriptor // Interface endpoints
}

// EndpointDescriptor represents the USB endpoint descriptor.
type EndpointDescriptor struct {
	Type           EndpointType       // Endpoint type
	BMAttributes   EndpointAttributes // Endpoint attribute bits
	WMaxPacketSize uint16             // Max packet size, bytes
}

// CntEndpoints returns InterfaceDescriptor's count of endpoints.
// Please notice that the [EndpointInOut] endpoints are counted twice.
func (iff InterfaceDescriptor) CntEndpoints() int {
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

// SetupRequestType is the request type bits. It is used
// in the [SetupPacket].
type SetupRequestType uint8

// SetupRequestType bits:
const (
	// Request direction, input if set
	SetupIn SetupRequestType = 0x80

	// Request types
	RequestTypeStandard SetupRequestType = 0x00
	RequestTypeClass    SetupRequestType = 0x20
	RequestTypeVendor   SetupRequestType = 0x40
	RequestTypeTypeMask SetupRequestType = 0x60

	// Recipient
	RecipientDevice    SetupRequestType = 0x00
	RecipientInterface SetupRequestType = 0x01
	RecipientEndpoint  SetupRequestType = 0x02
	RecipientOther     SetupRequestType = 0x03
	RecipientMask      SetupRequestType = 0x03
)

// String returns string representation of SetupRequest, for logging.
func (t SetupRequestType) String() string {
	dir := "->"
	if t&SetupIn != 0 {
		dir = "<-"
	}

	ty := "unknown"
	switch t & RequestTypeTypeMask {
	case RequestTypeStandard:
		ty = "standard"
	case RequestTypeClass:
		ty = "class"
	case RequestTypeVendor:
		ty = "vendor"
	}

	rec := "unknown"
	switch t & RecipientMask {
	case RecipientDevice:
		rec = "device"
	case RecipientInterface:
		rec = "interface"
	case RecipientEndpoint:
		rec = "endpoint"
	case RecipientOther:
		rec = "other"
	}

	return fmt.Sprintf("%s%s%s", ty, dir, rec)
}

// SetupRequest is the request code
type SetupRequest uint8

// SetupRequest assigned values:
const (
	/** Request status of the specific recipient */
	RequestGetStatus SetupRequest = 0x00

	// Clear or disable a specific feature
	RequestClearFeature SetupRequest = 0x01

	// Set or enable a specific feature
	RequestSetFeature SetupRequest = 0x03

	// Set device address for all future accesses
	RequestSetAddress SetupRequest = 0x05

	// Get the specified descriptor
	RequestGetDescriptor SetupRequest = 0x06

	// Used to update existing descriptors or add new descriptors
	RequestSetDescriptor SetupRequest = 0x07

	// Get the current device configuration value
	RequestGetConfiguration SetupRequest = 0x08

	// Set device configuration
	RequestSetConfiguration SetupRequest = 0x09

	// Return the selected alternate setting for the specified interface
	RequestGetInterface SetupRequest = 0x0a

	// Select an alternate interface for the specified interface
	RequestSetInterface SetupRequest = 0x0b
)

// String returns string representation of SetupRequest, for logging.
func (r SetupRequest) String() string {
	switch r {
	case RequestGetStatus:
		return "GET_STATUS"
	case RequestClearFeature:
		return "CLEAR_FEATURE"
	case RequestSetFeature:
		return "SET_FEATURE"
	case RequestSetAddress:
		return "SET_ADDRESS"
	case RequestGetDescriptor:
		return "GET_DESCRIPTOR"
	case RequestSetDescriptor:
		return "SET_DESCRIPTOR"
	case RequestGetConfiguration:
		return "GET_CONFIGURATION"
	case RequestSetConfiguration:
		return "SET_CONFIGURATION"
	case RequestGetInterface:
		return "SET_INTERFACE"
	case RequestSetInterface:
		return "SET_INTERFACE"
	}

	return fmt.Sprintf("Unknown(0x%2.2x)", uint8(r))
}

// SetupPacket is the USB Setup Packet
type SetupPacket struct {
	RequestType SetupRequestType
	Request     SetupRequest
	WValue      uint16
	WIndex      uint16
	WLength     uint16
}

// String returns string representation of the SetupPacket, for logging
func (p SetupPacket) String() string {
	name := fmt.Sprintf("%s (%s)", p.Request, p.RequestType)

	switch p.Request {
	case RequestGetDescriptor:
		t := DescriptorType(p.WValue >> 8)
		i := p.WValue & 255
		return fmt.Sprintf("%s: %s[%d]", name, t, i)

	case RequestSetConfiguration, RequestSetInterface:
		return fmt.Sprintf("%s: %d", name, p.WValue)
	}

	return name
}

// Encode returns the binary representation of the setup packet.
func (p SetupPacket) Encode() [8]byte {
	var data [8]byte

	data[0] = uint8(p.RequestType)
	data[1] = uint8(p.Request)
	binary.LittleEndian.PutUint16(data[2:4], p.WValue)
	binary.LittleEndian.PutUint16(data[4:6], p.WIndex)
	binary.LittleEndian.PutUint16(data[6:8], p.WLength)

	return data
}

// Decode decodes setup packet from the binary representation.
func (p *SetupPacket) Decode(data [8]byte) {
	p.RequestType = SetupRequestType(data[0])
	p.Request = SetupRequest(data[1])
	p.WValue = binary.LittleEndian.Uint16(data[2:4])
	p.WIndex = binary.LittleEndian.Uint16(data[4:6])
	p.WLength = binary.LittleEndian.Uint16(data[6:8])
}
