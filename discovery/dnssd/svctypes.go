// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Service types we are interested in

package dnssd

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
