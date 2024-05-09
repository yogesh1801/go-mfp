// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command -- a command definition.

package argv

import (
	"errors"
	"fmt"
)

// Command defines a command.
//
// Every command MUST have a name and MAY have Options,
// positional Parameters and SubCommands
//
// It corresponds to the following usage syntax:
//
//   command [options] [params]
//   command [options] sub-command ...
//
// Parameters and SubCommands are mutually exclusive.
type Command struct {
	// Command name.
	Name string

	// Help string, a single-line description.
	Help string

	// Description contains a long command explanation.
	Description string

	// Options, if any.
	Options []Option

	// Positional parameters, if any.
	Parameters []Parameter

	// Sub-commands, if any.
	SubCommands []Command

	// The main function
	Main func(argv []string) error
}

// Verify checks correctness of Command definition. It fails if any
// error is found and returns description of the first caught error
func (cmd *Command) Verify() error {
	// Command must have a name
	if cmd.Name == "" {
		return errors.New("missed command name")
	}

	// Parameters and SubCommands are mutually exclusive
	if cmd.hasParameters() && cmd.hasSubCommands() {
		return fmt.Errorf(
			"%s: Parameters and SubCommands are mutually exclusive",
			cmd.Name)
	}

	// Verify Options and Parameters
	err := cmd.verifyOptions()
	if err == nil {
		err = cmd.verifyParameters()
	}

	if err != nil {
		return fmt.Errorf("%s: %s", cmd.Name, err)
	}

	// Verify SubCommands
	err = cmd.verifySubCommands()
	if err != nil {
		return fmt.Errorf("%s.%s", cmd.Name, err)
	}

	return err
}

// verifyOptions verifies command options
func (cmd *Command) verifyOptions() error {
	optnames := make(map[string]struct{})
	for _, opt := range cmd.Options {
		err := opt.verify()
		if err != nil {
			return err
		}

		names := append([]string{opt.Name}, opt.Aliases...)
		for _, name := range names {
			if _, found := optnames[name]; found {
				return fmt.Errorf(
					"duplicated option %q", name)
			}

			optnames[name] = struct{}{}
		}
	}

	return nil
}

// verifyParameters verifies command parameters
func (cmd *Command) verifyParameters() error {
	// Verify each parameter individually
	paramnames := make(map[string]struct{})
	for _, param := range cmd.Parameters {
		err := param.verify()
		if err != nil {
			return err
		}

		if _, found := paramnames[param.Name]; found {
			return fmt.Errorf(
				"duplicated parameter %q", param.Name)
		}

		paramnames[param.Name] = struct{}{}
	}

	// Verify parameters disposition
	var repeated, optional *Parameter

	for i := range cmd.Parameters {
		param := &cmd.Parameters[i]

		if param.optional() {
			if repeated != nil {
				return fmt.Errorf(
					"optional parameter %q used after repeated %q",
					param.Name, repeated.Name)
			}

			optional = param
		} else {
			if optional != nil {
				return fmt.Errorf(
					"required parameter %q used after optional %q",
					param.Name, optional.Name)
			}
		}

		if param.repeated() {
			if repeated != nil {
				return fmt.Errorf(
					"repeated parameter used twice (%q and %q)",
					repeated.Name, param.Name)
			}

			repeated = param
		}
	}

	return nil
}

// verifySubCommands verifies command SubCommands
func (cmd *Command) verifySubCommands() error {
	subcmdnames := make(map[string]struct{})
	for _, subcmd := range cmd.SubCommands {
		if _, found := subcmdnames[subcmd.Name]; found {
			return fmt.Errorf(
				"duplicated subcommand %q", subcmd.Name)
		}

		subcmdnames[subcmd.Name] = struct{}{}

		err := subcmd.Verify()
		if err != nil {
			return err
		}
	}

	return nil
}

// Parse parses Command's arguments and returns either
// Invocation or error.
func (cmd *Command) Parse(argv []string) (*Invocation, error) {
	prs := newParser(cmd, argv)

	err := prs.parse()
	if err != nil {
		return nil, err
	}

	return newInvocation(prs), nil
}

// Complete returns array of completion suggestions for
// the Command when used with specified (probably incomplete)
// command line.
func (cmd *Command) Complete(cmdline string) []string {
	return nil
}

// hasOptions tells if Command has Options
func (cmd *Command) hasOptions() bool {
	return len(cmd.Options) != 0
}

// hasParameters tells if Command has Parameters
func (cmd *Command) hasParameters() bool {
	return len(cmd.Parameters) != 0
}

// hasSubCommands tells if Command has SubCommands
func (cmd *Command) hasSubCommands() bool {
	return len(cmd.SubCommands) != 0
}
