// MFP - Miulti-Function Printers and scanners toolkit
// Utility functions and data BLOBs for testing
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP BLOB parser

package testutils

import "github.com/OpenPrinting/goipp"

func ippMustParse(blob []byte) *goipp.Message {
	msg := &goipp.Message{}
	err := msg.DecodeBytesEx(blob,
		goipp.DecoderOptions{EnableWorkarounds: true})

	if err != nil {
		panic(err)
	}

	return msg
}
