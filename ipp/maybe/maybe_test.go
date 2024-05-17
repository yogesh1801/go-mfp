// MFP   - Miulti-Function Printers and scanners toolkit
// MAYBE - Go Maybe type for IPP values
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// maybe types test

package maybe_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/OpenPrinting/goipp"
	"github.com/alexpevzner/mfp/ipp/maybe"
)

var (
	// Check that NoValue, Unknown and Unsupported are
	// assignable to everybody.
	_ maybe.Binary       = maybe.NoValue
	_ maybe.Boolean      = maybe.NoValue
	_ maybe.Collection   = maybe.NoValue
	_ maybe.Integer      = maybe.NoValue
	_ maybe.Range        = maybe.NoValue
	_ maybe.Resolution   = maybe.NoValue
	_ maybe.String       = maybe.NoValue
	_ maybe.TextWithLang = maybe.NoValue
	_ maybe.Time         = maybe.NoValue

	_ maybe.Binary       = maybe.Unknown
	_ maybe.Boolean      = maybe.Unknown
	_ maybe.Collection   = maybe.Unknown
	_ maybe.Integer      = maybe.Unknown
	_ maybe.Range        = maybe.Unknown
	_ maybe.Resolution   = maybe.Unknown
	_ maybe.String       = maybe.Unknown
	_ maybe.TextWithLang = maybe.Unknown
	_ maybe.Time         = maybe.Unknown

	_ maybe.Binary       = maybe.Unsupported
	_ maybe.Boolean      = maybe.Unsupported
	_ maybe.Collection   = maybe.Unsupported
	_ maybe.Integer      = maybe.Unsupported
	_ maybe.Range        = maybe.Unsupported
	_ maybe.Resolution   = maybe.Unsupported
	_ maybe.String       = maybe.Unsupported
	_ maybe.TextWithLang = maybe.Unsupported
	_ maybe.Time         = maybe.Unsupported

	// Check that constructors compiles as expected.
	_ maybe.Binary       = maybe.NewBinary([]byte{'h', 'e', 'l', 'l', 'o'})
	_ maybe.Boolean      = maybe.NewBoolean(true)
	_ maybe.Collection   = maybe.NewCollection(goipp.Collection{})
	_ maybe.Integer      = maybe.NewInteger(12345)
	_ maybe.Range        = maybe.NewRange(goipp.Range{})
	_ maybe.Resolution   = maybe.NewResolution(goipp.Resolution{})
	_ maybe.String       = maybe.NewString("hello")
	_ maybe.TextWithLang = maybe.NewTextWithLang(goipp.TextWithLang{})
	_ maybe.Time         = maybe.NewTime(time.Now())
)

// TestVoidValues tests predefined void values
func TestVoidValues(t *testing.T) {
	type testData struct {
		val interface{} // NoValue/Unknown/Unsupported value
		err error       // Expected error
	}

	tests := []testData{
		{maybe.NoValue, maybe.ErrNoValue},
		{maybe.Unknown, maybe.ErrUnknown},
		{maybe.Unsupported, maybe.ErrUnsupported},
	}

	methods := []string{
		"Binary", "Boolean", "Collection", "Integer",
		"Range", "Resolution", "String", "Time",
	}

	for _, test := range tests {
		v := reflect.ValueOf(test.val)
		name := v.Type().String()

		for _, m := range methods {
			meth := v.MethodByName(m)
			if meth == (reflect.Value{}) {
				t.Errorf("%s: lacks method %s", name, m)
			} else {
				title := fmt.Sprintf("(%s) %s()", name, m)
				ret := meth.Call(nil)

				switch {
				case len(ret) != 2:
					t.Errorf("%s:\n"+
						"expected 2 return values, received %d",
						title, len(ret))
				case !ret[0].IsZero():
					t.Errorf("%s:\n"+
						"non-Zero first return value: "+
						"%#v",
						title, ret[0].Interface())
				case !reflect.DeepEqual(ret[1].Interface(), test.err):
					t.Errorf("%s:\n"+
						"error expected %q present %q",
						title,
						test.err, ret[1].Interface())
				}
			}
		}
	}
}
