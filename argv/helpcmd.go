// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The standard 'help' command and '--help' option.

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
		Name:      "-h",
		Aliases:   []string{"--help"},
		Help:      "print help page",
		Immediate: HelpHandler,
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
	// If we run in the immediate mode (i.e., as result of -h option being
	// used with current command, our "target" is the current command.
	//
	// Otherwise, we are called as a sub-command of parent command, so we
	// must switch to that parent command.
	//
	// In the later case, if parent cannot be figured, it is an error. Looks
	// that somebody is calling HelpCommand directly, not as a part of
	// sub-commands hierarchy.
	var cmd *Command
	if inv.IsImmediate() {
		cmd = inv.Cmd()
	} else {
		parent := inv.Parent()
		if parent == nil {
			return errors.New("HelpHandler must be used in sub-command")
		}
		cmd = parent.Cmd()
	}

	// The 'help' command may have an optional parameter,
	// catch and handle it now
	name, ok := inv.Get("[command]")
	if ok {
		subcmd, err := cmd.FindSubCommand(name)
		if err != nil {
			return err
		}
		cmd = subcmd
	}

	// And if it is OK so far, it's a time to generate a help page
	Help(cmd, HelpOutput)

	return nil
}
