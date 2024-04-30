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
	cmd     *Command                  // Command being parsed
	argv    []string                  // Arguments being parsed
	nextarg int                       // Index of the next argument
	options map[*Option]*parserOptVal // Actually parsed options
}

// parserOptVal represents parsed option with value
type parserOptVal struct {
	opt    *Option  // Option description
	name   string   // Actual name being used
	values []string // Option values
}

// newParser creates a new parser
func newParser(cmd *Command, argv []string) *parser {
	return &parser{
		cmd:     cmd,
		argv:    argv,
		options: make(map[*Option]*parserOptVal),
	}
}

// parse parses the argv
func (prs *parser) parse() error {
	// Parse mix of options and other arguments

	for !prs.done() {
		arg := prs.next()
		if arg == "--" {
			break
		}

		var err error

		switch {
		case prs.isShortOption(arg):
			err = prs.handleShortOption(arg)

		case prs.isLongOption(arg):
			err = prs.handleLongOption(arg)
		}

		if err != nil {
			return err
		}
	}

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

// done returns true if all arguments are consumed
func (prs *parser) done() bool {
	return prs.nextarg == len(prs.argv)
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

	case prs.isLongOption(name):
		// For --name=value, pick out the name
		idx := strings.IndexByte(name, '=')
		if idx >= 0 {
			name = arg[:idx]
			val = arg[idx:]
			novalue = false
		} else {
			name = arg
			novalue = true
		}
	}

	return
}

// findOption find Command's Option by name.
func (prs *parser) findOption(name string) *Option {
	// If option name and value mixed in a same argument,
	// pick out the name:
	//
	//   -cVAL     - short option case
	//   -long=val - long option case
	switch {
	case prs.isShortOption(name):
		// In a short option case, name is a dash plus
		// single character
		name = name[:2]

	case prs.isLongOption(name):
		// For --name=value, pick out the name
		idx := strings.IndexByte(name, '=')
		if idx >= 0 {
			name = name[:idx]
		}
	}

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
