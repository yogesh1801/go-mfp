// MFP - Miulti-Function Printers and scanners toolkit
// UUID mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UUID tests

package uuid

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

// TestParse tests UUID parser
func TestParse(t *testing.T) {
	type testData struct {
		s    string // Input string
		uuid UUID   // Expected output
		err  string // Expected error (in a string form)
	}

	tests := []testData{
		// Lower case
		{
			s: "01234567-89ab-cdef-0123-456789abcdef",
			uuid: UUID{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			},
		},

		// Upper case
		{
			s: "01234567-89AB-CDEF-0123-456789ABCDEF",
			uuid: UUID{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			},
		},

		// Some random UUID
		{
			s: "c69fe12a-1491-46c3-9083-48035aa4d749",
			uuid: UUID{
				0xc6, 0x9f, 0xe1, 0x2a, 0x14, 0x91, 0x46, 0xc3,
				0x90, 0x83, 0x48, 0x03, 0x5a, 0xa4, 0xd7, 0x49,
			},
		},

		// URN style
		{
			s: "urn:uuid:c69fe12a-1491-46c3-9083-48035aa4d749",
			uuid: UUID{
				0xc6, 0x9f, 0xe1, 0x2a, 0x14, 0x91, 0x46, 0xc3,
				0x90, 0x83, 0x48, 0x03, 0x5a, 0xa4, 0xd7, 0x49,
			},
		},

		// uuid: prefix
		{
			s: "uuid:c69fe12a-1491-46c3-9083-48035aa4d749",
			uuid: UUID{
				0xc6, 0x9f, 0xe1, 0x2a, 0x14, 0x91, 0x46, 0xc3,
				0x90, 0x83, 0x48, 0x03, 0x5a, 0xa4, 0xd7, 0x49,
			},
		},

		// No dashes
		{
			s: "c69fe12a149146c3908348035aa4d749",
			uuid: UUID{
				0xc6, 0x9f, 0xe1, 0x2a, 0x14, 0x91, 0x46, 0xc3,
				0x90, 0x83, 0x48, 0x03, 0x5a, 0xa4, 0xd7, 0x49,
			},
		},

		// Microsoft style
		{
			s: "{c69fe12a149146c3908348035aa4d749}",
			uuid: UUID{
				0xc6, 0x9f, 0xe1, 0x2a, 0x14, 0x91, 0x46, 0xc3,
				0x90, 0x83, 0x48, 0x03, 0x5a, 0xa4, 0xd7, 0x49,
			},
		},

		// Too short
		{
			s:   "c69fe12a149146c3908348035aa4d74",
			err: "UUID is too short (31 digits)",
		},

		// Too long
		{
			s:   "c69fe12a149146c3908348035aa4d7490",
			err: "UUID is too long (33 digits)",
		},

		// Invalid character
		{
			s:   "?c69fe12a149146c3908348035aa4d7490",
			err: `UUID contains invalid character: "?"`,
		},
	}

	for _, test := range tests {
		uuid, err := Parse(test.s)

		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("%s: error mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.s, test.err, err.Error())
		}

		if uuid != test.uuid {
			t.Errorf("%s: value mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.s, test.uuid, uuid)
		}
	}
}

// TestFormat tests UUID formatters
func TestFormat(t *testing.T) {
	type testData struct {
		uuid   UUID   // The UUID
		format string // Formatter name
		out    string // Expected output
	}

	tests := []testData{
		{
			uuid:   Must(Parse("c69fe12a149146c3908348035aa4d749")),
			format: "string",
			out:    "c69fe12a-1491-46c3-9083-48035aa4d749",
		},

		{
			uuid:   Must(Parse("c69fe12a149146c3908348035aa4d749")),
			format: "URN",
			out:    "urn:uuid:c69fe12a-1491-46c3-9083-48035aa4d749",
		},

		{
			uuid:   Must(Parse("c69fe12a149146c3908348035aa4d749")),
			format: "Microsoft",
			out:    "{c69fe12a-1491-46c3-9083-48035aa4d749}",
		},
	}

	for _, test := range tests {
		var out string
		switch strings.ToLower(test.format) {
		case "string":
			out = test.uuid.String()
		case "urn":
			out = test.uuid.URN()
		case "microsoft":
			out = test.uuid.Microsoft()
		default:
			panic(fmt.Sprintf(
				"unhandled formatter %q", test.format))
		}

		if out != test.out {
			t.Errorf("%s: formatting mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.format, test.out, out)
		}
	}
}

// TestRandomFrom tests RandomFrom function
func TestRandomFrom(t *testing.T) {
	type testData struct {
		in   []byte // Input data
		uuid UUID   // Expected output
		err  string // Expected error string
	}

	tests := []testData{
		// Normal input
		{
			in: []byte{
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
				0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
			},
			uuid: Must(Parse("000102030405460788090a0b0c0d0e0f")),
		},

		// All zeroes input
		{
			in: []byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			uuid: Must(Parse("00000000000040008000000000000000")),
		},

		// All ones input
		{
			in: []byte{
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			},
			uuid: Must(Parse("ffffffffffff4fffbfffffffffffffff")),
		},

		// Truncated input
		{
			in: []byte{
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
			},
			err: "unexpected EOF",
		},
	}

	for _, test := range tests {
		r := bytes.NewReader(test.in)
		uuid, err := RandomFrom(r)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("%x: error mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.in, test.err, err.Error())
		}

		if uuid != test.uuid {
			t.Errorf("%x: value mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.in, test.uuid, uuid)
		}

	}
}

// TestVerVar tests UUID.Version and UUID.Variant
// At the same time it tests UUID, generated by Random function
func TestVerVar(t *testing.T) {
	type testData struct {
		uuid UUID    // Input UUID
		ver  Version // Expected Version
		vrnt Variant // Expected Variant
	}

	tests := []testData{
		// Version 0, Variant 0b_100_00000,
		{
			uuid: Must(Parse("000000000000-00-00-80-00000000000000")),
			ver:  0,
			vrnt: VariantRFC4122,
		},

		// Version 1, Variant 0b_100_00000,
		{
			uuid: Must(Parse("000000000000-10-00-80-00000000000000")),
			ver:  1,
			vrnt: VariantRFC4122,
		},

		// Version 2, Variant 0b_100_00000,
		{
			uuid: Must(Parse("000000000000-20-00-80-00000000000000")),
			ver:  2,
			vrnt: VariantRFC4122,
		},

		// Version 4, Variant 0b_100_00000,
		{
			uuid: Must(Parse("000000000000-40-00-80-00000000000000")),
			ver:  4,
			vrnt: VariantRFC4122,
		},

		// Version 0, Variant 0b_000_00000,
		{
			uuid: Must(Parse("000000000000-80-00-00-00000000000000")),
			ver:  8,
			vrnt: VariantNCS,
		},

		// Version 0, Variant 0b_001_00000,
		{
			uuid: Must(Parse("000000000000-00-00-20-00000000000000")),
			ver:  0,
			vrnt: VariantNCS,
		},

		// Version 0, Variant 0b_010_00000,
		{
			uuid: Must(Parse("000000000000-00-00-40-00000000000000")),
			ver:  0,
			vrnt: VariantNCS,
		},

		// Version 0, Variant 0b_011_00000,
		{
			uuid: Must(Parse("000000000000-00-00-60-00000000000000")),
			ver:  0,
			vrnt: VariantNCS,
		},

		// Version 0, Variant 0b_100_00000,
		{
			uuid: Must(Parse("000000000000-00-00-80-00000000000000")),
			ver:  0,
			vrnt: VariantRFC4122,
		},

		// Version 0, Variant 0b_101_00000,
		{
			uuid: Must(Parse("000000000000-00-00-a0-00000000000000")),
			ver:  0,
			vrnt: VariantRFC4122,
		},

		// Version 0, Variant 0b_110_00000,
		{
			uuid: Must(Parse("000000000000-00-00-c0-00000000000000")),
			ver:  0,
			vrnt: VariantMicrosoft,
		},

		// Version 0, Variant 0b_111_00000,
		{
			uuid: Must(Parse("000000000000-00-00-e0-00000000000000")),
			ver:  0,
			vrnt: VariantFuture,
		},

		// Random
		{
			uuid: Must(Random()),
			ver:  VersionRandom,
			vrnt: VariantRFC4122,
		},
	}

	for _, test := range tests {
		ver := test.uuid.Version()
		vrnt := test.uuid.Variant()

		if ver != test.ver {
			t.Errorf("%s: version mismatch:\n"+
				"expected: %d\n"+
				"present:  %d\n",
				test.uuid, test.ver, ver)
		}

		if vrnt != test.vrnt {
			t.Errorf("%s: variant mismatch:\n"+
				"expected: %d\n"+
				"present:  %d\n",
				test.uuid, test.vrnt, vrnt)
		}
	}
}

// TestSHA1 tests Name-Based UUIDs generation
func TestNameBased(t *testing.T) {
	type testData struct {
		space UUID   // The namespace
		name  string // String to use to generate UUID from
		uuid  UUID   // Expected output
	}

	tests := []testData{
		// Based on RFC9562, A.4
		{
			space: NameSpaceDNS,
			name:  "www.example.com",
			uuid:  Must(Parse("2ed6657de927568b95e12665a8aea6a2")),
		},
	}

	for _, test := range tests {
		uuid := SHA1(test.space, test.name)
		if uuid != test.uuid {
			t.Errorf("%s: output mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.name, test.uuid, uuid)
		}
	}
}

// TestMust tests how Must panics in a case of error
func TestMust(t *testing.T) {
	defer func() {
		reason := recover()
		if reason == nil {
			return
		}

		err, ok := reason.(error)
		if !ok || err == nil {
			panic(reason)
		}
	}()

	Must(Parse(""))
	t.Errorf("Must didn't panic!")
}
