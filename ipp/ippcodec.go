// IPPX - High-level implementation of IPP printing protocol on Go
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
	offset      uintptr            // Field offset within structure
	attrName    string             // IPP attribute name
	attrTag     goipp.Tag          // IPP attribute tag
	zeroTag     goipp.Tag          // How to encode zero value
	slice       bool               // It's a slice of values
	conformance ippAttrConformance // Attribute conformance
	min, max    int                // Range limits for integers

	// Encode/decode functions
	encode  func(p unsafe.Pointer) []goipp.Value
	decode  func(p unsafe.Pointer, v goipp.Values) error
	enctag  func(v goipp.Value) goipp.Tag
	iszero  func(p unsafe.Pointer) bool
	setzero func(p unsafe.Pointer)
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

		tagStr, found := fld.Tag.Lookup("ipp")
		if !found {
			if strings.HasPrefix(string(fld.Tag), "ipp:") {
				err := fmt.Errorf("%s.%s: invalid tag %q",
					diagTypeName(t), fld.Name, fld.Tag)
				return nil, err
			}

			continue
		}

		// Field must be exported
		if !fld.IsExported() {
			err := fmt.Errorf("%s.%s: ipp: tag used with unexported field",
				diagTypeName(t), fld.Name)
			return nil, err
		}

		// Parse ipp: struct tag
		tag, err := ippStructTagParse(tagStr)
		if err != nil {
			err := fmt.Errorf("%s.%s: %w",
				diagTypeName(t), fld.Name, err)
			return nil, err
		}

		// Check for duplicates
		if found := attrNames[tag.name]; found != "" {
			err := fmt.Errorf("%s.%s: attribute %q already used by %s",
				diagTypeName(t), fld.Name, tag.name, found)
			return nil, err
		}
		attrNames[tag.name] = fld.Name

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

		// Generate encoding/decoding step
		step := ippCodecStep{
			offset:      fld.Offset,
			attrName:    tag.name,
			attrTag:     tag.ippTag,
			zeroTag:     tag.zeroTag,
			slice:       slice,
			conformance: tag.conformance,
			min:         tag.min,
			max:         tag.max,

			enctag: methods.enctag,
			iszero: func(p unsafe.Pointer) bool {
				return reflect.NewAt(fldType, p).Elem().IsZero()
			},
			setzero: func(p unsafe.Pointer) {
				reflect.NewAt(fldType, p).Elem().SetZero()
			},
		}

		if step.attrTag == 0 {
			step.attrTag = methods.ippTag
		}

		t1 := step.attrTag.Type()
		t2 := methods.ippTag.Type()

		ok := t1 == t2 ||
			t1 == goipp.TypeBinary && t2 == goipp.TypeString ||
			t2 == goipp.TypeBinary && t1 == goipp.TypeString

		if !ok {
			err := fmt.Errorf("%s.%s: can't represent %s as %s",
				diagTypeName(t), fld.Name, fld.Type, step.attrTag)

			return nil, err
		}

		if slice {
			t := reflect.SliceOf(fldType)

			step.encode = func(p unsafe.Pointer) []goipp.Value {
				return ippEncSlice(p, t, methods.encode)
			}

			step.decode = func(p unsafe.Pointer,
				vals goipp.Values) error {
				return ippDecSlice(p, vals, t, methods.decode)
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

// Embed nesded codec
func (codec *ippCodec) embed(offset uintptr, nested *ippCodec) {
	for _, step := range nested.steps {
		step.offset += offset
		codec.steps = append(codec.steps, step)
	}
}

// Encode structure into the goipp.Attributes
func (codec ippCodec) encode(in interface{}, attrs *goipp.Attributes) {
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
	newattrs := make(goipp.Attributes, 0, len(codec.steps))

	for _, step := range codec.steps {
		attr := goipp.Attribute{Name: step.attrName}
		ptr := unsafe.Pointer(uintptr(p) + step.offset)

		// Check for zero value, if it requires special handling
		doZero := step.conformance == ipAttrOptional ||
			step.zeroTag != goipp.TagZero

		if doZero && step.iszero(ptr) {
			if step.zeroTag != goipp.TagZero {
				attr.Values.Add(step.zeroTag, goipp.Void{})
				newattrs.Add(attr)
			}
			continue
		}

		// Normal encode
		val := step.encode(ptr)

		if len(val) != 0 {
			for _, v := range val {
				tag := step.attrTag
				if tag == goipp.TagZero {
					tag = step.enctag(v)
				}

				attr.Values.Add(tag, v)
			}

			newattrs.Add(attr)
		}
	}

	// Now export newly encoded attributes
	*attrs = append(*attrs, newattrs...)
}

// Decode structure from the goipp.Attributes
//
// This function wraps (ippCodec) doDecode, adding some common
// error handling and so on
func (codec ippCodec) decode(out interface{}, attrs goipp.Attributes) error {
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

// ippStructTag represents parsed ipp: struct tag
type ippStructTag struct {
	name        string             // Attribute name
	ippTag      goipp.Tag          // IPP tag
	zeroTag     goipp.Tag          // How to encode zero value
	conformance ippAttrConformance // Attribute conformance
	min, max    int                // Range limits for integers
}

// ippStructTagToIppTag maps ipp: struct tag keyword to the
// corresponding IPP tag
var ippStructTagToIppTag = map[string]goipp.Tag{
	"boolean":          goipp.TagBoolean,
	"charset":          goipp.TagCharset,
	"datetime":         goipp.TagDateTime,
	"enum":             goipp.TagEnum,
	"integer":          goipp.TagInteger,
	"keyword":          goipp.TagKeyword,
	"mimemediatype":    goipp.TagMimeType,
	"name":             goipp.TagName,
	"namewithlanguage": goipp.TagNameLang,
	"naturallanguage":  goipp.TagLanguage,
	"rangeofinteger":   goipp.TagRange,
	"resolution":       goipp.TagResolution,
	"string":           goipp.TagString,
	"text":             goipp.TagText,
	"textwithlanguage": goipp.TagTextLang,
	"uri":              goipp.TagURI,
	"urischeme":        goipp.TagURIScheme,
}

// ippStructTagParse parses ipp: struct tag into the
// ippStructTag structure
func ippStructTagParse(s string) (*ippStructTag, error) {
	// split struct tag into parts.
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	// first part must be attribute name
	if len(parts) < 1 || parts[0] == "" {
		return nil, errors.New("missed attribute name")
	}

	// Initialize ippStructTag
	stag := &ippStructTag{
		name: parts[0],
		min:  math.MinInt32,
		max:  math.MaxInt32,
	}

	// Parse attribute conformance:
	//   ?name - optional attribute
	//   !name - required attribute
	//   name  - recommended attribute
	switch stag.name[0] {
	case '?':
		stag.conformance = ipAttrOptional
		stag.name = stag.name[1:]

	case '!':
		stag.conformance = ipAttrRequired
		stag.name = stag.name[1:]

	default:
		stag.conformance = ipAttrRecommended
	}

	if stag.name == "" {
		return nil, errors.New("missed attribute name")
	}

	// Parse remaining parameters
	for _, part := range parts[1:] {
		if part == "" {
			continue
		}

		// Apply available parsers until OK or error
		ok, err := stag.parseKeyword(part)
		if !ok && err == nil {
			ok, err = stag.parseMinMax(part)
		}
		if !ok && err == nil {
			ok, err = stag.parseRange(part)
		}

		// Check for result
		if !ok && err == nil {
			err = errors.New("unknown keyword")
		}

		if err != nil {
			err = fmt.Errorf("%q: %s", part, err)
			return nil, err
		}
	}

	return stag, nil
}

// parseKeyword parses keyword parameter of the ipp: struct tag:
//    - IPP tag ("int", "name" etc)
//    - May be, some flags in a future
//
// Return value:
//    - true, nil  - parameter was parsed and applied
//    - false, nil - parameter is not keyword
//    - false, err - invalid parameter
func (stag *ippStructTag) parseKeyword(s string) (bool, error) {
	kw := strings.ToLower(s)
	zeroTag := goipp.TagZero

	if strings.HasSuffix(kw, "|unknown") {
		zeroTag = goipp.TagUnknown
		kw = kw[:len(kw)-8]
	}

	if tag, ok := ippStructTagToIppTag[kw]; ok {
		stag.ippTag = tag
		stag.zeroTag = zeroTag
		return true, nil
	}

	return false, nil
}

// parseRange parses min/max limit constraints:
//   >NNN - min range
//   <NNN - max range
//
// Return value:
//    - true, nil  - parameter was parsed and applied
//    - false, nil - parameter is not keyword
//    - false, err - invalid parameter
func (stag *ippStructTag) parseMinMax(s string) (bool, error) {
	// Limit starts with '<' or '>'
	pfx := s[0]
	if pfx != '<' && pfx != '>' {
		return false, nil
	}

	// Parse limit
	v, err := strconv.ParseInt(s[1:], 10, 64)
	if err != nil {
		err = errors.New("invalid limit")
		return false, err
	}

	// Save limit; check for range
	switch pfx {
	case '>':
		if math.MinInt32-1 <= v && v <= math.MaxInt32-1 {
			stag.min = int(v + 1)
		} else {
			err = errors.New("limit out of range")
		}

	case '<':
		if math.MinInt32+1 <= v && v <= math.MaxInt32+1 {
			stag.max = int(v - 1)
		} else {
			err = errors.New("limit out of range")
		}
	}

	return true, err
}

// parseRange parses range constraints:
//   MIN:MAX
//
// Return value:
//    - true, nil  - parameter was parsed and applied
//    - false, nil - parameter is not keyword
//    - false, err - invalid parameter
func (stag *ippStructTag) parseRange(s string) (bool, error) {
	fields := strings.Split(s, ":")
	if len(fields) != 2 || fields[0] == "" || fields[1] == "" {
		return false, nil
	}

	var min, max int64
	var err error

	// Parse min/max. Don't propagate syntax here, just
	// reject the parameter
	min, err = strconv.ParseInt(fields[0], 10, 64)
	if err == nil {
		if fields[1] == "MAX" {
			max = math.MaxInt32
		} else {
			max, err = strconv.ParseInt(fields[1], 10, 64)
		}
	}

	if err != nil {
		return false, nil
	}

	// Check range
	switch {
	case min < math.MinInt32 || min > math.MaxInt32:
		err = fmt.Errorf("%v out of range", min)
	case max < math.MinInt32 || max > math.MaxInt32:
		err = fmt.Errorf("%v out of range", max)
	case min > max:
		err = fmt.Errorf("range min>max")
	}

	if err == nil {
		stag.min = int(min)
		stag.max = int(max)
	}

	return true, err
}

// ippCodecMethods contains per-reflect.Kind encode and decode
// functions
type ippCodecMethods struct {
	ippTag goipp.Tag
	encode func(p unsafe.Pointer) []goipp.Value
	decode func(p unsafe.Pointer, v goipp.Values) error
	enctag func(v goipp.Value) goipp.Tag
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
func ippCodecMethodsCollection(t reflect.Type, slice bool) (
	*ippCodecMethods, error) {

	codec, err := ippCodecGenerate(t)
	if err != nil {
		return nil, err
	}

	m := &ippCodecMethods{
		ippTag: goipp.TagBeginCollection,

		encode: func(p unsafe.Pointer) []goipp.Value {
			return ippEncCollection(p, codec)
		},

		decode: func(p unsafe.Pointer, v goipp.Values) error {
			return ippDecCollection(p, v, codec)
		},
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
		ippTag: goipp.TagZero,
		encode: ippEncIntegerOrRange,
		decode: ippDecIntegerOrRange,
		enctag: ippTagIntegerOrRange,
	},

	reflect.TypeOf(goipp.Range{}): &ippCodecMethods{
		ippTag: goipp.TagRange,
		encode: ippEncRange,
		decode: ippDecRange,
	},

	reflect.TypeOf(goipp.Resolution{}): &ippCodecMethods{
		ippTag: goipp.TagResolution,
		encode: ippEncResolution,
		decode: ippDecResolution,
	},

	reflect.TypeOf(goipp.TextWithLang{}): &ippCodecMethods{
		ippTag: goipp.TagTextLang,
		encode: ippEncTextWithLang,
		decode: ippDecTextWithLang,
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

// Encode: goipp.IntegerOrRange
func ippEncIntegerOrRange(p unsafe.Pointer) []goipp.Value {
	in := *(*goipp.IntegerOrRange)(p)
	out := []goipp.Value{in}
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

// Encode tag: goipp.IntegerOrRange
func ippTagIntegerOrRange(v goipp.Value) (tag goipp.Tag) {
	switch v.(type) {
	case goipp.Integer:
		tag = goipp.TagInteger
	case goipp.Range:
		tag = goipp.TagRange
	}
	return
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
		return ippErrConvertMake(vals[0], goipp.TypeRange)
	}

	*(*goipp.Range)(p) = res
	return nil
}

// Encode: goipp.Resolution
func ippEncResolution(p unsafe.Pointer) []goipp.Value {
	in := *(*goipp.Resolution)(p)
	out := []goipp.Value{goipp.Resolution(in)}
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
func ippEncTextWithLang(p unsafe.Pointer) []goipp.Value {
	in := *(*goipp.TextWithLang)(p)
	out := []goipp.Value{goipp.TextWithLang(in)}
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
func ippEncVersion(p unsafe.Pointer) []goipp.Value {
	in := *(*goipp.Version)(p)
	out := []goipp.Value{goipp.String(in.String())}
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
func ippEncDateTime(p unsafe.Pointer) []goipp.Value {
	in := *(*time.Time)(p)
	out := []goipp.Value{goipp.Time{Time: in}}
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
		return ippErrConvertMake(vals[0], goipp.TypeBoolean)
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
		return ippErrConvertMake(vals[0], goipp.TypeInteger)
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
func ippEncUint16(p unsafe.Pointer) []goipp.Value {
	in := *(*uint16)(p)
	out := []goipp.Value{goipp.Integer(in)}
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
