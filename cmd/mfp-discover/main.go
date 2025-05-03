// MFP - Miulti-Function Printers and scanners toolkit
// mfp-discover: device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The main() function.

package main

import "github.com/OpenPrinting/go-mfp/cmd/mfp-discover/discover"

// main function for the mfp-discover command
func main() {
	discover.Command.Main(nil)
}
