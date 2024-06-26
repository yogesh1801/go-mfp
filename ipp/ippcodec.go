// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Encoding and decoding structures to/from IPP attributes

package ipp

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/OpenPrinting/goipp"
)

// ippAttrConformance represents conformance level of each
// particular attribute
type ippAttrConformance int

const (
	ipAttrOptional    ippAttrConformance = iota // Optional attribute
	ipAttrRecommended                           // Recommended attribute
	ipAttrRequired                              // Required attribute
)

// ippEncodeAttrs encodes attributes defined by particular structure
// into goipp.Attributes.
//
// The obj parameter must be pointer to structure that implements
// the Object interface. Its codec will be generated on demand.
//
// This function will panic, if codec cannot be generated.
func ippEncodeAttrs(obj Object) goipp.Attributes {
	codec := ippCodecGet(obj)
	return codec.encodeAttrs(obj)
}

// ippDecodeAttrs encodes attributes defined by particular structure
// into goipp.Attributes.
//
// The obj parameter must be pointer to structure that implements
// the Object interface. Its codec will be generated on demand.
//
// This function will panic, if codec cannot be generated.
func ippDecodeAttrs(obj Object, attrs goipp.Attributes) error {
	codec := ippCodecGet(obj)

	err := codec.decodeAttrs(obj, attrs)
	if err == nil {
		obj.Attrs().set(attrs)
	}

	return err
}

// ippKnownAttrs returns information about Object's known attributes
//
// This function will panic, if codec cannot be generated.
func ippKnownAttrs(obj Object) []AttrInfo {
	codec := ippCodecGet(obj)
	return codec.knownAttrs
}

// ippCodec represents actions required to encode/decode structures
// of the particular type. Codecs are generated at initialization and
// then reused, to minimize performance overhead associated with
// reflection
type ippCodec struct {
	t          reflect.Type   // Type of structure
	steps      []ippCodecStep // Encoding/decoding steps
	knownAttrs []AttrInfo     // Known attributes
}

// ippCodecStep represents a single encoding/decoding step for the
// ippCodec
type ippCodecStep struct {
	offset      uintptr            // Field offset within structure
	attrName    string             // IPP attribute name
	attrTag     goipp.Tag          // IPP attribute tag
	zeroTag     goipp.Tag          // How to encode zero value
	slice       bool               // It's a slice of values
	conformance ippAttrConformance // Attribute conformance
	min, max    int                // Range limits for integers

	// Encode/decode functions
	encode  func(p unsafe.Pointer) goipp.Values
	decode  func(p unsafe.Pointer, v goipp.Values) error
	iszero  func(p unsafe.Pointer) bool
	setzero func(p unsafe.Pointer)
}

// Cache of generated codecs
// Works as map[reflect.Type]*ippCodec
var ippCodecCache = sync.Map{}

// ippCodecMustGenerate calls ippCodecGenerate for the particular
// type and panics if it fails
func ippCodecMustGenerate(t reflect.Type) *ippCodec {
	codec, err := ippCodecGenerate(t)
	if err != nil {
		panic(err)
	}
	return codec
}

// ippCodecGet returns codec for the Object. Codecs are generated
// on demand and cached.
//
// This function will panic, if codec generation fails.
func ippCodecGet(obj Object) *ippCodec {
	// Check obj type. It MUST be pointer to structure
	// that implements Object interface.
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Pointer {
		err := fmt.Errorf("%s is not pointer to structure",
			diagTypeName(t))
		panic(err)
	}

	// We need reflect.Type of the structure, not pointer
	t = t.Elem()

	// Lookup cache
	if cached, ok := ippCodecCache.Load(t); ok {
		return cached.(*ippCodec)
	}

	// Generate now
	codec, err := ippCodecGenerate(t)
	if err != nil {
		panic(err)
	}

	// Update cache and return
	ippCodecCache.Store(t, codec)

	return codec
}

// ippCodecGenerate generates codec for the particular type.
// It manages and uses a cache of successfully generated codecs.
func ippCodecGenerate(t reflect.Type) (*ippCodec, error) {
	// Compile new codec
	attrNames := make(map[string]string)
	codec, err := ippCodecGenerateInternal(t, attrNames)

	// At least 1 step must be generated
	if err == nil && len(codec.steps) == 0 {
		err = fmt.Errorf("%s: contains no IPP fields",
			diagTypeName(t))
	}

	if err != nil {
		return nil, err
	}

	// Build knownAttrs
	codec.knownAttrs = make([]AttrInfo, len(codec.steps))
	for i := range codec.steps {
		codec.knownAttrs[i].Name = codec.steps[i].attrName
		codec.knownAttrs[i].Tag = codec.steps[i].attrTag
	}

	// Done for now!
	return codec, nil
}

// ippCodecGenerateInternal is the internal function
// behind the ippCodecGenerate()
//
// It calls itself recursively, to implement support of
// nested (embedded) structures
//
// attrNames is the map of IPP attribute names into
// field names, used for detection and reporting of
// duplicate usage of attribute names
func ippCodecGenerateInternal(t reflect.Type,
	attrNames map[string]string) (*ippCodec, error) {

	if t.Kind() != reflect.Struct {
		err := fmt.Errorf("%s: is not struct", diagTypeName(t))
		return nil, err
	}

	codec := &ippCodec{
		t: t,
	}

	maybeCodecType := reflect.TypeOf((*maybeCodecInterface)(nil)).Elem()

	for i := 0; i < t.NumField(); i++ {
		// Fetch field by field
		//
		// - Ignore anonymous fields
		// - Ignore fields without ipp: tag
		fld := t.Field(i)

		// Handle embedded structures
		if fld.Anonymous {
			if fld.IsExported() &&
				fld.Type.Kind() == reflect.Struct {

				nested, err := ippCodecGenerateInternal(
					fld.Type, attrNames)

				if err != nil {
					return nil, err
				}

				codec.embed(fld.Offset, nested)
			}
			continue
		}

		// Fetch ipp: struct tag string. Ignore fields without this tags.
		tagStr, found := fld.Tag.Lookup("ipp")
		if !found {
			if strings.HasPrefix(string(fld.Tag), "ipp:") {
				err := fmt.Errorf("%s.%s: invalid tag %q",
					diagTypeName(t), fld.Name, fld.Tag)
				return nil, err
			}

			continue
		}

		// Field with ipp: tag must be exported.
		if !fld.IsExported() {
			err := fmt.Errorf("%s.%s: ipp: tag used with unexported field",
				diagTypeName(t), fld.Name)
			return nil, err
		}

		// Parse ipp: struct tag.
		tag, err := ippStructTagParse(tagStr)
		if err != nil {
			err := fmt.Errorf("%s.%s: %w",
				diagTypeName(t), fld.Name, err)
			return nil, err
		}

		// Check attribute name for duplicates.
		if found := attrNames[tag.name]; found != "" {
			err := fmt.Errorf("%s.%s: attribute %q already used by %s",
				diagTypeName(t), fld.Name, tag.name, found)
			return nil, err
		}
		attrNames[tag.name] = fld.Name

		// Obtain fldType.
		//
		// Field may be wrapped into Maybe[T], handle it here.
		fldType := fld.Type
		var maybe maybeCodecInterface

		if reflect.PointerTo(fldType).Implements(maybeCodecType) {
			// If type implements maybeCodecType, it is value
			// wrapped into Maybe[T].
			//
			// We need the underlying type to generate encode/decode steps,
			// but Maybe[T] wrapper will be saved as nil pointer to
			// maybeCodecType. We will need it later to generate Maybe[T]
			// wrappers.
			maybe = reflect.NewAt(fldType, nil).Interface().(maybeCodecInterface)
			fldType = maybe.typeof()
		}

		// Handle slices.
		//
		// Like Maybe[T], slices are handled as wrapper of actual data type.
		// So if we are at slice, move fldType down to the underlying data
		// type and remember the fact; we will need it later to generate
		// slice wrappers.
		fldKind := fldType.Kind()
		slice := fldKind == reflect.Slice
		if slice {
			fldType = fldType.Elem()
			fldKind = fldType.Kind()
		}

		// Now fldType points to the actual type to be encoded and decoded.
		// Obtain its ippCodecMethods.
		methods := ippCodecMethodsByType[fldType]
		if methods == nil {
			methods = ippCodecMethodsByKind[fldKind]
		}
		if methods == nil && fldKind == reflect.Struct {
			methods, err = ippCodecMethodsCollection(fldType, slice)
			if err != nil {
				err = fmt.Errorf("%s.%s: %w",
					diagTypeName(t), fld.Name, err)
				return nil, err
			}
		}

		if methods == nil {
			err := fmt.Errorf("%s.%s: %s type not supported",
				diagTypeName(t), fld.Name, fldKind)

			return nil, err
		}

		// Generate encoding/decoding step for underlying type.
		step := ippCodecStep{
			offset:      fld.Offset,
			attrName:    tag.name,
			attrTag:     tag.ippTag,
			zeroTag:     tag.zeroTag,
			slice:       slice,
			conformance: tag.conformance,
			min:         tag.min,
			max:         tag.max,

			encode: methods.encode,
			decode: methods.decode,

			iszero: func(p unsafe.Pointer) bool {
				return reflect.NewAt(fldType, p).Elem().IsZero()
			},
			setzero: func(p unsafe.Pointer) {
				reflect.NewAt(fldType, p).Elem().SetZero()
			},
		}

		// Guess actual attribute Tag.
		if _, found := kwRegisteredTypes[fldType]; found {
			// Underlying type registered as keyword.
			// Use goipp.TagKeyword.
			step.attrTag = goipp.TagKeyword
		} else if _, found := enRegisteredTypes[fldType]; found {
			// Use goipp.TagEnum for registered enum types.
			step.attrTag = goipp.TagEnum
		}

		if step.attrTag == 0 {
			// There is no Tag override from the struct tag.
			// Use default tag for the data type.
			step.attrTag = methods.defaultIppTag
		}

		// Check for compatibility between IPP representation
		// for the Tag being chosen and Tag implied by the
		// underlying field type.
		//
		// They must be the same (i.e., both are goipp.TypeInteger or
		// both are goipp.TypeResolution and so on).
		//
		// The only exception, goipp.TypeString can be converted
		// into goipp.TypeBinary and visa versa.
		t1 := step.attrTag.Type()
		t2 := methods.defaultIppTag.Type()

		ok := t1 == t2 ||
			t1 == goipp.TypeBinary && t2 == goipp.TypeString ||
			t2 == goipp.TypeBinary && t1 == goipp.TypeString

		if !ok {
			err := fmt.Errorf("%s.%s: can't represent %s as %s",
				diagTypeName(t), fld.Name, fld.Type, step.attrTag)

			return nil, err
		}

		// Generate slice wrapper for slice fields.
		if slice {
			t := reflect.SliceOf(fldType)

			encode := step.encode
			decode := step.decode

			step.encode = func(p unsafe.Pointer) goipp.Values {
				return ippEncSlice(p, t, encode)
			}

			step.decode = func(p unsafe.Pointer,
				vals goipp.Values) error {
				return ippDecSlice(p, vals, t, decode)
			}
		}

		// Generate Maybe[T] wrapper where appropriate.
		if maybe != nil {
			encode := step.encode
			decode := step.decode

			step.encode = func(p unsafe.Pointer) goipp.Values {
				return maybe.encode(p, encode)
			}

			step.decode = func(p unsafe.Pointer,
				vals goipp.Values) error {
				return maybe.decode(p, vals, decode)
			}
		}

		// Append step to the codec
		codec.steps = append(codec.steps, step)
	}

	return codec, nil
}

// Embed nested codec
//
// Embedded structures are handled this way:
//  1. The separate nested codec is generated for the embedded
//     structure
//  2. Its steps are moved to the outer codec, up the stack.
//     Offsets are corrected.
//
// This function does the work of moving nested codec's steps
// to the outer codec.
func (codec *ippCodec) embed(offset uintptr, nested *ippCodec) {
	for _, step := range nested.steps {
		step.offset += offset
		codec.steps = append(codec.steps, step)
	}
}

// Encode structure into the goipp.Attributes
func (codec *ippCodec) encodeAttrs(in interface{}) (attrs goipp.Attributes) {
	// Check for type mismatch
	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Pointer || v.Elem().Type() != codec.t {
		err := fmt.Errorf("Encoder for %q applied to %q",
			diagTypeName(reflect.PtrTo(codec.t)),
			diagTypeName(reflect.TypeOf(in)))
		panic(err)
	}

	// Now encode, step by step
	p := unsafe.Pointer(v.Pointer())
	attrs = make(goipp.Attributes, 0, len(codec.steps))

	for _, step := range codec.steps {
		attr := goipp.Attribute{Name: step.attrName}
		ptr := unsafe.Pointer(uintptr(p) + step.offset)

		// Check for zero value, if it requires special handling
		doZero := step.conformance == ipAttrOptional ||
			step.zeroTag != goipp.TagZero

		if doZero && step.iszero(ptr) {
			if step.zeroTag != goipp.TagZero {
				attr.Values.Add(step.zeroTag, goipp.Void{})
				attrs.Add(attr)
			}
			continue
		}

		// Normal encode
		values := step.encode(ptr)

		if len(values) != 0 {
			for _, v := range values {
				tag := step.attrTag
				if v.T != goipp.TagZero {
					tag = v.T
				}

				attr.Values.Add(tag, v.V)
			}

			attrs.Add(attr)
		}
	}

	return
}

// Decode structure from the goipp.Attributes
//
// This function wraps (ippCodec) doDecode, adding some common
// error handling and so on
func (codec *ippCodec) decodeAttrs(out interface{},
	attrs goipp.Attributes) error {

	err := codec.doDecode(out, attrs)
	if err != nil {
		err = fmt.Errorf("IPP decode %s: %w",
			diagTypeName(codec.t), err)
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
		// If we see some attribute, the second occurrence
		// Silently ignored. Note, CUPS does the same
		//
		// For details, see discussion here:
		//   https://lore.kernel.org/printing-architecture/84EEF38C-152E-4779-B1E8-578D6BB896E6@msweet.org/
		if _, found := attrByName[attr.Name]; !found {
			attrByName[attr.Name] = attr
		}
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
		if errors.As(err, &conv) && conv.from == goipp.TypeVoid {
			step.setzero(unsafe.Pointer(uintptr(p) + step.offset))
			err = nil
		}

		if err != nil {
			return fmt.Errorf("%q: %w", step.attrName, err)
		}
	}

	return nil
}

// ippCodecMethods contains per-type encode and decode functions
type ippCodecMethods struct {
	// Default IPP tag. Used if there is no per-attribute tag
	// override (encoded as a part of struct tag).
	defaultIppTag goipp.Tag

	// Encode function
	//
	// IPP, at the wire protocol level, doesn't distinguish between
	// scalar and vector values (scalars are encoded as a single-value
	// vectors), so encode always returns slice of values.
	//
	// IPP tag, returned by encode (as a part of goipp.Values slice),
	// if not goipp.TagZero, overrides per-attribute tag setting.
	encode func(p unsafe.Pointer) goipp.Values

	// Decode function
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
	t reflect.Type, encode func(unsafe.Pointer) goipp.Values) goipp.Values {

	slice := reflect.NewAt(t, p).Elem()
	if slice.IsNil() {
		return nil
	}

	vals := make(goipp.Values, slice.Len())
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

// ippCodecMethodsCollection creates ippCodecMethods for encoding
// nested structure or slice of structures as IPP Collection
func ippCodecMethodsCollection(t reflect.Type, slice bool) (
	*ippCodecMethods, error) {

	codec, err := ippCodecGenerate(t)
	if err != nil {
		return nil, err
	}

	m := &ippCodecMethods{
		defaultIppTag: goipp.TagBeginCollection,

		encode: func(p unsafe.Pointer) goipp.Values {
			return ippEncCollection(p, codec)
		},

		decode: func(p unsafe.Pointer, v goipp.Values) error {
			return ippDecCollection(p, v, codec)
		},
	}

	return m, nil
}

// Encode: nested structure as collection
func ippEncCollection(p unsafe.Pointer, codec *ippCodec) goipp.Values {

	ss := reflect.NewAt(codec.t, p).Interface()

	attrs := codec.encodeAttrs(ss)

	return goipp.Values{{goipp.TagBeginCollection, goipp.Collection(attrs)}}
}

// Decode: single nested structure from collection
func ippDecCollection(p unsafe.Pointer, vals goipp.Values,
	codec *ippCodec) error {

	slice, err := ippDecCollectionInternal(p, vals, codec)
	if err != nil {
		return err
	}

	ss := reflect.NewAt(codec.t, p)
	ss.Elem().Set(slice.Index(0))

	return nil
}

// ippDecCollectionInternal decodes collection as a slice
// of decoded structures of type codec.t
//
// It returns slice of decodes structures, represented
// as reflect.Value.
func ippDecCollectionInternal(p unsafe.Pointer, vals goipp.Values,
	codec *ippCodec) (reflect.Value, error) {

	slice := reflect.Zero(reflect.SliceOf(codec.t))
	for i := range vals {
		coll, ok := vals[i].V.(goipp.Collection)
		if !ok {
			err := ippErrConvertMake(vals[i], goipp.TypeCollection)
			return reflect.Value{}, err
		}

		attrs := goipp.Attributes(coll)
		ss := reflect.New(codec.t)

		err := codec.doDecode(ss.Interface(), attrs)
		if err != nil {
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
	reflect.TypeOf((*goipp.IntegerOrRange)(nil)).Elem(): &ippCodecMethods{
		defaultIppTag: goipp.TagZero,
		encode:        ippEncIntegerOrRange,
		decode:        ippDecIntegerOrRange,
	},

	reflect.TypeOf(goipp.Range{}): &ippCodecMethods{
		defaultIppTag: goipp.TagRange,
		encode:        ippEncRange,
		decode:        ippDecRange,
	},

	reflect.TypeOf(goipp.Resolution{}): &ippCodecMethods{
		defaultIppTag: goipp.TagResolution,
		encode:        ippEncResolution,
		decode:        ippDecResolution,
	},

	reflect.TypeOf(goipp.TextWithLang{}): &ippCodecMethods{
		defaultIppTag: goipp.TagTextLang,
		encode:        ippEncTextWithLang,
		decode:        ippDecTextWithLang,
	},

	reflect.TypeOf(goipp.Version(0)): &ippCodecMethods{
		defaultIppTag: goipp.TagKeyword,
		encode:        ippEncVersion,
		decode:        ippDecVersion,
	},

	reflect.TypeOf(time.Time{}): &ippCodecMethods{
		defaultIppTag: goipp.TagDateTime,
		encode:        ippEncDateTime,
		decode:        ippDecDateTime,
	},
}

// Encode: goipp.IntegerOrRange
func ippEncIntegerOrRange(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.IntegerOrRange)(p)
	var tag goipp.Tag
	switch in.(type) {
	case goipp.Integer:
		tag = goipp.TagInteger
	case goipp.Range:
		tag = goipp.TagRange
	}
	out := goipp.Values{{tag, in}}
	return out
}

// Decode: goipp.IntegerOrRange
func ippDecIntegerOrRange(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.IntegerOrRange)
	if !ok {
		err := fmt.Errorf("can't convert %s to Integer or RangeOfInteger",
			vals[0].T)
		return err
	}

	*(*goipp.IntegerOrRange)(p) = res
	return nil
}

// Encode: goipp.Range
func ippEncRange(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.Range)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Range(in)}}
	return out
}

// Decode: goipp.Range
func ippDecRange(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Range)
	if !ok {
		return ippErrConvertMake(vals[0], goipp.TypeRange)
	}

	*(*goipp.Range)(p) = res
	return nil
}

// Encode: goipp.Resolution
func ippEncResolution(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.Resolution)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Resolution(in)}}
	return out
}

// Decode: goipp.Resolution
func ippDecResolution(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Resolution)
	if !ok {
		return ippErrConvertMake(vals[0], goipp.TypeResolution)
	}

	*(*goipp.Resolution)(p) = res
	return nil
}

// Encode: goipp.TextWithLang
func ippEncTextWithLang(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.TextWithLang)(p)
	out := goipp.Values{{goipp.TagZero, goipp.TextWithLang(in)}}
	return out
}

// Decode: goipp.TextWithLang
func ippDecTextWithLang(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.TextWithLang)
	if !ok {
		return ippErrConvertMake(vals[0], goipp.TypeTextWithLang)
	}

	*(*goipp.TextWithLang)(p) = res
	return nil
}

// Encode: goipp.Version
func ippEncVersion(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.Version)(p)
	out := goipp.Values{{goipp.TagZero, goipp.String(in.String())}}
	return out
}

// Decode: goipp.Version
func ippDecVersion(p unsafe.Pointer, vals goipp.Values) error {
	s, ok := vals[0].V.(goipp.String)
	if !ok {
		return ippErrConvertMake(vals[0], goipp.TypeString)
	}

	ver, err := ippDecVersionString(string(s))
	if err != nil {
		return err
	}

	*(*goipp.Version)(p) = ver
	return nil
}

// Decode: IPP version string (X.Y).
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
func ippEncDateTime(p unsafe.Pointer) goipp.Values {
	in := *(*time.Time)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Time{Time: in}}}
	return out
}

// Decode: time.Time
func ippDecDateTime(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Time)
	if !ok {
		return ippErrConvertMake(vals[0], goipp.TypeDateTime)
	}

	*(*time.Time)(p) = res.Time
	return nil

}

// ----- ippCodecMethods for particular reflect.Kind-s -----

// ippCodecMethodsByKind maps reflect.Kind to the particular
// ippCodecMethods structure
var ippCodecMethodsByKind = map[reflect.Kind]*ippCodecMethods{
	reflect.Bool: &ippCodecMethods{
		defaultIppTag: goipp.TagBoolean,
		encode:        ippEncBool,
		decode:        ippDecBool,
	},

	reflect.Int: &ippCodecMethods{
		defaultIppTag: goipp.TagInteger,
		encode:        ippEncInt,
		decode:        ippDecInt,
	},

	reflect.String: &ippCodecMethods{
		defaultIppTag: goipp.TagText,
		encode:        ippEncString,
		decode:        ippDecString,
	},

	reflect.Uint16: &ippCodecMethods{
		defaultIppTag: goipp.TagInteger,
		encode:        ippEncUint16,
		decode:        ippDecUint16,
	},
}

// Encode: bool
func ippEncBool(p unsafe.Pointer) goipp.Values {
	in := *(*bool)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Boolean(in)}}
	return out
}

// Decode: bool
func ippDecBool(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Boolean)
	if !ok {
		return ippErrConvertMake(vals[0], goipp.TypeBoolean)
	}

	*(*bool)(p) = bool(res)
	return nil
}

// Encode: int
func ippEncInt(p unsafe.Pointer) goipp.Values {
	in := *(*int)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Integer(in)}}
	return out
}

// Decode: int
func ippDecInt(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		return ippErrConvertMake(vals[0], goipp.TypeInteger)
	}

	*(*int)(p) = int(res)
	return nil
}

// Encode: string
func ippEncString(p unsafe.Pointer) goipp.Values {
	in := *(*string)(p)
	out := goipp.Values{{goipp.TagZero, goipp.String(in)}}
	return out
}

// Decode: string
func ippDecString(p unsafe.Pointer, vals goipp.Values) error {
	switch res := vals[0].V.(type) {
	case goipp.String:
		*(*string)(p) = string(res)

	case goipp.Binary:
		*(*string)(p) = string(res)

	default:
		return ippErrConvertMake(vals[0], goipp.TypeString)
	}

	return nil
}

// Encode: uint16
func ippEncUint16(p unsafe.Pointer) goipp.Values {
	in := *(*uint16)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Integer(in)}}
	return out
}

// Decode: uint16
func ippDecUint16(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		return ippErrConvertMake(vals[0], goipp.TypeInteger)
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
	fromTag  goipp.Tag
	from, to goipp.Type
}

func ippErrConvertMake(fromval struct {
	T goipp.Tag
	V goipp.Value
}, to goipp.Type) ippErrConvert {
	return ippErrConvert{
		fromTag: fromval.T,
		from:    fromval.V.Type(),
		to:      to,
	}
}

// Convert ippErrConvert to string.
// Implements error interface.
func (err ippErrConvert) Error() string {
	return fmt.Sprintf("can't convert %s to %s", err.fromTag, err.to)
}
