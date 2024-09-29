// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common device information

package discovery

import (
	"net/netip"
	"slices"

	"github.com/alexpevzner/mfp/uuid"
)

// Device consist of the multiple functional units. There are
// three types of units:
//   - [PrintUnit], for printing
//   - [ScanUnit], for scanning
//   - [FaxoutUnit], for sending faxes.
//
// Multiple units of each type may exist, and depending on the device,
// they may have different parameters.
//
// Each unit has its unique [UnitID], the combination of parameters,
// that uniquely identifies the unit.
type Device struct {
	// Device metadata
	MakeModel string    // Device make and model
	DNSSDName string    // DNS-SD name, "" if none
	DNSSDUUID uuid.UUID // DNS-SD UUID, uuid.NilUUID if n/a

	// USBManufacturer and USBModel may be available for DNS-DS
	// devices too, as 'usb_MFG' and 'usb_MDL' TXT records. Apple use
	// these parameters for PDL file selection as a fallback, if
	// 'product' is not available in the TXT.
	//
	// Please notice, it is not necessary true that MakeModel
	// is the concatenation of these two strings.
	USBManufacturer string // Manufacturer name
	USBModel        string // Model name

	// USBSerial may be available for the ipp-usb devices too.
	USBSerial string // USB serial number, "" if n/a

	// Connectivity
	Addrs []netip.Addr // Device's IP addresses

	// Device units
	PrintUnits  []PrintUnit  // Print units
	ScanUnits   []ScanUnit   // Scan units
	FaxoutUnits []FaxoutUnit // Faxout units
}

// device is the internal representation of the Device
type device struct {
	units []unit       // Device's units
	addrs []netip.Addr // Device's IP addresses
}

// Export exports device as Device
func (dev device) Export() Device {
	out := Device{Addrs: dev.addrs}

	// Classify units
	var ippPrinters []*unit
	var lpdPrinters []*unit
	var appsockPrinters []*unit
	var usbPrinters []*unit
	var ippScanners []*unit
	var esclSanners []*unit
	var wsdScanners []*unit
	var ippFaxes []*unit

	for i := range dev.units {
		un := &dev.units[i]
		switch un.id.SvcType {
		case ServicePrinter:
			switch un.id.SvcProto {
			case ServiceIPP:
				ippPrinters = append(ippPrinters, un)
			case ServiceLPD:
				lpdPrinters = append(lpdPrinters, un)
			case ServiceAppSocket:
				appsockPrinters = append(appsockPrinters, un)
			case ServiceUSB:
				usbPrinters = append(usbPrinters, un)
			}

		case ServiceScanner:
			switch un.id.SvcProto {
			case ServiceIPP:
				ippScanners = append(ippScanners, un)
			case ServiceESCL:
				esclSanners = append(esclSanners, un)
			case ServiceWSD:
				wsdScanners = append(wsdScanners, un)
			}

		case ServiceFaxout:
			switch un.id.SvcProto {
			case ServiceIPP:
				ippFaxes = append(ippFaxes, un)
			}
		}
	}

	// Merge by classes
	printUnits := slices.Concat(
		ippPrinters,
		lpdPrinters,
		appsockPrinters,
		usbPrinters,
	)

	scanUnits := slices.Concat(
		ippScanners,
		esclSanners,
		wsdScanners,
	)

	faxoutUnits := ippFaxes

	// Convert units to external representation and save to device.
	for _, un := range printUnits {
		exp := un.Export().(PrintUnit)
		out.PrintUnits = append(out.PrintUnits, exp)
	}

	for _, un := range scanUnits {
		exp := un.Export().(ScanUnit)
		out.ScanUnits = append(out.ScanUnits, exp)
	}

	for _, un := range faxoutUnits {
		exp := un.Export().(FaxoutUnit)
		out.FaxoutUnits = append(out.FaxoutUnits, exp)
	}

	// Extract metadata
	dnssdUnits := slices.Concat(
		ippPrinters,
		lpdPrinters,
		appsockPrinters,
		ippScanners,
		esclSanners,
		ippFaxes,
	)

	allUnits := slices.Concat(
		ippPrinters,
		lpdPrinters,
		appsockPrinters,
		usbPrinters,
		ippScanners,
		esclSanners,
		wsdScanners,
		ippFaxes,
	)

	for _, un := range dnssdUnits {
		if un.meta.MakeModel != "" {
			out.MakeModel = un.meta.MakeModel
			break
		}
	}

	for _, un := range dnssdUnits {
		if un.meta.Manufacturer != "" && un.meta.Model != "" {
			out.USBManufacturer = un.meta.Manufacturer
			out.USBModel = un.meta.Model
			break
		}
	}

	for _, un := range dnssdUnits {
		if un.id.DNSSDName != "" && un.id.UUID != uuid.NilUUID {
			out.DNSSDName = un.id.DNSSDName
			out.DNSSDUUID = un.id.UUID
			break
		}
	}

	for _, un := range allUnits {
		if un.id.USBSerial != "" {
			out.USBSerial = un.id.USBSerial
		}
	}

	return out
}
