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
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/OpenPrinting/goipp"
)

// ippCodec represents actions required to encode/decode structures
// of the particular type. Codecs are generated at initialization and
// then reused, to minimize performance overhead associated with
// reflection
type ippCodec struct {
	t     reflect.Type   // Type of structure
	steps []ippCodecStep // Encoding/decoding steps
}

// ippCodecStep represents a single encoding/decoding step for the
// ippCodec
type ippCodecStep struct {
	offset               uintptr   // Field offset within the structure
	attrName             string    // IPP attribute name
	attrTag              goipp.Tag // IPP attribute tag
	slice                bool      // It's a slice of values
	flgRange, flgNorange bool      // Range/NoRange classification

	// Encode/decode functions
	encode func(p unsafe.Pointer) []goipp.Value
	decode func(p unsafe.Pointer, v goipp.Values) error
}

// Standard codecs, precompiled
var (
	// ippCodecPrinterAttributes is PrinterAttributes codec
	ippCodecPrinterAttributes = ippCodecMustGenerate(
		reflect.TypeOf(PrinterAttributes{}))
)

// ippCodecMustGenerate calls ippCodecGenerate for the particular
// type and panics if it fails
func ippCodecMustGenerate(t reflect.Type) *ippCodec {
	codec, err := ippCodecGenerate(t)
	if err != nil {
		panic(err)
	}
	return codec
}

// ippCodecGenerate generates codec for the particular type.
func ippCodecGenerate(t reflect.Type) (*ippCodec, error) {
	if t.Kind() != reflect.Struct {
		err := fmt.Errorf("%s: is not struct", t)
		return nil, err
	}

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

		// Parse ipp: struct tag
		tag, err := ippStructTagParse(tagStr)
		if err != nil {
			return nil, fmt.Errorf("%s.%s: %w", t, fld.Name, err)
		}

		// Obtain ippCodecMethods
		fldType := fld.Type
		fldKind := fldType.Kind()
		slice := fldKind == reflect.Slice
		if slice {
			fldType = fldType.Elem()
			fldKind = fldType.Kind()
		}

		collection := false

		methods := ippCodecMethodsByType[fldType]
		if methods == nil {
			methods = ippCodecMethodsByKind[fldKind]
		}
		if methods == nil && fldKind == reflect.Struct {
			var skip func(error) bool

			if tag.flgRange {
				skip = ippErrIsIntegerToRangeConversion
			} else if tag.flgNorange {
				skip = ippErrIsRangeToIntegerConversion
			}

			methods, err = ippCodecMethodsCollection(fldType,
				slice, skip)

			if err != nil {
				return nil, err
			}

			collection = true
		}

		if methods == nil {
			err := fmt.Errorf("%s.%s: %s type not supported",
				t, fld.Name, fldKind)

			return nil, err
		}

		// Generate encoding/decoding step
		step := ippCodecStep{
			offset:     fld.Offset,
			attrName:   tag.name,
			attrTag:    tag.ippTag,
			slice:      slice,
			flgRange:   tag.flgRange,
			flgNorange: tag.flgNorange,
		}

		if step.attrTag == 0 {
			step.attrTag = methods.ippTag
		}

		if slice {
			t := reflect.SliceOf(fldType)
			step.encode = func(p unsafe.Pointer) []goipp.Value {
				return ippEncSlice(p, t, methods.encode)
			}

			if !collection {
				step.decode = func(p unsafe.Pointer,
					vals goipp.Values) error {
					return ippDecSlice(p, vals, t,
						methods.decode)
				}
			} else {
				step.decode = methods.decode
			}
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
func (codec ippCodec) encode(in interface{}, attrs *goipp.Attributes) {
	// Check for type mismatch
	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Pointer || v.Elem().Type() != codec.t {
		err := fmt.Errorf("Encoder for %q applied to %q",
			reflect.PtrTo(codec.t), reflect.TypeOf(in))
		panic(err)
	}

	// Now encode, step by step
	//
	// Note, attribute names may duplicate, because the same
	// name can be used twice in Go structure: first time for
	// the Integer values and second time for the RangeOfInteger
	// values. In IPP it is encoded under the same attribute name,
	// but in Go structure there will be two fields
	//
	// So we cannot encode into the linear sequence of attributes
	// directly (otherwise we may have duplicate attributes on
	// output). Instead, first we collect attributes into the
	// map, indexed by attribute name, them export this map
	// into the final linear sequence
	p := unsafe.Pointer(v.Pointer())
	attrByName := make(map[string]goipp.Attribute)

	for _, step := range codec.steps {
		attr := goipp.Attribute{Name: step.attrName}
		val := step.encode(unsafe.Pointer(uintptr(p) + step.offset))

		if len(val) != 0 {
			for _, v := range val {
				attr.Values.Add(step.attrTag, v)
			}

			if prev, found := attrByName[step.attrName]; found {
				attr.Values = append(prev.Values,
					attr.Values...)
			}
			attrByName[step.attrName] = attr
		}
	}

	// Export encoded attributes from the attrByName
	//
	// We sort resulting attributes by name with the purpose
	// ho have always reproducible result. IPP doesn't dictate
	// any particular order of attributes, but Go map traversal
	// will always produce a different order of attributes, which
	// we want to avoid
	newattrs := make(goipp.Attributes, 0, len(attrByName))
	for _, attr := range attrByName {
		attrs.Add(attr)
	}
	sort.Slice(newattrs, func(i, j int) bool {
		return newattrs[i].Name < newattrs[j].Name
	})
	*attrs = append(*attrs, newattrs...)
}

// Decode structure from the goipp.Attributes
//
// This function wraps (ippCodec) doDecode, adding some common
// error handling and so on
func (codec ippCodec) decode(out interface{}, attrs goipp.Attributes) error {
	err := codec.doDecode(out, attrs)
	if err != nil {
		err = fmt.Errorf("IPP decode %s: %w", codec.t, err)
	}
	return err
}

// doDecode performs the actual work of decoding structure
// from goipp.Attributes
func (codec ippCodec) doDecode(out interface{}, attrs goipp.Attributes) error {
	// Check for type mismatch
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Pointer || v.Elem().Type() != codec.t {
		err := fmt.Errorf("Decoder for %q applied to %q",
			reflect.PtrTo(codec.t), reflect.TypeOf(out))
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
		// IPP protocol doesn't allow attributes without values
		// and github.com/OpenPrinting/goipp will never return
		// them.
		//
		// So if Values slice is empty, this is artificial construct.
		// Reject it in this case.
		if len(attr.Values) == 0 {
			err := fmt.Errorf("%q: at least 1 value required",
				step.attrName)
			return err
		}

		// Call decoder
		err := step.decode(unsafe.Pointer(uintptr(p)+step.offset),
			attr.Values)

		var conv ippErrConvert
		if errors.As(err, &conv) {
			switch {
			case step.flgRange &&
				(conv.from == goipp.TypeInteger) &&
				(conv.to == goipp.TypeRange):
				err = nil

			case step.flgNorange &&
				(conv.to == goipp.TypeInteger) &&
				(conv.from == goipp.TypeRange):
				err = nil
			}
		}

		if err != nil {
			return fmt.Errorf("%q: %w", step.attrName, err)
		}
	}

	return nil
}

// ippStructTag represents parsed ipp: struct tag
type ippStructTag struct {
	name                 string    // Attribute name
	ippTag               goipp.Tag // IPP tag
	flgRange, flgNorange bool      // "range"/"norange" flags
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
		switch strings.ToLower(part) {
		case "boolean":
			tag.ippTag = goipp.TagBoolean
		case "charset":
			tag.ippTag = goipp.TagCharset
		case "datetime":
			tag.ippTag = goipp.TagDateTime
		case "enum":
			tag.ippTag = goipp.TagEnum
		case "integer":
			tag.ippTag = goipp.TagInteger
		case "keyword":
			tag.ippTag = goipp.TagKeyword
		case "mimemediatype":
			tag.ippTag = goipp.TagMimeType
		case "name":
			tag.ippTag = goipp.TagName
		case "namewithlanguage":
			tag.ippTag = goipp.TagNameLang
		case "naturallanguage":
			tag.ippTag = goipp.TagLanguage
		case "rangeofinteger":
			tag.ippTag = goipp.TagRange
		case "resolution":
			tag.ippTag = goipp.TagResolution
		case "string":
			tag.ippTag = goipp.TagString
		case "text":
			tag.ippTag = goipp.TagText
		case "textwithlanguage":
			tag.ippTag = goipp.TagTextLang
		case "uri":
			tag.ippTag = goipp.TagURI
		case "urischeme":
			tag.ippTag = goipp.TagURIScheme

		case "range":
			tag.flgRange = true
		case "norange":
			tag.flgNorange = true
		}
	}

	return tag, nil
}

// ippCodecMethods contains per-reflect.Kind encode and decode
// functions
type ippCodecMethods struct {
	ippTag goipp.Tag
	encode func(p unsafe.Pointer) []goipp.Value
	decode func(p unsafe.Pointer, v goipp.Values) error
}

// ----- ippCodecMethods for slices -----

// ippEncSlice encodes slice of values
//
// p is pointer to slice, t is slice type (so for []string t will be
// []string while p's value will be of type *[]string, represented
// as unsafe.Pointer)
//
// encode is the single-value encoder (i.e., ippEncString for slice
// of strings)
func ippEncSlice(p unsafe.Pointer,
	t reflect.Type, encode func(unsafe.Pointer) []goipp.Value) []goipp.Value {

	slice := reflect.NewAt(t, p).Elem()
	if slice.IsNil() {
		return nil
	}

	vals := make([]goipp.Value, slice.Len())
	for i := range vals {
		p2 := unsafe.Pointer(slice.Index(i).Addr().Pointer())
		v := encode(p2)
		vals[i] = v[0]
	}

	return vals
}

// ippDecSlice decodes slice of values
func ippDecSlice(p unsafe.Pointer, vals goipp.Values,
	t reflect.Type, decode func(unsafe.Pointer, goipp.Values) error) error {

	slice := reflect.MakeSlice(t, 0, len(vals))
	tmp := reflect.New(t.Elem())

	for i := range vals {
		tmp.Elem().SetZero()
		err := decode(unsafe.Pointer(tmp.Pointer()), vals[i:i+1])
		if err != nil {
			return err
		}

		slice = reflect.Append(slice, tmp.Elem())
	}

	out := reflect.NewAt(t, p)
	out.Elem().Set(slice)

	return nil
}

// ----- ippCodecMethods for nested structures -----

// ippCodecMethodsCollection creates ippCodecMethods for encoding
// nested structure or slice of structures as IPP Collection
func ippCodecMethodsCollection(t reflect.Type, slice bool,
	skip func(error) bool) (*ippCodecMethods, error) {

	codec, err := ippCodecGenerate(t)
	if err != nil {
		return nil, err
	}

	m := &ippCodecMethods{
		ippTag: goipp.TagBeginCollection,

		encode: func(p unsafe.Pointer) []goipp.Value {
			return ippEncCollection(p, codec)
		},
	}

	if slice {
		m.decode = func(p unsafe.Pointer, v goipp.Values) error {
			return ippDecCollectionSlice(p, v, codec, skip)
		}
	} else {
		m.decode = func(p unsafe.Pointer, v goipp.Values) error {
			return ippDecCollection(p, v, codec, skip)
		}
	}

	return m, nil
}

// Encode: nested structure as collection
func ippEncCollection(p unsafe.Pointer, codec *ippCodec) []goipp.Value {

	ss := reflect.NewAt(codec.t, p).Interface()

	var attrs goipp.Attributes
	codec.encode(ss, &attrs)

	return []goipp.Value{goipp.Collection(attrs)}
}

// Decode: single nested structure from collection
func ippDecCollection(p unsafe.Pointer, vals goipp.Values,
	codec *ippCodec, skip func(error) bool) error {

	slice, err := ippDecCollectionInternal(p, vals, codec, skip)
	if err != nil {
		return err
	}

	ss := reflect.NewAt(codec.t, p)
	ss.Elem().Set(slice.Index(0))

	return nil
}

// Decode: nested slice of structures from collection
func ippDecCollectionSlice(p unsafe.Pointer, vals goipp.Values,
	codec *ippCodec, skip func(error) bool) error {

	slice, err := ippDecCollectionInternal(p, vals, codec, skip)
	if err != nil {
		return err
	}

	out := reflect.NewAt(reflect.SliceOf(codec.t), p).Elem()
	out.Set(slice)

	return nil
}

// ippDecCollectionInternal decodes collection as a slice
// of decoded structures of type codec.t
//
// If some of attribute values decodes with error and skip
// is not nil and skip(err) == true, this value simply skipped
// instead of propagation the error up to the stack.

// The purpose of this mechanism is to allow attributes with either
// Integer or RangeOfInteger values within underlying collections,
// like "media-size-supported" in printer attributes. In Go structure
// such an attribute take 2 fields with different range/noRange types,
// and this mechanism is used to filter values between these two
// fields when decoding.
//
// It returns slice of decodes structures, represented
// as reflect.Value.
func ippDecCollectionInternal(p unsafe.Pointer, vals goipp.Values,
	codec *ippCodec, skip func(error) bool) (reflect.Value, error) {

	slice := reflect.Zero(reflect.SliceOf(codec.t))
	for i := range vals {
		coll, ok := vals[i].V.(goipp.Collection)
		if !ok {
			err := ippErrConvert{
				from: vals[i].V.Type(),
				to:   goipp.TypeCollection,
			}
			return reflect.Value{}, err
		}

		attrs := goipp.Attributes(coll)
		ss := reflect.New(codec.t)

		err := codec.doDecode(ss.Interface(), attrs)
		if err != nil {
			if skip != nil && skip(err) {
				continue
			}

			return reflect.Value{}, err
		}

		slice = reflect.Append(slice, ss.Elem())
	}

	return slice, nil
}

// ----- ippCodecMethods for particular types -----

// ippCodecMethodsByType maps reflect.Type to the particular
// ippCodecMethods structure
var ippCodecMethodsByType = map[reflect.Type]*ippCodecMethods{
	reflect.TypeOf(goipp.Range{}): &ippCodecMethods{
		ippTag: goipp.TagRange,
		encode: ippEncRange,
		decode: ippDecRange,
	},

	reflect.TypeOf(goipp.Version(0)): &ippCodecMethods{
		ippTag: goipp.TagKeyword,
		encode: ippEncVersion,
		decode: ippDecVersion,
	},

	reflect.TypeOf(time.Time{}): &ippCodecMethods{
		ippTag: goipp.TagDateTime,
		encode: ippEncDateTime,
		decode: ippDecDateTime,
	},
}

// Encode: goipp.Range
func ippEncRange(p unsafe.Pointer) []goipp.Value {
	in := *(*goipp.Range)(p)
	out := []goipp.Value{goipp.Range(in)}
	return out
}

// Decode: goipp.Range
func ippDecRange(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Range)
	if !ok {
		err := ippErrConvert{
			from: vals[0].V.Type(),
			to:   goipp.TypeRange,
		}
		return err
	}

	*(*goipp.Range)(p) = res
	return nil
}

// Encode: goipp.Version
func ippEncVersion(p unsafe.Pointer) []goipp.Value {
	in := *(*goipp.Version)(p)
	out := []goipp.Value{goipp.String(in.String())}
	return out
}

// Decode: goipp.Version
func ippDecVersion(p unsafe.Pointer, vals goipp.Values) error {
	s, ok := vals[0].V.(goipp.String)
	if !ok {
		err := ippErrConvert{
			from: vals[0].V.Type(),
			to:   goipp.TypeString,
		}
		return err
	}

	ver, err := ippDecVersionString(string(s))
	if err != nil {
		return err
	}

	*(*goipp.Version)(p) = ver
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

// Encode: time.Time
func ippEncDateTime(p unsafe.Pointer) []goipp.Value {
	in := *(*time.Time)(p)
	out := []goipp.Value{goipp.Time{Time: in}}
	return out
}

// Decode: time.Time
func ippDecDateTime(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Time)
	if !ok {
		err := ippErrConvert{
			from: vals[0].V.Type(),
			to:   goipp.TypeDateTime,
		}
		return err
	}

	*(*time.Time)(p) = res.Time
	return nil

}

// ----- ippCodecMethods for particular reflect.Kind-s -----

// ippCodecMethodsByKind maps reflect.Kind to the particular
// ippCodecMethods structure
var ippCodecMethodsByKind = map[reflect.Kind]*ippCodecMethods{
	reflect.Bool: &ippCodecMethods{
		ippTag: goipp.TagBoolean,
		encode: ippEncBool,
		decode: ippDecBool,
	},

	reflect.Int: &ippCodecMethods{
		ippTag: goipp.TagInteger,
		encode: ippEncInt,
		decode: ippDecInt,
	},

	reflect.String: &ippCodecMethods{
		ippTag: goipp.TagText,
		encode: ippEncString,
		decode: ippDecString,
	},

	reflect.Uint16: &ippCodecMethods{
		ippTag: goipp.TagInteger,
		encode: ippEncUint16,
		decode: ippDecUint16,
	},
}

// Encode: bool
func ippEncBool(p unsafe.Pointer) []goipp.Value {
	in := *(*bool)(p)
	out := []goipp.Value{goipp.Boolean(in)}
	return out
}

// Decode: bool
func ippDecBool(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Boolean)
	if !ok {
		err := ippErrConvert{
			from: vals[0].V.Type(),
			to:   goipp.TypeBoolean,
		}
		return err
	}

	*(*bool)(p) = bool(res)
	return nil
}

// Encode: int
func ippEncInt(p unsafe.Pointer) []goipp.Value {
	in := *(*int)(p)
	out := []goipp.Value{goipp.Integer(in)}
	return out
}

// Decode: int
func ippDecInt(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		err := ippErrConvert{
			from: vals[0].V.Type(),
			to:   goipp.TypeInteger,
		}
		return err
	}

	*(*int)(p) = int(res)
	return nil
}

// Encode: string
func ippEncString(p unsafe.Pointer) []goipp.Value {
	in := *(*string)(p)
	out := []goipp.Value{goipp.String(in)}
	return out
}

// Decode: string
func ippDecString(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.String)
	if !ok {
		err := ippErrConvert{
			from: vals[0].V.Type(),
			to:   goipp.TypeString,
		}
		return err
	}

	*(*string)(p) = string(res)
	return nil
}

// Encode: uint16
func ippEncUint16(p unsafe.Pointer) []goipp.Value {
	in := *(*uint16)(p)
	out := []goipp.Value{goipp.Integer(in)}
	return out
}

// Decode: uint16
func ippDecUint16(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		err := ippErrConvert{
			from: vals[0].V.Type(),
			to:   goipp.TypeInteger,
		}
		return err
	}

	if res < 0 || res > math.MaxUint16 {
		err := fmt.Errorf("Value %d out of range", res)
		return err
	}

	*(*uint16)(p) = uint16(res)
	return nil
}

// ----- Errors, specific to IPP codec -----

// Can't convert XXX to YYY
type ippErrConvert struct {
	from, to goipp.Type
}

// Convert ippErrConvert to string.
// Implements error interface.
func (err ippErrConvert) Error() string {
	return fmt.Sprintf("can't convert %s to %s", err.from, err.to)
}

// ippErrIsIntegerToRangeConversion returns true if error is
// ippErrConvert caused by Integer->Range conversion
func ippErrIsIntegerToRangeConversion(err error) bool {
	var conv ippErrConvert
	if errors.As(err, &conv) {
		return conv.from == goipp.TypeInteger &&
			conv.to == goipp.TypeRange
	}
	return false
}

// ippErrIsIntegerToRangeConversion returns true if error is
// ippErrConvert caused by Range->Integer conversion
func ippErrIsRangeToIntegerConversion(err error) bool {
	var conv ippErrConvert
	if errors.As(err, &conv) {
		return conv.to == goipp.TypeInteger &&
			conv.from == goipp.TypeRange
	}
	return false
}
