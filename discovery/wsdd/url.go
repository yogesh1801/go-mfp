// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// URLs handling

package wsdd

import (
	"net"
	"net/netip"
	"net/url"
	"strings"
)

// urlParse parses and validates HTTP URL.
// On success, it returns parsed URL. Otherwise, it returns nil.
//
// URL, to be accepted, needs to fit the following criteria:
//  1. It must be syntactically correct
//  2. It must have http or https scheme
//  3. It must have explicit hostname
func urlParse(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		return nil
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return nil
	}

	if u.Host == "" {
		return nil
	}

	return u
}

// urlIsLiteral reports if URL has a literal IP address
func urlIsLiteral(u *url.URL) bool {
	_, err := netip.ParseAddr(u.Hostname())
	return err == nil
}

// urlIsLiteral reports if URL has a literal IP4 address
func urlIsIP4(u *url.URL) bool {
	addr, err := netip.ParseAddr(u.Hostname())
	return err == nil && addr.Unmap().Is4()
}

// urlIsLiteral reports if URL has a literal IP6 address
func urlIsIP6(u *url.URL) bool {
	addr, err := netip.ParseAddr(u.Hostname())
	return err == nil && addr.Unmap().Is6()
}

// urlWithZone adds or replaces IP6 zone suffix in the URL's hostname,
// if hostname is the link-local IPv6 literal address. If zone is empty,
// the zone is removed.
//
// Otherwise, it returns the original URL.
func urlWithZone(u *url.URL, zone string) *url.URL {
	addr, err := netip.ParseAddr(u.Hostname())
	if err == nil {
		u = urlShallowCopy(u)

		if addr.Is6() && addr.IsLinkLocalUnicast() {
			addr = addr.WithZone(zone)
		}

		host := addr.String()
		if port := u.Port(); port != "" {
			host = net.JoinHostPort(host, port)
		} else if strings.IndexByte(host, ':') >= 0 {
			host = "[" + host + "]"
		}

		u.Host = host
	}
	return u
}

// urlZone extracts zone out of the IP6 literal URL.
// If URL is not literal, or not IP6 link-local unicast,
// the empty string is returned.
func urlZone(u *url.URL) (zone string) {
	addr, err := netip.ParseAddr(u.Hostname())
	if err == nil && addr.Is6() && addr.IsLinkLocalUnicast() {
		zone = addr.Zone()
	}
	return
}

// urlShallowCopy creates a shallow copy of the URL
func urlShallowCopy(u *url.URL) *url.URL {
	u2 := *u
	return &u2
}
