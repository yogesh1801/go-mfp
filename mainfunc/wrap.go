// MFP      - Miulti-Function Printers and scanners toolkit
// mainfunc - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Main function wrapper

package mainfunc

import (
	"fmt"
	"os"
)

// Wrap is the universal wrapper for command's main functions
// which converts such a function into the body of the standalone
// command.
//
// Usage:
//
//   // main function for the mfp-cups command
//   func main() {
//           mainfunc.Wrap(mainfunc.MainCups)
//   }
func Wrap(main func(argv []string) error) {
	err := main(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
