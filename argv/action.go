// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Action -- contains parsed Command's arguments.

package argv

// Action defines action to be taken when Command is
// applied to the command line.
type Action struct {
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
}

// newAction creates and feels a new Action.
func newAction(prs *parser) *Action {
	act := &Action{
		byName:  prs.byName,
		subcmd:  prs.subcmd,
		subargv: prs.subargv,
	}

	act.parameters = make([]string, len(prs.parameters))
	for i := range prs.parameters {
		act.parameters[i] = prs.parameters[i].value
	}

	return act
}

// Get returns the first value of option or parameter by its name.
//
// The value of flag options (options that don't expect explicit
// value) considered to be an empty string.
func (act *Action) Get(name string) (val string, found bool) {
	vals, found := act.byName[name]
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
func (act *Action) Values(name string) []string {
	return act.byName[name]
}

// ParamCount returns count of positional parameters.
func (act *Action) ParamCount() int {
	return len(act.parameters)
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
func (act *Action) ParamGet(n int) string {
	if 0 <= n && n < len(act.parameters) {
		return act.parameters[n]
	}

	return ""
}

// SubCommand returns Command's SubCommand and its arguments. If Command
// doesn't have SubCommands, this function returns (nil, nil).
func (act *Action) SubCommand() (*Command, []string) {
	return act.subcmd, act.subargv
}
