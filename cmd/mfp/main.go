// MFP   - Miulti-Function Printers and scanners toolkit
// mains - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// mfp-cups command implementation

package main

import (
	"fmt"
	"os"

	"github.com/alexpevzner/mfp/mainfunc"
)

// main function for the mfp-cups command
func main() {
	// Check command line
	if len(os.Args) != 2 {
		fmt.Printf("Usage: mfp command [args...]\n")
		fmt.Printf("Commands are:\n")
		for _, cmd := range mainfunc.Commands {
			fmt.Printf("    %s\n", cmd.Name)
		}
		os.Exit(1)
	}

	// Lookup the command
	cmd := mainfunc.CommandByName(os.Args[1])
	if cmd == nil {
		fmt.Printf("%s: command not found\n", os.Args[1])
		os.Exit(1)
	}

	// Call command's main function
	cmd.Main(os.Args[1:])
}
