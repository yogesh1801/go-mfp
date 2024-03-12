// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Encoding and decoding structures to/from IPP attributes

package ippx

import (
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
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
	slice    bool
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

	println("=============================")
	p2 := &PrinterAttributes{}
	err := ippCodecPrinterAttributes.decode(p2, msg.Printer)
	if err != nil {
		panic(err)
	}

	v := reflect.ValueOf(*p2)
	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		fld := v.Field(i)
		fmt.Printf("%s: %#v\n", name, fld.Interface())
	}
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

		// Obtain ippCodecMethods
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
		if methods == nil && fldKind == reflect.Struct {
			methods, err = ippCodecMethodsCollection(fldType)
			if err != nil {
				return nil, err
			}
		}

		if methods == nil {
			err := fmt.Errorf("%s: %s type not supported",
				fld.Name, fldKind)

			return nil, err
		}

		// Generate encoding/decoding step
		step := ippCodecStep{
			offset:   fld.Offset,
			attrName: tag.name,
			attrTag:  tag.ippTag,
			slice:    slice,
		}

		if step.attrTag == 0 {
			step.attrTag = methods.ippTag
		}

		if slice {
			step.encode = methods.encodeSlice
			step.decode = methods.decodeSlice
		} else {
			step.encode = methods.encode
			step.decode = methods.decode
		}

		// Append step to the codec
		codec.steps = append(codec.steps, step)
	}

	return codec, nil
}

// Encode structure into the goipp.Attributes
//
// This function wraps (ippCodec) doEncode, adding some common
// error handling and so on
func (codec ippCodec) encode(in interface{}, attrs *goipp.Attributes) error {
	err := codec.doEncode(in, attrs)
	if err != nil {
		err = fmt.Errorf("IPP encode %s: %s", codec.t.Name(), err)
	}
	return err
}

// Decode structure from the goipp.Attributes
//
// This function wraps (ippCodec) doDecode, adding some common
// error handling and so on
func (codec ippCodec) decode(out interface{}, attrs goipp.Attributes) error {
	err := codec.doDecode(out, attrs)
	if err != nil {
		err = fmt.Errorf("IPP decode %s: %s", codec.t.Name(), err)
	}
	return err
}

// doEncode performs the actual work of encoding structure
// into goipp.Attributes
func (codec ippCodec) doEncode(in interface{}, attrs *goipp.Attributes) error {
	// Check for type mismatch
	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Pointer || v.Elem().Type() != codec.t {
		err := fmt.Errorf("Encoder for %q applied to %q",
			"*"+codec.t.Name(), reflect.TypeOf(in).Name())
		panic(err)
	}

	// Now encode, step by step
	p := unsafe.Pointer(v.Pointer())
	for _, step := range codec.steps {
		attr := goipp.Attribute{Name: step.attrName}
		val, err := step.encode(unsafe.Pointer(uintptr(p) + step.offset))
		if err != nil {
			return fmt.Errorf("%q: %s", step.attrName, err)
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

// doDecode performs the actual work of decoding structure
// from goipp.Attributes
func (codec ippCodec) doDecode(out interface{}, attrs goipp.Attributes) error {
	// Check for type mismatch
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Pointer || v.Elem().Type() != codec.t {
		err := fmt.Errorf("Decoder for %q applied to %q",
			"*"+codec.t.Name(), reflect.TypeOf(out).Name())
		panic(err)
	}

	// Build map of attributes
	attrByName := make(map[string]goipp.Attribute)
	for _, attr := range attrs {
		// If we see some attribute twice, we simply concatenate
		// values
		//
		// FIXME: check against IPP specs what is better to do
		// here
		if prev, found := attrByName[attr.Name]; found {
			attr.Values = append(prev.Values, attr.Values...)
		}
		attrByName[attr.Name] = attr
	}

	// Now decode, step by step
	p := unsafe.Pointer(v.Pointer())
	for _, step := range codec.steps {
		// Lookup the attribute
		attr, found := attrByName[step.attrName]
		if !found {
			// FIXME: place to handle required attributes
			continue
		}

		// If not slice, at least one value must be present
		if !step.slice && len(attr.Values) == 0 {
			err := fmt.Errorf("%s: at least 1 value required",
				step.attrName)
			return err
		}

		// Call decoder
		err := step.decode(unsafe.Pointer(uintptr(p)+step.offset), attr.Values)
		if err != nil {
			return fmt.Errorf("%q: %s", step.attrName, err)
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
	decode, decodeSlice func(p unsafe.Pointer, v goipp.Values) error
}

// ----- ippCodecMethods for nested structures -----

// ippCodecMethodsCollection creates ippCodecMethods for encoding
// nested structure or slice of structures as IPP Collection
func ippCodecMethodsCollection(t reflect.Type) (*ippCodecMethods, error) {
	codec, err := ippCodecGenerate(t)
	if err != nil {
		return nil, err
	}

	m := &ippCodecMethods{
		ippTag: goipp.TagBeginCollection,

		encode: func(p unsafe.Pointer) ([]goipp.Value, error) {
			return ippEncCollection(p, codec)
		},

		encodeSlice: func(p unsafe.Pointer) ([]goipp.Value, error) {
			return ippEncCollectionSlice(p, codec)
		},

		decode: func(p unsafe.Pointer, v goipp.Values) error {
			return ippDecCollection(p, v, codec)
		},

		decodeSlice: func(p unsafe.Pointer, v goipp.Values) error {
			return ippDecCollectionSlice(p, v, codec)
		},
	}

	return m, nil
}

// Encode: single nested structure as collection
func ippEncCollection(p unsafe.Pointer,
	codec *ippCodec) ([]goipp.Value, error) {

	ss := reflect.NewAt(codec.t, p).Interface()

	var attrs goipp.Attributes
	err := codec.doEncode(ss, &attrs)
	if err != nil {
		return nil, err
	}

	return []goipp.Value{goipp.Collection(attrs)}, nil
}

// Encode: nested slice of structure as collection
func ippEncCollectionSlice(p unsafe.Pointer,
	codec *ippCodec) ([]goipp.Value, error) {
	return nil, nil
}

// Decode: single nested structure from collection
func ippDecCollection(p unsafe.Pointer, v goipp.Values,
	codec *ippCodec) error {
	return errors.New("NOT IMPLEMENTED")
}

// Decode: nested nested slice of structures from collection
func ippDecCollectionSlice(p unsafe.Pointer, v goipp.Values,
	codec *ippCodec) error {
	return errors.New("NOT IMPLEMENTED")
}

// ----- ippCodecMethods for particular types -----

// ippCodecMethodsByType maps reflect.Type to the particular
// ippCodecMethods structure
var ippCodecMethodsByType = map[reflect.Type]*ippCodecMethods{
	reflect.TypeOf(goipp.Range{}): &ippCodecMethods{
		ippTag:      goipp.TagRange,
		encode:      ippEncRange,
		encodeSlice: ippEncRangeSlice,
		decode:      ippDecRange,
		decodeSlice: ippDecRangeSlice,
	},

	reflect.TypeOf(goipp.Version(0)): &ippCodecMethods{
		ippTag:      goipp.TagKeyword,
		encode:      ippEncVersion,
		encodeSlice: ippEncVersionSlice,
		decode:      ippDecVersion,
		decodeSlice: ippDecVersionSlice,
	},
}

// Encode: single goipp.Range
func ippEncRange(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*goipp.Range)(p)
	out := []goipp.Value{goipp.Range(in)}
	return out, nil
}

// Encode: slice of goipp.Range
func ippEncRangeSlice(p unsafe.Pointer) ([]goipp.Value, error) {
	in := *(*[]goipp.Range)(p)
	out := make([]goipp.Value, len(in))

	for i := range in {
		out[i] = goipp.Range(in[i])
	}

	return out, nil
}

// Decode: single goipp.Range
func ippDecRange(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Range)
	if !ok {
		err := fmt.Errorf("can't convert %s to RangeOfInteger",
			vals[0].V.Type().String())
		return err
	}

	*(*goipp.Range)(p) = goipp.Range(res)
	return nil
}

// Decode: slice of goipp.Range
func ippDecRangeSlice(p unsafe.Pointer, vals goipp.Values) error {
	out := make([]goipp.Range, len(vals))
	for i, val := range vals {
		res, ok := val.V.(goipp.Range)
		if !ok {
			err := fmt.Errorf("can't convert %s to RangeOfInteger",
				val.V.Type().String())
			return err
		}
		out[i] = goipp.Range(res)
	}

	*(*[]goipp.Range)(p) = out
	return nil
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

// Decode: single goipp.Version
func ippDecVersion(p unsafe.Pointer, vals goipp.Values) error {
	s, ok := vals[0].V.(goipp.String)
	if !ok {
		err := fmt.Errorf("can't convert %s to String",
			vals[0].V.Type().String())
		return err
	}

	ver, err := ippDecVersionString(string(s))
	if err != nil {
		return err
	}

	*(*goipp.Version)(p) = ver
	return nil
}

// Decode: slice of goipp.Version
func ippDecVersionSlice(p unsafe.Pointer, vals goipp.Values) error {
	out := make([]goipp.Version, len(vals))
	for i, val := range vals {
		s, ok := val.V.(goipp.String)
		if !ok {
			err := fmt.Errorf("can't convert %s to String",
				val.V.Type().String())
			return err
		}

		var err error
		out[i], err = ippDecVersionString(string(s))
		if err != nil {
			return err
		}
	}

	*(*[]goipp.Version)(p) = out
	return nil
}

// Decode: IPP version string (X.Y).
// Common function for ippDecVersion() and ippDecVersionSlice()
func ippDecVersionString(s string) (goipp.Version, error) {
	var major, minor uint64
	var err error

	ss := strings.Split(s, ".")
	if len(ss) != 2 {
		goto ERROR
	}

	major, err = strconv.ParseUint(ss[0], 10, 8)
	if err != nil {
		goto ERROR
	}

	minor, err = strconv.ParseUint(ss[1], 10, 8)
	if err != nil {
		goto ERROR
	}

	return goipp.MakeVersion(uint8(major), uint8(minor)), nil

ERROR:
	return 0, fmt.Errorf("%q: invalid version string", s)
}

// ----- ippCodecMethods for particular reflect.Kind-s -----

// ippCodecMethodsByKind maps reflect.Kind to the particular
// ippCodecMethods structure
var ippCodecMethodsByKind = map[reflect.Kind]*ippCodecMethods{
	reflect.Bool: &ippCodecMethods{
		ippTag:      goipp.TagBoolean,
		encode:      ippEncBool,
		encodeSlice: ippEncBoolSlice,
		decode:      ippDecBool,
		decodeSlice: ippDecBoolSlice,
	},

	reflect.Int: &ippCodecMethods{
		ippTag:      goipp.TagInteger,
		encode:      ippEncInt,
		encodeSlice: ippEncIntSlice,
		decode:      ippDecInt,
		decodeSlice: ippDecIntSlice,
	},

	reflect.String: &ippCodecMethods{
		ippTag:      goipp.TagText,
		encode:      ippEncString,
		encodeSlice: ippEncStringSlice,
		decode:      ippDecString,
		decodeSlice: ippDecStringSlice,
	},

	reflect.Uint16: &ippCodecMethods{
		ippTag:      goipp.TagInteger,
		encode:      ippEncUint16,
		encodeSlice: ippEncUint16Slice,
		decode:      ippDecUint16,
		decodeSlice: ippDecUint16Slice,
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

// Decode: single bool
func ippDecBool(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Boolean)
	if !ok {
		err := fmt.Errorf("can't convert %s to Boolean",
			vals[0].V.Type().String())
		return err
	}

	*(*bool)(p) = bool(res)
	return nil
}

// Decode: slice of bool
func ippDecBoolSlice(p unsafe.Pointer, vals goipp.Values) error {
	out := make([]bool, len(vals))
	for i, val := range vals {
		res, ok := val.V.(goipp.Boolean)
		if !ok {
			err := fmt.Errorf("can't convert %s to Boolean",
				val.V.Type().String())
			return err
		}
		out[i] = bool(res)
	}

	*(*[]bool)(p) = out
	return nil
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

// Decode: single int
func ippDecInt(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		err := fmt.Errorf("can't convert %s to Integer",
			vals[0].V.Type().String())
		return err
	}

	*(*int)(p) = int(res)
	return nil
}

// Decode: slice of int
func ippDecIntSlice(p unsafe.Pointer, vals goipp.Values) error {
	out := make([]int, len(vals))
	for i, val := range vals {
		res, ok := val.V.(goipp.Integer)
		if !ok {
			err := fmt.Errorf("can't convert %s to Integer",
				val.V.Type().String())
			return err
		}
		out[i] = int(res)
	}

	*(*[]int)(p) = out
	return nil
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

// Decode: single string
func ippDecString(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.String)
	if !ok {
		err := fmt.Errorf("can't convert %s to String",
			vals[0].V.Type().String())
		return err
	}

	*(*string)(p) = string(res)
	return nil
}

// Decode: slice of string
func ippDecStringSlice(p unsafe.Pointer, vals goipp.Values) error {
	out := make([]string, len(vals))
	for i, val := range vals {
		res, ok := val.V.(goipp.String)
		if !ok {
			err := fmt.Errorf("can't convert %s to String",
				val.V.Type().String())
			return err
		}
		out[i] = string(res)
	}

	*(*[]string)(p) = out
	return nil
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

// Decode: single uint16
func ippDecUint16(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		err := fmt.Errorf("can't convert %s to Integer",
			vals[0].V.Type().String())
		return err
	}

	if res < 0 || res > math.MaxUint16 {
		err := fmt.Errorf("Value %d out of range", res)
		return err
	}

	*(*uint16)(p) = uint16(res)
	return nil
}

// Decode: slice of uint16
func ippDecUint16Slice(p unsafe.Pointer, vals goipp.Values) error {
	out := make([]uint16, len(vals))
	for i, val := range vals {
		res, ok := val.V.(goipp.Integer)
		if !ok {
			err := fmt.Errorf("can't convert %s to Integer",
				val.V.Type().String())
			return err
		}

		if res < 0 || res > math.MaxUint16 {
			err := fmt.Errorf("Value %d out of range", res)
			return err
		}

		out[i] = uint16(res)
	}

	*(*[]uint16)(p) = out
	return nil
}
