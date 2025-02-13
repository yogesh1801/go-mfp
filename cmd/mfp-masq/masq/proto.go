// MFP - Miulti-Function Printers and scanners toolkit
// The "masq" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package masq

// proto identifies proxy protocol
type proto int

const (
	protoIPP proto = iota
	protoESCL
	protoWSD
)
