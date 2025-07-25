// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by Mohammed Imaduddin (mdimad005@gmail.com)
// See LICENSE for license terms and conditions
//
// IPP over USB emulation logic

package main

import (
	"fmt"
	"net"
)

// Server implements the USBIP server.
type Server struct {
	USBDevices []USBDevice
}

// AddUSBDevice adds a USBDevice to the container.
func (srv *Server) AddUSBDevice(device USBDevice) {
	srv.USBDevices = append(srv.USBDevices, device)
}

// HandleAttach handles a device import request from the USB/IP client.
func (srv *Server) HandleAttach() *OPREPImport {
	usbDev := srv.USBDevices[0]
	deviceDescriptor := usbDev.GetDeviceDescriptor()
	configurations := usbDev.GetConfigurations()

	var usbPath [256]byte
	copy(usbPath[:], "/sys/devices/pci0000:00/0000:00:01.2/usb1/1-1")

	var busID [32]byte
	copy(busID[:], "1-1")

	return &OPREPImport{
		Base:                USBIPHeader{Version: ProtocolVersion, Command: OpRepImport, Status: 0},
		UsbPath:             usbPath,
		BusID:               busID,
		Busnum:              1,
		Devnum:              2,
		Speed:               2,
		IDVendor:            deviceDescriptor.IDVendor,
		IDProduct:           deviceDescriptor.IDProduct,
		BcdDevice:           deviceDescriptor.BcdDevice,
		BDeviceClass:        deviceDescriptor.BDeviceClass,
		BDeviceSubClass:     deviceDescriptor.BDeviceSubClass,
		BDeviceProtocol:     deviceDescriptor.BDeviceProtocol,
		BConfigurationValue: configurations[0].BConfigurationValue,
		BNumConfigurations:  deviceDescriptor.BNumConfigurations,
		BNumInterfaces:      configurations[0].BNumInterfaces,
	}
}

// HandleDeviceList responds with a list of available USB devices.
func (srv *Server) HandleDeviceList() *OPREPDevList {
	usbDev := srv.USBDevices[0]
	deviceDescriptor := usbDev.GetDeviceDescriptor()
	configurations := usbDev.GetConfigurations()

	var usbPath [256]byte
	copy(usbPath[:], "/sys/devices/pci0000:00/0000:00:01.2/usb1/1-1")

	var busID [32]byte
	copy(busID[:], "1-1")

	list := &OPREPDevList{
		Base:                USBIPHeader{Version: ProtocolVersion, Command: OpRepDevlist, Status: 0},
		NExportedDevice:     1,
		UsbPath:             usbPath,
		BusID:               busID,
		Busnum:              1,
		Devnum:              2,
		Speed:               2,
		IDVendor:            deviceDescriptor.IDVendor,
		IDProduct:           deviceDescriptor.IDProduct,
		BcdDevice:           deviceDescriptor.BcdDevice,
		BDeviceClass:        deviceDescriptor.BDeviceClass,
		BDeviceSubClass:     deviceDescriptor.BDeviceSubClass,
		BDeviceProtocol:     deviceDescriptor.BDeviceProtocol,
		BConfigurationValue: configurations[0].BConfigurationValue,
		BNumConfigurations:  deviceDescriptor.BNumConfigurations,
		Interfaces: []USBInterface{
			{
				BInterfaceClass:    configurations[0].Interfaces[0][0].BInterfaceClass,
				BInterfaceSubClass: configurations[0].Interfaces[0][0].BInterfaceSubClass,
				BInterfaceProtocol: configurations[0].Interfaces[0][0].BInterfaceProtocol,
				Align:              0,
			},
		},
	}

	return list
}

// Run starts the USB/IP server and handles incoming connections.
func (srv *Server) Run(ip string, port int) {
	if ip == "" {
		ip = "0.0.0.0"
	}
	if port == 0 {
		port = 3240
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("Connection address: %s\n", conn.RemoteAddr().String())
		go srv.serve(conn)
	}
}

// serve manages a client connection.
func (srv *Server) serve(conn net.Conn) {
	req := USBIPHeader{}

	defer func() {
		fmt.Println("Close connection")
		conn.Close()
	}()

	var buf [64]byte // Enough for both USBIPHeader and USBIPCMDSubmit
	var busid [32]byte

	// Handle OpReqDevlist or OpReqImport commands
	err := connReadAll(conn, buf[:8])
	if err != nil {
		return
	}
	req.Unpack(buf[:8])
	fmt.Println("Header Packet")
	fmt.Printf("command: %x\n", req.Command)

	switch req.Command {
	case OpReqDevlist:
		fmt.Println("list of devices")
		connWriteAll(conn, srv.HandleDeviceList().Pack())
		return

	case OpReqImport:
		fmt.Println("attach device")
		err = connReadAll(conn, busid[:])
		if err != nil {
			return
		}

		err = connWriteAll(conn, srv.HandleAttach().Pack())
		if err != nil {
			return
		}
	}

	// Now handle USB I/O requests
	for {
		fmt.Println("----------------")
		fmt.Println("handles requests")
		cmd := USBIPCMDSubmit{}
		cmdHeaderData := buf[:cmd.Size()]
		err := connReadAll(conn, cmdHeaderData)
		if err != nil {
			break
		}
		cmd.Unpack(cmdHeaderData)

		var transferBuffer []byte
		if cmd.Direction == DirectionOut && cmd.TransferBufferLength > 0 {
			transferBuffer = make([]byte, cmd.TransferBufferLength)
			err = connReadAll(conn, transferBuffer)
			if err != nil {
				return
			}
		}

		fmt.Printf("usbip cmd %x\n", cmd.Command)
		fmt.Printf("usbip seqnum %x\n", cmd.Seqnum)
		fmt.Printf("usbip devid %x\n", cmd.Devid)
		fmt.Printf("usbip direction %x\n", cmd.Direction)
		fmt.Printf("usbip ep %x\n", cmd.Ep)
		fmt.Printf("usbip flags %x\n", cmd.TransferFlags)
		fmt.Printf("usbip transfer buffer length %x\n", cmd.TransferBufferLength)
		fmt.Printf("usbip start %x\n", cmd.StartFrame)
		fmt.Printf("usbip number of packets %x\n", cmd.NumberOfPackets)
		fmt.Printf("usbip interval %x\n", cmd.Interval)
		fmt.Printf("usbip setup %s\n", BytesToString(cmd.Setup[:]))
		fmt.Printf("usbip transfer buffer %s\n", BytesToString(transferBuffer))

		usbReq := USBRequest{
			Seqnum:               cmd.Seqnum,
			Devid:                cmd.Devid,
			Direction:            cmd.Direction,
			Ep:                   cmd.Ep,
			Flags:                cmd.TransferFlags,
			TransferBufferLength: cmd.TransferBufferLength,
			NumberOfPackets:      cmd.NumberOfPackets,
			Interval:             cmd.Interval,
			Setup:                cmd.Setup,
			TransferBuffer:       transferBuffer,
		}

		srv.USBDevices[0].SetConnection(conn)
		if baseDevice, ok := srv.USBDevices[0].(interface{ HandleUSBRequest(USBDevice, USBRequest) }); ok {
			baseDevice.HandleUSBRequest(srv.USBDevices[0], usbReq)
		}
	}
}
