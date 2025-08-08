// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// The virtual device

package usbip

import (
	"errors"
	"fmt"
	"math"
	"syscall"
)

// Device represents the virtual USB device.
//
// Device lifetime is following:
//   - use [NewDevice] to create the new Device.
//   - fill required callbacks
//   - use [Server.AddDevice] to add device to the server
type Device struct {
	// Device description
	Descriptor USBDeviceDescriptor // Device descriptor

	// Device state (read-only)
	Configuration int   // Current configuration
	AltSettings   []int // Per-interface current alt setting

	// User-defined callbacks
	OnConfigurationChange func() // Called when configuration changes

	// Internal state
	maxAltSettings []int          // Max alt setting, by interface
	endpoints      []*Endpoint    // Endpoints of current configuration
	strings        []string       // Content of string descriptors
	stringsmap     map[string]int // Indices of strings descriptors
}

// NewDevice creates a new device, based on the provided device descriptor.
func NewDevice(desc USBDeviceDescriptor) (*Device, error) {
	// Validate device descriptor
	confs := len(desc.Configurations)

	switch {
	case confs == 0:
		return nil, errors.New("Device has no configurations")

	case confs > USBMaxConfigurations:
		err := fmt.Errorf("Device has too many (%d) configurations",
			confs)
		return nil, err
	}

	for confno, conf := range desc.Configurations {
		iffs := len(conf.Interfaces)

		switch {
		case iffs == 0:
			err := fmt.Errorf(
				"Configuration %d has no interfaces",
				confno)
			return nil, err

		case iffs > USBMaxInterfaces:
			err := fmt.Errorf(
				"Configuration %d has too many [%d] interfaces",
				confno, iffs)
			return nil, err
		}

		endpoints := 1 // Reserved for configuration endpoint
		for _, iff := range conf.Interfaces {
			endpoints += iff.CntEndpoints()
			if endpoints > USBMaxEndpoints {
				err := fmt.Errorf(
					"Configuration %d has too many (%d) endpoints",
					confno, endpoints)
				return nil, err
			}
		}
	}

	// Initialize the Device structure
	dev := &Device{
		Descriptor: desc,
		stringsmap: make(map[string]int),
	}

	// Populate strings
	dev.addstring("") // Reserve index 0
	dev.addstring(desc.IManufacturer)
	dev.addstring(desc.IProduct)
	dev.addstring(desc.ISerialNumber)

	for _, conf := range desc.Configurations {
		dev.addstring(conf.IConfiguration)

		for _, iff := range conf.Interfaces {
			for _, alt := range iff.AltSettings {
				dev.addstring(alt.IInterface)
			}
		}
	}

	// Initialize current configuration
	dev.setConfiguration(1)

	return dev, nil
}

// GetStatus returns the USB device status.
func (dev *Device) getStatus() ([]byte, syscall.Errno) {
	data := make([]byte, 2)

	conf := dev.Descriptor.Configurations[dev.Configuration-1]

	if conf.BMAttributes&USBConfAttrSelfPowered != 0 {
		data[1] |= 0x01
	}

	if conf.BMAttributes&USBConfAttrRemoteWakeup != 0 {
		data[1] |= 0x02
	}

	return data, 0
}

// GetDescriptor returns USB descriptor, in the USB wire representation.
// If type or index is invalid, it returns nil and syscall.EPIPE.
func (dev *Device) getDescriptor(t USBDescriptorType, i int) ([]byte, syscall.Errno) {
	switch t {
	case USBDescriptorDevice:
		if i > 0 {
			return nil, syscall.EPIPE
		}

		desc := dev.Descriptor
		enc := newEncoder(18)

		enc.PutU8(18)                         // bLength
		enc.PutU8(uint8(USBDescriptorDevice)) // bDescriptorType
		enc.PutLE16(uint16(desc.BCDUSB))      // bcdUSB
		enc.PutU8(desc.BDeviceClass)          // bDeviceClass
		enc.PutU8(desc.BDeviceSubClass)       // bDeviceSubClass
		enc.PutU8(desc.BDeviceProtocol)       // bDeviceProtocol
		enc.PutU8(desc.BMaxPacketSize)        // bMaxPacketSize
		enc.PutLE16(desc.IDVendor)            // idVendor
		enc.PutLE16(desc.IDProduct)           // idProduct
		enc.PutLE16(uint16(desc.BCDDevice))   // bcdDevice

		i := dev.getstring(desc.IManufacturer)
		enc.PutU8(uint8(i)) // iManufacturer

		i = dev.getstring(desc.IProduct)
		enc.PutU8(uint8(i)) // iProduct

		i = dev.getstring(desc.ISerialNumber)
		enc.PutU8(uint8(i)) // iSerialNumber

		enc.PutU8(uint8(len(desc.Configurations))) // bNumConfigurations

		return enc.Bytes(), 0

	case USBDescriptorConfiguration:
		if i > len(dev.Descriptor.Configurations) {
			return nil, syscall.EPIPE
		}

		// Encode Configuration Descriptor
		conf := dev.Descriptor.Configurations[i]
		enc := newEncoder(256)

		enc.PutU8(9)                                 // bLength
		enc.PutU8(uint8(USBDescriptorConfiguration)) // bDescriptorType
		enc.PutLE16(0)                               // wTotalLength, reserved
		enc.PutU8(uint8(len(conf.Interfaces)))       // bNumInterfaces
		enc.PutU8(uint8(i + 1))                      // bConfigurationValue

		i := dev.getstring(conf.IConfiguration)
		enc.PutU8(uint8(i)) // iConfiguration

		attrs := conf.BMAttributes
		attrs |= USBConfAttrReserved
		enc.PutU8(uint8(attrs)) // bmAttributes

		enc.PutU8(conf.MaxPower) // bMaxPower

		// Encode interfaces and endpoints
		epnum := 1
		for iffno, iff := range conf.Interfaces {
			for altno, alt := range iff.AltSettings {
				endpoints := iff.CntEndpoints()

				enc.PutU8(9)                             // bLength
				enc.PutU8(uint8(USBDescriptorInterface)) // bDescriptorType
				enc.PutU8(uint8(iffno))                  // bInterfaceNumber
				enc.PutU8(uint8(altno))                  // bAlternateSetting
				enc.PutU8(uint8(endpoints))              // bNumEndpoints
				enc.PutU8(alt.BInterfaceClass)           // bInterfaceClass
				enc.PutU8(alt.BInterfaceSubClass)        // bInterfaceSubClass
				enc.PutU8(alt.BInterfaceProtocol)        // bInterfaceProtocol

				i := dev.getstring(alt.IInterface)
				enc.PutU8(uint8(i)) // iInterface

				for _, ep := range alt.Endpoints {
					ty := ep.Type()
					if ty == EndpointIn || ty == EndpointInOut {
						enc.PutU8(7) // bLength
						enc.PutU8(   // bDescriptorType
							uint8(USBDescriptorEndpoint))

						addr := epnum | 0x80
						epnum++

						enc.PutU8(uint8(addr))            // bEndpointAddress
						enc.PutU8(uint8(ep.Attrs()))      // bmAttributes
						enc.PutLE16(uint16(ep.PktSize())) // bmAttributes
						enc.PutU8(0)                      // bInterval
					}

					if ty == EndpointOut || ty == EndpointInOut {
						enc.PutU8(7) // bLength
						enc.PutU8(   // bDescriptorType
							uint8(USBDescriptorEndpoint))

						addr := epnum
						epnum++

						enc.PutU8(uint8(addr))            // bEndpointAddress
						enc.PutU8(uint8(ep.Attrs()))      // bmAttributes
						enc.PutLE16(uint16(ep.PktSize())) // bmAttributes
						enc.PutU8(0)                      // bInterval
					}
				}
			}
		}

		// Patch wTotalLength
		data := enc.Bytes()
		length := len(data)
		data[2] = uint8(length)
		data[3] = uint8(length >> 8)

		return data, 0

	case USBDescriptorString:
		// String Descriptor Zero is special: it contains list of
		// supported languages.
		if i == 0 {
			enc := newEncoder(4)
			enc.PutU8(8)                          // bLength
			enc.PutU8(uint8(USBDescriptorString)) // bDescriptorType
			enc.PutLE16(0x0409)                   // wLANGID[0]
			return enc.Bytes(), 0
		}

		if i >= len(dev.strings) {
			return nil, syscall.EPIPE
		}

		s := []byte(dev.strings[i])
		enc := newEncoder(4)
		enc.PutU8(uint8(len(s)*2 + 2))        // bLength
		enc.PutU8(uint8(USBDescriptorString)) // bDescriptorType
		for _, c := range s {
			enc.PutLE16(uint16(c)) // bString
		}
		return enc.Bytes(), 0

	}

	return nil, syscall.EPIPE
}

// GetConfiguration returns the current configuration.
func (dev *Device) getConfiguration() ([]byte, syscall.Errno) {
	return []byte{uint8(dev.Configuration)}, 0
}

// SetConfiguration selects the current configuration.
func (dev *Device) setConfiguration(n int) ([]byte, syscall.Errno) {
	if n > len(dev.Descriptor.Configurations) {
		return nil, syscall.EPIPE
	}

	if n != 0 {
		// Save Configuration
		dev.Configuration = n
		conf := dev.Descriptor.Configurations[n-1]

		// Update dev.AltSettings
		dev.AltSettings = make([]int, len(conf.Interfaces))
		dev.maxAltSettings = make([]int, len(conf.Interfaces))

		for i, iff := range conf.Interfaces {
			dev.maxAltSettings[i] = len(iff.AltSettings)
		}

		// Update dev.Endpoints
		// dev.endpoints[0] is reserved.
		dev.endpoints = make([]*Endpoint, 0, 16)
		dev.endpoints = append(dev.endpoints, nil)

		for _, iff := range conf.Interfaces {
			for _, alt := range iff.AltSettings {
				for _, ep := range alt.Endpoints {
					switch ep.Type() {
					case EndpointIn, EndpointOut:
						dev.endpoints = append(dev.endpoints,
							ep)
					case EndpointInOut:
						dev.endpoints = append(dev.endpoints,
							ep, ep)
					}
				}
			}
		}

		// Call user callback
		if dev.OnConfigurationChange != nil {
			dev.OnConfigurationChange()
		}
	}

	return []byte{}, 0
}

// submit routes protoIOSubmitRequest to the endpoint
func (dev *Device) submit(rq *protoIOSubmitRequest) syscall.Errno {
	// Lookup the endpoint
	var ep *Endpoint
	if int(rq.Endpoint) <= len(dev.endpoints) {
		ep = dev.endpoints[rq.Endpoint]
	}

	if ep == nil {
		return syscall.EPIPE
	}

	// Check for direction mismatch
	if (rq.Input && ep.Type() == EndpointOut) ||
		(!rq.Input && ep.Type() == EndpointIn) {
		return syscall.EPIPE
	}

	// Route the request
	return ep.submit(rq)
}

// unlink routes protoIOUnlinkRequest to the endpoint
func (dev *Device) unlink(rq *protoIOUnlinkRequest) syscall.Errno {
	// Lookup the endpoint
	var ep *Endpoint
	if int(rq.Endpoint) <= len(dev.endpoints) {
		ep = dev.endpoints[rq.Endpoint]
	}

	if ep == nil {
		return syscall.EPIPE
	}

	// Route the request
	return ep.unlink(rq)
}

// shutdown cancels app pending protoIOSubmitRequest at all endpoints.
func (dev *Device) shutdown() {
	for _, ep := range dev.endpoints[1:] {
		ep.shutdown()
	}
}

// GetInterfaceStatus returns the USB interface status.
func (dev *Device) getInterfaceStatus(ifn int) ([]byte, syscall.Errno) {
	if ifn >= len(dev.AltSettings) {
		return nil, syscall.EPIPE
	}

	return make([]byte, 2), 0
}

// GetInterface returns alternate setting selection on the given interface.
func (dev *Device) getInterface(ifn int) ([]byte, syscall.Errno) {
	if ifn >= len(dev.AltSettings) {
		return nil, syscall.EPIPE
	}

	return []byte{uint8(dev.AltSettings[ifn])}, 0
}

// SetInterface selects alternate setting selection on the given interface.
func (dev *Device) setInterface(ifn, alt int) ([]byte, syscall.Errno) {
	if ifn >= len(dev.AltSettings) {
		return nil, syscall.EPIPE
	}

	if alt > dev.maxAltSettings[ifn] {
		return nil, syscall.EPIPE
	}

	dev.AltSettings[ifn] = alt
	return []byte{}, 0
}

// addstring adds string to strings descriptor.
func (dev *Device) addstring(s string) {
	if len(dev.strings) < math.MaxUint8 {
		if len(s) > USBMaxStringLength {
			s = s[:USBMaxStringLength]
		}

		if _, found := dev.stringsmap[s]; !found {
			dev.stringsmap[s] = len(dev.strings)
			dev.strings = append(dev.strings, s)
		}
	}
}

// getstring returns the string descriptor index for the string.
// If index is not available, 0 will be returned (which is OK for USB).
func (dev *Device) getstring(s string) int {
	if len(s) > USBMaxStringLength {
		s = s[:USBMaxStringLength]
	}

	return dev.stringsmap[s]
}
