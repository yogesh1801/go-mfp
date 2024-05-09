// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The 'help' command

package argv

import (
	"errors"
	"io"
	"os"
)

var (
	// HelpOption intended to be used as Option in Command definition
	// to indicate that the Command implements commonly used -h and --help
	// flags.
	HelpOption = Option{
		Name:    "-h",
		Aliases: []string{"--help"},
		Help:    "print help page",
	}

	// HelpCommand to be used as SubCommand in Command definition
	// to indicate that the Command implements commonly used "help"
	// sub-command.
	HelpCommand = Command{
		Name: "help",
		Help: "print help page",
		Parameters: []Parameter{
			{
				Name: "[command]",
			},
		},
		Handler: HelpHandler,
	}

	// HelpOutput is where help output is written
	HelpOutput io.Writer = os.Stdout
)

// HelpHandler is the standard Handler for 'help' Command
func HelpHandler(inv *Invocation) error {
	parent := inv.Parent()
	if parent == nil {
		return errors.New("HelpHandler must be used in sub-command")
	}

	cmd := inv.parent.Cmd()
	name, ok := inv.Get("[command]")
	if ok {
		subcmd, err := cmd.FindSubCommand(name)
		if err != nil {
			return err
		}
		cmd = subcmd
	}

	Help(cmd, HelpOutput)

	return nil
}
