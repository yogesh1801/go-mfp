// MFP - Miulti-Function Printers and scanners toolkit
// UUID mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UUID type and functions.

package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
	"strings"
)

// Predefined UUID values:
var (
	NilUUID = UUID{}
	MaxUUID = Must(Parse("ffffffff-ffff-ffff-ffff-ffffffffffff"))
)

// UUID represents a parsed UUID. This type is comparable and can be
// used as the map key.
type UUID [16]byte

// Parse is the universal UUID parser.
//
// It recognizes at least the following UUID formats:
//   - urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//   - uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//   - xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
//   - {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
//   - xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//
// It is very forgiving and should not be used as a validating
// parser.
func Parse(s string) (UUID, error) {
	// Strip decrations
	switch {
	case strings.HasPrefix(s, "urn:uuid:"):
		s = s[9:]
	case strings.HasPrefix(s, "uuid:"):
		s = s[5:]
	case strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}"):
		s = s[1 : len(s)-1]
	}

	// Now parse
	var uuid UUID
	cnt := 0
	for _, c := range s {
		var v rune

		// Parse next character. Ignore dashes ('-')
		switch {
		case c == '-':
			// Ignore dash characters
			continue

		case '0' <= c && c <= '9':
			v = c - '0'

		case 'a' <= c && c <= 'f':
			v = c - 'a' + 10

		case 'A' <= c && c <= 'F':
			v = c - 'A' + 10

		default:
			err := fmt.Errorf(
				"UUID contains invalid character: %q",
				string(c))
			return NilUUID, err
		}

		if cnt < 32 {
			if cnt&1 == 0 {
				uuid[cnt/2] = byte(v << 4)
			} else {
				uuid[cnt/2] |= byte(v)
			}
		}

		cnt++
	}

	// We must have exactly 32 digits
	switch {
	case cnt < 32:
		err := fmt.Errorf("UUID is too short (%d digits)", cnt)
		return NilUUID, err
	case cnt > 32:
		err := fmt.Errorf("UUID is too long (%d digits)", cnt)
		return NilUUID, err
	}

	return uuid, nil
}

// Random generates a random UUID.
// It uses [rand.Reader] as the source of entropy.
func Random() (UUID, error) {
	return RandomFrom(rand.Reader)
}

// Must returns UUID if err is nil and panics otherwise.
func Must(uuid UUID, err error) UUID {
	if err != nil {
		panic(err)
	}
	return uuid
}

// RandomFrom generates a random UUID.
// It uses provided [io.Reader] as a source of entropy.
func RandomFrom(reader io.Reader) (UUID, error) {
	var uuid UUID

	_, err := io.ReadFull(reader, uuid[:])
	if err != nil {
		return NilUUID, err
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4 (VersionRandom)
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is VariantRFC4122

	return uuid, nil
}

// String returns the string form of UUID:
//
//	xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (uuid UUID) String() string {
	const format = "" +
		"%.2x%.2x%.2x%.2x-" +
		"%.2x%.2x-" +
		"%.2x%.2x-" +
		"%.2x%.2x-" +
		"%.2x%.2x%.2x%.2x%.2x%.2x"

	return fmt.Sprintf(
		format,
		uuid[0], uuid[1], uuid[2], uuid[3],
		uuid[4], uuid[5], uuid[6], uuid[7],
		uuid[8], uuid[9], uuid[10], uuid[11],
		uuid[12], uuid[13], uuid[14], uuid[15])
}

// URN returns the URN form of UUID, per [RFC 2141]:
//
//	urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//
// [RFC 2141]: https://www.rfc-editor.org/rfc/rfc2141.html
func (uuid UUID) URN() string {
	return "urn:uuid:" + uuid.String()
}

// Microsoft returns the Microsoft style form of UUID:
//
//	{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
func (uuid UUID) Microsoft() string {
	return "{" + uuid.String() + "}"
}

// Variant returns the [Variant] encoded in uuid.
func (uuid UUID) Variant() Variant {
	switch uuid[8] & 0b_111_00000 {
	case 0b_100_00000, 0b_101_00000:
		return VariantRFC4122
	case 0b_110_00000:
		return VariantMicrosoft
	case 0b_111_00000:
		return VariantFuture
	default: // 0b_0xx_00000
		return VariantNCS
	}
}

// Version returns the [Version] of uuid.
func (uuid UUID) Version() Version {
	return Version(uuid[6] >> 4)
}
