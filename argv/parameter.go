// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Parameter -- defines a command's positional parameter.

package argv

import (
	"errors"
	"fmt"
	"strings"
)

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
// Not all combinations of required, optional and repeated
// parameters are valid.
//
// Valid combinations:
//
//   cmd param1 param2 [param3] [param4]      - OK
//   cmd param1 param2 [param3] [param4...]   - OK
//   cmd param1 param2 param3... param4       - OK
//   cmd param1 param2... param3 param4       - OK
//
// Inlaid combinations:
//
//   Required parameter       cmd param1 [param2] param3
//   can't follow optional
//   parameter
//
//   Optional parameter       cmd param1 param2... [param3]
//   can't follow repeated    cmd param1 [param2...] [param3]
//   parameter
//
//   Only one repeated        cmd param1 param2... param3...
//   parameter is allowed
//
// These rules exist so simplify unambiguous matching of actual
// parameters against formal (declared) ones.
type Parameter struct {
	// Name is the parameter name.
	Name string

	// Help string, a single-line description.
	Help string

	// Validate callback called to validate parameter
	Validate func(string) error

	// Complete callback called for auto-completion.
	// It receives the prefix, already typed by user
	// (which may be empty) and must return completion
	// suggestions without that prefix.
	Complete func(string) []string
}

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
			err := fmt.Errorf("missed closing ']' character in parameter %q",
				param.Name)
			return err
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

// optional returns true if parameter is required
func (param *Parameter) required() bool {
	return !param.optional()
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

// complete is the convenience wrapper around Parameter.Complete
// callback. It call callback only if one is not nil.
func (param *Parameter) complete(prefix string) (compl []string) {
	if param.Complete != nil {
		compl = param.Complete(prefix)
	}
	return
}
