// MFP  - Miulti-Function Printers and scanners toolkit
// DEST - Internet Printing Protocol implementation
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
