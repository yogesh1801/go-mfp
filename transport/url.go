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
	"net/netip"
	"net/url"
	"path"
	"strconv"
	"strings"
)

// Default ports, by protocol:
const (
	DefaultPortHTTP  = 80
	DefaultPortHTTPS = 443
	DefaultPortIPP   = 631
	DefaultPortIPPS  = 631
)

// URL errors:
var (
	ErrURLInvalid       = errors.New(`URL: syntax error`)
	ErrURLSchemeMissed  = errors.New(`URL: missed scheme`)
	ErrURLSchemeInvalid = errors.New(`URL: invalid scheme`)
	ErrURLUNIXHost      = errors.New(`URL: host must be "localhost" or empty`)
)

// ParseAddr parses address string and returns result as a destination URL.
//
// This function intended for parsing addresses entered by users (interactively
// or via configuration files) and less restrictive that [ParseURL].
//
// The following additional forms are accepted:
//
//	Input                	Meaning
//	=====			=======
//	192.168.0.1		scheme://192.168.0.1:port/path
//	2001:db8::1		scheme://[2001:db8::1]:port/path
//	[2001:db8::1]		scheme://[2001:db8::1]:port/path
//	192.168.0.1:1234	scheme://192.168.0.1:1234/path
//	[2001:db8::1]:1234	scheme://[2001:db8::1]:1234/path
//	/path			unix:/path
//
// Missed scheme, port and path are taken from template. The template
// must be either URL string that MUST parse, or empty string.
//
// If template is empty, the following defaults are used:
//
//   - scheme: if port is known and equal to 80, 443, 631, the scheme will
//     be http", "https" or "ipp", respectively. Otherwise, it will be "http"
//   - port: 80
//   - path: "/"
//
// If template is not empty and it doesn't parse, this function
// panics instead of returning an error.
func ParseAddr(addr, template string) (*url.URL, error) {
	// Setup default scheme, port and path. Use template, if provided.
	scheme := "http"
	port := ""
	path := "/"

	if template != "" {
		templateURL := MustParseURL(template)
		scheme = templateURL.Scheme
		port = templateURL.Port()
		if port != "" {
			port = ":" + port
		}
		path = templateURL.Path
	}

	// Try IP addr, IP addr with port, UNIX path. Rebuild
	// URL string, if anything of above does match.
	if host := parseIPAddr(addr); host != "" {
		addr = scheme + "://" + host + port + path
	} else if host := parseIPAddrPort(addr); host != "" {
		if template == "" {
			switch {
			case strings.HasSuffix(host, ":80"):
				scheme = "http"
			case strings.HasSuffix(host, ":443"):
				scheme = "https"
			case strings.HasSuffix(host, ":631"):
				scheme = "ipp"
			}
		}

		addr = scheme + "://" + host + path
	} else if strings.HasPrefix(addr, "/") {
		addr = "unix:" + addr
	}

	// Now parse whatever we have.
	u, err := ParseURL(addr)
	if err != nil {
		return nil, errors.New("invalid address or URL")
	}

	return u, nil
}

// parseIPAddr parses IP address. It accepts the following forms of addresses:
//
//   - IPv4 dotted decimal ("192.0.2.1")
//   - IPv6 ("2001:db8::1")
//   - IPv6 in square brackets ("[2001:db8::1]")
//
// The returned string is "" on a error, or suitable as the URL.Host on success
func parseIPAddr(addr string) string {
	if strings.HasPrefix(addr, "[") && strings.HasSuffix(addr, "]") &&
		strings.IndexByte(addr, ':') >= 0 {
		addr = addr[1 : len(addr)-1]
	}

	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return ""
	}

	host := ip.String()
	if strings.IndexByte(host, ':') >= 0 {
		return "[" + host + "]"
	}

	return host
}

// parseIPAddrPort parses IP address with port. It accepts the following forms
// of addresses:
//
//   - IPv4 dotted decimal ("192.0.2.1:80")
//   - IPv6 ("[2001:db8::1]:80")
//
// The returned string is "" on a error, or suitable as the URL.Host on success
func parseIPAddrPort(addr string) string {
	ip, err := netip.ParseAddrPort(addr)
	if err != nil {
		return ""
	}

	return ip.String()
}

// ParseURL is the URL parser. In comparison to the [url.Parse] from the
// standard library it adds the following functionality:
//
//   - allowed schemes are "http", "https", "ipp", "ipps" and "unix"
//   - "unix" URL specifies HTTP request via UNIX domain sockets.
//   - Path part of the URL normalized, multiple slashed are replaced
//     with a single slash, "." and ".." segments are processed.
//
// If port number is explicitly set and matches default for the schema
// (i.e., :80 // for "http" and so on), it is stripped.
//
// The "unix" URL schema is similar to the "file" schema, as defined
// in the [RFC 8089] (surprisingly, there are still no official registration
// for the "unix" schema), with the following notes:
//
//   - "authority" part of URL may be set or omitted. If set, it
//     must be either empty or "localhost" (case-insensitive). So
//     valid forms are: "unix:/path" (no authority), "unix:///path" (empty
//     authority) or "unix://localhost/path" (localhost authority).
//   - in any case, the "unix" URL is normalized into "no authority"
//     short form (i.e., "unix:/path")
//
// Unlike [url.Parse], any unknown schemes are rejected.
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

// MustParseAddr uses [ParseAddr] to parse the address string.
// It panics in a case of errors.
func MustParseAddr(in, template string) *url.URL {
	u, err := ParseAddr(in, template)
	if err != nil {
		panic(err)
	}
	return u
}

// MustParseURL uses [ParseURL] to parse the URL string.
// It panics in a case of errors.
func MustParseURL(in string) *url.URL {
	u, err := ParseURL(in)
	if err != nil {
		panic(err)
	}
	return u
}

// ValidateAddr is the argv.Option and argv.Parameter Validate callback.
//
// It accepts any string that successfully parses with [ParseAddr] function.
func ValidateAddr(in string) error {
	_, err := ParseAddr(in, "")
	return err
}

// ValidateURL is the argv.Option and argv.Parameter Validate callback.
//
// It accepts any string that successfully parses with [ParseURL] function.
func ValidateURL(in string) error {
	_, err := ParseURL(in)
	return err
}

// URLPort returns a port number for the URL.
// If port is not set within the URL explicitly, the URL.Scheme
// will be consulted.
//
// If port number cannot be obtained, -1 will be returned.
func URLPort(u *url.URL) int {
	s := u.Port()
	if s != "" {
		port, err := strconv.Atoi(s)
		if err == nil && port >= 0 && port < 65536 {
			return port
		}

		return -1
	}

	switch u.Scheme {
	case "http":
		return DefaultPortHTTP
	case "https":
		return DefaultPortHTTPS
	case "ipp":
		return DefaultPortIPP
	case "ipps":
		return DefaultPortIPPS
	}

	return -1
}
