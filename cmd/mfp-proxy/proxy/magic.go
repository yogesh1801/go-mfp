// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package proxy

import "bytes"

// magic returns file extension for known image formats.
// If format cannot be recognized, it returns "bin"
func magic(data []byte) string {
	for _, m := range magicTable {
		if bytes.HasPrefix(data, m.prefix) {
			return m.ext
		}
	}

	return "bin"
}

var magicTable = []struct {
	prefix []byte
	ext    string
}{
	{[]byte("BM"), "bmp"},
	{[]byte{0xff, 0xd8}, "jpeg"},
	{[]byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a}, "png"},
	{[]byte{'I', 'I', '*', 0}, "tiff"},
	{[]byte{'M', 'M', 0, '*'}, "tiff"},
	{[]byte("%PDF"), "pdf"},
	{[]byte("%!PS"), "pdf"},
}
