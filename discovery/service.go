// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Service type and protocol

package discovery

import "fmt"

// ServiceType represents a service type
type ServiceType int

// ServiceType constants:
const (
	ServicePrinter ServiceType = iota // Printer
	ServiceScanner                    // Scanner
	ServiceFaxout                     // Fax
)

// String returns ServiceType name, for debugging
func (t ServiceType) String() string {
	switch t {
	case ServicePrinter:
		return "printer"
	case ServiceScanner:
		return "scanner"
	case ServiceFaxout:
		return "fax"
	}

	return fmt.Sprintf("unknown (%d)", int(t))
}

// ServiceProto represents service protocol
type ServiceProto int

// ServiceProto constants:
const (
	ServiceIPP       ServiceProto = iota // IPP/IPPS printer or scanner
	ServiceESCL                          // ESCL scanner
	ServiceLPD                           // LPD printer
	ServiceAppSocket                     // AppSocket (JetDirect) printer
	ServiceWSD                           // WSD printer or scanner
	ServiceUSB                           // USB printer
)

// String returns ServiceProto name, for debugging
func (p ServiceProto) String() string {
	switch p {
	case ServiceIPP:
		return "IPP"
	case ServiceESCL:
		return "ESCL"
	case ServiceLPD:
		return "LPD"
	case ServiceAppSocket:
		return "AppSocket"
	case ServiceWSD:
		return "WSD"
	case ServiceUSB:
		return "USB"
	}

	return fmt.Sprintf("unknown (%d)", int(p))
}
