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

// Forward performs URL translation in the forward (remote->local) direction.
func (ux *URLXlat) Forward(u *url.URL) *url.URL {
	return ux.translate(u, ux.local, ux.remote)
}

// Reverse performs URL translation in the reverse (local->remote) direction.
func (ux *URLXlat) Reverse(u *url.URL) *url.URL {
	return ux.translate(u, ux.remote, ux.local)
}

// ForwardPath translates Path part of the URL in the forward
// (local->remote) direction.
func (ux *URLXlat) ForwardPath(path string) string {
	pathOut, _ := ux.translatePath(path, ux.local, ux.remote)
	return pathOut
}

// ReversePath translates Path part of the URL in the reverse
// (remote->local) direction.
func (ux *URLXlat) ReversePath(path string) string {
	pathOut, _ := ux.translatePath(path, ux.remote, ux.local)
	return pathOut
}

// translate translates URL u in the (from->to) direction,
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
	pathOut, ok := ux.translatePath(u.Path, from, to)
	if !ok {
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

// translatePath translates path part of the URL in the (from->to)
// direction.
func (ux *URLXlat) translatePath(pathIn string,
	from, to *url.URL) (pathOut string, ok bool) {

	// Input path must be prefixed by the path we are
	// translating from (from.Path).
	//
	// If this is true, we replace the prefix with the
	// path we are translating to (to.Path).
	switch {
	case !strings.HasPrefix(pathIn, from.Path):
		// u.Path must be prefixed by from.Path.
		// Otherwise, don't translate
		pathOut = pathIn

	case pathIn == from.Path:
		pathOut = to.Path
		ok = true

	case strings.HasSuffix(from.Path, "/") || pathIn[len(from.Path)] == '/':
		// if pathIn is longer that from.Path, they must
		// diverge at the path separator
		//
		// Translate pathIn at this case
		pathOut = path.Join(to.Path, pathIn[len(from.Path):])
		ok = true

	default:
		// Otherwise, don't translate
		pathOut = pathIn
	}

	return
}
