// MFP - Miulti-Function Printers and scanners toolkit
// IPP registrations to Go converter.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Attribute syntax parser

package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/goipp"
)

// Syntax represents the parsed syntax of the attribute
//
// The syntax looks as follows:
//
//	integer
//	integer (0:MAX)
//	collection | no-value
//	1setOf (integer(MIN:MAX))
//	1setOf collection | no-value
//	1setOf type2 keyword | name(MAX)
//
// and so on.
type Syntax struct {
	SetOf      bool        // True for 1setOf attributes
	Collection bool        // This is collection
	Tags       []goipp.Tag // Allowed value tags
	Min, Max   int32       // Value limits
}

// MIN and MAX values for value limits
const (
	MIN = math.MinInt32
	MAX = math.MaxInt32
)

type tok1setOF struct{}
type tokValue struct {
	tags     []goipp.Tag
	min, max int32
}

// ParseSyntax parses attribute syntax.
func ParseSyntax(s string) (syntax Syntax, err error) {
	// Tokenize and decode syntax string
	tokens, err := syntax.decodeTokens(s)
	if err != nil {
		return
	}

	// Parse syntax
	syntax.Min = MIN
	syntax.Max = MAX

	for _, tok := range tokens {
		switch tok := tok.(type) {
		case tok1setOF:
			syntax.SetOf = true

		case tokValue:
			syntax.Tags = append(syntax.Tags, tok.tags...)
			syntax.Min = generic.Max(syntax.Min, tok.min)
			syntax.Max = generic.Min(syntax.Max, tok.max)
		}
	}

	// Sort and dedup tags, to make equal syntaxes comparable
	sort.Slice(syntax.Tags, func(i, j int) bool {
		t1 := syntax.Tags[i]
		t2 := syntax.Tags[j]

		switch {
		case t1.Type() != goipp.TypeVoid && t2.Type() == goipp.TypeVoid:
			return true
		case t1.Type() == goipp.TypeVoid && t2.Type() != goipp.TypeVoid:
			return false
		default:
			return t1 < t2
		}
	})

	end := 1
	for i := 1; i < len(syntax.Tags); i++ {
		if syntax.Tags[i] != syntax.Tags[i-1] {
			syntax.Tags[end] = syntax.Tags[i]
			end++
		}
	}

	syntax.Tags = syntax.Tags[:end]

	for _, tag := range syntax.Tags {
		if tag == goipp.TagBeginCollection {
			syntax.Collection = true
			break
		}
	}

	return
}

// Equal reports if two syntaxes are equal
func (syntax Syntax) Equal(syntax2 Syntax) bool {
	return reflect.DeepEqual(syntax, syntax2)
}

// FormatMin string returns Min as a string, either the value or "MIN"
// if syntax.Min == MIN
func (syntax Syntax) FormatMin() string {
	if syntax.Min == MIN {
		return "MIN"
	}
	return strconv.FormatInt(int64(syntax.Min), 10)
}

// FormatMax string returns Max as a string, either the value or "MAX"
// if syntax.Max == MAX
func (syntax Syntax) FormatMax() string {
	if syntax.Max == MAX {
		return "MAX"
	}
	return strconv.FormatInt(int64(syntax.Max), 10)
}

// decodeTokens splits syntax string into tokens and decodes
// these tokens
func (syntax Syntax) decodeTokens(s string) ([]any, error) {
	strtok := syntax.tokenize(s)
	tokens := make([]any, 0, len(strtok))

	for i := 0; i < len(strtok); i++ {
		tok := strings.ToLower(strtok[i])
		switch tok {
		case "1setof":
			tokens = append(tokens, tok1setOF{})
		// These are ignored
		case "type1", "type2":
		case "(", ")":
		case "|":
		case "'":

		default:
			tags := tags[tok]
			if tags == nil {
				err := fmt.Errorf("%q: invalid token %q",
					s, strtok[i])
				return nil, err
			}

			min, max, skip, err := syntax.decodeLimits(strtok[i+1:])
			if err != nil {
				err := fmt.Errorf("%q: %w", s, err)
				return nil, err
			}

			tokval := tokValue{
				tags: tags,
				min:  min,
				max:  max,
			}

			tokens = append(tokens, tokval)
			i += skip
		}
	}

	return tokens, nil
}

// decodeLimits decodes MIN/MAX limits after the
// sequence of string tokens and returns decoded
// values and count of consumed tokens
func (syntax Syntax) decodeLimits(strtok []string) (
	min, max int32, consumed int, err error) {

	min, max = MIN, MAX

	switch {
	// value(MAX)
	case len(strtok) >= 3 &&
		strtok[0] == "(" && strtok[2] == ")":

		max, err = syntax.decodeMinMax(strtok[1])
		if err == nil {
			consumed = 3
		}

	// value(MIN:MAX)
	case len(strtok) >= 5 &&
		strtok[0] == "(" && strtok[2] == ":" && strtok[4] == ")":

		min, err = syntax.decodeMinMax(strtok[1])
		if err == nil {
			max, err = syntax.decodeMinMax(strtok[3])
		}
		if err == nil {
			consumed = 5
		}
	}

	return
}

// decodeMinMax decodes min or max value.
func (syntax Syntax) decodeMinMax(s string) (v int32, err error) {
	switch strings.ToLower(s) {
	case "min":
		return MIN, nil
	case "max":
		return MAX, nil
	}

	var tmp int64
	tmp, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		err = fmt.Errorf("invalid limit %q", s)
		return
	}

	return int32(tmp), nil
}

// tokenize splits syntax string into token strings
func (syntax Syntax) tokenize(s string) []string {
	strtok := []string{}
	in := []byte(s)

	for len(in) != 0 {
		c := rune(in[0])
		switch {
		case unicode.IsLetter(c) || unicode.IsDigit(c):
			token := ""
			for len(in) > 0 &&
				(unicode.IsLetter(rune(in[0])) ||
					unicode.IsDigit(rune(in[0])) ||
					in[0] == '-') {

				token += string(in[0])
				in = in[1:]
			}
			strtok = append(strtok, token)

		case c == '-':
			token := string(c)
			in = in[1:]

			for len(in) > 0 && (unicode.IsDigit(rune(in[0]))) {
				token += string(in[0])
				in = in[1:]
			}
			strtok = append(strtok, token)

		case unicode.IsSpace(rune(in[0])):
			in = in[1:]
		default:
			strtok = append(strtok, string(rune(in[0])))
			in = in[1:]
		}
	}

	return strtok
}

// tags maps tag names (e.g., "integer") to tag values
var tags = map[string][]goipp.Tag{}

func init() {
	tags["collection"] = []goipp.Tag{goipp.TagBeginCollection}
	tags["name"] = []goipp.Tag{
		goipp.TagName,
		goipp.TagNameLang,
	}
	tags["text"] = []goipp.Tag{
		goipp.TagText,
		goipp.TagTextLang,
	}

	for tag := goipp.TagUnsupportedValue; tag < goipp.TagExtension; tag++ {
		switch tag {
		// These are handled the special way
		case goipp.TagBeginCollection:
		case goipp.TagEndCollection:
		case goipp.TagMemberName:
		// By default tag name maps to the tag value
		default:
			tags[strings.ToLower(tag.String())] = []goipp.Tag{tag}
		}
	}
}
