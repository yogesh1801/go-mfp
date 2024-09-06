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
	svcTypeLPR       = "_printer._tcp"        // LPR printer
	svcTypeESCL      = "_uscan._tcp"          // eSCL scan
	svcTypeESCLS     = "_uscans._tcp"         // eSCL scan over HTTPS
)

// Service types we are interested in
var svcTypes = []string{
	svcTypeAppSocket,
	svcTypeIPP,
	svcTypeIPPS,
	svcTypeLPR,
	svcTypeESCL,
	svcTypeESCLS,
}

// svcTypeToKind returns discovery.UnitKind for the service type
func svcTypeToKind(svcType string) discovery.UnitKind {
	return svcTypeToKindMap[svcType]
}

// svcTypeIsSecure reports if service type uses encrypted connection
func svcTypeIsSecure(svcType string) bool {
	switch svcType {
	case svcTypeIPPS, svcTypeESCLS:
		return true
	}
	return false
}

// svcTypeIsScan reports if service type is the scan service
func svcTypeIsScan(svcType string) bool {
	switch svcType {
	case svcTypeESCL, svcTypeESCLS:
		return true
	}
	return false
}

// svcTypeToKindMap maps svcType to discovery.UnitKind
var svcTypeToKindMap = map[string]discovery.UnitKind{
	svcTypeAppSocket: discovery.KindAppSocketPrinter,
	svcTypeIPP:       discovery.KindIPPPrinter,
	svcTypeIPPS:      discovery.KindIPPPrinter,
	svcTypeLPR:       discovery.KindLPRPrinter,
	svcTypeESCL:      discovery.KindESCLScanner,
	svcTypeESCLS:     discovery.KindESCLScanner,
}
