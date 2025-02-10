// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// URL Proxying

package transport

import (
	"net/url"
	"path"
	"strings"
)

// URLXlat performs HTTP URL translation for proxying purpose.
type URLXlat struct {
	local, remote *url.URL // Local and remote base URLs
}

// NewURLXlat creates a new URLXlat
func NewURLXlat(local, remote *url.URL) *URLXlat {
	local = URLClone(local)
	remote = URLClone(remote)

	URLForcePort(local)
	URLForcePort(remote)

	return &URLXlat{local, remote}
}

// Forward performs URL translation into local->remote direction
func (ux *URLXlat) Forward(u *url.URL) *url.URL {
	return ux.translate(u, ux.local, ux.remote)
}

// Reverse performs URL translation into reverse->local direction
func (ux *URLXlat) Reverse(u *url.URL) *url.URL {
	return ux.translate(u, ux.remote, ux.local)
}

// translate provides an actual translation.
func (ux *URLXlat) translate(u, from, to *url.URL) *url.URL {
	// Match schemes
	switch {
	case (u.Scheme == "http" || u.Scheme == "ipp" || u.Scheme == "unix") &&
		(from.Scheme == "http" || from.Scheme == "ipp" || from.Scheme == "unix"):
	case (u.Scheme == "https" || u.Scheme == "ipps") &&
		(from.Scheme == "https" || from.Scheme == "ipps"):

	default:
		// Schemes mismatch, don't translate
		return u
	}

	// Match host names
	if u.Hostname() != from.Hostname() {
		// Host names mismatch, don't translate
		return u
	}

	// Match ports
	if URLPort(u) != URLPort(from) {
		// Ports mismatch, don't translate
		return u
	}

	pathIn := u.Path
	pathFrom := from.Path

	switch {
	case !strings.HasPrefix(pathIn, pathFrom):
		// u.Path must be prefixed by from.Path
		return u

	case pathIn == pathFrom:
	case pathFrom == "/" || pathIn[len(pathFrom)] == '/':
		// if pathIn is longer that pathFrom, they must
		// diverge at the path separator
		//
		// Translate pathIn at this case
		pathIn = pathIn[len(pathFrom):]

	default:
		// Otherwise, don't translate
		return u
	}

	// Perform a translation
	u = URLClone(u)
	if to.Scheme == "unix" {
		u.Scheme = to.Scheme
		u.Host = ""
		u.OmitHost = true
	} else if u.Scheme == "unix" {
		u.Scheme = to.Scheme
	}

	u.User = to.User
	u.Host = to.Host
	u.Path = path.Join(to.Path, pathIn)

	URLStripPort(u)

	return u
}
