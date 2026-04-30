// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// WS-Scan units and conversion constants

package wsscan

// wsscanDPI is the resolution at which WS-Scan dimensions are expressed
// (thousandths of an inch, i.e. 1/1000").
const wsscanDPI = 1000

// Brightness and Contrast adjustment range, as mandated by the WS-Scan spec.
// All scan services must support the full range; 0 means no adjustment.
const (
	brightnessMin = -1000
	brightnessMax = 1000
	contrastMin   = -1000
	contrastMax   = 1000
)
