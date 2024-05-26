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
//
//	MIN:MAX
//
// Return value:
//   - true, nil  - parameter was parsed and applied
//   - false, nil - parameter is not keyword
//   - false, err - invalid parameter
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
