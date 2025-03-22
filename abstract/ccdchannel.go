// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan CCD color channel

package abstract

// CCDChannel specifies which CCD color channel to use for grayscale
// and monochrome scannig.
type CCDChannel int

// Known CCD Channels.
const (
	CCDChannelUnset           CCDChannel = iota // Not set
	CCDChannelRed                               // Use the RED channel
	CCDChannelGreen                             // Use the Green channel
	CCDChannelBlue                              // Use the Blue channel
	CCDChannelNTSC                              // NTSC-standard mix
	CCDChannelGrayCcd                           // Gray channel in hawdware
	CCDChannelGrayCcdEmulated                   // Emulated Gray
	ccdChannelMax
)
