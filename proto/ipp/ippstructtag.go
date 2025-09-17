// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// "ipp:" struct tag parser

package ipp

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/OpenPrinting/goipp"
)

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
	"collection":       goipp.TagBeginCollection,
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
			err = fmt.Errorf("unknown keyword: %q", part)
		}

		if err != nil {
			return nil, err
		}
	}

	return stag, nil
}

// parseKeyword parses keyword parameter of the ipp: struct tag:
//   - IPP tag ("int", "name" etc)
//   - May be, some flags in a future
//
// Return value:
//   - true, nil  - parameter was parsed and applied
//   - false, nil - parameter is not keyword
//   - false, err - invalid parameter
func (stag *ippStructTag) parseKeyword(s string) (bool, error) {
	kw := strings.ToLower(s)
	zeroTag := goipp.TagZero

	switch {
	case strings.HasSuffix(kw, "|unknown"):
		zeroTag = goipp.TagUnknown
		kw = kw[:len(kw)-8]
	case strings.HasSuffix(kw, "|no-value"):
		zeroTag = goipp.TagNoValue
		kw = kw[:len(kw)-9]
	}

	if tag, ok := ippStructTagToIppTag[kw]; ok {
		stag.ippTag = tag
		stag.zeroTag = zeroTag
		return true, nil
	}

	return false, nil
}

// parseRange parses min/max limit constraints:
//
//	>NNN - min range
//	<NNN - max range
//
// Return value:
//   - true, nil  - parameter was parsed and applied
//   - false, nil - parameter is not keyword
//   - false, err - invalid parameter
func (stag *ippStructTag) parseMinMax(s string) (bool, error) {
	// Limit starts with '<' or '>'
	pfx := s[0]
	if pfx != '<' && pfx != '>' {
		return false, nil
	}

	// Parse limit
	v, err := strconv.ParseInt(s[1:], 10, 64)
	if err != nil {
		err = fmt.Errorf("%q: invalid limit", s)
		return false, err
	}

	// Save limit; check for range
	switch pfx {
	case '>':
		if math.MinInt32-1 <= v && v <= math.MaxInt32-1 {
			stag.min = int(v + 1)
		} else {
			err = fmt.Errorf("%q: limit out of range", s)
		}

	case '<':
		if math.MinInt32+1 <= v && v <= math.MaxInt32+1 {
			stag.max = int(v - 1)
		} else {
			err = fmt.Errorf("%q: limit out of range", s)
		}
	}

	return true, err
}

// parseRange parses range constraints:
//
//	MAX
//	MIN:MAX
//
// Return value:
//   - true, nil  - parameter was parsed and applied
//   - false, nil - parameter is not range
//   - false, err - invalid parameter
func (stag *ippStructTag) parseRange(s string) (bool, error) {
	// Range starts with '('
	if s[0] != '(' {
		return false, nil
	}

	// Range ends with ')'
	if s[len(s)-1] != ')' {
		return false, fmt.Errorf("range: missed ')'")
	}

	// Strip brackets
	s = s[1 : len(s)-1]

	// Range consists of 1 or 2 fields
	fields := strings.Split(s, ":")
	if (len(fields) != 1 && len(fields) != 2) ||
		fields[0] == "" || fields[1] == "" {
		err := fmt.Errorf("range (%s): invalid syntax", s)
		return false, err
	}

	// Parse fields.
	parsed := make([]int32, len(fields))
	for i, fld := range fields {
		switch fld {
		case "MIN":
			parsed[i] = math.MinInt32
		case "MAX":
			parsed[i] = math.MaxInt32
		default:
			v, err := strconv.ParseInt(fld, 10, 32)
			parsed[i] = int32(v)

			if err != nil {
				err = fmt.Errorf("range (%s): invalid value", s)
				return false, err
			}
		}
	}

	// Check range
	var min, max int32
	if len(fields) == 2 {
		min = parsed[0]
		max = parsed[1]
	} else {
		min = math.MinInt32
		max = parsed[0]
	}

	if min > max {
		return false, fmt.Errorf("range (%d:%d): min>max", min, max)
	}

	stag.min = int(min)
	stag.max = int(max)

	return true, nil
}
