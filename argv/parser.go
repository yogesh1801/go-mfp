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
	cmd     *Command       // Command being parsed
	argv    []string       // Arguments being parsed
	options []parserOptVal // Actually parsed options
}

// parserOptVal represents parsed option with value
type parserOptVal struct {
	opt   *Option  // Option description
	name  string   // Actual name being used
	value []string // Option values
}

// newParser creates a new parser
func newParser(cmd *Command, argv []string) *parser {
	return &parser{
		cmd:  cmd,
		argv: argv,
	}
}

// parse parses the argv
func (prs *parser) parse() error {
	// Parse mix of options and other arguments
	for i := 0; i < len(prs.argv); i++ {
		arg := prs.argv[i]
		if arg == "--" {
			break
		}

		switch {
		case prs.isShortOption(arg):
			// Split into name and value and try to find Option
			name, val, novalue := prs.splitOptVal(arg)
			opt := prs.findOption(name)
			if opt == nil {
				err := fmt.Errorf("unknown option: %q", name)
				return err
			}

			// Two simple cases:
			//   - no value
			//   - option requires a value
			//
			// prs.appendOptVal() will handle all errors checking
			// here, so just call it and we are done.
			if novalue || opt.withValue() {
				err := prs.appendOptVal(opt, name, val, novalue)
				if err != nil {
					return err
				}
				break
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
		}
	}

	return nil
}

// isOption tells if argument is option name
func (prs *parser) isOption(arg string) bool {
	return prs.isShortOption(arg) || prs.isLongOption(arg)
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

	if novalue && opt.withValue() {
		err := fmt.Errorf("option requires "+"an argument: %q", name)
		return err
	}

	return nil
}
