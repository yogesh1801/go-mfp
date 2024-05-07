// MFP   - Miulti-Function Printers and scanners toolkit
// mains - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// mfp-shell command implementation

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/mainfunc"
	"github.com/peterh/liner"
)

// main function for the mfp-shell command
func main() {
	// Setup liner library
	editline := liner.NewLiner()
	defer editline.Close()

	editline.SetCtrlCAborts(true)

	// Setup history
	historyPath := mainfunc.PathUserConfDir("mfp")
	os.MkdirAll(historyPath, 0755)

	historyPath = filepath.Join(historyPath, "mfp-shell.history")

	if file, err := os.Open(historyPath); err == nil {
		editline.ReadHistory(file)
		file.Close()
	}

	// Read and execute line by line
	for {
		line, err := editline.Prompt("MFP> ")
		if err != nil {
			fmt.Printf("\n")
			break
		}

		savehistory, err := exec(line)
		if savehistory {
			editline.AppendHistory(strings.Trim(line, " "))
			if file, err := os.Create(historyPath); err == nil {
				editline.WriteHistory(file)
				file.Close()
			}

		}

		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
}

// exec parses and executes the command.
//
// Returned savehistory is true if line is "good enough" to
// be saved to the history file.
func exec(line string) (savehistory bool, err error) {
	// Tokenize string
	tokens, err := argv.Tokenize(line)
	if err != nil {
		return false, err
	}

	// Ignore empty lines
	if len(tokens) == 0 {
		return false, nil
	}

	// Update history

	// Lookup the command
	cmd := mainfunc.CommandByName(tokens[0])
	if cmd == nil {
		err = fmt.Errorf("%q: command not found", tokens[0])
		return true, err
	}

	cmd.Main(tokens)

	return true, nil
}
