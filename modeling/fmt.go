// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Model file formatter

package modeling

import (
	"fmt"
	"strings"

	"github.com/OpenPrinting/go-mfp/cpython"
)

// formatter saves Model data to files
type formatter struct {
	buf strings.Builder // Output buffer
	err error           // Sticky error
}

// formatPython formats the [cpython.Object] as a string.
func formatPython(obj *cpython.Object) (string, error) {
	var f formatter

	f.format(obj)
	return f.buf.String(), f.err
}

// Format writes Python object into the io.Writer.
func (f *formatter) format(obj *cpython.Object) error {
	f.formatValue(obj, 0)
	f.write("\n")
	return f.err
}

// formatValue writes Python value into the io.Writer.
// It's a helper function for the formatter.Format
func (f *formatter) formatValue(obj *cpython.Object, indent int) {
	// Check for error object
	if f.err == nil {
		f.err = obj.Err()
	}

	// Check for the sticky error
	if f.err != nil {
		return
	}

	// Format the value
	switch {
	case obj.IsDict():
		f.formatDict(obj, indent)

	case obj.IsSeq():
		f.formatArray(obj, indent)

	default:
		var s string
		s, f.err = obj.Repr()
		f.Printf("%s", s)
	}
}

// formatDict writes Python dict-like object into the io.Writer.
// It's a helper function for the formatter.Format
func (f *formatter) formatDict(dict *cpython.Object, indent int) {
	// Retrieve dictionary keys
	var keyobjs []*cpython.Object
	keyobjs, f.err = dict.Keys()
	if f.err != nil {
		return
	}

	if len(keyobjs) == 0 {
		f.Printf("{}\n")
	}

	// Retrieve key strings and item objects
	keys := make([]string, len(keyobjs))
	vals := make([]*cpython.Object, len(keyobjs))

	for i := range keyobjs {
		keys[i], f.err = keyobjs[i].Repr()
		if f.err == nil {
			vals[i] = dict.Get(keyobjs[i])
			f.err = vals[i].Err()
		}

		if f.err != nil {
			return
		}
	}

	// Format the dictionary
	//
	// Only the small dictionaries of simple object are
	// formatted horizontally.
	horizontal := len(keys) <= 2
	for i := 0; i < len(vals) && horizontal; i++ {
		val := vals[i]
		switch {
		case val.IsBool():
		case val.IsFloat():
		case val.IsLong():
		case val.IsUnicode():
		default:
			horizontal = false
		}
	}

	if horizontal {
		// Format dictionary horizontally
		f.Printf("{")

		for i := range keys {
			key, val := keys[i], vals[i]

			f.Printf("%s: ", key)
			f.formatValue(val, indent)

			last := i == len(keys)-1
			if !last {
				f.write(", ")
			}
		}

		f.Printf("}")
	} else {
		// Format dictionary vertically
		f.Printf("{\n")
		for i := range keys {
			key, val := keys[i], vals[i]

			f.indent(indent + 1)
			f.Printf("%s: ", key)

			f.formatValue(val, indent+1)

			last := i == len(keys)-1
			if !last {
				f.write(",")
			}

			f.write("\n")
		}

		f.indent(indent)
		f.write("}")
	}
}

// formatArray writes Python array-like object into the io.Writer.
// It's a helper function for the formatter.Format
func (f *formatter) formatArray(obj *cpython.Object, indent int) {
	// Obtain array length
	var length int
	length, f.err = obj.Len()
	if f.err != nil {
		return
	}

	if length == 0 {
		f.Printf("[]\n")
	}

	// Retrieve all object in the array
	//
	// Here we also decide, will we use horizontal o vertical
	// output format.
	vals := make([]*cpython.Object, length)
	width := 0
	for i := 0; i < length; i++ {
		// Fill array of values
		vals[i] = obj.Get(i)
		f.err = vals[i].Err()
		if f.err != nil {
			return
		}

		// Compute total width
		var s string
		s, f.err = vals[i].Repr()
		if f.err != nil {
			return
		}

		width += len(s) + 3
	}

	horizontal := width <= 60

	// Format the array
	f.write("[")

	for i := 0; i < length; i++ {
		val := vals[i]

		if !horizontal {
			f.write("\n")
			f.indent(indent + 1)
		}

		f.formatValue(val, indent+1)

		last := i == length-1
		if !last {
			f.write(",")
			if horizontal {
				f.write(" ")
			}
		}
	}

	if !horizontal {
		f.write("\n")
		f.indent(indent)
	}

	f.write("]")
}

// write writes the string.
func (f *formatter) write(s string) {
	if f.err == nil {
		_, f.err = f.buf.Write([]byte(s))
	}
}

// indent writes the indentation space.
func (f *formatter) indent(indent int) {
	if f.err == nil {
		_, f.err = fmt.Fprintf(&f.buf, "%*s", indent*2, "")
	}
}

// printf writes the formatted string.
func (f *formatter) Printf(format string, args ...any) {
	if f.err == nil {
		_, f.err = fmt.Fprintf(&f.buf, format, args...)
	}
}
