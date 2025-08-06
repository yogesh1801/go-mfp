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
	"errors"
	"fmt"
	"net"
	"sync"
	"syscall"

	"github.com/OpenPrinting/go-mfp/log"
)

// Server implements the USBIP server.
type Server struct {
	devices    [USBMaxDevices + 1]devslot
	nextslot   int
	byid       map[devBusID]*devslot    // Slots by devBusID
	bylocation map[devLocation]*devslot // Slots by devLocation
	lock       sync.Mutex
}

// devslot represents a device slot.
type devslot struct {
	dev      *Device
	busid    devBusID
	location devLocation
	pconn    *protoConn
}

// NewServer creates a new [Server]
func NewServer() *Server {
	srv := &Server{
		nextslot:   1,
		byid:       make(map[devBusID]*devslot, USBMaxDevices+1),
		bylocation: make(map[devLocation]*devslot, USBMaxDevices+1),
	}

	for i := range srv.devices {
		slot := &srv.devices[i]
		slot.location.Bus = 1 // Hardcoded for now
		slot.location.Dev = i
		slot.busid = slot.location.BusID()

		srv.byid[slot.busid] = slot
		srv.bylocation[slot.location] = slot
	}

	return srv
}

// AddDevice adds a USBDevice to the container.
func (srv *Server) AddDevice(dev *Device) error {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	// Look for the empty slot.
	for i := range srv.devices {
		next := (srv.nextslot + i) % len(srv.devices)
		if next != 0 && srv.devices[next].dev == nil {
			srv.devices[next].dev = dev
			srv.nextslot = (next + 1) & len(srv.devices)
			return nil
		}
	}

	return errors.New("Can't add new device. All slots are busy")
}

// doImport handles a device import request from the USB/IP client.
// It returns true if device was actually attached and the protocol
// response message in any case.
func (srv *Server) doImport(pconn *protoConn, rq *protoImportRequest) *protoImportResponse {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	// Find the device
	slot := srv.byid[rq.BusID]
	if slot == nil || slot.dev == nil || slot.pconn != nil {
		return &protoImportResponse{Status: protoHsError}
	}

	// Attach device to connection
	slot.pconn = pconn

	// Build the response
	rsp := &protoImportResponse{
		Status: protoHsOK,
		DevInfo: &devInfo{
			Location: slot.location,
			Device:   slot.dev,
		},
	}

	return rsp
}

// doDevlist responds with a list of available USB devices.
// It returns the protocol response message.
func (srv *Server) doDevlist() *protoDevlistResponse {
	// Gather connected but not yet exported slots
	srv.lock.Lock()

	slots := make([]devslot, 0, len(srv.devices))
	for _, slot := range srv.devices {
		if slot.dev != nil && slot.pconn == nil {
			slots = append(slots, slot)
		}
	}

	srv.lock.Unlock()

	// Generate a response
	rsp := &protoDevlistResponse{
		Status:  protoHsOK,
		Devlist: make([]devInfo, len(slots)),
	}

	for i := range slots {
		rsp.Devlist[i].Location = slots[i].location
		rsp.Devlist[i].Device = slots[i].dev
	}

	return rsp
}

// doDisconnect handles the client disconnect event.
func (srv *Server) doDisconnect(pconn *protoConn) {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	for i := range srv.devices {
		if srv.devices[i].pconn == pconn {
			srv.devices[i].pconn = nil
		}
	}
}

// Run starts the USB/IP server and handles incoming connections.
func (srv *Server) Run(ctx context.Context, ip string, port int) {
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
		go srv.Serve(ctx, conn)
	}
}

// Serve manages a client connection.
func (srv *Server) Serve(ctx context.Context, conn net.Conn) {
	log.Debug(ctx, "Connection: from %s", conn.RemoteAddr())

	// Create protocol connection
	pconn := newProtoConn(conn)

	defer func() {
		srv.doDisconnect(pconn)
		log.Debug(ctx, "Connection: closed")
		pconn.Close()
	}()

	devinfo := srv.handshake(ctx, pconn)
	if devinfo == nil {
		return
	}

	srv.io(ctx, pconn, devinfo)
}

// handshake performs the USBIP handshake exchange.
func (srv *Server) handshake(ctx context.Context, pconn *protoConn) *devInfo {
	rq, err := pconn.RecvHandshake()
	if err != nil {
		log.Debug(ctx, "Handshake: %s", err)
		return nil
	}

	log.Debug(ctx, "Handshake: < %s", rq)

	var rsp protoHandshakeResponse

	switch rq := rq.(type) {
	case *protoDevlistRequest:
		rsp = srv.doDevlist()
	case *protoImportRequest:
		rsp = srv.doImport(pconn, rq)
	}

	if rsp != nil {
		log.Debug(ctx, "Handshake: > %s", rsp)
		err = pconn.SendHandshake(rsp)
		if err != nil {
			log.Debug(ctx, "Handshake: %s", err)
			return nil
		}
	}

	if rsp, ok := rsp.(*protoImportResponse); ok && rsp.Status == protoHsOK {
		return rsp.DevInfo
	}

	return nil
}

// io performs USBIO I/O exchange.
func (srv *Server) io(ctx context.Context, pconn *protoConn, devinfo *devInfo) {
	dev := devinfo.Device
	defer dev.shutdown()

	// Map of the pending protoIOSubmitRequest by Seqnum.
	//
	// We need it, because protoIOUnlinkRequest lacks some
	// important information (endpoint number, for example).
	submitsBySeqnum := make(map[uint32]*protoIOSubmitRequest)

	// Completion callback for protoIOSubmitRequest
	completion := func(rq *protoIOSubmitRequest, err syscall.Errno) {
		delete(submitsBySeqnum, rq.Seqnum)

		rsp := rq.Response(err)
		log.Debug(ctx, "IO:   > %s", rsp)
		pconn.SendIO(rsp)
	}

	for {
		rq, err := pconn.RecvIO()
		if err != nil {
			log.Debug(ctx, "IO:   < %s", err)
			return
		}

		log.Debug(ctx, "IO:   < %s", rq)

		hdr := rq.Header()
		if hdr.Location != devinfo.Location {
			log.Debug(ctx, "IO:   < Device: expected %s, present %s",
				devinfo.Location, hdr.Location)

			// FIXME - reject request and continue
			return
		}

		switch rq := rq.(type) {
		case *protoIOSubmitRequest:
			// Control requests are processed instantly.
			if rq.Endpoint == 0 {
				data, err := srv.control(ctx, dev, rq)

				rq.Buffer = data
				rq.actualLength = len(data)

				rsp := rq.Response(err)
				pconn.SendIO(rsp)
				break
			}

			submitsBySeqnum[rq.Seqnum] = rq

			rq.completion = completion
			err := dev.submit(rq)
			if err != 0 {
				rsp := rq.Response(err)
				pconn.SendIO(rsp)
			}

		case *protoIOUnlinkRequest:
			err := syscall.Errno(0)

			submit := submitsBySeqnum[rq.Seqnum]
			if submit != nil {
				// This information is missed in
				// the USBIP_CMD_SUBMIT.
				rq.Endpoint = submit.Endpoint

				err = dev.unlink(rq)
			}

			rsp := rq.Response(err)
			pconn.SendIO(rsp)
		}
	}
}

// control handles USB control requests.
func (srv *Server) control(ctx context.Context,
	dev *Device, rq *protoIOSubmitRequest) ([]byte, syscall.Errno) {

	var setup USBSetupPacket
	setup.Decode(rq.Setup)

	log.Debug(ctx, "CTRL: < %s", setup)
	var data []byte
	var err syscall.Errno

	switch setup.RequestType & USBRecipientMask {
	case USBRecipientDevice:
		switch setup.Request {
		case USBRequestGetStatus:
			data, err = dev.getStatus()

		case USBRequestGetDescriptor:
			t := USBDescriptorType(setup.WValue >> 8)
			i := int(setup.WValue & 255)
			data, err = dev.getDescriptor(t, i)

		case USBRequestGetConfiguration:
			data, err = dev.getConfiguration()

		case USBRequestSetConfiguration:
			n := setup.WValue
			data, err = dev.setConfiguration(int(n))
		}

	case USBRecipientInterface:
		ifn := int(setup.WIndex)
		alt := int(setup.WValue)

		switch setup.Request {
		case USBRequestGetStatus:
			data, err = dev.getInterfaceStatus(ifn)

		case USBRequestGetInterface:
			data, err = dev.getInterface(ifn)

		case USBRequestSetInterface:
			data, err = dev.setInterface(ifn, alt)
		}
	}

	if data != nil {
		log.Debug(ctx, "CTRL: > %d bytes returned", len(data))
	} else {
		err = syscall.EPIPE
		log.Debug(ctx, "CTRL: > %s", err)
	}

	return data, err
}
