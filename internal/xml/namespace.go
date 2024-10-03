// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML namespace

package xml

import "strings"

// Namespace maps XML namespace URLs to short prefixes.
//
// Namespace initialization may look as follows:
//
//	var ns = Namespace{
//	        {"http://www.w3.org/2003/05/soap-envelope", "s"},
//	        {"https://www.w3.org/2003/05/soap-envelope", "s"},
//	        {"http://schemas.xmlsoap.org/ws/2005/04/discovery", "d"},
//	        {"https://schemas.xmlsoap.org/ws/2005/04/discovery", "d"},
//	}
//
// When used with the [Decode] function, this namespace specifies
// that names all elements and attributes whose prefix in the original
// document maps to the "http://www.w3.org/2003/05/soap-envelope" URL
// will be rewritten to use prefix "s" and so on.
//
// When used with the [Encode] function, Namespace will be automatically
// written to output as a set of root element's xmlns attributes. Moreover,
// only actually used entries will be written to output.
//
// If Namespace contains multiple entries with the same Prefix and different
// URLs, all entries will be taken in account by Decode, and Encode will
// use the first matching entry in the list.
type Namespace []struct {
	URL    string // Namespace URL
	Prefix string // Namespace prefix
}

// Append appends item to the Namespace
func (ns *Namespace) Append(url string, prefix string) {
	item := struct{ URL, Prefix string }{url, prefix}
	*ns = append(*ns, item)
}

// ByURL searches Namespace by URL.
//
// It returns (Prefix, true) if requested element was found,
// or ("", false) otherwise.
func (ns Namespace) ByURL(u string) (string, bool) {
	for _, ent := range ns {
		if u == ent.URL {
			return ent.Prefix, true
		}
	}

	return "", false
}

// ByPrefix searches Namespace by name prefix.
//
// It returns (URK, true) if requested element was found,
// or ("", false) otherwise.
func (ns Namespace) ByPrefix(p string) (string, bool) {
	for _, ent := range ns {
		if p == ent.Prefix {
			return ent.URL, true
		}
	}

	return "", false
}

// nsPrefix returns namespace prefix for the name
//
// The second returned value will be false, if name doesn't contain
// a namespace prefix.
func nsPrefix(name string) (string, bool) {
	i := strings.IndexByte(name, ':')
	if i >= 0 {
		return name[:i], true
	}

	return "", false
}
