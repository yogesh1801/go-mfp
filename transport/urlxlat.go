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

// Local returns the local-side URL
func (ux *URLXlat) Local() *url.URL {
	return ux.local
}

// Remote returns the remote-side URL
func (ux *URLXlat) Remote() *url.URL {
	return ux.remote
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

	// Translate path.
	//
	// Input path must be prefixed by the path we are
	// translating from (from.Path).
	//
	// If this is true, we replace the prefix with the
	// path we are translating to (to.Path).
	pathIn := u.Path
	pathOut := ""

	switch {
	case !strings.HasPrefix(pathIn, from.Path):
		// u.Path must be prefixed by from.Path
		return u

	case pathIn == from.Path:
		pathOut = to.Path

	case from.Path == "/" || pathIn[len(from.Path)] == '/':
		// if pathIn is longer that from.Path, they must
		// diverge at the path separator
		//
		// Translate pathIn at this case
		pathOut = path.Join(to.Path, pathIn[len(from.Path):])

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
	u.Path = pathOut

	URLStripPort(u)

	return u
}
