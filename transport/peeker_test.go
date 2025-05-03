// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Peeker test

package transport

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// TestPeeker performs Peeker test
func TestPeeker(t *testing.T) {
	type testData struct {
		in      []byte // Input stream
		peek    []byte // Will peek and compare
		replace []byte // nil - Rewind, not nil - Replace
		out     []byte // Output after Rewind or replace
	}

	tests := []testData{
		{
			in:      []byte("123456789"),
			peek:    []byte("12345"),
			replace: nil,
			out:     []byte("123456789"),
		},

		{
			in:      []byte("123456789"),
			peek:    []byte("12345"),
			replace: []byte("abcdef"),
			out:     []byte("abcdef6789"),
		},

		{
			in:      []byte("123456789"),
			peek:    []byte("12345"),
			replace: []byte{},
			out:     []byte("6789"),
		},
	}

	for _, test := range tests {
		p := NewPeeker(io.NopCloser(bytes.NewReader(test.in)))
		d := make([]byte, len(test.peek))
		p.Read(d)

		if !generic.EqualSlices(test.peek, d) {
			t.Errorf("in=%q, peek %d bytes:\n"+
				"expected: %q\n"+
				"present:  %q",
				test.in, len(test.peek),
				test.peek, d)
			p.Close()
			continue
		}

		if test.replace == nil {
			p.Rewind()
		} else {
			p.Replace(test.replace)
		}

		var buf bytes.Buffer
		io.Copy(&buf, p)

		d = buf.Bytes()
		if !generic.EqualSlices(test.out, d) {
			repl := "nil"
			if test.replace != nil {
				repl = fmt.Sprintf("%q", test.replace)
			}

			t.Errorf("in=%q, peek %d bytes, replace=%s:\n"+
				"expected: %q\n"+
				"present:  %q",
				test.in, len(test.peek), repl,
				test.out, d)
			p.Close()
			continue
		}

		p.Close()
	}
}
