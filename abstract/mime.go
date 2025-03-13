// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common MIME types

package abstract

// Common MIME Types:
const (
	// MIMETypeNetbpm represents the "Portable anymap" format, also
	// known as Netpbm.  This format is very simple, as it does not
	// use any image compression and consists of a basic header
	// followed by a sequence of raw image bytes.
	//
	// This format is used as an intermediate step for image format
	// conversion.  Additionally, the abstract [Scanner]
	// implementation, backed by SANE, is expected to return images
	// in this format. This is because SANE's internal
	// representation of raw images can easily be converted into
	// Netpbm by adding the appropriate header.
	//
	// https://en.wikipedia.org/wiki/Netpbm
	MIMETypeNetbpm = "image/x-portable-anymap"

	// Other common image formats:
	MIMETypeJPEG = "image/jpeg"
	MIMETypeTIFF = "image/tiff"
	MIMETypePNG  = "image/png"
	MIMETypePDF  = "application/pdf"
)
