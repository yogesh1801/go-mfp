// MFP - Miulti-Function Printers and scanners toolkit
// UUID mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UUID version and variant

package uuid

// Version represents UUID verson.
// See [RFC 4122, 4.1.3.] for details.
//
// [RFC 4122, 4.1.3.]: https://www.rfc-editor.org/rfc/rfc4122.html#section-4.1.3
type Version int

// Standard Version values:
const (
	VersionTimeBased     Version = 1
	VersionDCESecurity   Version = 2
	VersionNameBasedMD5  Version = 3
	VersionRandom        Version = 4
	VersionNameBasedSHA1 Version = 5
)

// Variant represents UUID variant.
// See [RFC 4122, 4.1.1.] for details.
//
// [RFC 4122, 4.1.1.]: https://www.rfc-editor.org/rfc/rfc4122.html#section-4.1.1
type Variant int

// Standard Variant values:
const (
	VariantRFC4122   Variant = 1 // Per RFC 4122
	VariantNCS       Variant = 2 // Reserved, NCS backward compatibility.
	VariantMicrosoft Variant = 3 // Reserved, MS backward compatibility.
	VariantFuture    Variant = 4 // Reserved for future definition.
)
