// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP-specific URL parsing

package transport

import (
	"errors"
	"net/url"
	"path"
	"strings"
)

// Default ports, by protocol
const (
	DefaultPortHTTP  = 80
	DefaultPortHTTPS = 443
	DefaultPortIPP   = 631
	DefaultPortIPPS  = 631
)

// URL errors
var (
	ErrURLInvalid       = errors.New(`URL: syntax error`)
	ErrURLSchemeMissed  = errors.New(`URL: missed scheme`)
	ErrURLSchemeInvalid = errors.New(`URL: invalid scheme`)
	ErrURLUNIXHost      = errors.New(`URL: host must be "localhost" or empty`)
)

// ParseUserURL parses URL, entered by user. The 'addr'
// parameter must be either URL with schema set to "ipp",
// "ipps", "http" or "https" or IPv4 or IPv6 literal address
// optionally followed by port. If IPv6 address is used,
// it must be enclosed into the square brackets. IPv6
// may specify a zone suffix, as defined in [RFC 4007].
//
// The 'path' parameter is only used if URL is specified as
// IP literal. At this case, it defines path part of the
// URL. In other cases, the 'path' parameter is ignored.
//
// The path part of parsed URL is normalized according
// rules defined by the [RFC 3986].
//
// Examples:
//
//	ipp://host/...         IPP URL, port 631
//	ipp://host:port/...    IPP URL, specified port
//	ipps://host/...        IPPS URL, port 631
//	ipps://host:port/...   IPPS URL, specified port
//	http://host/...        IPP URL, port 80
//	http://host:port/...   IPP URL, specified port
//	https://host/...       IPPS URL, port 443
//	https://host:port/...  IPPS URL, specified port
//	IP4-ADDR               ipp://IP4-ADDR:631/path
//	IP4-ADDR:PORT          ipp://IP4-ADDR:PORT/path
//	[IP4-ADDR]             ipp://IP6-ADDR:631/path
//	[IP6-ADDR]:PORT        ipp://IP6-ADDR:PORT/path
//
// [RFC 3986]: https://www.rfc-editor.org/rfc/rfc3986.html
// [RFC 4007]: https://www.rfc-editor.org/rfc/rfc4007.html
func ParseUserURL(addr, path string) (*url.URL, error) {
	return nil, nil
}

// ParseURL is the URL parser. In comparison to the [url.Parse] from the
// standard library, it adds the following functionality:
//
//   - allowed schemes are "http", "https", "ipp", "ipps" and "unix"
//   - "unix" URL specifies HTTP request via UNIX domain sockets.
//   - Path part of the URL normalized, multiple slashed are replaced
//     with a single slash, "." and ".." segments are processed.
//
// The port number is stripped, if it explicitly set and matches the
// desired scheme.
//
// The "unix" URL schema is similar to the "file" schema, as defined
// in the [RFC 8089] (surprisingly, there are still no official registration
// for the "unix" schema). With the following notes:
//
//   - "authority" part of URL may be set or omitted. If set, it
//     must be either or "localhost" (case-insensitive). So valid
//     forms are: "unix:/path" (no authority), "unix:///path" (empty
//     authority) or "unix://localhost/path" (localhost authority).
//   - in any case, the "unix" URL is normalized into "no authority"
//     short form (i.e., "unix:/path")
//
// [RFC 8089]: https://www.rfc-editor.org/rfc/rfc8089.html
func ParseURL(in string) (*url.URL, error) {
	// Parse the URL string
	u, err := url.Parse(in)
	if err != nil {
		return nil, ErrURLInvalid
	}

	// Do schema-specific checks and postprocessing
	port := ""

	switch u.Scheme {
	case "http":
		port = "80"
	case "https":
		port = "443"
	case "ipp", "ipps":
		port = "631"

	case "unix":
		switch strings.ToLower(u.Host) {
		case "", "localhost":
		default:
			return nil, ErrURLUNIXHost
		}

		u.Host = ""
		u.OmitHost = true

	case "":
		return nil, ErrURLSchemeMissed
	default:
		return nil, ErrURLSchemeInvalid
	}

	if port != "" && u.Port() == port {
		u.Host, _ = strings.CutSuffix(u.Host, ":"+port)
	}

	// Normalize path
	endSlash := strings.HasSuffix(u.Path, "/")

	u.Path = path.Clean(u.Path)
	switch u.Path {
	case "", ".":
		u.Path = "/"
	}

	if endSlash && !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}

	return u, nil
}
