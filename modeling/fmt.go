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
	"io"

	"github.com/OpenPrinting/go-mfp/cpython"
)

// formatter saves Model data to files
type formatter struct {
	w   io.Writer // Output io.Writer
	err error     // Sticky error
}

// newFormatter creates a new formatter
func newFormatter(w io.Writer) *formatter {
	return &formatter{w: w}
}

// Format writes Python object into the io.Writer.
func (f *formatter) Format(obj *cpython.Object) error {
	f.formatValue(obj, 0)
	f.write("\n")
	return f.err
}

// formatValue writes Python value into the io.Writer.
// It's a helper function for the formatter.Format
func (f *formatter) formatValue(obj *cpython.Object, indent int) {
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
		f.printf("%s", s)
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
		f.printf("{}\n")
	}

	// Retrieve key strings and item objects
	keys := make([]string, len(keyobjs))
	vals := make([]*cpython.Object, len(keyobjs))

	for i := range keyobjs {
		keys[i], f.err = keyobjs[i].Repr()
		if f.err == nil {
			vals[i], f.err = dict.Get(keyobjs[i])
		}

		if f.err != nil {
			return
		}
	}

	// Format the dictionary
	f.printf("{\n")
	for i := range keys {
		key, val := keys[i], vals[i]

		f.indent(indent + 1)
		f.printf("%s: ", key)

		f.formatValue(val, indent+1)
		if f.err != nil {
			return
		}

		last := i == len(keys)-1
		if !last {
			f.write(",")
		}

		f.write("\n")
	}

	f.indent(indent)
	f.write("}")
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
		f.printf("[]\n")
	}

	// Retrieve all object in the array
	//
	// Here we also decide, will we use horizontal o vertical
	// output format.
	vals := make([]*cpython.Object, length)
	horizontal := true
	for i := 0; i < length; i++ {
		vals[i], f.err = obj.Get(i)
		if f.err != nil {
			return
		}

		// Use horizontal format for a sequence of keywords,
		// they all are Unicode strings from the Python point
		// of view.
		horizontal = horizontal && vals[i].IsUnicode()
	}

	// Format the array
	f.write("[")

	for i := 0; i < length; i++ {
		val := vals[i]

		if !horizontal {
			f.write("\n")
			f.indent(indent + 1)
		}

		f.formatValue(val, indent+1)
		if f.err != nil {
			return
		}

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

	f.write("}")
}

// write writes the string.
func (f *formatter) write(s string) {
	if f.err == nil {
		_, f.err = f.w.Write([]byte(s))
	}
}

// indent writes the indentation space.
func (f *formatter) indent(indent int) {
	if f.err == nil {
		_, f.err = fmt.Fprintf(f.w, "%*s", indent*2, "")
	}
}

// printf writes the formatted string.
func (f *formatter) printf(format string, args ...any) {
	if f.err == nil {
		_, f.err = fmt.Fprintf(f.w, format, args...)
	}
}
