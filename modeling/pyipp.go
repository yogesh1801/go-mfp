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
	"fmt"
	"strings"
	"time"

	"github.com/OpenPrinting/go-mfp/cpython"
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

	objs := make([]*cpython.Object, 0, len(vals))
	for _, v := range vals {
		obj, err := model.pyExportIPPValue(v.T, v.V)
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}

	if len(objs) == 1 {
		return objs[0], nil
	}

	return model.py.NewObject(objs)
}

// pyExportIPPValue exports IPP value as [cpython.Object].
func (model *Model) pyExportIPPValue(tag goipp.Tag, val goipp.Value) (
	*cpython.Object, error) {

	// Collections handled the special way
	if v, ok := val.(goipp.Collection); ok {
		return model.pyExportIPPAttrs(goipp.Attributes(v))
	}

	// Obtain name of the Python type
	pytypename := pyIPPTagName[tag]
	if pytypename == "" {
		return nil, fmt.Errorf("invalid IPP tag %d", int(tag))
	}

	pytypename = "ipp." + pytypename

	// Obtain constructor
	pytype, err := model.py.Eval(pytypename)
	if err != nil {
		return nil, err
	}

	// Encode the value
	switch v := val.(type) {
	case goipp.Void:
		return pytype.Call()
	case goipp.Integer:
		return pytype.Call(v)
	case goipp.Boolean:
		return pytype.Call(bool(v))
	case goipp.String:
		return pytype.Call(v)
	case goipp.Time:
		return pytype.Call(v.Format(time.RFC3339))
	case goipp.Resolution:
		return pytype.Call(v.Xres, v.Yres, v.Units.String())
	case goipp.Range:
		return pytype.Call(v.Lower, v.Upper)
	case goipp.TextWithLang:
		return pytype.Call(v.Text, v.Lang)
	case goipp.Binary:
		return pytype.Call(string(v))
	}

	return model.py.None(), nil
}

// pyImportIPP imports ipp.Object from the Python representation
func (model *Model) pyImportIPP(s ipp.Object, obj *cpython.Object) error {
	attrs, err := model.pyImportIPPAttrs(obj)
	if err != nil {
		return err
	}

	_ = attrs
	return errors.New("not implemented")
}

// pyImportIPPAttrs imports IPP attributes from the [cpython.Object].
func (model *Model) pyImportIPPAttrs(obj *cpython.Object) (
	attrs goipp.Attributes, err error) {

	// Retrieve dictionary keys
	var keyobjs []*cpython.Object
	keyobjs, err = obj.Keys()
	if err != nil {
		return
	}

	for i := range keyobjs {
		var key string
		var valobj *cpython.Object

		// Obtain key name and value
		key, err = keyobjs[i].Repr()
		if err == nil {
			valobj, err = obj.Get(keyobjs[i])
		}

		if err != nil {
			return
		}

		// Decode the value
		var vals goipp.Values
		vals, err = model.pyImportIPPValues(valobj)
		if err != nil {
			return nil, pyIPPImportErrorWrap(key, err)
		}

		// Append the attribute
		attrs.Add(goipp.Attribute{Name: key, Values: vals})
	}

	return
}

// pyImportIPPAttrs imports IPP values from the [cpython.Object].
func (model *Model) pyImportIPPValues(obj *cpython.Object) (
	goipp.Values, error) {

	// If obj is the list object, expand it
	var objs []*cpython.Object

	if obj.IsSeq() {
		sz, err := obj.Len()
		if err != nil {
			return nil, err
		}

		objs = make([]*cpython.Object, sz)
		for i := 0; i < sz; i++ {
			objs[i], err = obj.Get(i)
			if err != nil {
				return nil, err
			}
		}
	} else {
		objs = []*cpython.Object{obj}
	}

	// Now decode each value
	vals := make(goipp.Values, len(objs))
	for i := 0; i < len(objs); i++ {
		tag, val, err := model.pyImportIPPValue(objs[i])
		if err != nil {
			return nil, err
		}

		vals[i].T = tag
		vals[i].V = val
	}

	return vals, nil
}

func (model *Model) pyImportIPPValue(obj *cpython.Object) (
	tag goipp.Tag, val goipp.Value, err error) {

	if obj.TypeModuleName() == "ipp" {
		typename := obj.TypeName()
		tag = pyIPPTagByName[typename]

		switch tag.Type() {
		case goipp.TypeVoid:
			val = goipp.Void{}
		case goipp.TypeInteger:
			var data int64
			data, err = obj.Int()
			val = goipp.Integer(data)
		case goipp.TypeBoolean:
			var data bool
			data, err = obj.Bool()
			val = goipp.Boolean(data)
		case goipp.TypeString:
			var data string
			data, err = obj.Str()
			val = goipp.String(data)
		case goipp.TypeDateTime:
			var data string
			data, err = obj.Str()
			if err != nil {
				return
			}

			var t time.Time
			t, err = time.Parse(time.RFC3339, data)
			if err != nil {
				return
			}

			val = goipp.Time{Time: t}
		case goipp.TypeResolution:
		case goipp.TypeRange:
		case goipp.TypeTextWithLang:
		case goipp.TypeBinary:
			var data []byte
			data, err = obj.Bytes()
			val = goipp.Binary(data)
		default:
			err = fmt.Errorf("ipp.%s: unknown type", typename)
		}

		return
	}

	if obj.IsDict() {
		var attrs goipp.Attributes
		attrs, err = model.pyImportIPPAttrs(obj)
		val = goipp.Collection(attrs)
		return
	}

	err = fmt.Errorf("%s cannot be converted to IPP value", obj.TypeName())
	return
}

// pyIPPTagName maps goipp.Tag to its Python name
var pyIPPTagName = map[goipp.Tag]string{
	// Delimiters
	goipp.TagZero: "ZERO",
	goipp.TagEnd:  "END",

	// Groups of attributes
	goipp.TagOperationGroup:         "OPERATION",
	goipp.TagJobGroup:               "JOB",
	goipp.TagPrinterGroup:           "PRINTER",
	goipp.TagUnsupportedGroup:       "UNSUPPORTED_GROUP",
	goipp.TagSubscriptionGroup:      "SUBSCRIPTION",
	goipp.TagEventNotificationGroup: "EVENT_NOTIFICATION",
	goipp.TagResourceGroup:          "RESOURCE",
	goipp.TagDocumentGroup:          "DOCUMENT",
	goipp.TagSystemGroup:            "SYSTEM",

	// Special values
	goipp.TagUnsupportedValue: "UNSUPPORTED_VALUE",
	goipp.TagDefault:          "DEFAULT",
	goipp.TagUnknown:          "UNKNOWN",
	goipp.TagNoValue:          "NOVALUE",
	goipp.TagNotSettable:      "NOTSETTABLE",
	goipp.TagDeleteAttr:       "DELETEATTR",
	goipp.TagAdminDefine:      "ADMINDEFINE",

	// Values
	goipp.TagInteger:    "INTEGER",
	goipp.TagBoolean:    "BOOLEAN",
	goipp.TagEnum:       "ENUM",
	goipp.TagString:     "STRING",
	goipp.TagDateTime:   "DATE",
	goipp.TagResolution: "RESOLUTION",
	goipp.TagRange:      "RANGE",
	goipp.TagTextLang:   "TEXTLANG",
	goipp.TagNameLang:   "NAMELANG",
	goipp.TagText:       "TEXT",
	goipp.TagName:       "NAME",
	goipp.TagKeyword:    "KEYWORD",
	goipp.TagURI:        "URI",
	goipp.TagURIScheme:  "URISCHEME",
	goipp.TagCharset:    "CHARSET",
	goipp.TagLanguage:   "LANGUAGE",
	goipp.TagMimeType:   "MIMETYPE",
	goipp.TagExtension:  "EXTENSION",

	// Collections
	goipp.TagBeginCollection: "BEGIN_COLLECTION",
	goipp.TagEndCollection:   "END_COLLECTION",
	goipp.TagMemberName:      "MEMBERNAME",
}

// pyIPPTagByName maps goipp.Tag's Python name to its value
var pyIPPTagByName = map[string]goipp.Tag{}

// init populates the pyIPPTagByName name
func init() {
	for tag, name := range pyIPPTagName {
		pyIPPTagByName[name] = tag
	}
}

// pyIPPImportError represents the error that happens during
// importing the Python object into IPP structures
type pyIPPImportError struct {
	path []string // Path over attribute names
	err  error    // Underlying error
}

// pyIPPImportErrorWrap wraps error into the pyIPPImportError.
// name is the name of the attribute the error is related to.
func pyIPPImportErrorWrap(name string, err error) error {
	if e, ok := err.(pyIPPImportError); ok {
		return pyIPPImportError{
			path: append([]string{name}, e.path...),
			err:  e.err,
		}
	}

	return pyIPPImportError{
		path: []string{name},
		err:  err,
	}
}

// Error returns the error message
func (e pyIPPImportError) Error() string {
	return fmt.Sprintf("%s: %s", strings.Join(e.path, "."), e.err)
}

// Unwrap "unwraps" the error.
func (e pyIPPImportError) Unwrap() error {
	return e.err
}
