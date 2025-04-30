// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP enums

package ipp

import "reflect"

// kwRegisteredTypes lists all registered keyword types for IPP codec.
var enRegisteredTypes = map[reflect.Type]struct{}{
	reflect.TypeOf(EnPrinterType(0)): struct{}{},
}
