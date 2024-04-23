// MFP   - Miulti-Function Printers and scanners toolkit
// mains - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// mfp-cups command implementation

package main

import (
	"os"

	"github.com/alexpevzner/mfp/mains"
)

// main function for the mfp-cups command
func main() {
	mains.MainCups(os.Args)
}
