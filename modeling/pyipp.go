// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversion between ipp.Object and cpython.Object

package modeling

import (
	"fmt"
	"strings"
	"time"

	"github.com/OpenPrinting/go-mfp/cpython"
	"github.com/OpenPrinting/go-mfp/proto/ipp"
	"github.com/OpenPrinting/goipp"
)

// pyExportIPP converts the [ipp.Object] into the [cpython.Object].
func (model *Model) pyExportIPP(s ipp.Object) (*cpython.Object, error) {
	obj, err := model.pyExportIPPAttrs(s.RawAttrs().All())
	return obj, err
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

// pyImportPrinterAppributes imports IPP printer attributes from the
// Python representation
func (model *Model) pyImportPrinterAppributes(obj *cpython.Object) (
	*ipp.PrinterAttributes, error) {
	attrs, err := model.pyImportIPPAttrs(obj)
	if err != nil {
		return nil, err
	}

	return ipp.DecodePrinterAttributes(attrs)
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
		key, err = keyobjs[i].Str()
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

// pyImportIPPValue imports IPP value from the Python object
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
		case goipp.TypeString, goipp.TypeBinary:
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
			val, err = model.pyImportIPPResolution(obj)
		case goipp.TypeRange:
			val, err = model.pyImportIPPRange(obj)
		case goipp.TypeTextWithLang:
			val, err = model.pyImportIPPTextWithLang(obj, tag)
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

// pyImportIPPResolution imports IPP resolution from the Python object
func (model *Model) pyImportIPPResolution(obj *cpython.Object) (
	res goipp.Resolution, err error) {

	var x, y int64
	var units goipp.Units

	// Load Xres
	obj2, err := obj.GetAttr("X")
	if err == nil {
		x, err = obj2.Int()
	}

	if err != nil {
		err = pyIPPImportErrorWrap("X", err)
		return
	}

	// Load Yres
	obj2, err = obj.GetAttr("Y")
	if err == nil {
		y, err = obj2.Int()
	}

	if err != nil {
		err = pyIPPImportErrorWrap("Y", err)
		return
	}

	// Load Units
	obj2, err = obj.GetAttr("Units")
	if err == nil {
		var s string
		s, err = obj2.Str()

		switch {
		case err != nil:
		case s == "dpi":
			units = goipp.UnitsDpi
		case s == "dpcm":
			units = goipp.UnitsDpcm
		default:
			err = fmt.Errorf("%s: invalid resolution units", s)
		}
	}

	if err != nil {
		err = pyIPPImportErrorWrap("Units", err)
		return
	}

	res = goipp.Resolution{
		Xres:  int(x),
		Yres:  int(y),
		Units: units,
	}

	return
}

// pyImportIPPRange imports IPP range from the Python object
func (model *Model) pyImportIPPRange(obj *cpython.Object) (
	rng goipp.Range, err error) {

	var lower, upper int64

	// Load Lower
	obj2, err := obj.GetAttr("Lower")
	if err == nil {
		lower, err = obj2.Int()
	}

	if err != nil {
		err = pyIPPImportErrorWrap("Lower", err)
		return
	}

	// Load Upper
	obj2, err = obj.GetAttr("Upper")
	if err == nil {
		upper, err = obj2.Int()
	}

	if err != nil {
		err = pyIPPImportErrorWrap("Upper", err)
		return
	}

	rng = goipp.Range{
		Lower: int(lower),
		Upper: int(upper),
	}

	return
}

// pyImportIPPTextWithLang imports IPP text with language from the Python object
func (model *Model) pyImportIPPTextWithLang(obj *cpython.Object, tag goipp.Tag) (
	txt goipp.TextWithLang, err error) {

	var lang, text string

	// Load lang
	obj2, err := obj.GetAttr("Lang")
	if err == nil {
		lang, err = obj2.Str()
	}

	if err != nil {
		err = pyIPPImportErrorWrap("Lang", err)
		return
	}

	// Load Text or Name
	nm := "Text"
	if tag == goipp.TagNameLang {
		nm = "Name"
	}

	obj2, err = obj.GetAttr(nm)
	if err == nil {
		text, err = obj2.Str()
	}

	if err != nil {
		err = pyIPPImportErrorWrap(nm, err)
		return
	}

	txt = goipp.TextWithLang{
		Lang: lang,
		Text: text,
	}

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
