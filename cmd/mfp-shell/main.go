// MFP           - Miulti-Function Printers and scanners toolkit
// cmd/mfp-shell - Interactive shell.
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
	"github.com/alexpevzner/mfp/commands"
	"github.com/peterh/liner"
)

// main function for the mfp-shell command
func main() {
	// Setup liner library
	editline := liner.NewLiner()
	defer editline.Close()

	editline.SetCtrlCAborts(true)

	// Setup completion
	editline.SetCompleter(completer)

	// Setup history
	historyPath := commands.PathUserConfDir("mfp")
	os.MkdirAll(historyPath, 0755)

	historyPath = filepath.Join(historyPath, "mfp-shell.history")

	if file, err := os.Open(historyPath); err == nil {
		editline.ReadHistory(file)
		file.Close()
	}

	// Read and execute line by line
	fmt.Println("MFP interactive console.")
	fmt.Println("Confused? Say help!")
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

// completer handles command-line completion
func completer(line string) (out []string) {
	args, tail, tailspace, _ := argv.TokenizeEx(line)
	if tailspace != "" {
		args = append(args, "")
	}

	strip := len(tail)
	if len(args) > 0 {
		strip += len(args[len(args)-1])
	}

	line = line[:len(line)-strip]
	compl, flags := commands.CmdMfp.Complete(args)

	space := " "
	if flags&argv.CompleterNoSpace != 0 {
		space = ""
	}

	for _, c := range compl {
		out = append(out, line+c+space)
	}

	return
}

// exec parses and executes the command.
//
// Returned savehistory is true if line is "good enough" to
// be saved to the history file.
func exec(line string) (savehistory bool, err error) {
	// Tokenize string
	argv, err := argv.Tokenize(line)
	if err != nil {
		return false, err
	}

	// Ignore empty lines
	if len(argv) == 0 {
		return false, nil
	}

	// Execute the command
	err = commands.CmdMfp.Run(argv)

	return true, err
}
