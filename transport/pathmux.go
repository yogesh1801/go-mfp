// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Path-based HTTP request multiplexer

package transport

import (
	"net/http"
	"sort"
	"strings"
	"sync"
)

// PathMux is an HTTP request multiplexer that makes routing decisions
// based on the URL path.
//
// The [PathMux.Add] method registers a handler for a path and all its
// subpaths. [PathMux.Del] removes a registered handler, and
// [PathMux.Contains] checks if a handler is already installed for
// a given path.
//
// These methods normalize their path parameter before use, making
// "path", "/path" and "foo/../path" equivalent.
//
// Requests processed by the multiplexer have their paths normalized.
// Normalization removes duplicate slashes and processes "." and ".."
// elements, resulting in a clean, canonical path. The trailing slash
// is preserved during this process.
//
// When matching a request path against a pattern, PathMux verifies
// that the request path belongs to the subtree defined by the pattern.
// For example, "/foo" matches "/foo" and "/foo/bar", but not "/foobar".
//
// If a pattern ends with a slash ("/"), the matched request path must
// also include the slash. For example, "/foo/" matches "/foo/1" and
// "/foo/2", but not "/foo" (while "/foo" without trailing slash
// matches all three cases).
//
// In case of multiple matches, the most specific (longest) pattern wins.
//
// If PathMux.Add is called twice with the same pattern, the second call
// overrides the previous handler rather than adding a duplicate.
type PathMux struct {
	mappings []pathMuxMapping
	lock     sync.RWMutex
}

// pathMuxMapping represents a path->handler mapping.
type pathMuxMapping struct {
	path    string
	handler http.Handler
}

// compare compares pathMuxMapping.path with path for sorting.
//
// It either pathMuxMapping.path or path is prefix of another,
// it places the longest path first (so most specific match
// will be listed first).
//
// Otherwise, paths are sorted in the lexicographical order,
// just for predictability.
func (m pathMuxMapping) compare(path string) int {
	switch {
	case m.path == path:
		return 0
	case strings.HasPrefix(m.path, path):
		return -1
	case strings.HasPrefix(path, m.path):
		return 1
	}

	return strings.Compare(m.path, path)
}

// match matches pathMuxMapping against the request path.
func (m pathMuxMapping) match(path string) bool {
	if !strings.HasPrefix(path, m.path) {
		return false
	}

	matched := len(m.path)
	return len(path) == matched ||
		strings.HasSuffix(m.path, "/") ||
		path[matched] == '/'
}

// NewPathMux creates the new [PathMux].
func NewPathMux() *PathMux {
	return &PathMux{}
}

// Add adds path to handler mapping to the [PathMux].
//
// It returns true if the new mapping was added, false if
// the previous mapping was updated.
func (mux *PathMux) Add(path string, handler http.Handler) bool {
	mux.lock.Lock()
	defer mux.lock.Unlock()

	path = CleanURLPath(path)

	n, found := sort.Find(len(mux.mappings), func(i int) int {
		return -mux.mappings[i].compare(path)
	})

	if found {
		mux.mappings[n].handler = handler
		return false
	}

	m := pathMuxMapping{path, handler}
	if n == len(mux.mappings) {
		mux.mappings = append(mux.mappings, m)
	} else {
		mux.mappings = append(mux.mappings, pathMuxMapping{})
		copy(mux.mappings[n+1:], mux.mappings[n:])
		mux.mappings[n] = m
	}

	return true
}

// Del deletes path to handler mapping from the [PathMux].
//
// It returns true if the mapping was found and removed, false otherwise.
func (mux *PathMux) Del(path string) bool {
	mux.lock.Lock()
	defer mux.lock.Unlock()

	path = CleanURLPath(path)

	n, found := sort.Find(len(mux.mappings), func(i int) int {
		return -mux.mappings[i].compare(path)
	})

	if found {
		copy(mux.mappings[n:], mux.mappings[n+1:])
		mux.mappings = mux.mappings[:len(mux.mappings)-1]
	}

	return found
}

// Contains reports if [PathMux] already has mapping for the path
func (mux *PathMux) Contains(path string) bool {
	mux.lock.Lock()
	defer mux.lock.Unlock()

	path = CleanURLPath(path)

	_, found := sort.Find(len(mux.mappings), func(i int) int {
		return -mux.mappings[i].compare(path)
	})

	return found
}

// Empty reports if PathMux contains no registered mappings.
func (mux *PathMux) Empty() bool {
	mux.lock.Lock()
	defer mux.lock.Unlock()

	return len(mux.mappings) == 0
}

// ServeHTTP dispatches the request to the handler, based
// on the request URI path.
func (mux *PathMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Canonicalize URL path
	r.URL.Path = CleanURLPath(r.URL.Path)

	// Lookup and call the handler
	handler := mux.handler(r.URL.Path)
	if handler != nil {
		handler.ServeHTTP(w, r)
		return
	}

	query := NewServerQuery(w, r)
	query.Reject(http.StatusNotFound, nil)
}

// handler returns the request handler, based on the request URL path.
func (mux *PathMux) handler(path string) http.Handler {
	mux.lock.RLock()
	defer mux.lock.RUnlock()

	for _, m := range mux.mappings {
		if m.match(path) {
			return m.handler
		}
	}

	return nil
}
