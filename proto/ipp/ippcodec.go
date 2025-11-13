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

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/goipp"
)

// ippEncoder maintains context for encoding Object into the
// goipp.Attributes.
type ippEncoder struct {
}

// Encode encodes Object into goipp.Attributes.
//
// The obj parameter must be pointer to structure that implements
// the Object interface. Its codec will be generated on demand.
//
// This function will panic, if codec cannot be generated.
func (enc *ippEncoder) Encode(obj Object) goipp.Attributes {
	codec := ippCodecGet(obj)
	return codec.encodeAttrs(enc, obj)
}

// Encode: goipp.IntegerOrRange
func (enc *ippEncoder) encIntegerOrRange(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.IntegerOrRange)(p)
	var tag goipp.Tag
	switch in.(type) {
	case goipp.Integer:
		tag = goipp.TagInteger
	case goipp.Range:
		tag = goipp.TagRange
	default:
		return nil
	}
	out := goipp.Values{{tag, in}}
	return out
}

// Encode: goipp.Range
func (enc *ippEncoder) encRange(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.Range)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Range(in)}}
	return out
}

// Encode: goipp.Resolution
func (enc *ippEncoder) encResolution(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.Resolution)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Resolution(in)}}
	return out
}

// Encode: goipp.TextWithLang
func (enc *ippEncoder) encTextWithLang(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.TextWithLang)(p)
	out := goipp.Values{{goipp.TagZero, goipp.TextWithLang(in)}}
	return out
}

// Encode: goipp.Version
func (enc *ippEncoder) encVersion(p unsafe.Pointer) goipp.Values {
	in := *(*goipp.Version)(p)
	out := goipp.Values{{goipp.TagZero, goipp.String(in.String())}}
	return out
}

// Encode: time.Time
func (enc *ippEncoder) encDateTime(p unsafe.Pointer) goipp.Values {
	in := *(*time.Time)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Time{Time: in}}}
	return out
}

// Encode: bool
func (enc *ippEncoder) encBool(p unsafe.Pointer) goipp.Values {
	in := *(*bool)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Boolean(in)}}
	return out
}

// Encode: int
func (enc *ippEncoder) encInt(p unsafe.Pointer) goipp.Values {
	in := *(*int)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Integer(in)}}
	return out
}

// Encode: string
func (enc *ippEncoder) encString(p unsafe.Pointer) goipp.Values {
	in := *(*string)(p)
	out := goipp.Values{{goipp.TagZero, goipp.String(in)}}
	return out
}

// Encode: uint16
func (enc *ippEncoder) encUint16(p unsafe.Pointer) goipp.Values {
	in := *(*uint16)(p)
	out := goipp.Values{{goipp.TagZero, goipp.Integer(in)}}
	return out
}

// Encode: nested structure as collection
func (enc *ippEncoder) encCollection(p unsafe.Pointer, codec *ippCodec) goipp.Values {

	ss := reflect.NewAt(codec.t, p).Interface()

	attrs := codec.encodeAttrs(enc, ss)

	return goipp.Values{{goipp.TagBeginCollection, goipp.Collection(attrs)}}
}

// encSlice encodes slice of values
//
// p is pointer to slice, t is slice type (so for []string t will be
// []string while p's value will be of type *[]string, represented
// as unsafe.Pointer)
//
// encode is the single-value encoder (i.e., (*ippEncoder).encString for slice
// of strings)
func (enc *ippEncoder) encSlice(
	p unsafe.Pointer, t reflect.Type, encode encodeFunc) goipp.Values {

	slice := reflect.NewAt(t, p).Elem()
	if slice.IsNil() {
		return nil
	}

	vals := make(goipp.Values, slice.Len())
	for i := range vals {
		p2 := unsafe.Pointer(slice.Index(i).Addr().Pointer())
		v := encode(enc, p2)
		vals[i] = v[0]
	}

	return vals
}

// encPtr encodes pointer to value
//
// p is pointer to pointer, t is the type of value the pointer
// points to, (so for *string t will be string while p's value will
// be of type *string, represented as unsafe.Pointer)
//
// encode is the single-value encoder (i.e., (*ippEncoder).encString for slice
// of strings)
func (enc *ippEncoder) encPtr(
	p unsafe.Pointer, t reflect.Type, encode encodeFunc) goipp.Values {

	ptr := reflect.NewAt(t, p).Elem()
	if ptr.IsNil() {
		return nil
	}

	return encode(enc, unsafe.Pointer(ptr.Pointer()))
}

// ippDecoder maintains context for decocing Object from the
// goipp.Attributes.
type ippDecoder struct {
	opt      DecodeOptions // Decoder options
	codec    *ippCodec     // Codec for the object
	typename string        // Type name being decoding
	path     []any         // Path to current attr (string/int indices)
	errors   []error       // Decode errors
}

// begin initializes ippDecoder before decoding the object.
func (dec *ippDecoder) begin(obj Object) {
	dec.codec = ippCodecGet(obj)
	dec.typename = diagTypeName(dec.codec.t)
	dec.errors = nil

	// Reserve path slots.
	// 8 should be enough in most cases to avoid re-allocation.
	dec.path = make([]any, 0, 8)
}

// Decode decodes Object from the goipp.Attributes.
//
// The obj parameter must be pointer to structure that implements
// the Object interface. Its codec will be generated on demand.
//
// This function will panic, if codec cannot be generated.
func (dec *ippDecoder) Decode(obj Object, attrs goipp.Attributes) error {
	dec.begin(obj)

	err := dec.codec.decodeAttrs(dec, obj, attrs)
	if err == nil {
		obj.RawAttrs().save(attrs, dec.errors)
	}

	return err
}

// DecodeSingle decodes (updates) a single attribute of the Object.
func (dec *ippDecoder) DecodeSingle(obj Object, attr goipp.Attribute) error {
	dec.begin(obj)
	return dec.codec.decodeAttrs(dec, obj, goipp.Attributes{attr})
}

// Decode: goipp.IntegerOrRange
func (dec *ippDecoder) decIntegerOrRange(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.IntegerOrRange)
	if !ok {
		err := fmt.Errorf("can't convert %s to Integer or RangeOfInteger",
			vals[0].T)
		return dec.errWrap(err)
	}

	*(*goipp.IntegerOrRange)(p) = res
	return nil
}

// Decode: goipp.Range
func (dec *ippDecoder) decRange(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Range)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeRange)
	}

	*(*goipp.Range)(p) = res
	return nil
}

// Decode: goipp.Resolution
func (dec *ippDecoder) decResolution(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Resolution)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeResolution)
	}

	*(*goipp.Resolution)(p) = res
	return nil
}

// Decode: goipp.TextWithLang
func (dec *ippDecoder) decTextWithLang(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.TextWithLang)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeTextWithLang)
	}

	*(*goipp.TextWithLang)(p) = res
	return nil
}

// Decode: goipp.Version
func (dec *ippDecoder) decVersion(p unsafe.Pointer, vals goipp.Values) error {
	s, ok := vals[0].V.(goipp.String)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeString)
	}

	ver, err := dec.decVersionString(string(s))
	if err != nil {
		return err
	}

	*(*goipp.Version)(p) = ver
	return nil
}

// Decode: IPP version string (X.Y).
func (dec *ippDecoder) decVersionString(s string) (goipp.Version, error) {
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
	return 0, dec.errWrap(fmt.Errorf("%q: invalid version string", s))
}

// Decode: time.Time
func (dec *ippDecoder) decDateTime(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Time)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeDateTime)
	}

	*(*time.Time)(p) = res.Time
	return nil

}

// Decode: bool
func (dec *ippDecoder) decBool(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Boolean)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeBoolean)
	}

	*(*bool)(p) = bool(res)
	return nil
}

// Decode: int
func (dec *ippDecoder) decInt(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeInteger)
	}

	*(*int)(p) = int(res)
	return nil
}

// Decode: string
func (dec *ippDecoder) decString(p unsafe.Pointer, vals goipp.Values) error {
	switch res := vals[0].V.(type) {
	case goipp.String:
		*(*string)(p) = string(res)

	case goipp.Binary:
		*(*string)(p) = string(res)

	default:
		return dec.errConvert(vals[0], goipp.TypeString)
	}

	return nil
}

// Decode: uint16
func (dec *ippDecoder) decUint16(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeInteger)
	}

	if res < 0 || res > math.MaxUint16 {
		return dec.errWrap(fmt.Errorf("Value %d out of range", res))
	}

	*(*uint16)(p) = uint16(res)
	return nil
}

// Decode: single nested structure from collection
func (dec *ippDecoder) decCollection(
	p unsafe.Pointer, vals goipp.Values,
	codec *ippCodec) error {

	slice, err := dec.decCollectionInternal(p, vals, codec)
	if err != nil {
		return err
	}

	ss := reflect.NewAt(codec.t, p)
	ss.Elem().Set(slice.Index(0))

	return nil
}

// decCollectionInternal decodes collection as a slice
// of decoded structures of type codec.t
//
// It returns slice of decodes structures, represented
// as reflect.Value.
func (dec *ippDecoder) decCollectionInternal(
	p unsafe.Pointer, vals goipp.Values,
	codec *ippCodec) (reflect.Value, error) {

	slice := reflect.Zero(reflect.SliceOf(codec.t))
	for i := range vals {
		coll, ok := vals[i].V.(goipp.Collection)
		if !ok {
			err := dec.errConvert(vals[i], goipp.TypeCollection)
			return reflect.Value{}, err
		}

		attrs := goipp.Attributes(coll)
		ss := reflect.New(codec.t)

		err := codec.decodeAttrs(dec, ss.Interface(), attrs)
		if err != nil {
			return reflect.Value{}, err
		}

		slice = reflect.Append(slice, ss.Elem())
	}

	return slice, nil
}

// ippDecSlice decodes slice of values
func (dec *ippDecoder) decSlice(
	p unsafe.Pointer, vals goipp.Values, t reflect.Type, decode decodeFunc) error {

	// Update current path
	dec.pathEnter()
	defer dec.pathLeave()

	// Setup things
	slice := reflect.MakeSlice(t, 0, len(vals))
	tmp := reflect.New(t.Elem())
	zero := reflect.Zero(t.Elem())

	// Handle OOB
	//
	// See RFC8010, 3.5.2. for details
	// https://datatracker.ietf.org/doc/html/rfc8010#section-3.5.2
	if vals[0].T.Type() == goipp.TypeVoid {
		reflect.NewAt(t, p).Elem().Set(reflect.Zero(t))
		return nil
	}

	// Now decode, step by step
	for i := range vals {
		dec.pathSet(i)

		tmp.Elem().Set(zero)
		err := decode(dec, unsafe.Pointer(tmp.Pointer()), vals[i:i+1])
		if err != nil {
			if dec.opt.KeepTrying {
				// Skip the value
				dec.errPush(err)
				continue
			}

			return err
		}

		slice = reflect.Append(slice, tmp.Elem())
	}

	out := reflect.NewAt(t, p)
	out.Elem().Set(slice)

	return nil
}

// ippDecSlice decodes pointer to value.
func (dec *ippDecoder) decPtr(
	p unsafe.Pointer, vals goipp.Values, t reflect.Type, decode decodeFunc) error {

	assert.Must(len(vals) > 0)

	// Handle OOB
	//
	// See RFC8010, 3.5.2. for details
	// https://datatracker.ietf.org/doc/html/rfc8010#section-3.5.2
	if vals[0].T.Type() == goipp.TypeVoid {
		reflect.NewAt(t, p).Elem().Set(reflect.Zero(t))
		return nil
	}

	// Decode the value
	tmp := reflect.New(t.Elem())
	err := decode(dec, unsafe.Pointer(tmp.Pointer()), vals[:1])

	if err != nil {
		if dec.opt.KeepTrying {
			dec.errPush(err)
			reflect.NewAt(t, p).Elem().Set(reflect.Zero(t))
			return nil
		}

		return err
	}

	ptr := reflect.NewAt(t, p).Elem()
	ptr.Set(tmp)

	return nil
}

// pathEnter advises the path to the next nesting level.
func (dec *ippDecoder) pathEnter() {
	dec.path = append(dec.path, nil)
}

// pathLeave shrinks the path to reflect leaving the current nesting level.
func (dec *ippDecoder) pathLeave() {
	dec.path = dec.path[:len(dec.path)-1]
}

// Set attribute name or index at the current path level
func (dec *ippDecoder) pathSet(p any) {
	dec.path[len(dec.path)-1] = p
}

// pathString returns the current path as a string.
func (dec *ippDecoder) pathString() string {
	// Estimate buffer size
	sz := 0
	for _, p := range dec.path {
		switch p := p.(type) {
		case string:
			sz += len(p) + 1
		case int:
			sz += 4 // "[NN]"
		}
	}

	// Format the path
	path := make([]byte, 0, sz)
	for _, p := range dec.path {
		switch p := p.(type) {
		case string:
			path = append(path, '/')
			path = append(path, ([]byte)(p)...)
		case int:
			path = append(path, '[')
			path = append(path, ([]byte)(strconv.Itoa(p))...)
			path = append(path, ']')
		}
	}

	// Strip leading '/' and convert path to the string
	return string(path[1:])
}

// errWrap wraps the decode error so it includes common information,
// like path to the problematic attribute.
func (dec *ippDecoder) errWrap(err error) error {
	err = fmt.Errorf("IPP decode %s: %q: %w",
		dec.typename, dec.pathString(), err)
	return err
}

// errConvert returns type conversion error
func (dec *ippDecoder) errConvert(fromval struct {
	T goipp.Tag
	V goipp.Value
}, to goipp.Type) error {

	var err error = ippErrConvert{
		fromTag: fromval.T,
		from:    fromval.V.Type(),
		to:      to,
	}

	return dec.errWrap(err)
}

// errPush adds the error into the ippDecoder.errors slice.
func (dec *ippDecoder) errPush(err error) {
	dec.errors = append(dec.errors, err)
}

// ippKnownAttrs returns information about Object's known attributes
//
// This function will panic, if codec cannot be generated.
func ippKnownAttrs(obj Object) []AttrInfo {
	codec := ippCodecGet(obj)
	return codec.knownAttrs
}

// ippKnownAttrsType returns information about known attributes of
// structure, defined by its type.
//
// t must be pointer to structure.
func ippKnownAttrsType(t reflect.Type) []AttrInfo {
	codec := ippCodecGetType(t)
	return codec.knownAttrs
}

// ippCodec represents actions required to encode/decode structures
// of the particular type. Codecs are generated at initialization and
// then reused, to minimize performance overhead associated with
// reflection
type ippCodec struct {
	t           reflect.Type             // Type of structure
	steps       []ippCodecStep           // Encoding/decoding steps
	stepsByName map[string]*ippCodecStep // Steps indexed by attribute name
	knownAttrs  []AttrInfo               // Known attributes
}

// encodeFunc encodes attribute values into the IPP representation
type encodeFunc func(enc *ippEncoder, p unsafe.Pointer) goipp.Values

// decodeFunc decodes attribute values from the IPP representation
type decodeFunc func(dec *ippDecoder, p unsafe.Pointer, v goipp.Values) error

// ippCodecStep represents a single encoding/decoding step for the
// ippCodec
type ippCodecStep struct {
	offset   uintptr   // Field offset within structure
	attrName string    // IPP attribute name
	attrTag  goipp.Tag // IPP attribute tag
	zeroTag  goipp.Tag // How to encode zero value
	min, max int       // Range limits for integers
	isSlice  bool      // Slice output, multiple values expected

	// Encode/decode functions
	encode  func(enc *ippEncoder, p unsafe.Pointer) goipp.Values
	decode  func(dec *ippDecoder, p unsafe.Pointer, v goipp.Values) error
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

	return ippCodecGetType(t)
}

// ippCodecGetType returns coded for structure, defined by its type.
//
// t must be pointer to structure.
func ippCodecGetType(t reflect.Type) *ippCodec {
	// We need reflect.Type of the structure, not pointer
	assert.Must(t.Kind() == reflect.Pointer)
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

	// Type must either implement the Object interface or
	// contain at least 1 IPP field.
	if err == nil && len(codec.steps) == 0 {
		var obj Object
		objtype := reflect.TypeOf(&obj).Elem()

		if !reflect.PointerTo(t).AssignableTo(objtype) {
			err = fmt.Errorf("%s: contains no IPP fields",
				diagTypeName(t))
		}
	}

	if err != nil {
		return nil, err
	}

	// Build stepsByName and knownAttrs
	codec.stepsByName = make(map[string]*ippCodecStep, len(codec.steps))
	codec.knownAttrs = make([]AttrInfo, len(codec.steps))

	for i := range codec.steps {
		codec.stepsByName[codec.steps[i].attrName] = &codec.steps[i]
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
		// Handle special cases: pointers and slices.
		fldType := fld.Type
		fldKind := fldType.Kind()

		isOptional := false
		isSlice := false

		switch fldKind {
		case reflect.Pointer:
			isOptional = true
			fldType = fldType.Elem()
			fldKind = fldType.Kind()

		case reflect.Slice:
			isSlice = true
			fldType = fldType.Elem()
			fldKind = fldType.Kind()
		}

		// Now fldType points to the actual type to be encoded and
		// decoded.  Obtain its ippCodecMethods.
		methods := ippCodecMethodsByType[fldType]
		if methods == nil {
			methods = ippCodecMethodsByKind[fldKind]
		}
		if methods == nil && fldKind == reflect.Struct {
			methods, err = ippCodecMethodsCollection(fldType)
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
		zero := reflect.Zero(fldType)
		step := ippCodecStep{
			offset:   fld.Offset,
			attrName: tag.name,
			attrTag:  tag.ippTag,
			zeroTag:  tag.zeroTag,
			min:      tag.min,
			max:      tag.max,
			isSlice:  isSlice,

			encode: methods.encode,
			decode: methods.decode,

			setzero: func(p unsafe.Pointer) {
				reflect.NewAt(fldType, p).Elem().Set(zero)
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
		if isSlice {
			t := reflect.SliceOf(fldType)

			encode := step.encode
			decode := step.decode

			step.encode = func(enc *ippEncoder, p unsafe.Pointer) goipp.Values {
				return enc.encSlice(p, t, encode)
			}

			step.decode = func(dec *ippDecoder, p unsafe.Pointer,
				vals goipp.Values) error {
				return dec.decSlice(p, vals, t, decode)
			}
		}

		// Generate optional.Val[T] wrapper where appropriate.
		if isOptional {
			t := reflect.PointerTo(fldType)

			encode := step.encode
			decode := step.decode

			step.encode = func(enc *ippEncoder, p unsafe.Pointer) goipp.Values {
				return enc.encPtr(p, t, encode)
			}

			step.decode = func(dec *ippDecoder,
				p unsafe.Pointer, vals goipp.Values) error {
				return dec.decPtr(p, vals, t, decode)
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
func (codec *ippCodec) encodeAttrs(enc *ippEncoder,
	in interface{}) (attrs goipp.Attributes) {

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
		// Encode attribute values
		ptr := unsafe.Pointer(uintptr(p) + step.offset)

		values := step.encode(enc, ptr)

		if values == nil && step.zeroTag != goipp.TagZero {
			values = goipp.Values{{step.zeroTag, goipp.Void{}}}
		}

		// If we have value, encode the whole attribute
		if len(values) != 0 {
			attr := goipp.Attribute{Name: step.attrName}

			for _, v := range values {
				// Use default step.attrTag, if value tag
				// is not set explicitly.
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
func (codec *ippCodec) decodeAttrs(dec *ippDecoder,
	out interface{}, attrs goipp.Attributes) error {

	// Update path
	dec.pathEnter()
	defer dec.pathLeave()

	// Check for type mismatch
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Pointer || v.Elem().Type() != codec.t {
		err := fmt.Errorf("Decoder for %q applied to %q",
			reflect.PtrTo(codec.t), reflect.TypeOf(out))
		panic(err)
	}

	p := unsafe.Pointer(v.Pointer())

	// Now decode, attribute by attribute
	attersSeen := generic.NewSet[string]()
	for _, attr := range attrs {
		if !attersSeen.TestAndAdd(attr.Name) {
			// If we see the same attribute, the second occurrence
			// is silently ignored. Note, CUPS does the same
			//
			// For details, see discussion here:
			//   https://lore.kernel.org/printing-architecture/84EEF38C-152E-4779-B1E8-578D6BB896E6@msweet.org/
			continue
		}

		step, found := codec.stepsByName[attr.Name]
		if found {
			err := codec.doDecodeStep(dec, p, step, attr)
			if err != nil {
				dec.errPush(err)
				return err
			}
		}
	}

	return nil
}

// doDecodeStep decodes a single attribute.
// p is the unsafe.Pointer to the outer structure.
func (codec ippCodec) doDecodeStep(dec *ippDecoder,
	p unsafe.Pointer, step *ippCodecStep, attr goipp.Attribute) error {

	// Update path
	dec.pathSet(step.attrName)

	// At least one attribute value must be present.
	//
	// IPP protocol doesn't allow attributes without values
	// and github.com/OpenPrinting/goipp will never returns
	// them.
	//
	// So if Values slice is empty, this is artificial construct.
	// Reject it in this case.
	if len(attr.Values) == 0 {
		err := dec.errWrap(errors.New("at least 1 value required"))
		dec.errPush(err)
		return err
	}

	// If not slice and we have more that 1 value, generate a
	// warning and continue.
	if !step.isSlice && len(attr.Values) > 1 {
		err := fmt.Errorf("1 value expected, %d present", len(attr.Values))
		dec.errPush(dec.errWrap(err))
	}

	// Call decoder
	err := step.decode(dec,
		unsafe.Pointer(uintptr(p)+step.offset), attr.Values)

	if err != nil {
		if dec.opt.KeepTrying {
			dec.errPush(err)
			step.setzero(unsafe.Pointer(uintptr(p) + step.offset))
			return nil
		}
	}

	return err
}

// ippCodecMethods contains per-type encode and decode functions
type ippCodecMethods struct {
	// Default IPP tag. Used if there is no per-attribute tag
	// override (encoded as a part of struct tag).
	defaultIppTag goipp.Tag

	// encode function is called to encode the Go value
	// into the IPP value.
	//
	// IPP, at the wire protocol level, doesn't distinguish between
	// scalar and vector values (scalars are encoded as a single-value
	// vectors), so encode always returns slice of values.
	//
	// IPP tag, returned by encode (as a part of goipp.Values slice),
	// if not goipp.TagZero, overrides per-attribute tag setting.
	encode encodeFunc

	// decode function is called to decode IPP value into
	// the Go value.
	decode decodeFunc
}

// ippCodecMethodsCollection creates ippCodecMethods for encoding
// nested structure or slice of structures as IPP Collection
func ippCodecMethodsCollection(t reflect.Type) (
	*ippCodecMethods, error) {

	codec, err := ippCodecGenerate(t)
	if err != nil {
		return nil, err
	}

	m := &ippCodecMethods{
		defaultIppTag: goipp.TagBeginCollection,

		encode: func(enc *ippEncoder, p unsafe.Pointer) goipp.Values {
			return enc.encCollection(p, codec)
		},

		decode: func(dec *ippDecoder, p unsafe.Pointer, v goipp.Values) error {
			return dec.decCollection(p, v, codec)
		},
	}

	return m, nil
}

// ippCodecMethodsByType maps reflect.Type to the particular
// ippCodecMethods structure
var ippCodecMethodsByType = map[reflect.Type]*ippCodecMethods{
	reflect.TypeOf((*goipp.IntegerOrRange)(nil)).Elem(): &ippCodecMethods{
		defaultIppTag: goipp.TagZero,
		encode:        (*ippEncoder).encIntegerOrRange,
		decode:        (*ippDecoder).decIntegerOrRange,
	},

	reflect.TypeOf(goipp.Range{}): &ippCodecMethods{
		defaultIppTag: goipp.TagRange,
		encode:        (*ippEncoder).encRange,
		decode:        (*ippDecoder).decRange,
	},

	reflect.TypeOf(goipp.Resolution{}): &ippCodecMethods{
		defaultIppTag: goipp.TagResolution,
		encode:        (*ippEncoder).encResolution,
		decode:        (*ippDecoder).decResolution,
	},

	reflect.TypeOf(goipp.TextWithLang{}): &ippCodecMethods{
		defaultIppTag: goipp.TagTextLang,
		encode:        (*ippEncoder).encTextWithLang,
		decode:        (*ippDecoder).decTextWithLang,
	},

	reflect.TypeOf(goipp.Version(0)): &ippCodecMethods{
		defaultIppTag: goipp.TagKeyword,
		encode:        (*ippEncoder).encVersion,
		decode:        (*ippDecoder).decVersion,
	},

	reflect.TypeOf(time.Time{}): &ippCodecMethods{
		defaultIppTag: goipp.TagDateTime,
		encode:        (*ippEncoder).encDateTime,
		decode:        (*ippDecoder).decDateTime,
	},
}

// ippCodecMethodsByKind maps reflect.Kind to the particular
// ippCodecMethods structure
var ippCodecMethodsByKind = map[reflect.Kind]*ippCodecMethods{
	reflect.Bool: &ippCodecMethods{
		defaultIppTag: goipp.TagBoolean,
		encode:        (*ippEncoder).encBool,
		decode:        (*ippDecoder).decBool,
	},

	reflect.Int: &ippCodecMethods{
		defaultIppTag: goipp.TagInteger,
		encode:        (*ippEncoder).encInt,
		decode:        (*ippDecoder).decInt,
	},

	reflect.String: &ippCodecMethods{
		defaultIppTag: goipp.TagText,
		encode:        (*ippEncoder).encString,
		decode:        (*ippDecoder).decString,
	},

	reflect.Uint16: &ippCodecMethods{
		defaultIppTag: goipp.TagInteger,
		encode:        (*ippEncoder).encUint16,
		decode:        (*ippDecoder).decUint16,
	},
}

// ippErrConvert represents conversion error between expected
// and present goipp.Type
type ippErrConvert struct {
	fromTag  goipp.Tag  // Attribute tag
	from, to goipp.Type // Source and destination types
}

// Convert ippErrConvert to string.
// Implements error interface.
func (err ippErrConvert) Error() string {
	return fmt.Sprintf("can't convert %s to %s", err.fromTag, err.to)
}
