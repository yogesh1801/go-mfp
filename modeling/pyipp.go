// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversion between ipp.Object and cpython.Object

package modeling

import (
	"errors"

	"github.com/OpenPrinting/go-mfp/cpython"
	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/proto/ipp"
	"github.com/OpenPrinting/goipp"
)

// pyExportIPP converts the [ipp.Object] into the [cpython.Object].
func (model *Model) pyExportIPP(s ipp.Object) (*cpython.Object, error) {
	return model.pyExportIPPAttrs(s.RawAttrs().All())
}

// pyExportIPPAttrs exports IPP attributes into the [cpython.Object].
func (model *Model) pyExportIPPAttrs(attrs goipp.Attributes) (
	*cpython.Object, error) {

	// Create output cpython.Object (the empty dict).
	dict, err := model.py.NewObject(map[any]any(nil))
	if err != nil {
		return nil, err
	}

	// Roll over all IPP attributes
	for _, attr := range attrs {
		vals, err := model.pyExportIPPValues(attr.Values)
		if err != nil {
			return nil, err
		}

		err = dict.Set(keywordNormalize(attr.Name), vals)
		if err != nil {
			return nil, err
		}
	}

	return dict, nil
}

// pyExportIPPValues exports IPP attribute values into the [cpython.Object].
func (model *Model) pyExportIPPValues(vals goipp.Values) (
	*cpython.Object, error) {

	objs := make([]*cpython.Object, len(vals))
	for i := range objs {
		var err error
		objs[i], err = model.pyExportIPPValue(vals[i].T, vals[i].V)
		if err != nil {
			return nil, err
		}
	}

	if len(objs) == 1 {
		return objs[0], nil
	}

	return model.py.NewObject(objs)
}

// pyExportIPPValue exports a single IPP value into the [cpython.Object].
func (model *Model) pyExportIPPValue(tag goipp.Tag, val goipp.Value) (
	*cpython.Object, error) {

	switch v := val.(type) {
	case goipp.Void:
		return model.py.None(), nil
	case goipp.Integer:
		return model.py.NewObject(v)
	case goipp.Boolean:
		return model.py.Bool(bool(v)), nil
	case goipp.String:
		return model.py.NewObject(v)
	case goipp.Time:
		return model.clsDateTimeFromISO.Call(v.String())
	case goipp.Resolution:
		return model.py.NewObject(v.String())
	case goipp.Range:
		return model.py.NewObject(v.String())
	case goipp.TextWithLang:
		return model.py.NewObject(v.String())
	case goipp.Binary:
		return model.py.NewObject(string(v))
	case goipp.Collection:
		return model.pyExportIPPAttrs(goipp.Attributes(v))
	}

	return model.py.None(), nil
	assert.NoError(errors.New("internal error"))
	return nil, nil
}
