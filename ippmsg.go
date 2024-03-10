// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP messages encoding decoding

package ippx

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unsafe"

	"github.com/OpenPrinting/goipp"
)

// ippCodec represents actions required to encode/decode structures
// of the particular type. Codecs are generated at initialization and
// then reused, to minimize performance overhead associated with
// reflection
type ippCodec struct {
	t     reflect.Type
	steps []ippCodecStep
}

// ippCodecStep represents a single encoding/decoding step for the
// ippCodec
type ippCodecStep struct {
	offset   uintptr
	attrName string
	attrTag  goipp.Tag
	encode   func(p unsafe.Pointer) ([]goipp.Value, error)
	decode   func(p unsafe.Pointer, v goipp.Values) error
}

// Standard codecs, precompiled
var (
	// ippCodecPrinterAttributes is PrinterAttributes codec
	ippCodecPrinterAttributes = ippCodecMustGenerate(
		reflect.TypeOf(PrinterAttributes{}))
)

func init() {
	println("=============================")
	p := &PrinterAttributes{
		CharsetConfigured:    DefaultCharsetConfigured,
		CharsetSupported:     DefaultCharsetSupported,
		CompressionSupported: []string{"none"},
		IppFeaturesSupported: []string{
			"airprint-1.7",
			"airprint-1.6",
			"airprint-1.5",
			"airprint-1.4",
		},
		IppVersionsSupported: DefaultIppVersionsSupported,
		OperationsSupported: []goipp.Op{
			goipp.OpGetPrinterAttributes,
		},
	}

	msg := goipp.NewResponse(goipp.DefaultVersion, 0, 0)
	ippCodecPrinterAttributes.encode(p, &msg.Printer)
	msg.Print(os.Stdout, false)
}

// ippCodecMustGenerate calls ippCodecGenerate for the particular
// type and panics if it fails
func ippCodecMustGenerate(t reflect.Type) *ippCodec {
	codec, err := ippCodecGenerate(t)
	if err != nil {
		err = fmt.Errorf("%s: %s", t.Name(), err)
		panic(err)
	}
	return codec
}

// ippCodecGenerate generates codec for the particular type.
func ippCodecGenerate(t reflect.Type) (*ippCodec, error) {
	codec := &ippCodec{
		t: t,
	}

	for i := 0; i < t.NumField(); i++ {
		// Fetch field by field
		//
		// - Ignore anonymous fields
		// - Ignore unexported fields
		// - Ignore fields without ipp: tag
		fld := t.Field(i)

		if fld.Anonymous {
			continue
		}

		if !fld.IsExported() {
			continue
		}

		tagStr, found := fld.Tag.Lookup("ipp")
		if !found {
			continue
		}
		println(fld.Name, tagStr)

		// Parse ipp: struct tag
		tag, err := ippStructTagParse(tagStr)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", fld.Name, err)
		}

		// Generate encoding/decoding step
		step := ippCodecStep{
			offset:   fld.Offset,
			attrName: tag.name,
			attrTag:  tag.ippTag,
		}

		fldType := fld.Type
		fldKind := fldType.Kind()
		slice := fldKind == reflect.Slice
		if slice {
			fldType = fldType.Elem()
			fldKind = fldType.Kind()
		}

		methods := ippCodecMethodsByType[fldType]
		if methods == nil {
			methods = ippCodecMethodsByKind[fldKind]
		}

		if methods != nil {
			if step.attrTag == 0 {
				step.attrTag = methods.ippTag
			}

			if slice {
				step.encode = methods.encodeSlice
			} else {
				step.encode = methods.encode
			}
		} else if fldKind == reflect.Struct {
			// FIXME: skip for now
			continue
		} else {
			err := fmt.Errorf("%s: %s type not supported",
				fld.Name, fldKind)

			return nil, err
		}

		// Append step to the codec
		codec.steps = append(codec.steps, step)
	}

	return codec, nil
}

// Encode structure into the goipp.Attributes
func (codec ippCodec) encode(in interface{}, attrs *goipp.Attributes) error {
	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Pointer || v.Elem().Type() != codec.t {
		err := fmt.Errorf("Encoder for %q applied to %q",
			"*"+codec.t.Name(), reflect.TypeOf(in).Name())
		panic(err)
	}

	p := v.Pointer()
	for _, step := range codec.steps {
		attr := goipp.Attribute{Name: step.attrName}
		val, err := step.encode(unsafe.Pointer(p + step.offset))
		if err != nil {
			return fmt.Errorf("%s: %s", step.attrName, err)
		}

		if len(val) != 0 {
			for _, v := range val {
				attr.Values.Add(step.attrTag, v)
			}
			attrs.Add(attr)
		}
	}

	return nil
}

// ippStructTag represents parsed ipp: struct tag
type ippStructTag struct {
	name   string    // Attribute name
	ippTag goipp.Tag // IPP tag
}

// ippStructTagParse parses ipp: struct tag into the
// ippStructTag structure
func ippStructTagParse(s string) (*ippStructTag, error) {
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	if len(parts) < 1 || parts[0] == "" {
		return nil, errors.New("missed attribute name")
	}

	tag := &ippStructTag{
		name: parts[0],
	}

	for _, part := range parts[1:] {
		switch part {
		case "boolean":
			tag.ippTag = goipp.TagBoolean
		case "charset":
			tag.ippTag = goipp.TagCharset
		case "dateTime":
			tag.ippTag = goipp.TagDateTime
		case "enum":
			tag.ippTag = goipp.TagEnum
		case "integer":
			tag.ippTag = goipp.TagInteger
		case "keyword":
			tag.ippTag = goipp.TagKeyword
		case "mimeMediaType":
			tag.ippTag = goipp.TagMimeType
		case "name":
			tag.ippTag = goipp.TagName
		case "nameWithLanguage":
			tag.ippTag = goipp.TagNameLang
		case "naturalLanguage":
			tag.ippTag = goipp.TagLanguage
		case "rangeOfInteger":
			tag.ippTag = goipp.TagRange
		case "resolution":
			tag.ippTag = goipp.TagResolution
		case "string":
			tag.ippTag = goipp.TagString
		case "text":
			tag.ippTag = goipp.TagText
		case "textWithLanguage":
			tag.ippTag = goipp.TagTextLang
		case "uri":
			tag.ippTag = goipp.TagURI
		case "uriScheme":
			tag.ippTag = goipp.TagURIScheme
		}
	}

	return tag, nil
}

// ippCodecMethods contains per-reflect.Kind encode and decode
// functions
type ippCodecMethods struct {
	ippTag              goipp.Tag
	encode, encodeSlice func(p unsafe.Pointer) ([]goipp.Value, error)
	decode              func(p unsafe.Pointer, v goipp.Values) error
}

// ippCodecMethodsByType maps reflect.Type to the particular
// ippCodecMethods structure
var ippCodecMethodsByType = map[reflect.Type]*ippCodecMethods{
	reflect.TypeOf(goipp.Version(0)): &ippCodecMethods{
		ippTag:      goipp.TagKeyword,
		encode:      ippEncVersion,
		encodeSlice: ippEncVersionSlice,
	},
}

// Encode: single goipp.Version
func ippEncVersion(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*goipp.Version)(p)
	out := []goipp.Value{goipp.String(in.String())}
	return out, nil
}

// Encode: slice of goipp.Version
func ippEncVersionSlice(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*[]goipp.Version)(p)
	out := make([]goipp.Value, len(in))

	for i := range in {
		out[i] = goipp.String(in[i].String())
	}

	return out, nil
}

// ippCodecMethodsByKind maps reflect.Kind to the particular
// ippCodecMethods structure
var ippCodecMethodsByKind = map[reflect.Kind]*ippCodecMethods{
	reflect.Bool: &ippCodecMethods{
		ippTag:      goipp.TagBoolean,
		encode:      ippEncBool,
		encodeSlice: ippEncBoolSlice,
	},

	reflect.Int: &ippCodecMethods{
		ippTag:      goipp.TagInteger,
		encode:      ippEncInt,
		encodeSlice: ippEncIntSlice,
	},

	reflect.String: &ippCodecMethods{
		ippTag:      goipp.TagText,
		encode:      ippEncString,
		encodeSlice: ippEncStringSlice,
	},

	reflect.Uint16: &ippCodecMethods{
		ippTag:      goipp.TagInteger,
		encode:      ippEncUint16,
		encodeSlice: ippEncUint16Slice,
	},
}

// Encode: single bool
func ippEncBool(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*bool)(p)
	out := []goipp.Value{goipp.Boolean(in)}
	return out, nil
}

// Encode: slice of bool
func ippEncBoolSlice(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*[]bool)(p)
	out := make([]goipp.Value, len(in))

	for i := range in {
		out[i] = goipp.Boolean(in[i])
	}

	return out, nil
}

// Encode: single int
func ippEncInt(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*int)(p)
	out := []goipp.Value{goipp.Integer(in)}
	return out, nil
}

// Encode: slice of int
func ippEncIntSlice(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*[]int)(p)
	out := make([]goipp.Value, len(in))

	for i := range in {
		out[i] = goipp.Integer(in[i])
	}

	return out, nil
}

// Encode: single string
func ippEncString(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*string)(p)
	out := []goipp.Value{goipp.String(in)}
	return out, nil
}

// Encode: slice of string
func ippEncStringSlice(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*[]string)(p)
	out := make([]goipp.Value, len(in))

	for i := range in {
		out[i] = goipp.String(in[i])
	}

	return out, nil
}

// Encode: single uint16
func ippEncUint16(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*uint16)(p)
	out := []goipp.Value{goipp.Integer(in)}
	return out, nil
}

// Encode: slice of uint16
func ippEncUint16Slice(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*[]uint16)(p)
	out := make([]goipp.Value, len(in))

	for i := range in {
		out[i] = goipp.Integer(in[i])
	}

	return out, nil
}
