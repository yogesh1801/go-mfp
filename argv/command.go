// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command definition structures.

package argv

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

	// Sub-commands, if any.
	SubCommands []Command

	// Positional parameters, if any.
	Parameters []Parameter
}

// Option defines an option.
//
// Option MUST have a name and MAY have one or more aliases.
//
// Name and Aliases of all Options MUST be unique within a scope
// of Command that defines them (sub-commands have their own scopes).
//
// Name and Aliases may use either short or long syntax. Short Name
// consist of single dash (-) character, long Name starts with double
// dash (--):
//
//   -c           - the short name
//   --long-name  - the long name
//
// And if used with value:
//
//   -c XXX                             - the short name with value
//   --long-name XXX or --long-name=XXX - the long name with value
//
// Option MAY have a value. Presence of name is indicated by the
// non-nil ValDef field.
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
//   cmd param1 param2 [param3] [param4]     - OK
//   cmd param1 param2 [param3] [param4...]  - OK
//   cmd param1 param2 [param3] param4       - error
//
// In the last case, if we have only 3 parameter values,
// we can't tell unambiguously, if it param1 param2 param3
// or param1 param2 param4.
//
// After repeated (optional or not) parameter, more
// non-optional may follow:
//
//   cmd param1 param2 param3... param4    - OK
//   cmd param1 param2 [param3...] param4  - OK
//
// At this case, if we have N parameter values, we first
// assign values to the non-optional ones, the remaining
// values assigned to the repeated parameter.
//
// Only one parameter may be repeated:
//   cmd param1 param2 param3... param4    - OK
//   cmd param1 param2 param3... param4..  - error
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

// Validate verifies Command definition. It fails if any
// error is found and returns description of the first caught
// error
func (cmd *Command) Validate() error {
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
