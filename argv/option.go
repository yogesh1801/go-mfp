// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Option -- defines a Command's option.

package argv

import (
	"errors"
	"fmt"
	"strings"
)

// Option defines an option.
//
// Option MUST have a name and MAY have one or more aliases.
//
// Name and Aliases of all Options MUST be unique within a scope
// of Command that defines them (sub-commands have their own scopes).
//
// Option MAY have a value. Presence of name is indicated by the
// non-nil Validate field.
//
// Option may have either short or long syntax:
//
//	-c                                 - short option without value
//	--long-name                        - long option without value
//	-c XXX or -cXXX                    - short name with value
//	--long-name XXX or --long-name=XXX - long name with value
//
// Short options without value can be combined:
//
//	-cru equals to -c -r -u
//
// Short name starts with a single dash (-) character followed
// by a single alphanumeric character.
//
// Long name starts with a double dash (--) characters followed
// by an alphanumeric character, optionally followed by a sequence
// of characters, that include only alphanumeric characters and dashes.
//
//	-x            - valid
//	-abc          - invalid; short option with a long name
//	--x           - valid (the long option, though name is 1-character)
//	--long        - valid
//	--long-option - valid
//	---long       - invalid; character after -- must be alphanumerical
//
// This naming convention is consistent with [GNU extensions] to the POSIX
// recommendations for command-line options:
//
// [GNU extensions]: https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html
type Option struct {
	// Name is the option name.
	Name string

	// Aliases are the option aliases, if any.
	Aliases []string

	// Help string, a single-line description.
	Help string

	// Conflicts, if not nit, contains names of other Options
	// that MUST NOT be used together with this option.
	Conflicts []string

	// Requires, if not nil, contains names of other Options,
	// that MUST be used together with this option.
	Requires []string

	// Validate callback called to validate parameter.
	//
	// Use nil to indicate that this option has no value.
	Validate func(string) error

	// Complete is the callback called for auto-completion.
	//
	// See description of the Completer type for details.
	Complete Completer

	// Immediate, if not nil and option was encountered in
	// the Command's argv, overrides the Command's handler
	// and check for missed parameters and options is suppressed
	// if Immediate is used.
	//
	// It is intended to implement options, like --help, that
	// works as "immediate sub-command override" (even for commands
	// without sub-commands).
	//
	// If there are multiple "immediate" options in the Command's
	// Invocation, the first one always wins.
	Immediate func(*Invocation) error
}

// verify checks correctness of Option definition. It fails if any
// error is found and returns description of the first caught error
func (opt *Option) verify() error {
	// Option must have a name
	if opt.Name == "" {
		return errors.New("option must have a name")
	}

	// Verify syntax of option Name and Aliases
	names := append([]string{opt.Name}, opt.Aliases...)
	for _, name := range names {
		err := opt.verifyNameSyntax(name)
		if err != nil {
			return err
		}
	}

	// Verify name syntax of Conflicts and Requires
	for _, name := range opt.Conflicts {
		err := opt.verifyNameSyntax(name)
		if err != nil {
			return fmt.Errorf("Conflicts: %w", err)
		}
	}

	for _, name := range opt.Requires {
		err := opt.verifyNameSyntax(name)
		if err != nil {
			return fmt.Errorf("Requires: %w", err)
		}
	}

	return nil
}

// verifyNameSyntax verifies option name syntax
func (opt *Option) verifyNameSyntax(name string) error {
	var check string
	var short bool

	switch {
	case strings.HasPrefix(name, "--"):
		check = name[2:]
	case strings.HasPrefix(name, "-"):
		short = true
		check = name[1:]

	default:
		return fmt.Errorf(
			"option must start with dash (-): %q", name)
	}

	if check == "" {
		return fmt.Errorf("empty option name: %q", name)
	}

	if c := nameCheck(check); c >= 0 {
		return fmt.Errorf(
			"invalid char '%c' in option: %q", c, name)
	}

	if short && len(check) > 1 {
		return fmt.Errorf(
			"short option with long name: %q", name)
	}

	return nil
}

// withValue tells if Option has a value
func (opt *Option) withValue() bool {
	return opt.Validate != nil
}

// complete is the convenience wrapper around Option.Complete
// callback. It call callback only if one is not nil.
func (opt *Option) complete(prefix string) ([]string, CompleterFlags) {
	var compl []string
	var flags CompleterFlags

	if opt.Complete != nil {
		compl, flags = opt.Complete(prefix)
	}

	return compl, flags
}
