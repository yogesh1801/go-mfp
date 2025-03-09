// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The scanner interface

package abstract

import (
	"context"
	"io"
)

// Scanner represents the abstract (implementation-independent)
// interface to the document scanner.
//
// The interface is greatly inspired by the [eSCL], [WS-Scan] and [IPP-Scan]
// specifications.
//
// [eSCL]: https://mopria.org/mopria-escl-specification
// [WS-Scan]: https://github.com/MicrosoftDocs/windows-driver-docs/blob/staging/windows-driver-docs-pr/image/scan-service--ws-scan--schema.md
// [IPP-Scan]: https://ftp.pwg.org/pub/pwg/candidates/cs-ippscan10-20140918-5100.17.pdf
type Scanner interface {
	// Capabilities returns the [ScannerCapabilities].
	// Caller should not modify the returned structure.
	Capabilities() *ScannerCapabilities

	// Scan supplies the scan request. Request parameters are
	// defined via [ScannerRequest] structure.
	//
	// Requests are cancelable via provided [context.Context].
	Scan(context.Context, ScannerRequest) (io.ReadCloser, error)

	// Close closes the scanner connection.
	Close() error
}
