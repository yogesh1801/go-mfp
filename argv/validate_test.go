// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Test of validators

package argv

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

// TestValidators implements test of value validators
func TestValidators(t *testing.T) {
	type testData struct {
		name     string
		input    string
		validate func(string) error
		err      string
	}

	tests := []testData{
		// ValidateAny test
		{
			name:     "ValidateAny",
			input:    "hello",
			validate: ValidateAny,
			err:      "",
		},

		// ValidateStrings tests
		{
			name:     "ValidateStrings",
			input:    "one",
			validate: ValidateStrings([]string{"one", "two", "three"}),
			err:      "",
		},

		{
			name:     "ValidateStrings",
			input:    "two",
			validate: ValidateStrings([]string{"one", "two", "three"}),
			err:      "",
		},

		{
			name:     "ValidateStrings",
			input:    "hello",
			validate: ValidateStrings([]string{"one", "two", "three"}),
			err:      `"hello": invalid argument`,
		},

		// ValidateIntRange tests
		{
			name:     "ValidateIntRange",
			input:    "0x123",
			validate: ValidateIntRange(0, -1000, 1000),
			err:      "",
		},

		{
			name:     "ValidateIntRange",
			input:    "+0x123",
			validate: ValidateIntRange(0, -1000, 1000),
			err:      "",
		},

		{
			name:     "ValidateIntRange",
			input:    "-0x123",
			validate: ValidateIntRange(0, -1000, 1000),
			err:      "",
		},

		{
			name:     "ValidateIntRange",
			input:    "0x123",
			validate: ValidateIntRange(10, -1000, 1000),
			err:      `"0x123": invalid integer`,
		},

		{
			name:     "ValidateIntRange",
			input:    "hello",
			validate: ValidateIntRange(0, -1000, 1000),
			err:      `"hello": invalid integer`,
		},

		{
			name:     "ValidateIntRange",
			input:    "10000",
			validate: ValidateIntRange(0, -1000, 1000),
			err:      `"10000": value out of range (-1000...1000)`,
		},

		// ValidateUintRange tests
		{
			name:     "ValidateUintRange",
			input:    "0x123",
			validate: ValidateUintRange(0, 100, 1000),
			err:      "",
		},

		{
			name:     "ValidateUintRange",
			input:    "+0x123",
			validate: ValidateUintRange(0, 100, 1000),
			err:      `"+0x123": invalid integer`,
		},

		{
			name:     "ValidateUintRange",
			input:    "-0x123",
			validate: ValidateUintRange(0, 100, 1000),
			err:      `"-0x123": invalid integer`,
		},

		{
			name:     "ValidateUintRange",
			input:    "0x123",
			validate: ValidateUintRange(10, 100, 1000),
			err:      `"0x123": invalid integer`,
		},

		{
			name:     "ValidateUintRange",
			input:    "hello",
			validate: ValidateUintRange(0, 100, 1000),
			err:      `"hello": invalid integer`,
		},

		{
			name:     "ValidateUintRange",
			input:    "10000",
			validate: ValidateUintRange(0, 100, 1000),
			err:      `"10000": value out of range (100...1000)`,
		},

		// ValidateIntBits tests
		{
			name:     "ValidateIntBits",
			input:    "0",
			validate: ValidateIntBits(0, 1),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    "1",
			validate: ValidateIntBits(0, 1),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    "-1",
			validate: ValidateIntBits(0, 1),
			err:      `"-1": value doesn't fit 1 bits`,
		},

		{
			name:     "ValidateIntBits",
			input:    "2",
			validate: ValidateIntBits(0, 1),
			err:      `"2": value doesn't fit 1 bits`,
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MinInt8),
			validate: ValidateIntBits(0, 8),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MinInt8-1),
			validate: ValidateIntBits(0, 8),
			err:      `"-129": value doesn't fit 8 bits`,
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MaxInt8),
			validate: ValidateIntBits(0, 8),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MaxInt8+1),
			validate: ValidateIntBits(0, 8),
			err:      `"128": value doesn't fit 8 bits`,
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MinInt16),
			validate: ValidateIntBits(0, 16),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MinInt16-1),
			validate: ValidateIntBits(0, 16),
			err:      `"-32769": value doesn't fit 16 bits`,
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MaxInt16),
			validate: ValidateIntBits(0, 16),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MaxInt16+1),
			validate: ValidateIntBits(0, 16),
			err:      `"32768": value doesn't fit 16 bits`,
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MinInt32),
			validate: ValidateIntBits(0, 32),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MinInt32-1),
			validate: ValidateIntBits(0, 32),
			err:      `"-2147483649": value doesn't fit 32 bits`,
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MaxInt32),
			validate: ValidateIntBits(0, 32),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MaxInt32+1),
			validate: ValidateIntBits(0, 32),
			err:      `"2147483648": value doesn't fit 32 bits`,
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MinInt64),
			validate: ValidateIntBits(0, 64),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    fmt.Sprintf("%d", math.MaxInt64),
			validate: ValidateIntBits(0, 64),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    "+5",
			validate: ValidateIntBits(0, 32),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    "-5",
			validate: ValidateIntBits(0, 32),
			err:      "",
		},

		{
			name:     "ValidateIntBits",
			input:    "hello",
			validate: ValidateIntBits(0, 32),
			err:      `"hello": invalid integer`,
		},

		//////////////////////
		// ValidateUintBits tests
		{
			name:     "ValidateUintBits",
			input:    "0",
			validate: ValidateUintBits(0, 1),
			err:      "",
		},

		{
			name:     "ValidateUintBits",
			input:    "1",
			validate: ValidateUintBits(0, 1),
			err:      "",
		},

		{
			name:     "ValidateUintBits",
			input:    "2",
			validate: ValidateUintBits(0, 1),
			err:      `"2": value doesn't fit 1 bits`,
		},

		{
			name:     "ValidateUintBits",
			input:    fmt.Sprintf("%d", math.MaxUint8),
			validate: ValidateUintBits(0, 8),
			err:      "",
		},

		{
			name:     "ValidateUintBits",
			input:    fmt.Sprintf("%d", math.MaxUint8+1),
			validate: ValidateUintBits(0, 8),
			err:      `"256": value doesn't fit 8 bits`,
		},

		{
			name:     "ValidateUintBits",
			input:    fmt.Sprintf("%d", math.MaxUint16),
			validate: ValidateUintBits(0, 16),
			err:      "",
		},

		{
			name:     "ValidateUintBits",
			input:    fmt.Sprintf("%d", math.MaxUint16+1),
			validate: ValidateUintBits(0, 16),
			err:      `"65536": value doesn't fit 16 bits`,
		},

		{
			name:     "ValidateUintBits",
			input:    fmt.Sprintf("%d", math.MaxUint32),
			validate: ValidateUintBits(0, 32),
			err:      "",
		},

		{
			name:     "ValidateUintBits",
			input:    fmt.Sprintf("%d", math.MaxUint32+1),
			validate: ValidateUintBits(0, 32),
			err:      `"4294967296": value doesn't fit 32 bits`,
		},

		{
			name:     "ValidateUintBits",
			input:    "0",
			validate: ValidateUintBits(0, 64),
			err:      "",
		},

		{
			name:     "ValidateUintBits",
			input:    "18446744073709551615",
			validate: ValidateUintBits(0, 64),
			err:      "",
		},

		{
			name:     "ValidateUintBits",
			input:    "+5",
			validate: ValidateUintBits(0, 32),
			err:      `"+5": invalid integer`,
		},

		{
			name:     "ValidateUintBits",
			input:    "-5",
			validate: ValidateUintBits(0, 32),
			err:      `"-5": invalid integer`,
		},

		{
			name:     "ValidateUintBits",
			input:    "hello",
			validate: ValidateUintBits(0, 32),
			err:      `"hello": invalid integer`,
		},
	}

	for _, test := range tests {
		err := test.validate(test.input)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("%s: %q: expected %q, received %q",
				test.name, test.input, test.err, err)
		}
	}
}

// TestValidatorsPanic implements test of expected
// value validators panics
func TestValidatorsPanic(t *testing.T) {
	type testData struct {
		name string
		call func()
		err  string
	}

	tests := []testData{
		{
			name: "ValidateIntBits(0,0)",
			call: func() { ValidateIntBits(0, 0) },
			err:  `ValidateIntBits: bits (0) out of range (1...64)`,
		},

		{
			name: "ValidateIntBits(0,65)",
			call: func() { ValidateIntBits(0, 65) },
			err:  `ValidateIntBits: bits (65) out of range (1...64)`,
		},

		{
			name: "ValidateUintBits(0,0)",
			call: func() { ValidateUintBits(0, 0) },
			err:  `ValidateUintBits: bits (0) out of range (1...64)`,
		},

		{
			name: "ValidateUintBits(0,65)",
			call: func() { ValidateUintBits(0, 65) },
			err:  `ValidateUintBits: bits (65) out of range (1...64)`,
		},
	}

	for _, test := range tests {
		err := func() (err error) {
			errp := &err
			defer func() {
				ex := recover()
				*errp = ex.(error)
			}()

			test.call()

			return *errp
		}()

		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("%s: expected %q, received %q",
				test.name, test.err, err)
		}
	}
}
