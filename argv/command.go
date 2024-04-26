// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command definition structures.

package argv

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
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

	// Usage string, a single-line description.
	Usage string

	// Help text, a long description.
	Help string

	// Options, if any.
	Options []Option

	// Positional parameters, if any.
	Parameters []Parameter

	// Sub-commands, if any.
	SubCommands []Command
}

// Option defines an option.
//
// Option MUST have a name and MAY have one or more aliases.
//
// Name and Aliases of all Options MUST be unique within a scope
// of Command that defines them (sub-commands have their own scopes).
//
// Option MAY have a value. Presence of name is indicated by the
// non-nil ValDef field.
//
// Option name must start with single or double dash (- or --), followed
// by alphanumeric character, optionally followed by a sequence of
// characters, that include only alphanumeric characters and dashes.
//
// Option names may be either short or long. Name that consist of a
// single dash, followed by a single alphanumeric character considered
// short:
//
//   -c           - the short name
//   --long-name  - the long name
//
// And if used with value:
//
//   -c XXX                             - the short name with value
//   --long-name XXX or --long-name=XXX - the long name with value
type Option struct {
	// Name is the option name.
	Name string

	// Aliases are the option aliases, if any.
	Aliases []string

	// Usage string, a single-line description.
	Usage string

	// Requires, if not nil, contains names of other Options,
	// that MUST be used together with this option.
	Requires []string

	// Conflicts, if not nit, contains names of other Options
	// that MUST NOT be used together with this option.
	Conflicts []string

	// Validate callback called to validate parameter.
	//
	// Use nil to indicate that this option has no value.
	Validate func(string) error

	// Complete callback called for auto-completion.
	// It receives the prefix, already typed by user
	// (which may be empty) and must return completion
	// suggestions without that prefix.
	Complete func(string) []string
}

// Parameter defines a positional parameter.
//
// Parameter MUST have a name, and names of all Parameters
// MUST be unique within a scope  of Command that defines them
// (sub-commands have their own scopes).
//
// Parameter names used to generate help messages and to
// access parameters by name, hence requirement of uniqueness.
//
// If name of the Parameter ends with ellipsis (...), this is
// repeated parameter:
//
//   copy source... destination
//
// If name of the Parameter is taken into square braces ([name]),
// this is optional parameter:
//
//   print document [format]
//
// Optional parameter may be omitted.
//
// Ellipses a square braces syntax may be combined:
//
//   list [file...]
//
// Non-optional repeated parameter will consume 1 or more
// parameter values. Optional repeated parameter will consume
// 0 or more parameter values.
//
// All parameters after non-repeated optional parameters
// must be optional:
//
//   cmd param1 param2 [param3] [param4]      - OK
//   cmd param1 param2 [param3] [param4...]   - OK
//   cmd param1 param2 [param3] param4        - error
//
// In the last case, if we have only 3 parameter values,
// we can't tell unambiguously, if it param1 param2 param3
// or param1 param2 param4.
//
// After repeated (optional or not) parameter, more
// non-optional may follow:
//
//   cmd param1 param2 param3... param4     - OK
//   cmd param1 param2 [param3...] param4   - OK
//
// At this case, if we have N parameter values, we first
// assign values to the non-optional ones, the remaining
// values assigned to the repeated parameter.
//
// But optional parameter after repeated is not allowed:
//
//   cmd param1 param2 param3... [param4]   - error
//   cmd param1 param2 [param3...] [param4] - error
//
// Only one parameter may be repeated:
//   cmd param1 param2 param3... param4     - OK
//   cmd param1 param2 param3... param4..   - error
//
// Without this rule, this is hard to say unambiguously,
// how to distribute values between param3... and param4...
type Parameter struct {
	// Name is the parameter name.
	Name string

	// Usage string, a single-line description.
	Usage string

	// Requires, if not nil, contains names of Options,
	// that MUST be used together with this option.
	Requires []string

	// Conflicts, if not nit, contains names of Options
	// that MUST NOT be used together with this option.
	Conflicts []string

	// Validate callback called to validate parameter
	Validate func(string) error

	// Complete callback called for auto-completion.
	// It receives the prefix, already typed by user
	// (which may be empty) and must return completion
	// suggestions without that prefix.
	Complete func(string) []string
}

// Action defines action to be taken when Command is
// applied to the command line.
type Action struct {
	options map[string][]string
}

// ----- Command methods -----

// Verify checks correctness of Command definition. It fails if any
// error is found and returns description of the first caught error
func (cmd *Command) Verify() error {
	// Command must have a name
	if cmd.Name == "" {
		return errors.New("missed command name")
	}

	// Verify options. Also check that option names doesn't duplicate
	optnames := make(map[string]struct{})
	for _, opt := range cmd.Options {
		err := opt.verify()
		if err != nil {
			return fmt.Errorf("%s: %s", cmd.Name, err)
		}

		names := append([]string{opt.Name}, opt.Aliases...)
		for _, name := range names {
			if _, found := optnames[name]; found {
				return fmt.Errorf("%s: duplicated option %q",
					cmd.Name, name)
			}

			optnames[name] = struct{}{}
		}
	}

	// Verify parameters
	if cmd.Parameters != nil && cmd.SubCommands != nil {
		return fmt.Errorf(
			"%s: Parameters and SubCommands are mutually exclusive",
			cmd.Name)
	}

	paramnames := make(map[string]struct{})
	for _, param := range cmd.Parameters {
		err := param.verify()
		if err != nil {
			return fmt.Errorf("%s: %s", cmd.Name, err)
		}

		if _, found := paramnames[param.Name]; found {
			return fmt.Errorf("%s: duplicated option %q",
				cmd.Name, param.Name)
		}

		paramnames[param.Name] = struct{}{}
	}

	return nil
}

// Apply applies Command to argument. On success
// it returns Action which defines further procession.
func (cmd *Command) Apply(argv []string) (*Action, error) {
	return nil, nil
}

// Complete returns array of completion suggestions for
// the Command when used with specified (probably incomplete)
// command line.
func (cmd *Command) Complete(cmdline string) []string {
	return nil
}

// ----- Option methods -----

// verify checks correctness of Option definition. It fails if any
// error is found and returns description of the first caught error
func (opt *Option) verify() error {
	// Option must have a name
	if opt.Name == "" {
		return errors.New("option must have a name")
	}

	// Verify name syntax
	names := append([]string{opt.Name}, opt.Aliases...)
	for _, name := range names {
		var check string

		switch {
		case strings.HasPrefix(name, "--"):
			check = name[2:]
		case strings.HasPrefix(name, "-"):
			check = name[1:]

		default:
			return fmt.Errorf("option must start with dash (-): %q",
				name)
		}

		if check == "" {
			return fmt.Errorf("empty option name: %q", name)
		}

		if c := nameCheck(check); c >= 0 {
			return fmt.Errorf("invalid char '%c' in option: %q",
				c, name)
		}
	}

	return nil
}

// ----- Parameter methods -----

// verify checks correctness of Parameter definition. It fails if any
// error is found and returns description of the first caught error
func (param *Parameter) verify() error {
	// Parameter must have a name
	if param.Name == "" {
		return errors.New("parameter must have a name")
	}

	// Verify name syntax
	check := param.Name
	if strings.HasPrefix(check, "[") {
		// If name starts with "[", this is optional parameter,
		// and it must end with "]"
		if strings.HasSuffix(check, "]") {
			check = check[1 : len(check)-1]
		} else {
			return errors.New("missed closing ']' character")
		}
	}

	if strings.HasSuffix(check, "...") {
		// Strip trailing "...", if any
		check = check[0 : len(check)-3]
	}

	// Check remaining name
	if check == "" {
		return fmt.Errorf("parameter name is empty: %q", param.Name)
	}

	if c := nameCheck(check); c >= 0 {
		return fmt.Errorf("invalid char '%c' in parameter: %q",
			c, param.Name)
	}

	return nil
}

// optional returns true if parameter is optional
func (param *Parameter) optional() bool {
	return strings.HasPrefix(param.Name, "[")
}

// repeated returns true if parameter is repeated
func (param *Parameter) repeated() bool {
	return strings.HasSuffix(param.Name, "...") ||
		strings.HasSuffix(param.Name, "...]")
}

// ----- Action methods -----

// Getopt returns value of option or parameter as a single string.
//
// For multi-value options and repeated parameters values
// are concatenated into the single string using CSV encoding.
func (act *Action) Getopt(name string) (val string, found bool) {
	return "", false
}

// GetoptSlice returns value of option or parameter as a slice of string.
// If option is not found, it returns nil
func (act *Action) GetoptSlice(name string) (val []string) {
	return nil
}

// ----- Miscellaneous functions -----

// nameCheck function verifies syntax of Options and
// Parameters names. Valid name starts with letter or
// digit and then consist of letters, digits and dash
// characters.
//
// It returns the first invalid character, if one is
// encountered, or -1 otherwise.
func nameCheck(name string) rune {
	for i, c := range name {
		switch {
		// Letters and digits always allowed
		case unicode.IsLetter(c) || unicode.IsDigit(c):

		// Dash allowed expect the very first character
		case i > 0 && c == '-':

		// Other characters not allowed
		default:
			return c
		}
	}

	return -1
}
