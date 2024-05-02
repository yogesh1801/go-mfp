// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Argv parser

package argv

import (
	"fmt"
	"strings"
)

// parser implements command line parsing.
type parser struct {
	cmd        *Command                  // Command being parsed
	argv       []string                  // Arguments being parsed
	nextarg    int                       // Index of the next argument
	options    map[*Option]*parserOptVal // Actually parsed options
	parameters []parserParamVal          // Parameters by number
	byName     map[string][]string       // Options and parameters by name
	subcmd     *Command                  // Sub-command discovered
	subargv    []string                  // Sub-command argv
}

// parserOptVal represents parsed option with value
type parserOptVal struct {
	opt    *Option  // Option description
	name   string   // Actual name being used
	values []string // Option values
}

// parserParamVal represents parsed positional parameter with value
type parserParamVal struct {
	param *Parameter
	value string
}

// newParser creates a new parser.
//
// It panics, if cmd.Verify() returns an error.
func newParser(cmd *Command, argv []string) *parser {
	err := cmd.Verify()
	if err != nil {
		panic(err)
	}

	return &parser{
		cmd:     cmd,
		argv:    argv,
		options: make(map[*Option]*parserOptVal),
		byName:  make(map[string][]string),
	}
}

// parse parses the argv
func (prs *parser) parse() error {
	// Parse arguments, one by one.
	var doneOptions bool
	var paramValues []string

	paramsMin, paramsMax := prs.cmd.paramsInfo()

	for !prs.done() {
		arg := prs.next()

		var err error

		switch {
		case !doneOptions && arg == "--":
			doneOptions = true

		case !doneOptions && prs.isShortOption(arg):
			err = prs.handleShortOption(arg)

		case !doneOptions && prs.isLongOption(arg):
			err = prs.handleLongOption(arg)

		case !doneOptions && prs.cmd.hasSubCommands():
			err = prs.handleSubCommand(arg)

		case len(paramValues) < paramsMax:
			paramValues = append(paramValues, arg)

		default:
			err = fmt.Errorf("unexpected parameter: %q", arg)
		}

		if err != nil {
			return err
		}
	}

	// Toss paramValues
	if len(paramValues) < paramsMin {
		missed := &prs.cmd.Parameters[len(paramValues)]
		return fmt.Errorf("missed parameter: %q", missed.Name)
	}

	if prs.cmd.hasSubCommands() && prs.subcmd == nil {
		return fmt.Errorf("missed sub-command name")
	}

	if prs.cmd.hasParameters() {
		err := prs.handleParameters(paramValues)
		if err != nil {
			return err
		}
	}

	// Export results
	prs.exportResults()

	return nil
}

// handleShortOption handles a short option
func (prs *parser) handleShortOption(arg string) error {
	// Split into name and value and try to find Option
	name, val, novalue := prs.splitOptVal(arg)
	opt := prs.findOption(name)
	if opt == nil {
		err := fmt.Errorf("unknown option: %q", name)
		return err
	}

	// Two simple cases:
	//   - option argument doesn't contain a value (i.e., -c, not -cXXX)
	//   - option requires a value, so argument cannot be treated as
	//     a multi-options argument
	//
	// These cases are handled the same way: we attempt to fetch
	// the next argument as option value, if value is required, and
	// let prs.appendOptVal() to do the rest.
	if novalue || opt.withValue() {
		if novalue && opt.withValue() {
			val, novalue = prs.nextValue()
		}

		return prs.appendOptVal(opt, name, val, novalue)
	}

	// Short options without value can be combined:
	//
	//   -cru equals to -c -r -u
	//
	// If we are here, we have a fist option without the value
	// and non-empty value.
	//
	// So try to consider value as a sequence of short options
	err := prs.appendOptVal(opt, name, "", true)
	if err != nil {
		return err
	}

	for _, c := range val {
		name2 := "-" + string(c)

		opt2 := prs.findOption(name2)
		if opt2 == nil {
			err := fmt.Errorf(
				"unknown option: %q",
				name2)
			return err
		}

		err := prs.appendOptVal(opt2, name2, "", true)
		if err != nil {
			return err
		}
	}

	return nil
}

// handleLongOption handles a long option
func (prs *parser) handleLongOption(arg string) error {
	name, val, novalue := prs.splitOptVal(arg)

	opt := prs.findOption(name)
	if opt == nil {
		err := fmt.Errorf("unknown option: %q", name)
		return err
	}

	if novalue && opt.withValue() {
		val, novalue = prs.nextValue()
	}

	err := prs.appendOptVal(opt, name, val, novalue)
	if err != nil {
		return err
	}

	return nil
}

// handleSubCommand handles a sub-command
func (prs *parser) handleSubCommand(arg string) error {
	subcommands := prs.findSubCommand(arg)

	switch {
	case len(subcommands) == 0:
		return fmt.Errorf("unknown sub-command: %q", arg)
	case len(subcommands) > 1:
		return fmt.Errorf("ambiguous sub-command: %q", arg)
	}

	prs.subcmd = subcommands[0]
	prs.subargv = prs.argv[prs.nextarg:]
	return nil
}

// exportResults rearranges parsing results into the usable form
func (prs *parser) exportResults() {
	// Save options values
	for _, optval := range prs.options {
		opt := optval.opt
		prs.byName[opt.Name] = optval.values

		for _, name := range opt.Aliases {
			prs.byName[name] = optval.values
		}
	}

	// Save parameters values
	//
	// Note, repeated parameters may have multiple values associated
	// with the same parameter
	for _, paramval := range prs.parameters {
		name := paramval.param.Name
		values := prs.byName[name]
		values = append(values, paramval.value)
		prs.byName[name] = values
	}
}

// handleParameters handles positional parameters
func (prs *parser) handleParameters(paramValues []string) error {
	// Build slice of parameters' descriptors
	paramDescs := make([]*Parameter, len(paramValues))
	rept := -1

	for i := 0; i < len(paramValues); i++ {
		paramDescs[i] = &prs.cmd.Parameters[i]
		if paramDescs[i].repeated() {
			rept = i
			break
		}
	}

	if rept >= 0 {
		i := len(paramDescs) - 1
		j := len(prs.cmd.Parameters) - 1

		for !prs.cmd.Parameters[j].repeated() {
			paramDescs[i] = &prs.cmd.Parameters[j]
			i--
			j--
		}

		for i := rept + 1; i < len(paramDescs); i++ {
			if paramDescs[i] == nil {
				paramDescs[i] = paramDescs[rept]
			}
		}
	}

	// Validate parameters one by one
	for i := range paramValues {
		val := paramValues[i]
		desc := paramDescs[i]

		if desc.Validate != nil {
			err := desc.Validate(val)
			if err != nil {
				return fmt.Errorf("%w: %q", err, desc.Name)
			}
		}
	}

	// Save parameters
	prs.parameters = make([]parserParamVal, len(paramValues))
	for i := range paramValues {
		prs.parameters[i].param = paramDescs[i]
		prs.parameters[i].value = paramValues[i]
	}

	return nil
}

// done returns true if all arguments are consumed
func (prs *parser) done() bool {
	return prs.nextarg == len(prs.argv) || prs.subcmd != nil
}

// next returns the next argument.
func (prs *parser) next() string {
	if prs.nextarg < len(prs.argv) {
		arg := prs.argv[prs.nextarg]
		prs.nextarg++
		return arg
	}

	return ""
}

// nextValue returns the next argument, of one exist.
func (prs *parser) nextValue() (val string, novalue bool) {
	if !prs.done() {
		return prs.next(), false
	}

	return "", true
}

// isShortOption tells if argument is a short option
func (prs *parser) isShortOption(arg string) bool {
	return len(arg) >= 2 && arg[0] == '-' && arg[1] != '-'
}

// isShortOption tells if argument is a long option
func (prs *parser) isLongOption(arg string) bool {
	return len(arg) >= 3 && arg[0] == '-' && arg[1] == '-'
}

// splitOptVal splits option argument into name and value in a case
// when they are placed into the single argument:
//
//  -cVAL     - short option case
//  -long=val - long option case
func (prs *parser) splitOptVal(arg string) (name, val string, novalue bool) {
	switch {
	case prs.isShortOption(arg):
		name = arg[:2]
		val = arg[2:]
		novalue = val == ""

	case prs.isLongOption(arg):
		// For --name=value, pick out the name
		idx := strings.IndexByte(arg, '=')
		if idx >= 0 {
			name = arg[:idx]
			val = arg[idx+1:]
			novalue = false
		} else {
			name = arg
			novalue = true
		}
	}

	return
}

// findOption finds Command's Option by name.
func (prs *parser) findOption(name string) *Option {
	for i := range prs.cmd.Options {
		opt := &prs.cmd.Options[i]
		if name == opt.Name {
			return opt
		}

		for i := range opt.Aliases {
			if name == opt.Aliases[i] {
				return opt
			}
		}
	}

	return nil
}

// findSubCommand finds Command's SubCommand by name.
//
// The name may be abbreviated, so in a case of inexact
// match it may return more that one possible candidates.
//
// If no matches found it will return nil and in a case
// of exact match it will return just a single command,
// even if more inexact matches exist
//
// This is up to the caller how to handle this ambiguity.
func (prs *parser) findSubCommand(name string) []*Command {
	var inexact []*Command
	for i := range prs.cmd.SubCommands {
		subcmd := &prs.cmd.SubCommands[i]

		if name == subcmd.Name {
			return []*Command{subcmd}
		}

		if strings.HasPrefix(subcmd.Name, name) {
			inexact = append(inexact, subcmd)
		}
	}

	return inexact
}

// appendOptVal validates option value and appends
// it to the prs.options
func (prs *parser) appendOptVal(opt *Option, name, value string,
	novalue bool) error {

	// Validate things
	if novalue && opt.withValue() {
		err := fmt.Errorf("option requires operand: %q", name)
		return err
	}

	if !novalue {
		err := opt.Validate(value)
		if err != nil {
			return fmt.Errorf("%w: %q", err, name)
		}
	}

	// Save the option
	optval := prs.options[opt]
	if optval == nil {
		optval = &parserOptVal{
			opt:  opt,
			name: name,
		}

		prs.options[opt] = optval
	}

	optval.values = append(optval.values, value)

	return nil
}
