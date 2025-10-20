// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversion between ipp.Object and cpython.Object

package modeling

import (
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
		return model.py.NewObject(v.String())
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
