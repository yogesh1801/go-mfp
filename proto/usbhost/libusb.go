// MFP - Miulti-Function Printers and scanners toolkit
// USB host API
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// USB host API, libusb version

package usbhost

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/OpenPrinting/go-mfp/proto/usb"
)

// #cgo pkg-config: libusb-1.0
// #include <libusb.h>
//
// // Note, libusb_strerror accepts enum libusb_error argument, which
// // unfortunately behaves differently depending on target OS and compiler
// // version (sometimes as C.int, sometimes as int32). Looks like cgo
// // bug. Wrapping this function into this simple wrapper should
// // fix the problem. See ipp-usb, #18 for details
// static inline const char*
// libusb_strerror_wrapper (int code) {
//     return libusb_strerror(code);
// }
import "C"

var (
	// libusbContextPtr keeps a pointer to libusb_context.
	// It is initialized on demand
	libusbContextPtr *C.libusb_context

	// libusbContextLock protects libusbContextPtr initialization
	// in the multithreaded context
	libusbContextLock sync.Mutex
)

// libusbError represents libusb error.
type libusbError struct {
	function string // function that caused the error
	code     int    // linusb error code
}

// Error implements error interface for the libusbError
func (err libusbError) Error() string {
	s := C.GoString(C.libusb_strerror_wrapper(C.int(err.code)))
	return err.function + ": " + s
}

// ListDevices returns list of all connected USB devices.
func ListDevices() ([]DeviceInfo, error) {
	// Obtain libusb context
	context, err := libusbContext()
	if err != nil {
		return nil, err
	}

	// Obtain list of devices
	var devlist **C.libusb_device
	cnt := C.libusb_get_device_list(context, &devlist)
	if cnt < 0 {
		return nil, libusbError{"libusb_get_device_list", int(cnt)}
	}
	defer C.libusb_free_device_list(devlist, 1)

	// Decode device list
	infos := make([]DeviceInfo, 0, cnt)
	for _, dev := range unsafe.Slice(devlist, cnt) {
		info, err := libusbDecodeDeviceInfo(dev)

		if err != nil {
			if err, ok := err.(libusbError); ok {
				if !libusbIsFatal(err.code) {
					continue
				}
			}

			return nil, err
		}

		infos = append(infos, info)
	}

	return infos, nil
}

// libusbDecodeDeviceInfo decodes DeviceInfo out of C.libusb_device
func libusbDecodeDeviceInfo(dev *C.libusb_device) (info DeviceInfo, err error) {
	// Obtain device handle.
	// We need it to read device strings etc.
	var handle *C.libusb_device_handle
	rc := C.libusb_open(dev, &handle)
	if rc < 0 {
		err = libusbError{"libusb_open", int(rc)}
		return
	}

	defer C.libusb_close(handle)

	// Obtain device descriptor
	var cdesc C.struct_libusb_device_descriptor
	rc = C.libusb_get_device_descriptor(dev, &cdesc)
	if rc < 0 {
		err = libusbError{"libusb_get_device_descriptor", int(rc)}
		return
	}

	// Decode usb.DeviceDescriptor
	desc := usb.DeviceDescriptor{
		BCDUSB:          usb.Version(cdesc.bcdUSB),
		Speed:           usb.Speed(C.libusb_get_device_speed(dev)),
		BDeviceClass:    uint8(cdesc.bDeviceClass),
		BDeviceSubClass: uint8(cdesc.bDeviceSubClass),
		BDeviceProtocol: uint8(cdesc.bDeviceProtocol),
		BMaxPacketSize:  uint8(cdesc.bMaxPacketSize0),
		IDVendor:        uint16(cdesc.idVendor),
		IDProduct:       uint16(cdesc.idProduct),
		BCDDevice:       usb.Version(cdesc.bcdDevice),
		IManufacturer:   libusbGetString(handle, cdesc.iManufacturer),
		IProduct:        libusbGetString(handle, cdesc.iProduct),
		ISerialNumber:   libusbGetString(handle, cdesc.iSerialNumber),
	}

	// Roll over all configurations
	for confno := 0; confno < int(cdesc.bNumConfigurations); confno++ {
		// Get configuration descriptor
		var cconf *C.struct_libusb_config_descriptor
		rc = C.libusb_get_config_descriptor(dev,
			C.uint8_t(confno), &cconf)

		if rc < 0 {
			err = libusbError{"libusb_get_device_descriptor", int(rc)}
			return
		}

		// Decode configuration descriptor
		var conf usb.ConfigurationDescriptor
		conf, err = libusbDecodeConfigurationDescriptor(dev, handle, cconf)

		if err != nil {
			return
		}

		desc.Configurations = append(desc.Configurations, conf)

		// Free configuration descriptor
		C.libusb_free_config_descriptor(cconf)
	}

	// Decode Location
	loc := Location{
		Bus: int(C.libusb_get_bus_number(dev)),
		Dev: int(C.libusb_get_device_address(dev)),
	}

	// Build and return the DeviceInfo
	return DeviceInfo{loc, desc}, nil
}

// libusbDecodeConfiguration decodes the USB configuration descriptor
func libusbDecodeConfigurationDescriptor(dev *C.libusb_device,
	handle *C.libusb_device_handle,
	cconf *C.struct_libusb_config_descriptor) (
	conf usb.ConfigurationDescriptor, err error) {

	// Decode configuration descriptor itself
	conf = usb.ConfigurationDescriptor{
		IConfiguration: libusbGetString(handle, cconf.iConfiguration),
		BMAttributes:   usb.ConfAttributes(cconf.bmAttributes),
		MaxPower:       uint8(cconf.MaxPower),
	}

	// Roll over all interfaces
	ifcnt := cconf.bNumInterfaces
	ifaces := (*[256]C.struct_libusb_interface)(
		unsafe.Pointer(cconf._interface))[:ifcnt:ifcnt]

	conf.Interfaces = make([]usb.Interface, 0, ifcnt)

	for _, ciff := range ifaces {
		var iff usb.Interface
		iff, err = libusbDecodeInterface(dev, handle, &ciff)
		if err != nil {
			return
		}

		conf.Interfaces = append(conf.Interfaces, iff)
	}

	return
}

// libusbDecodeInterface decodes the USB interface alt settings
func libusbDecodeInterface(dev *C.libusb_device,
	handle *C.libusb_device_handle,
	ciff *C.struct_libusb_interface) (
	iff usb.Interface, err error) {

	// Roll over all alt settings
	altcnt := ciff.num_altsetting
	alts := (*[256]C.struct_libusb_interface_descriptor)(
		unsafe.Pointer(ciff.altsetting))[:altcnt:altcnt]

	iff.AltSettings = make([]usb.InterfaceDescriptor, 0, altcnt)

	for _, calt := range alts {
		var alt usb.InterfaceDescriptor
		alt, err = libusbDecodeInterfaceDescriptor(dev, handle, &calt)
		if err != nil {
			return
		}

		iff.AltSettings = append(iff.AltSettings, alt)
	}

	return
}

// libusbDecodeInterfaceDescriptor decodes the USB interface descriptor
func libusbDecodeInterfaceDescriptor(dev *C.libusb_device,
	handle *C.libusb_device_handle,
	calt *C.struct_libusb_interface_descriptor) (
	alt usb.InterfaceDescriptor, err error) {

	// Decode interface descriptor body
	alt = usb.InterfaceDescriptor{
		BInterfaceClass:    uint8(calt.bInterfaceClass),
		BInterfaceSubClass: uint8(calt.bInterfaceSubClass),
		BInterfaceProtocol: uint8(calt.bInterfaceProtocol),
		IInterface:         libusbGetString(handle, calt.iInterface),
	}

	// Roll over endpoints
	epcnt := calt.bNumEndpoints
	endpoints := (*[256]C.struct_libusb_endpoint_descriptor)(
		unsafe.Pointer(calt.endpoint))[:epcnt:epcnt]

	for _, cep := range endpoints {
		var ep usb.EndpointDescriptor
		ep, err = libusbDecodeEndpointDescriptor(dev, handle, &cep)
		if err != nil {
			return
		}

		alt.Endpoints = append(alt.Endpoints, ep)
	}

	return
}

// libusbDecodeEndpointDescriptor decodes the USB endpoint descriptor
func libusbDecodeEndpointDescriptor(dev *C.libusb_device,
	handle *C.libusb_device_handle,
	cep *C.struct_libusb_endpoint_descriptor) (
	ep usb.EndpointDescriptor, err error) {

	ep = usb.EndpointDescriptor{
		BMAttributes:   usb.EndpointAttributes(cep.bmAttributes),
		WMaxPacketSize: uint16(cep.wMaxPacketSize),
	}

	ep.Type = usb.EndpointOut
	if cep.bEndpointAddress&0x80 != 0 {
		ep.Type = usb.EndpointIn
	}

	return
}

// libusbGetString returns string from the device by the string
// descriptor index.
func libusbGetString(handle *C.libusb_device_handle, i C.uint8_t) string {
	var buf [256]C.uchar

	rc := C.libusb_get_string_descriptor_ascii(handle, i,
		&buf[0], C.int(len(buf)))

	if rc < 0 {
		return ""
	}

	return C.GoString((*C.char)(unsafe.Pointer(&buf[0])))
}

// libusbIsFatal determines whether a libusb error should be considered fatal
// (i.e., whether it should interrupt a major operation such as enumerating
// devices).
//
// Since devices may be unplugged during descriptor decoding or I/O errors may
// occur, we silently ignore certain error conditions related to these
// scenarios.
func libusbIsFatal(rc int) bool {
	switch rc {
	case 0:
	case C.LIBUSB_ERROR_NO_DEVICE, C.LIBUSB_ERROR_IO:

	default:
		return true
	}
	return false
}

// libusbContext returns libusb_context. It initializes context on demand.
func libusbContext() (*C.libusb_context, error) {
	// Acquire libusbContextLock
	libusbContextLock.Lock()
	defer libusbContextLock.Unlock()

	// Initialize libusb_context on demand
	if libusbContextPtr == nil {
		// Obtain libusb_context
		rc := C.libusb_init(&libusbContextPtr)
		if rc != 0 {
			err := libusbError{"libusb_init", int(rc)}
			return nil, err
		}

		// Start libusb thread (required for hotplug and
		// asynchronous I/O)
		go func() {
			runtime.LockOSThread()
			for {
				C.libusb_handle_events(libusbContextPtr)
			}
		}()

	}

	return libusbContextPtr, nil
}
