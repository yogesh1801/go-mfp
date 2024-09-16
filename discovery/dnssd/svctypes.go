// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Service types we are interested in

package dnssd

import "github.com/alexpevzner/mfp/discovery"

// Service type names
const (
	svcTypeAppSocket = "_pdl-datastream._tcp" // AppSocket AKA JetDirect
	svcTypeIPP       = "_ipp._tcp"            // IPP over
	svcTypeIPPS      = "_ipps._tcp"           // IPP over HTTPS
	svcTypeLPD       = "_printer._tcp"        // LPD printer
	svcTypeESCL      = "_uscan._tcp"          // eSCL scan
	svcTypeESCLS     = "_uscans._tcp"         // eSCL scan over HTTPS
)

// Service types we are interested in
var svcTypes = []string{
	svcTypeAppSocket,
	svcTypeIPP,
	svcTypeIPPS,
	svcTypeLPD,
	svcTypeESCL,
	svcTypeESCLS,
}

// svcTypeToProto returns discovery.ServiceProto for the
// service type
func svcTypeToDiscoveryServiceProto(svcType string) discovery.ServiceProto {
	switch svcType {
	case svcTypeAppSocket:
		return discovery.ServiceAppSocket
	case svcTypeIPP, svcTypeIPPS:
		return discovery.ServiceIPP
	case svcTypeLPD:
		return discovery.ServiceLPD
	case svcTypeESCL, svcTypeESCLS:
		return discovery.ServiceESCL
	}

	panic("internal error")
}

// svcTypeIsPrinter reports if service type is the printer service
func svcTypeIsPrinter(svcType string) bool {
	return !svcTypeIsScanner(svcType)
}

// svcTypeIsScanner reports if service type is the scanner service
func svcTypeIsScanner(svcType string) bool {
	switch svcType {
	case svcTypeESCL, svcTypeESCLS:
		return true
	}
	return false
}
