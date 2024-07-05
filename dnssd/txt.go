// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// DNS-SD TXT records

package dnssd

// TxtRecord represents a single TXT Key=Value record
type TxtRecord struct {
	Key, Value string // TXT entry: Key=Value
}
