// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Invocation -- contains parsed Command's arguments.

package argv

// Invocation represents a particular Command invocation.
//
// It contains a whole Command execution context, like parsed
// options and arguments.
type Invocation struct {
	// parent is the upper-level Invocation for sub-command
	// Invocation, nil for the root Invocation.
	parent *Invocation

	// cmd contains back reference to invoked Command.
	cmd *Command

	// argv contains original argv being used when Command
	// was invoked.
	argv []string

	// byName contains options and parameters values,
	// indexed by name.
	//
	// Because options and parameters names are incompatible,
	// we can use a single structure for both.
	byName map[string][]string

	// parameters contains parameters values, indexed by numbers.
	parameters []string

	// subcmd is the Command's SubCommand and subargv is its arguments.
	subcmd  *Command
	subargv []string

	// immediate is the first Option's Immediate callback, if any
	immediate func(*Invocation) error
}

// Parent returns Invocation's parent, which is the upper-level
// Invocation in a case of sun-command execution, nil otherwise.
func (inv *Invocation) Parent() *Invocation {
	return inv.parent
}

// IsImmediate returns true, if Invocation contains some active Option
// with non-nil Immediate callback (see Option description for more
// information).
func (inv *Invocation) IsImmediate() bool {
	return inv.immediate != nil
}

// Cmd returns a reference to Command, invoked by this Invocation
func (inv *Invocation) Cmd() *Command {
	return inv.cmd
}

// Argv returns the original argv being used when Command was invoked.
func (inv *Invocation) Argv() []string {
	return inv.argv
}

// Get returns the first value of option or parameter by its name.
//
// The value of flag options (options that don't expect explicit
// value) considered to be an empty string.
func (inv *Invocation) Get(name string) (val string, found bool) {
	vals, found := inv.byName[name]
	if found && len(vals) > 0 {
		val = vals[0]
	}

	return
}

// Values returns a slice of all values of option or parameter by
// its name.
//
// The value of flag options (options that don't expect explicit
// value) considered to be an empty string.
//
// For repeated flag options, the returned slice will contain one
// empty string per each occurrence.
func (inv *Invocation) Values(name string) []string {
	return inv.byName[name]
}

// ParamCount returns count of positional parameters.
func (inv *Invocation) ParamCount() int {
	return len(inv.parameters)
}

// ParamGet returns value of the n-th positional parameter.
// If n is our of range, it returns empty string ("").
//
// Parameters are numbered by their actual position within command's
// arguments, not by their position withing Parameters slice in Command
// description. This difference becomes important when we come to
// repeated parameters. Repeated parameter will take only one slot
// in the Parameters slice, but may be repeated (and take many positions)
// in the Command's argument.
func (inv *Invocation) ParamGet(n int) (param string) {
	if 0 <= n && n < len(inv.parameters) {
		param = inv.parameters[n]
	}

	return
}

// SubCommand returns Command's SubCommand and its arguments. If Command
// doesn't have SubCommands, this function returns (nil, nil).
func (inv *Invocation) SubCommand() (*Command, []string) {
	return inv.subcmd, inv.subargv
}
