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
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/OpenPrinting/go-mfp/util/missed"
	"golang.org/x/net/idna"
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
	ErrURLHostMissed    = errors.New(`URL: missed host`)
	ErrURLUNIXHost      = errors.New(`URL: host must be "localhost" or empty`)
)

// ParseAddr parses address string and returns result as a destination URL.
//
// This function intended for parsing addresses entered by users (interactively
// or via configuration files) and less restrictive that [ParseURL].
//
// The following additional forms are accepted (assuming http scheme):
//
//	Input                   Meaning
//	=====                   =======
//	192.168.0.1             http://192.168.0.1
//	2001:db8::1             http://[2001:db8::1]
//	[2001:db8::1]           http://[2001:db8::1]
//	192.168.0.1:1234        http://192.168.0.1:1234
//	[2001:db8::1]:1234      http://[2001:db8::1]:1234
//	192.168.0.1/path        http://192.168.0.1/path
//	192.168.0.1:1234/path   http://192.168.0.1:1234/path
//	2001:db8::1/path        http://[2001:db8::1]/path
//	[2001:db8::1]:1234/path http://[2001:db8::1]:1234/path
//	/path                   unix:/path
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
	// Parse template, if provided
	var templateURL *url.URL
	if template != "" {
		templateURL = MustParseURL(template)
	}

	// Split address into parts
	scheme, host, path, ok := parseSplit(addr)
	if !ok {
		return nil, ErrURLInvalid
	}

	// If we have a scheme, parse as a normal URL
	if scheme != "" {
		return ParseURL(addr)
	}

	// Initialize defaults
	port := ""
	scheme = "http"
	if templateURL != nil {
		scheme = templateURL.Scheme
		port = templateURL.Port()
		if path == "" {
			path = templateURL.Path
		}
	}

	// Try IP addr, IP addr with port, UNIX path. Rebuild
	// URL string, if something of above does match.
	if tmp := parseIPAddr(host); tmp != "" {
		host = tmp
		if port != "" {
			host = net.JoinHostPort(host, port)
		}
	} else if tmp := parseIPAddrPort(host); tmp != "" {
		host = tmp
		if templateURL == nil {
			switch {
			case strings.HasSuffix(host, ":80"):
				scheme = "http"
			case strings.HasSuffix(host, ":443"):
				scheme = "https"
			case strings.HasSuffix(host, ":631"):
				scheme = "ipp"
			}
		}
	} else if strings.HasPrefix(addr, "/") {
		scheme = "unix"
	} else {
		return nil, ErrURLInvalid
	}

	addr = scheme + "://" + host + "/" + path

	// Now parse whatever we have.
	u, err := ParseURL(addr)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// parseSplit splits the url string into the scheme, host and path
func parseSplit(s string) (scheme, host, path string, ok bool) {
	if s == "" {
		return
	}

	// Cut off scheme
	if i := strings.Index(s, "://"); i >= 0 {
		scheme, host = s[:i], s[i+3:]
		if scheme == "" || host == "" {
			return
		}
	} else {
		host = s
	}

	// Cut off path
	if i := strings.Index(host, "/"); i >= 0 {
		host, path = host[:i], host[i:]
	}

	ok = true
	return
}

// parseIPAddr parses IP address. It accepts the following forms of addresses:
//
//   - IPv4 dotted decimal ("192.0.2.1")
//   - IPv6 ("2001:db8::1")
//   - IPv6 in square brackets ("[2001:db8::1]")
//   - domain name ("example.com")
//
// The returned string is "" on a error, or suitable as the URL.Host on success
func parseIPAddr(addr string) string {
	if strings.HasPrefix(addr, "[") && strings.HasSuffix(addr, "]") &&
		strings.IndexByte(addr, ':') >= 0 {
		addr = addr[1 : len(addr)-1]
	}

	ip, err := netip.ParseAddr(addr)
	if err != nil {
		ascii, err := idna.Lookup.ToASCII(addr)
		if err == nil {
			return ascii
		}
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
//   - domain with literal port ("example.com:80")
//
// The returned string is "" on a error, or suitable as the URL.Host on success
func parseIPAddrPort(addr string) string {
	// Try netip.ParseAddrPort, it handles literal addresses
	ip, err := netip.ParseAddrPort(addr)
	if err == nil {
		return ip.String()
	}

	// Try to split into host and port and parse separately
	i := strings.LastIndexByte(addr, ':')
	if i <= 0 || i == len(addr)-1 {
		return ""
	}

	host, port := addr[:i], addr[i+1:]

	host = parseIPAddr(host)
	if host == "" {
		return ""
	}

	portnum, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s:%d", host, portnum)
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
	// Test some corner cases
	if in == "" {
		return nil, ErrURLInvalid
	}

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

	if u.Scheme != "unix" && u.Host == "" {
		return nil, ErrURLHostMissed
	}

	if port != "" && u.Port() == port {
		u.Host, _ = missed.StringsCutSuffix(u.Host, ":"+port)
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
		err = fmt.Errorf("%w (%q)", err, in)
		panic(err)
	}
	return u
}

// MustParseURL uses [ParseURL] to parse the URL string.
// It panics in a case of errors.
func MustParseURL(in string) *url.URL {
	u, err := ParseURL(in)
	if err != nil {
		err = fmt.Errorf("%w (%q)", err, in)
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

// URLClone makes a shallow copy of the input URL.
func URLClone(u *url.URL) *url.URL {
	u2 := *u
	return &u2
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

// URLForcePort ensures that u.Host includes the explicit port number,
// if applicable.
func URLForcePort(u *url.URL) {
	port := URLPort(u)
	if port >= 0 && u.Port() == "" {
		u.Host += ":" + strconv.Itoa(port)
	}
}

// URLStripPort strips unneeded explicit :port in the u.Host
func URLStripPort(u *url.URL) {
	port := URLPort(u)
	if port >= 0 && port == DefaultPort(u.Scheme) {
		suffix := ":" + strconv.Itoa(port)
		u.Host, _ = missed.StringsCutSuffix(u.Host, suffix)
	}
}

// DefaultPort returns the default port number for the scheme.
// For unknown schemes it returns -1.
func DefaultPort(scheme string) int {
	switch scheme {
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
