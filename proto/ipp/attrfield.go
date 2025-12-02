// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Handling of IPP structure fields

package ipp

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/OpenPrinting/go-mfp/proto/ipp/iana"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/goipp"
)

// attrFieldAnalyze analyzes field of the IPP-encodable
// go structure for encoding as IPP attribute.
//
// It understands the "ipp:" reflection tag which can take
// one of two forms:
//   - `ipp:"charset-configured,charset"` -- IPP name and syntax
//   - `ipp:"charset-configured,charset"` -- just the name
//
// If structure field doesn't contain the "ipp:" structure
// tag, it is not treated as the IPP attribute, and return
// value is ("", nil, nil)
//
// Otherwise, if there is no error, non-empty IPP attribute
// name is always returned.
//
// If syntax is present, it is decoded and returned as the
// *[iana.DefAttr] structure.
func attrFieldAnalyze(fld reflect.StructField) (
	name string, def *iana.DefAttr, err error) {

	// Is it IPP attribute
	tag, ok := fld.Tag.Lookup("ipp")
	if !ok {
		return "", nil, nil
	}

	// IPP attribute must be exported field
	if !fld.IsExported() {
		err = errors.New("ipp:tag used with unexported field")
		return
	}

	// IPP attribute must not be anonymous field
	if fld.Anonymous {
		err = errors.New("ipp:tag used with anonymous field")
		return
	}

	// Parse the struct tag -- split it to name and syntax
	tag = strings.TrimSpace(tag)
	name = tag
	if i := strings.IndexByte(tag, ','); i >= 0 {
		name = strings.TrimSpace(tag[:i])
		syntax := strings.TrimSpace(tag[i+1:])
		def, err = attrSyntaxParse(syntax)
	}

	if name == "" {
		err = errors.New("ipp:missed attribute name")
	}

	if err != nil {
		return "", nil, err
	}

	return
}

// attrSyntaxParse parses attribute syntax
func attrSyntaxParse(s string) (*iana.DefAttr, error) {
	def := iana.DefAttr{
		Min: iana.MIN,
		Max: iana.MAX,
	}
	tokens := attrSyntaxTokenize(s)
	tags := generic.NewSet[goipp.Tag]()

	if len(tokens) > 0 && tokens[0] == "1setof" {
		def.SetOf = true
		tokens = tokens[1:]
	}

	for len(tokens) > 0 {
		// Parse value tag
		var tag goipp.Tag
		var tok string

		tok, tokens = tokens[0], tokens[1:]

		switch tok {
		case "type1", "type2", "type3", "(", ")", "|":
			continue

		default:
			tag = attrTagsByName[tok]
			if tag == goipp.TagZero {

				err := fmt.Errorf("ipp:%q: unexpected token", tok)
				return nil, err
			}
		}

		tags.Add(tag)

		// Parse limits
		min := ""
		max := ""
		switch {
		case len(tokens) >= 3 && tokens[0] == "(" && tokens[2] == ")":
			max = tokens[1]
			tokens = tokens[3:]

		case len(tokens) >= 5 &&
			tokens[0] == "(" && tokens[2] == ":" && tokens[4] == ")":
			min = tokens[1]
			max = tokens[3]
			tokens = tokens[5:]
		}

		if min != "" && min != "min" {
			v, err := strconv.ParseInt(min, 10, 32)
			if err != nil {
				err := fmt.Errorf("ipp:%q: invalid limit", min)
				return nil, err
			}

			def.Min = generic.Max(def.Min, int32(v))
		}

		if max != "" && max != "max" {
			v, err := strconv.ParseInt(max, 10, 32)
			if err != nil {
				err := fmt.Errorf("ipp:%q: invalid limit", max)
				return nil, err
			}

			def.Max = generic.Min(def.Max, int32(v))
		}

		tagmin, tagmax := tag.Limits()
		def.Min = generic.Max(def.Min, tagmin)
		def.Max = generic.Min(def.Max, tagmax)
	}

	// Populate DefAttr.Tags
	for _, tag := range attrTagsSortingOrder {
		if tags.Contains(tag) {
			def.Tags = append(def.Tags, tag)
		}
	}

	if len(def.Tags) == 0 {
		err := fmt.Errorf("ipp:%q: no tags defined", s)
		return nil, err
	}

	return &def, nil
}

// attrSyntaxTokenize splits syntax string into tokens.
// Alphanumerical tokens are lowercased.
func attrSyntaxTokenize(s string) []string {
	word := false
	tokens := []string{}

	for _, c := range s {
		switch {
		case unicode.IsSpace(c):
			word = false

		case c > unicode.MaxLatin1:
			fallthrough
		default:
			word = false
			tokens = append(tokens, string(c))

		case unicode.IsDigit(c) || unicode.IsLetter(c) || c == '-':
			if !word {
				word = true
				tokens = append(tokens, "")
			}
			tokens[len(tokens)-1] += string(unicode.ToLower(c))
		}
	}

	return tokens
}

// attrTagsByName maps tags names to tags numbers
var attrTagsByName = map[string]goipp.Tag{
	"admin-define":     goipp.TagAdminDefine,
	"boolean":          goipp.TagBoolean,
	"charset":          goipp.TagCharset,
	"collection":       goipp.TagBeginCollection,
	"datetime":         goipp.TagDateTime,
	"default":          goipp.TagDefault,
	"delete-attribute": goipp.TagDeleteAttr,
	"enum":             goipp.TagEnum,
	"integer":          goipp.TagInteger,
	"keyword":          goipp.TagKeyword,
	"mimemediatype":    goipp.TagMimeType,

	// Use goipp.TagName for both
	"name":             goipp.TagName,
	"namewithlanguage": goipp.TagName,

	"naturallanguage": goipp.TagLanguage,
	"not-settable":    goipp.TagNotSettable,
	"no-value":        goipp.TagNoValue,
	"rangeofinteger":  goipp.TagRange,
	"resolution":      goipp.TagResolution,
	"string":          goipp.TagString,

	// Use goipp.TagText for both
	"text":             goipp.TagText,
	"textwithlanguage": goipp.TagText,

	"unknown":     goipp.TagUnknown,
	"unsupported": goipp.TagUnsupportedValue,
	"uri":         goipp.TagURI,
	"urischeme":   goipp.TagURIScheme,
}

// attrTagsSortingOrder defines sorting order for DefAttr.Tags
var attrTagsSortingOrder = []goipp.Tag{
	// Prefer Enum over Integer
	goipp.TagEnum,
	goipp.TagInteger,

	// Prefer Keyword, then Name, then NameLang
	goipp.TagKeyword,
	goipp.TagName,
	goipp.TagNameLang,

	// Prefer Text over TextLang
	goipp.TagText,
	goipp.TagTextLang,

	// No special order for these tags
	goipp.TagBoolean,
	goipp.TagString,
	goipp.TagDateTime,
	goipp.TagResolution,
	goipp.TagRange,
	goipp.TagBeginCollection,
	goipp.TagReservedString,
	goipp.TagURI,
	goipp.TagURIScheme,
	goipp.TagCharset,
	goipp.TagLanguage,
	goipp.TagMimeType,

	// Put no-value (OOB) tags to the end
	goipp.TagUnsupportedValue,
	goipp.TagDefault,
	goipp.TagUnknown,
	goipp.TagNoValue,
	goipp.TagNotSettable,
	goipp.TagDeleteAttr,
	goipp.TagAdminDefine,
}
