// MFP  - Miulti-Function Printers and scanners toolkit
// DEST - Destination URLs hanling
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP-specific URL parsing

package dest

import (
	"errors"
	"net/url"
)

var (
	// URL errors
	errInvalidURL       = errors.New("Printer URL: invalid URL")
	errInvalidURLScheme = errors.New("Printer URL: scheme must be ipp or ipps")
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

// urlParse parses IPP URL ("ipp://..." or "ipps://...")
//
// It returns parsed HTTP URL (with "http" or "https" scheme"),
// normalized URL string (with IPP scheme) or an error
//
// URL normalization implies resolving absolute path per RFC 3986,
// removing port when not needed and so on.
func urlParse(printerURL string) (*url.URL, string, error) {
	// Parse URL
	parsedURL, err := url.Parse(printerURL)
	if err != nil {
		return nil, "", errInvalidURL
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, "", errInvalidURL
	}

	// Normalize URL
	parsedURL = parsedURL.ResolveReference(parsedURL)

	// Adjust HTTP URL
	httpURL := *parsedURL
	switch httpURL.Scheme {
	case "ipp":
		httpURL.Scheme = "http"
	case "ipps":
		httpURL.Scheme = "https"
	default:
		err = errInvalidURLScheme
		return nil, "", err
	}

	switch {
	case httpURL.Port() == "":
		httpURL.Host += ":631"
	case httpURL.Scheme == "http" && httpURL.Port() == "80":
		httpURL.Host = httpURL.Hostname()
	case httpURL.Scheme == "https" && httpURL.Port() == "443":
		httpURL.Host = httpURL.Hostname()
	}

	// Format normalized URL
	if parsedURL.Port() == "631" {
		parsedURL.Host = parsedURL.Hostname()
	}
	norm := parsedURL.String()

	return &httpURL, norm, nil
}
