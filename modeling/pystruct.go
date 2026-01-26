// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversion between Go and Python protocol structures

package modeling

import (
	"fmt"
	"reflect"

	"github.com/OpenPrinting/go-mfp/cpython"
	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/util/uuid"
)

// pyExportStruct converts the protocol object, represented as Go
// structure or pointer to structure, into the Python dictionary.
//
// s MUST be struct or pointer to struct.
func (model *Model) pyExportStruct(s any) *cpython.Object {
	// Create output cpython.Object (the empty dict).
	dict := model.py.NewObject(map[any]any(nil))

	// Normalize input parameter and obtain the reflect.Value for it.
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Pointer && v.Elem().Kind() == reflect.Struct {
		v = v.Elem()
	}
	assert.Must((v.Kind() == reflect.Struct))

	// Roll over all struct fields
	flgs := reflect.VisibleFields(v.Type())
	for _, fld := range flgs {
		// Skip non-exposed fields
		if !fld.IsExported() {
			continue
		}

		// Obtain and normalize field value
		f := v.FieldByName(fld.Name)
		switch f.Kind() {
		case reflect.Slice:
			// Skip nil slices
			if f.IsNil() {
				continue
			}
		case reflect.Pointer:
			// Skip nil pointers. Dereference others.
			if f.IsNil() {
				continue
			}
			f = f.Elem()
		}

		// Convert into the Python Object and add to the dict,
		item := model.pyExportValue(f)
		err := dict.Set(keywordNormalize(fld.Name), item)

		if err != nil {
			return model.py.NewError(err)
		}
	}

	return dict
}

// pyExportSlice exports slice of values as the Python object.
func (model *Model) pyExportSlice(v reflect.Value) *cpython.Object {
	list := make([]*cpython.Object, v.Len())
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		list[i] = model.pyExportValue(elem)
	}

	return model.py.NewObject(list)
}

// pyExportValue exports a value as the Python object.
func (model *Model) pyExportValue(v reflect.Value) *cpython.Object {
	// Handle known types
	data := v.Interface()
	switch v := data.(type) {
	// The following types have their own simple classes
	// at the Python side.
	case escl.Version:
		return model.py.NewObject(v.String())
	case uuid.UUID:
		return model.clsUUID.Call(v.String())

	// fmt.Stringer becomes Python string
	case fmt.Stringer:
		return model.py.NewObject(v.String())
	}

	// Switch by reflect.Kind
	switch v.Kind() {
	case reflect.Struct:
		return model.pyExportStruct(data)

	case reflect.Slice:
		return model.pyExportSlice(v)
	}

	// Let Python handle default case
	return model.py.NewObject(data)
}

// pyImportStruct converts the Python object into the Go structure,
// that expected to be the protocol object.
//
// p MUST be pointer to struct or pointer to pointer to struct.
func (model *Model) pyImportStruct(p any, obj *cpython.Object) error {
	// Validate argument
	t := reflect.TypeOf(p)

	msg := fmt.Sprintf("%s: invalid type", t)
	assert.MustMsg(t.Kind() == reflect.Pointer, msg)
	assert.MustMsg(p != nil, "nil pointer dereference")

	t = t.Elem()
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	assert.MustMsg(t.Kind() == reflect.Struct, msg)

	// Create a new instance of the target structure
	v := reflect.New(t).Elem()

	// Import, field by field
	for _, fld := range reflect.VisibleFields(t) {
		// Lookup python dictionary
		kw := keywordNormalize(fld.Name)
		item := obj.Get(kw)

		if err := item.Err(); err != nil {
			if item.NotFound() {
				continue
			}
			return fmt.Errorf("%s: %s", fld.Name, item.Err())
		}

		// Decode the item, if found
		fldval := v.FieldByIndex(fld.Index)
		err := model.pyImportValue(fldval, item)
		if err != nil {
			return err
		}
	}

	// Save output
	out := reflect.ValueOf(p).Elem()
	if out.Type().Kind() == reflect.Pointer {
		out.Set(v.Addr())
	} else {
		out.Set(v)
	}

	return nil
}

// pyImportSlice imports slice of values from the Python object.
func (model *Model) pyImportSlice(v reflect.Value, obj *cpython.Object) error {
	// Obtain Python object items
	slice, err := obj.Slice()
	if err != nil {
		return err
	}

	// Allocate output memory
	v.Set(reflect.MakeSlice(v.Type(), len(slice), len(slice)))

	// Decode item by item
	for i, item := range slice {
		err = model.pyImportValue(v.Index(i), item)
		if err != nil {
			return err
		}
	}

	return nil
}

// pyImportValue imports a value from the Python object.
func (model *Model) pyImportValue(v reflect.Value, obj *cpython.Object) error {
	// If we are decoding pointer to value, create a new
	// value instance and shift to it.
	if v.Kind() == reflect.Pointer {
		v2 := reflect.New(v.Type().Elem())
		v.Set(v2)
		v = v2.Elem()
	}

	// Handle known types
	switch v.Interface().(type) {
	case escl.ADFOption:
		opt, err := esclDecodeADFOption(obj)
		if err == nil {
			v.Set(reflect.ValueOf(opt))
		}
		return err

	case escl.ADFState:
		st, err := esclDecodeADFState(obj)
		if err == nil {
			v.Set(reflect.ValueOf(st))
		}
		return err

	case escl.BinaryRendering:
		rnd, err := esclDecodeBinaryRendering(obj)
		if err == nil {
			v.Set(reflect.ValueOf(rnd))
		}
		return err

	case escl.CCDChannel:
		ccd, err := esclDecodeCCDChannel(obj)
		if err == nil {
			v.Set(reflect.ValueOf(ccd))
		}
		return err

	case escl.ColorMode:
		cm, err := esclDecodeColorMode(obj)
		if err == nil {
			v.Set(reflect.ValueOf(cm))
		}
		return err

	case escl.ColorSpace:
		sps, err := esclDecodeColorSpace(obj)
		if err == nil {
			v.Set(reflect.ValueOf(sps))
		}
		return err

	case escl.ContentType:
		ct, err := esclDecodeContentType(obj)
		if err == nil {
			v.Set(reflect.ValueOf(ct))
		}
		return err

	case escl.FeedDirection:
		feed, err := esclDecodeFeedDirection(obj)
		if err == nil {
			v.Set(reflect.ValueOf(feed))
		}
		return err

	case escl.ImagePosition:
		pos, err := esclDecodeImagePosition(obj)
		if err == nil {
			v.Set(reflect.ValueOf(pos))
		}
		return err

	case escl.InputSource:
		src, err := esclDecodeInputSource(obj)
		if err == nil {
			v.Set(reflect.ValueOf(src))
		}
		return err

	case escl.Intent:
		intent, err := esclDecodeIntent(obj)
		if err == nil {
			v.Set(reflect.ValueOf(intent))
		}
		return err

	case escl.JobState:
		st, err := esclDecodeJobState(obj)
		if err == nil {
			v.Set(reflect.ValueOf(st))
		}
		return err

	case escl.JobStateReason:
		rsn, err := esclDecodeJobStateReason(obj)
		if err == nil {
			v.Set(reflect.ValueOf(rsn))
		}
		return err

	case escl.Units:
		un, err := esclDecodeUnits(obj)
		if err == nil {
			v.Set(reflect.ValueOf(un))
		}
		return err

	case escl.Version:
		ver, err := esclDecodeVersion(obj)
		if err == nil {
			v.Set(reflect.ValueOf(ver))
		}
		return err

	case uuid.UUID:
		s, err := obj.Str()
		if err != nil {
			return err
		}

		u, err := uuid.Parse(s)
		if err == nil {
			v.Set(reflect.ValueOf(u))
		}

		return err
	}

	// Switch by reflect.Kind
	switch v.Kind() {
	case reflect.Struct:
		return model.pyImportStruct(v.Addr().Interface(), obj)

	case reflect.Slice:
		return model.pyImportSlice(v, obj)

	case reflect.Int:
		i, err := obj.Int()
		if err == nil {
			v.Set(reflect.ValueOf(int(i)))
		}
		return err

	case reflect.String:
		s, err := obj.Str()
		if err == nil {
			v.Set(reflect.ValueOf(s))
		}
		return err
	}

	return nil
}
