// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by Mohammed Imaduddin (mdimad005@gmail.com)
// See LICENSE for license terms and conditions
//
// USB/IP emulation logic

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

// ProtocolVersion is the protocol version
const ProtocolVersion = 0x0111

// OpCode represents the operation code
type OpCode uint16

const (
	OpReqDevlist OpCode = 0x8005 // OP_REQ_DEVLIST request
	OpRepDevlist OpCode = 0x0005 // OP_REP_DEVLIST reply
	OpReqImport  OpCode = 0x8003 // OP_REQ_IMPORT request
	OpRepImport  OpCode = 0x0003 // OP_REP_IMPORT reply
)

// Direction indicates the transfer direction
type Direction uint32

// Direction values
const (
	DirectionOut Direction = 0 // Client->device
	DirectionIn  Direction = 1 // Device->client
)

// USBIPHeader represents the common USB/IP protocol header used in communication.
type USBIPHeader struct {
	Version uint16 // The protocol version
	Command OpCode // Operation code
	Status  uint32
}

// Pack serializes the USBIPHeader into a byte slice using big-endian encoding.
func (h *USBIPHeader) Pack() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, h.Version)
	binary.Write(buf, binary.BigEndian, h.Command)
	binary.Write(buf, binary.BigEndian, h.Status)
	return buf.Bytes()
}

// Unpack deserializes the given byte slice into the USBIPHeader fields using big-endian encoding.
func (h *USBIPHeader) Unpack(data []byte) {
	buf := bytes.NewReader(data)
	binary.Read(buf, binary.BigEndian, &h.Version)
	binary.Read(buf, binary.BigEndian, &h.Command)
	binary.Read(buf, binary.BigEndian, &h.Status)
}

// Size returns the fixed size in bytes of a packed USBIPHeader.
func (h *USBIPHeader) Size() int {
	return 8
}

// USBInterface represents a USB interface descriptor containing class and protocol information.
type USBInterface struct {
	BInterfaceClass    uint8
	BInterfaceSubClass uint8
	BInterfaceProtocol uint8
	Align              uint8
}

// Pack serializes the USBInterface into a byte slice using big-endian encoding.
func (u *USBInterface) Pack() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, u.BInterfaceClass)
	binary.Write(buf, binary.BigEndian, u.BInterfaceSubClass)
	binary.Write(buf, binary.BigEndian, u.BInterfaceProtocol)
	binary.Write(buf, binary.BigEndian, u.Align)
	return buf.Bytes()
}

// OPREPDevList represents the device list reply structure in the USB/IP protocol.
type OPREPDevList struct {
	Base                USBIPHeader
	NExportedDevice     uint32
	UsbPath             [256]byte
	BusID               [32]byte
	Busnum              uint32
	Devnum              uint32
	Speed               uint32
	IDVendor            uint16
	IDProduct           uint16
	BcdDevice           uint16
	BDeviceClass        uint8
	BDeviceSubClass     uint8
	BDeviceProtocol     uint8
	BConfigurationValue uint8
	BNumConfigurations  uint8
	Interfaces          []USBInterface
}

// Pack serializes the OPREPDevList into a byte slice using big-endian encoding.
func (o *OPREPDevList) Pack() []byte {
	buf := new(bytes.Buffer)
	buf.Write(o.Base.Pack())
	binary.Write(buf, binary.BigEndian, o.NExportedDevice)
	buf.Write(o.UsbPath[:])
	buf.Write(o.BusID[:])
	binary.Write(buf, binary.BigEndian, o.Busnum)
	binary.Write(buf, binary.BigEndian, o.Devnum)
	binary.Write(buf, binary.BigEndian, o.Speed)
	binary.Write(buf, binary.BigEndian, o.IDVendor)
	binary.Write(buf, binary.BigEndian, o.IDProduct)
	binary.Write(buf, binary.BigEndian, o.BcdDevice)
	binary.Write(buf, binary.BigEndian, o.BDeviceClass)
	binary.Write(buf, binary.BigEndian, o.BDeviceSubClass)
	binary.Write(buf, binary.BigEndian, o.BDeviceProtocol)
	binary.Write(buf, binary.BigEndian, o.BConfigurationValue)
	binary.Write(buf, binary.BigEndian, o.BNumConfigurations)
	binary.Write(buf, binary.BigEndian, uint8(len(o.Interfaces)))
	for _, iff := range o.Interfaces {
		buf.Write(iff.Pack())
	}

	return buf.Bytes()
}

// OPREPImport represents the structure used to respond to an import request in the USB/IP protocol.
type OPREPImport struct {
	Base                USBIPHeader
	UsbPath             [256]byte
	BusID               [32]byte
	Busnum              uint32
	Devnum              uint32
	Speed               uint32
	IDVendor            uint16
	IDProduct           uint16
	BcdDevice           uint16
	BDeviceClass        uint8
	BDeviceSubClass     uint8
	BDeviceProtocol     uint8
	BConfigurationValue uint8
	BNumConfigurations  uint8
	BNumInterfaces      uint8
}

// Pack serializes the OPREPImport into a byte slice using big-endian encoding.
func (o *OPREPImport) Pack() []byte {
	buf := new(bytes.Buffer)
	buf.Write(o.Base.Pack())
	buf.Write(o.UsbPath[:])
	buf.Write(o.BusID[:])
	binary.Write(buf, binary.BigEndian, o.Busnum)
	binary.Write(buf, binary.BigEndian, o.Devnum)
	binary.Write(buf, binary.BigEndian, o.Speed)
	binary.Write(buf, binary.BigEndian, o.IDVendor)
	binary.Write(buf, binary.BigEndian, o.IDProduct)
	binary.Write(buf, binary.BigEndian, o.BcdDevice)
	binary.Write(buf, binary.BigEndian, o.BDeviceClass)
	binary.Write(buf, binary.BigEndian, o.BDeviceSubClass)
	binary.Write(buf, binary.BigEndian, o.BDeviceProtocol)
	binary.Write(buf, binary.BigEndian, o.BConfigurationValue)
	binary.Write(buf, binary.BigEndian, o.BNumConfigurations)
	binary.Write(buf, binary.BigEndian, o.BNumInterfaces)
	return buf.Bytes()
}

// USBIPRETSubmit represents the USB/IP return submit packet structure.
type USBIPRETSubmit struct {
	Command         uint32
	Seqnum          uint32
	Devid           uint32
	Direction       uint32
	Ep              uint32
	Status          uint32
	ActualLength    uint32
	StartFrame      uint32
	NumberOfPackets uint32
	ErrorCount      uint32
	Padding         uint64
	Data            []byte
}

// Pack serializes the USBIPRETSubmit structure into a byte slice using big-endian encoding.
func (u *USBIPRETSubmit) Pack() []byte {

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, u.Command)
	binary.Write(buf, binary.BigEndian, u.Seqnum)
	binary.Write(buf, binary.BigEndian, u.Devid)
	binary.Write(buf, binary.BigEndian, u.Direction)
	binary.Write(buf, binary.BigEndian, u.Ep)
	binary.Write(buf, binary.BigEndian, u.Status)
	binary.Write(buf, binary.BigEndian, u.ActualLength)
	binary.Write(buf, binary.BigEndian, u.StartFrame)
	binary.Write(buf, binary.BigEndian, u.NumberOfPackets)
	binary.Write(buf, binary.BigEndian, u.ErrorCount)
	binary.Write(buf, binary.BigEndian, u.Padding)
	buf.Write(u.Data)
	return buf.Bytes()
}

// Size returns the fixed size (in bytes) of the USBIPRETSubmit header.
func (u *USBIPRETSubmit) Size() int {

	return 48
}

// USBIPCMDSubmit represents the USB/IP command submit packet structure.
type USBIPCMDSubmit struct {
	Command              uint32
	Seqnum               uint32
	Devid                uint32
	Direction            Direction
	Ep                   uint32
	TransferFlags        uint32
	TransferBufferLength uint32
	StartFrame           uint32
	NumberOfPackets      uint32
	Interval             uint32
	Setup                [8]byte
}

// Unpack deserializes the byte slice into a USBIPCMDSubmit structure using big-endian encoding.
func (u *USBIPCMDSubmit) Unpack(data []byte) {

	buf := bytes.NewReader(data)
	binary.Read(buf, binary.BigEndian, &u.Command)
	binary.Read(buf, binary.BigEndian, &u.Seqnum)
	binary.Read(buf, binary.BigEndian, &u.Devid)
	binary.Read(buf, binary.BigEndian, &u.Direction)
	binary.Read(buf, binary.BigEndian, &u.Ep)
	binary.Read(buf, binary.BigEndian, &u.TransferFlags)
	binary.Read(buf, binary.BigEndian, &u.TransferBufferLength)
	binary.Read(buf, binary.BigEndian, &u.StartFrame)
	binary.Read(buf, binary.BigEndian, &u.NumberOfPackets)
	binary.Read(buf, binary.BigEndian, &u.Interval)
	copy(u.Setup[:], data[40:48])
}

// Size returns the fixed size (in bytes) of a USBIPCMDSubmit structure.
func (u *USBIPCMDSubmit) Size() int {
	return 48
}

// StandardDeviceRequest represents a standard USB control request structure.
type StandardDeviceRequest struct {
	BmRequestType uint8
	BRequest      uint8
	WValue        uint16
	WIndex        uint16
	WLength       uint16
}

// Unpack parses a StandardDeviceRequest from a byte slice.
func (s *StandardDeviceRequest) Unpack(data []byte) {
	buf := bytes.NewReader(data)
	binary.Read(buf, binary.LittleEndian, &s.BmRequestType)
	binary.Read(buf, binary.LittleEndian, &s.BRequest)
	binary.Read(buf, binary.LittleEndian, &s.WValue)
	binary.Read(buf, binary.LittleEndian, &s.WIndex)
	binary.Read(buf, binary.LittleEndian, &s.WLength)
}

// DeviceDescriptor describes a USB device, including vendor/product IDs and configurations.
type DeviceDescriptor struct {
	BLength            uint8
	BDescriptorType    uint8
	BcdUSB             uint16
	BDeviceClass       uint8
	BDeviceSubClass    uint8
	BDeviceProtocol    uint8
	BMaxPacketSize0    uint8
	IDVendor           uint16
	IDProduct          uint16
	BcdDevice          uint16
	IManufacturer      uint8
	IProduct           uint8
	ISerialNumber      uint8
	BNumConfigurations uint8
}

// Pack serializes the DeviceDescriptor into a byte slice using little-endian format.
func (d DeviceDescriptor) Pack() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, d.BLength)
	binary.Write(buf, binary.LittleEndian, d.BDescriptorType)
	binary.Write(buf, binary.LittleEndian, d.BcdUSB)
	binary.Write(buf, binary.LittleEndian, d.BDeviceClass)
	binary.Write(buf, binary.LittleEndian, d.BDeviceSubClass)
	binary.Write(buf, binary.LittleEndian, d.BDeviceProtocol)
	binary.Write(buf, binary.LittleEndian, d.BMaxPacketSize0)
	binary.Write(buf, binary.LittleEndian, d.IDVendor)
	binary.Write(buf, binary.LittleEndian, d.IDProduct)
	binary.Write(buf, binary.LittleEndian, d.BcdDevice)
	binary.Write(buf, binary.LittleEndian, d.IManufacturer)
	binary.Write(buf, binary.LittleEndian, d.IProduct)
	binary.Write(buf, binary.LittleEndian, d.ISerialNumber)
	binary.Write(buf, binary.LittleEndian, d.BNumConfigurations)
	return buf.Bytes()
}

// DeviceConfiguration defines a configuration for a USB device including interfaces.
type DeviceConfiguration struct {
	BLength             uint8
	BDescriptorType     uint8
	WTotalLength        uint16
	BNumInterfaces      uint8
	BConfigurationValue uint8
	IConfiguration      uint8
	BmAttributes        uint8
	BMaxPower           uint8
	Interfaces          [][]InterfaceDescriptor
}

// Pack serializes the DeviceConfiguration into a byte slice.
func (d *DeviceConfiguration) Pack() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, d.BLength)
	binary.Write(buf, binary.LittleEndian, d.BDescriptorType)
	binary.Write(buf, binary.LittleEndian, d.WTotalLength)
	binary.Write(buf, binary.LittleEndian, d.BNumInterfaces)
	binary.Write(buf, binary.LittleEndian, d.BConfigurationValue)
	binary.Write(buf, binary.LittleEndian, d.IConfiguration)
	binary.Write(buf, binary.LittleEndian, d.BmAttributes)
	binary.Write(buf, binary.LittleEndian, d.BMaxPower)
	return buf.Bytes()
}

// BOSDescriptor defines a Binary Object Store descriptor for a USB device.
type BOSDescriptor struct {
	BLength         uint8
	BDescriptorType uint8
	WTotalLength    uint16
	BNumDeviceCaps  uint8
}

// Pack serializes the BOSDescriptor into a byte slice.
func (b *BOSDescriptor) Pack() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, b.BLength)
	binary.Write(buf, binary.LittleEndian, b.BDescriptorType)
	binary.Write(buf, binary.LittleEndian, b.WTotalLength)
	binary.Write(buf, binary.LittleEndian, b.BNumDeviceCaps)
	return buf.Bytes()
}

// DeviceQualifierDescriptor describes a high-speed capable device in other speeds.
type DeviceQualifierDescriptor struct {
	BLength            uint8
	BDescriptorType    uint8
	BcdUSB             uint16
	BDeviceClass       uint8
	BDeviceSubClass    uint8
	BDeviceProtocol    uint8
	BMaxPacketSize0    uint8
	BNumConfigurations uint8
	BReserved          uint8
}

// Pack serializes the DeviceQualifierDescriptor into a byte slice.
func (d *DeviceQualifierDescriptor) Pack() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, d.BLength)
	binary.Write(buf, binary.LittleEndian, d.BDescriptorType)
	binary.Write(buf, binary.LittleEndian, d.BcdUSB)
	binary.Write(buf, binary.LittleEndian, d.BDeviceClass)
	binary.Write(buf, binary.LittleEndian, d.BDeviceSubClass)
	binary.Write(buf, binary.LittleEndian, d.BDeviceProtocol)
	binary.Write(buf, binary.LittleEndian, d.BMaxPacketSize0)
	binary.Write(buf, binary.LittleEndian, d.BNumConfigurations)
	binary.Write(buf, binary.LittleEndian, d.BReserved)
	return buf.Bytes()
}

// InterfaceDescriptor describes a single interface within a configuration.
type InterfaceDescriptor struct {
	BLength            uint8
	BDescriptorType    uint8
	BInterfaceNumber   uint8
	BAlternateSetting  uint8
	BNumEndpoints      uint8
	BInterfaceClass    uint8
	BInterfaceSubClass uint8
	BInterfaceProtocol uint8
	IInterface         uint8
	ClassDescriptor    interface{ Pack() []byte }
	Endpoints          []EndpointDescriptor
}

// Pack serializes the InterfaceDescriptor into a byte slice.
func (i *InterfaceDescriptor) Pack() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i.BLength)
	binary.Write(buf, binary.LittleEndian, i.BDescriptorType)
	binary.Write(buf, binary.LittleEndian, i.BInterfaceNumber)
	binary.Write(buf, binary.LittleEndian, i.BAlternateSetting)
	binary.Write(buf, binary.LittleEndian, i.BNumEndpoints)
	binary.Write(buf, binary.LittleEndian, i.BInterfaceClass)
	binary.Write(buf, binary.LittleEndian, i.BInterfaceSubClass)
	binary.Write(buf, binary.LittleEndian, i.BInterfaceProtocol)
	binary.Write(buf, binary.LittleEndian, i.IInterface)
	return buf.Bytes()
}

// EndpointDescriptor describes a USB endpoint within an interface.
type EndpointDescriptor struct {
	BLength          uint8
	BDescriptorType  uint8
	BEndpointAddress uint8
	BmAttributes     uint8
	WMaxPacketSize   uint16
	BInterval        uint8
	ClassDescriptor  interface{ Pack() []byte }
}

// Pack serializes the EndpointDescriptor into a byte slice.
func (e *EndpointDescriptor) Pack() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, e.BLength)
	binary.Write(buf, binary.LittleEndian, e.BDescriptorType)
	binary.Write(buf, binary.LittleEndian, e.BEndpointAddress)
	binary.Write(buf, binary.LittleEndian, e.BmAttributes)
	binary.Write(buf, binary.LittleEndian, e.WMaxPacketSize)
	binary.Write(buf, binary.LittleEndian, e.BInterval)
	return buf.Bytes()
}

// USBRequest represents a USB transfer request received from the host.
type USBRequest struct {
	Seqnum               uint32
	Devid                uint32
	Direction            Direction
	Ep                   uint32
	Flags                uint32
	TransferBufferLength uint32
	NumberOfPackets      uint32
	Interval             uint32
	Setup                [8]byte
	TransferBuffer       []byte
}

// Device defines an interface for emulated USB devices.
type Device interface {
	GetConfigurations() []DeviceConfiguration
	GetDeviceDescriptor() DeviceDescriptor
	HandleData(usbReq USBRequest)
	HandleDeviceSpecificControl(controlReq StandardDeviceRequest, usbReq USBRequest)
	SetConnection(conn net.Conn)
}

// BaseUSBDevice provides a default implementation for core USB device behaviors.
type BaseUSBDevice struct {
	Connection        net.Conn
	AllConfigurations []byte
}

// SetConnection sets the TCP connection used to communicate with the host.
func (b *BaseUSBDevice) SetConnection(conn net.Conn) {
	b.Connection = conn
}

// GenerateRawConfiguration flattens and serializes the configuration descriptors.
func (b *BaseUSBDevice) GenerateRawConfiguration(device Device) {
	var allConfigurations []byte
	for _, configuration := range device.GetConfigurations() {
		allConfigurations = append(allConfigurations, configuration.Pack()...)
		for _, interfaceGroup := range configuration.Interfaces {
			for _, interfaceAlt := range interfaceGroup {
				allConfigurations = append(allConfigurations, interfaceAlt.Pack()...)
				if interfaceAlt.ClassDescriptor != nil {
					allConfigurations = append(allConfigurations, interfaceAlt.ClassDescriptor.Pack()...)
				}
				for _, endpoint := range interfaceAlt.Endpoints {
					allConfigurations = append(allConfigurations, endpoint.Pack()...)
					if endpoint.ClassDescriptor != nil {
						allConfigurations = append(allConfigurations, endpoint.ClassDescriptor.Pack()...)
					}
				}
			}
		}
	}
	b.AllConfigurations = allConfigurations
}

// SendUSBRet sends a USBIP_RET_Submit response back to the host.
func (b *BaseUSBDevice) SendUSBRet(usbReq USBRequest, usbRes []byte, usbLen int, status uint32) {
	fmt.Printf("Sending %s\n", BytesToString(usbRes))
	ret := &USBIPRETSubmit{
		Command:         0x3,
		Seqnum:          usbReq.Seqnum,
		Devid:           0,
		Direction:       0,
		Ep:              0,
		Status:          status,
		ActualLength:    uint32(usbLen),
		StartFrame:      0,
		NumberOfPackets: 0xffffffff,
		ErrorCount:      0,
		Padding:         0,
		Data:            usbRes,
	}
	b.Connection.Write(ret.Pack())
}

// HandleGetDescriptor processes a GET_DESCRIPTOR request from the host.
func (b *BaseUSBDevice) HandleGetDescriptor(device Device, controlReq StandardDeviceRequest, usbReq USBRequest) bool {
	descriptorType := uint8(controlReq.WValue >> 8)
	descriptorIndex := uint8(controlReq.WValue & 0xff)
	fmt.Printf("handle_get_descriptor %d %d\n", descriptorType, descriptorIndex)

	if descriptorType == 0x01 { // Device Descriptor
		ret := device.GetDeviceDescriptor().Pack()
		b.SendUSBRet(usbReq, ret, len(ret), 0)
		return true
	} else if descriptorType == 0x02 { // Configuration Descriptor
		ret := b.AllConfigurations
		if int(controlReq.WLength) < len(ret) {
			ret = ret[:controlReq.WLength]
		}
		b.SendUSBRet(usbReq, ret, len(ret), 0)
		return true
	}
	return false
}

// HandleSetConfiguration processes a SET_CONFIGURATION request from the host.
func (b *BaseUSBDevice) HandleSetConfiguration(device Device, controlReq StandardDeviceRequest, usbReq USBRequest) bool {
	fmt.Printf("handle_set_configuration %d\n", controlReq.WValue)
	b.SendUSBRet(usbReq, []byte{}, 0, 0)
	return true
}

// HandleUSBControl handles a standard USB control transfer on endpoint 0.
func (b *BaseUSBDevice) HandleUSBControl(device Device, usbReq USBRequest) {
	controlReq := StandardDeviceRequest{}
	controlReq.Unpack(usbReq.Setup[:])
	handled := false
	fmt.Printf("  UC Request Type %d\n", controlReq.BmRequestType)
	fmt.Printf("  UC Request %d\n", controlReq.BRequest)
	fmt.Printf("  UC Value  %d\n", controlReq.WValue)
	fmt.Printf("  UC Index  %d\n", controlReq.WIndex)
	fmt.Printf("  UC Length %d\n", controlReq.WLength)

	if controlReq.BmRequestType == 0x80 { // Data flows IN, from Device to Host
		if controlReq.BRequest == 0x00 { // GET_STATUS
			configurations := device.GetConfigurations()
			attributes := configurations[0].BmAttributes
			isSelfPowered := (attributes & (1 << 6)) != 0
			isRemoteWakeup := (attributes & (1 << 5)) != 0
			ret := uint16(0x0000)
			if isRemoteWakeup {
				ret |= (1 << 1)
			}
			if isSelfPowered {
				ret |= 1
			}
			retBytes := make([]byte, 2)
			binary.LittleEndian.PutUint16(retBytes, ret)
			b.SendUSBRet(usbReq, retBytes, 2, 0)
			handled = true
		} else if controlReq.BRequest == 0x06 { // GET_DESCRIPTOR
			handled = b.HandleGetDescriptor(device, controlReq, usbReq)
		}
	} else if controlReq.BmRequestType == 0x00 { // Data flows OUT, from Host to Device
		if controlReq.BRequest == 0x09 { // Set Configuration
			handled = b.HandleSetConfiguration(device, controlReq, usbReq)
		}
	}

	if !handled {
		device.HandleDeviceSpecificControl(controlReq, usbReq)
	}
}

// HandleUSBRequest routes a USB request to the appropriate handler.
func (b *BaseUSBDevice) HandleUSBRequest(device Device, usbReq USBRequest) {
	if usbReq.Ep == 0 { // Endpoint 0 is always the control endpoint
		b.HandleUSBControl(device, usbReq)
	} else {
		device.HandleData(usbReq)
	}
}

// BytesToString returns a hex-escaped string for a byte slice
func BytesToString(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	var result strings.Builder
	for _, b := range data {
		result.WriteString(fmt.Sprintf("\\x%02x", b))
	}
	return result.String()
}
