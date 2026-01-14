// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Object decoder

package ipp

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/goipp"
)

// Decoder decodes [Object]s from the [goipp.Attributes].
type Decoder struct {
	opt      DecoderOptions // Decoder options
	codec    *ippCodec      // Codec for the object
	typename string         // Type name being decoding
	path     []any          // Path to current attr (string/int indices)
	errors   []error        // Decode errors
}

// DecoderOptions represent options used when [Object] is being
// decoded from the [goipp.Attributes].
type DecoderOptions struct {
	// KeepTrying, if set, instructs decoder do not stop on
	// value decoding errors, but just skip problematic value
	// and continue.
	KeepTrying bool
}

// NewDecoder creates the new [Decoder].
// Of opt is nil, [DefaultDecoderOptions] will be used
func NewDecoder(opt *DecoderOptions) *Decoder {
	dec := decoderAlloc()

	dec.opt = DefaultDecoderOptions
	if opt != nil {
		dec.opt = *opt
	}

	return dec
}

// begin initializes Decoder before decoding the object.
func (dec *Decoder) begin(obj Object) {
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
func (dec *Decoder) Decode(obj Object, attrs goipp.Attributes) error {
	dec.begin(obj)

	err := dec.codec.decodeAttrs(dec, obj, attrs)
	if err == nil {
		obj.RawAttrs().save(attrs, dec.errors)
	}

	return err
}

// DecodeSingle decodes (updates) a single attribute of the Object.
func (dec *Decoder) DecodeSingle(obj Object, attr goipp.Attribute) error {
	dec.begin(obj)
	return dec.codec.decodeAttrs(dec, obj, goipp.Attributes{attr})
}

// Errors returns slice of decoder errors.
func (dec *Decoder) Errors() []error {
	return dec.errors
}

// Decode: goipp.IntegerOrRange
func (dec *Decoder) decIntegerOrRange(p unsafe.Pointer, vals goipp.Values) error {
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
func (dec *Decoder) decRange(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Range)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeRange)
	}

	*(*goipp.Range)(p) = res
	return nil
}

// Decode: goipp.Resolution
func (dec *Decoder) decResolution(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Resolution)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeResolution)
	}

	*(*goipp.Resolution)(p) = res
	return nil
}

// Decode: goipp.TextWithLang
func (dec *Decoder) decTextWithLang(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.TextWithLang)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeTextWithLang)
	}

	*(*goipp.TextWithLang)(p) = res
	return nil
}

// Decode: goipp.Version
func (dec *Decoder) decVersion(p unsafe.Pointer, vals goipp.Values) error {
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
func (dec *Decoder) decVersionString(s string) (goipp.Version, error) {
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
func (dec *Decoder) decDateTime(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Time)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeDateTime)
	}

	*(*time.Time)(p) = res.Time
	return nil

}

// Decode: bool
func (dec *Decoder) decBool(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Boolean)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeBoolean)
	}

	*(*bool)(p) = bool(res)
	return nil
}

// Decode: int
func (dec *Decoder) decInt(p unsafe.Pointer, vals goipp.Values) error {
	res, ok := vals[0].V.(goipp.Integer)
	if !ok {
		return dec.errConvert(vals[0], goipp.TypeInteger)
	}

	*(*int)(p) = int(res)
	return nil
}

// Decode: string
func (dec *Decoder) decString(p unsafe.Pointer, vals goipp.Values) error {
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
func (dec *Decoder) decUint16(p unsafe.Pointer, vals goipp.Values) error {
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
func (dec *Decoder) decCollection(
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
func (dec *Decoder) decCollectionInternal(
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
func (dec *Decoder) decSlice(
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
func (dec *Decoder) decPtr(
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
func (dec *Decoder) pathEnter() {
	dec.path = append(dec.path, nil)
}

// pathLeave shrinks the path to reflect leaving the current nesting level.
func (dec *Decoder) pathLeave() {
	dec.path = dec.path[:len(dec.path)-1]
}

// Set attribute name or index at the current path level
func (dec *Decoder) pathSet(p any) {
	dec.path[len(dec.path)-1] = p
}

// pathString returns the current path as a string.
func (dec *Decoder) pathString() string {
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
func (dec *Decoder) errWrap(err error) error {
	err = fmt.Errorf("IPP decode %s: %q: %w",
		dec.typename, dec.pathString(), err)
	return err
}

// errConvert returns type conversion error
func (dec *Decoder) errConvert(fromval struct {
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

// errPush adds the error into the Decoder.errors slice.
func (dec *Decoder) errPush(err error) {
	dec.errors = append(dec.errors, err)
}

// Free returns Decoder into the [sync.Pool] of free decoders.
// Calling this function is not required, but recommended when
// decoder is not used anymore as an optimization.
func (dec *Decoder) Free() {
	decoderFree(dec)
}

var decoderPool = sync.Pool{
	New: func() any { return new(Decoder) },
}

// decoderAlloc allocates a new Decoder from the decoderPool
func decoderAlloc() *Decoder {
	return decoderPool.Get().(*Decoder)
}

// decoderFree returns Decoder into the decoderFree
func decoderFree(dec *Decoder) {
	// Don't keep this memory in the Pool, return it to GC
	dec.codec = nil
	dec.path = nil
	dec.errors = nil

	// Put Decoder to the Pool
	decoderPool.Put(dec)
}
