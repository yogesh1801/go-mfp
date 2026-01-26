// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversion HTTP types to Python

package modeling

import (
	"net/http"
	"sort"

	"github.com/OpenPrinting/go-mfp/cpython"
)

// httpHeaderToPython converts [http.Header] to [cpython.Object].
// The returned object will be of the http.client.HTTPMessage type.
func (model *Model) httpHeaderToPython(h http.Header) *cpython.Object {
	// Obtain and sort header lines.
	//
	// Note, Go maps are unordered and map iteration order is
	// not deterministic, but Python dictionaries are ordered and
	// it is nice to see keys at least at the deterministic order
	// at that side.
	type hdrline struct{ name, val string }
	hdrlines := make([]hdrline, 0, len(h))

	for name, values := range h {
		for _, val := range values {
			hdrlines = append(hdrlines,
				hdrline{name: name, val: val})
		}
	}

	sort.SliceStable(hdrlines, func(i, j int) bool {
		return hdrlines[i].name < hdrlines[j].name
	})

	// Create and populate the target Python Object
	obj := model.clsHTTPMessage.Call()
	for _, line := range hdrlines {
		err := obj.Set(line.name, line.val)
		if err != nil {
			return model.py.NewError(err)
		}
	}

	return obj
}

// httpHeaderToPython converts [cpython.Object] into [http.Header].
// The Python object must be of the http.client.HTTPMessage type.
func (model *Model) httpHeaderFromPython(obj *cpython.Object) (
	http.Header, error) {

	// Obtain keys
	keys, err := obj.Keys()
	if err != nil {
		return nil, err
	}

	// Extract headers
	h := make(http.Header, len(keys))
	for _, key := range keys {
		if !key.IsUnicode() {
			// It should not be, but just in case...
			continue
		}

		val := obj.Get(key)
		if err := val.Err(); err != nil {
			return nil, err
		}

		if !val.IsUnicode() {
			// Also very unlikely...
			continue
		}

		keystr, err := key.Unicode()
		if err != nil {
			return nil, err
		}

		valstr, err := val.Unicode()
		if err != nil {
			return nil, err
		}

		h.Add(keystr, valstr)
	}

	return h, nil
}
