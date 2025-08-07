// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Path-based HTTP request multiplexer test

package transport

import (
	"net/http"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/generic"
)

// TestPathMuxMappingCompare tests the pathMuxMapping.compare function.
func TestPathMuxMappingCompare(t *testing.T) {
	type testData struct {
		mappingPath string // pathMuxMapping.path
		path        string // path to compare
		expected    int    // expected result
	}

	tests := []testData{
		{mappingPath: "", path: "", expected: 0},
		{mappingPath: "aaa", path: "aaa", expected: 0},
		{mappingPath: "aaa", path: "bbb", expected: -1},
		{mappingPath: "bbb", path: "aaa", expected: 1},
		{mappingPath: "123", path: "12345", expected: 1},
		{mappingPath: "12345", path: "123", expected: -1},
	}

	for _, test := range tests {
		m := pathMuxMapping{path: test.mappingPath}
		c := m.compare(test.path)
		if c != test.expected {
			t.Errorf("pathMuxMapping{%q}.compare(%q):\n"+
				"expected: %d\n"+
				"present:  %d\n",
				test.mappingPath, test.path, test.expected, c)
		}
	}
}

// TestPathMuxAdd tests the PathMux.Add function.
func TestPathMuxAdd(t *testing.T) {
	type testData struct {
		paths    []string // Added paths
		expected []string // Expected PathMux.mappings
	}

	tests := []testData{
		{
			paths:    []string{"a", "b", "c"},
			expected: []string{"/a", "/b", "/c"},
		},
		{
			paths:    []string{"a", "c", "b"},
			expected: []string{"/a", "/b", "/c"},
		},
		{
			paths:    []string{"b", "a", "c"},
			expected: []string{"/a", "/b", "/c"},
		},
		{
			paths:    []string{"b", "c", "a"},
			expected: []string{"/a", "/b", "/c"},
		},
		{
			paths:    []string{"c", "a", "b"},
			expected: []string{"/a", "/b", "/c"},
		},
		{
			paths:    []string{"c", "b", "a"},
			expected: []string{"/a", "/b", "/c"},
		},
		{
			paths:    []string{"a", "b", "c", "c", "b", "a"},
			expected: []string{"/a", "/b", "/c"},
		},
		{
			paths:    []string{"123", "12345"},
			expected: []string{"/12345", "/123"},
		},
		{
			paths:    []string{"12345", "123"},
			expected: []string{"/12345", "/123"},
		},
	}

	for _, test := range tests {
		mux := NewPathMux()
		added := generic.NewSet[string]()

		for _, p := range test.paths {
			ok := mux.Add(p, nil)
			mustok := !added.Contains(p)
			added.Add(p)

			if ok != mustok {
				t.Errorf("PathMux.Add(%q):\n"+
					"expected: %v\n"+
					"present:  %v\n",
					p, mustok, ok)
			}
		}

		present := []string{}
		for _, m := range mux.mappings {
			present = append(present, m.path)
		}

		diff := testutils.Diff(test.expected, present)
		if diff != "" {
			t.Errorf("PathMux.Add:\n%s", diff)
		}
	}
}

// TestPathMuxDel tests the PathMux.Del function.
func TestPathMuxDel(t *testing.T) {
	type testData struct {
		add      []string // Added paths
		del      []string // Added paths
		expected []string // Expected PathMux.mappings
	}

	tests := []testData{
		{
			add:      []string{"a", "b", "c"},
			del:      []string{"a"},
			expected: []string{"/b", "/c"},
		},

		{
			add:      []string{"a", "b", "c"},
			del:      []string{"b"},
			expected: []string{"/a", "/c"},
		},

		{
			add:      []string{"a", "b", "c"},
			del:      []string{"c"},
			expected: []string{"/a", "/b"},
		},

		{
			add:      []string{"a", "b", "c"},
			del:      []string{"a", "b", "c"},
			expected: []string{},
		},
	}

	for _, test := range tests {
		mux := NewPathMux()
		added := generic.NewSet[string]()

		for _, p := range test.add {
			mux.Add(p, nil)
			added.Add(p)
		}

		for _, p := range test.del {
			ok := mux.Del(p)
			mustok := added.Contains(p)
			added.Del(p)

			if ok != mustok {
				t.Errorf("PathMux.Del(%q):\n"+
					"expected: %v\n"+
					"present:  %v\n",
					p, mustok, ok)
			}
		}

		present := []string{}
		for _, m := range mux.mappings {
			present = append(present, m.path)
		}

		diff := testutils.Diff(test.expected, present)
		if diff != "" {
			t.Errorf("PathMux.Del:\n%s", diff)
		}
	}
}

// TestPathMuxServeHTTP tests the PathMux.ServeHTTP function.
func TestPathMuxServeHTTP(t *testing.T) {
	lastpath := ""
	handler := func(path string) http.Handler {
		callback := func(http.ResponseWriter, *http.Request) {
			lastpath = path
		}
		return http.HandlerFunc(callback)
	}

	mux := NewPathMux()

	add := func(path string) {
		mux.Add(path, handler(path))
	}

	add("/")
	add("/ipp")
	add("/ipp/print")
	add("/ipp/faxout")
	add("/eSCL")

	check := func(in, out string) {
		rq, err := http.NewRequest("GET", "http://localhost", nil)
		assert.NoError(err)

		rq.URL.Path = in
		lastpath = ""
		mux.ServeHTTP(nil, rq)

		if lastpath != out {
			t.Errorf("PathMux.ServeHTTP(%q):\n"+
				"expected: %q\n"+
				"present:  %q\n",
				in, out, lastpath)
		}
	}

	check("/", "/")
	check("/ip", "/")
	check("/ipp", "/ipp")
	check("/ippa", "/")
	check("/ipp/print", "/ipp/print")
	check("/ipp/print/", "/ipp/print")
	check("/ipp/print/1", "/ipp/print")
	check("/ipp/faxout", "/ipp/faxout")
	check("/ipp/faxout/1", "/ipp/faxout")
}
