// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML namespace

package xml

import (
	"slices"
	"strings"
)

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
	Used   bool   // Flag for [Namespace.MarkUsed] and friends
}

// Clone creates a copy of [Namespace].
func (ns Namespace) Clone() Namespace {
	return slices.Clone(ns)
}

// MarkUsed marks Namespace entries that actually used by
// the XML tree by setting [Namespace.Used] flag.
func (ns Namespace) MarkUsed(root Element) {
	// Extract all prefixes without duplicates
	prefixes := make(map[string]struct{})

	iter := root.Iterate()
	for iter.Next() {
		elem := iter.Elem()

		prefix, ok := nsPrefix(elem.Name)
		if ok {
			prefixes[prefix] = struct{}{}
		}

		for _, attr := range elem.Attrs {
			prefix, ok = nsPrefix(attr.Name)
			if ok {
				prefixes[prefix] = struct{}{}
			}
		}
	}

	// Mark all used prefixes
	for i := range ns {
		ent := &ns[i]
		if _, found := prefixes[ent.Prefix]; found {
			ent.Used = true

			// Namespace may contain multiple entries
			// with the same Prefix. Mark only first
			// of them; others are only for decoding.
			delete(prefixes, ent.Prefix)
		}
	}
}

// MarkUsedName accepts XML element name (prefix:name) and
// marks Namespace entry that this prefix refers to with
// the [Namespace.Used] flag.
func (ns Namespace) MarkUsedName(name string) {
	prefix, ok := nsPrefix(name)
	if ok {
		for i := range ns {
			ent := &ns[i]
			if ent.Prefix == prefix {
				ent.Used = true
				return
			}
		}
	}
}

// ExportUsed exports marked as used subset of the Namespace
// as a sequence of xmlns attributes.
func (ns Namespace) ExportUsed() []Attr {
	attrs := make([]Attr, 0, len(ns))

	for _, ent := range ns {
		if ent.Used {
			attrs = append(attrs, Attr{
				Name:  "xmlns:" + ent.Prefix,
				Value: ent.URL,
			})
		}
	}

	return attrs
}

// Append appends item to the Namespace
func (ns *Namespace) Append(url string, prefix string) {
	item := struct {
		URL, Prefix string
		Used        bool
	}{url, prefix, false}
	*ns = append(*ns, item)
}

// ByURL searches Namespace by namespace URL.
//
// It returns (Prefix, true) if requested element was found,
// or ("", false) otherwise.
//
// See also [Namespace.IndexByURL]
func (ns Namespace) ByURL(u string) (string, bool) {
	if i := ns.IndexByURL(u); i >= 0 {
		return ns[i].Prefix, true
	}

	return "", false
}

// IndexByURL searches Namespace by namespace URL.
//
// It returns index of the found element or -1, if there is
// no match.
//
// See also: [Namespace.ByURL]
func (ns Namespace) IndexByURL(u string) int {
	for i, ent := range ns {
		if nsEqualURLs(u, ent.URL) {
			return i
		}
	}

	return -1
}

// ByPrefix searches Namespace by name prefix.
//
// It returns (URK, true) if requested element was found,
// or ("", false) otherwise.
func (ns Namespace) ByPrefix(p string) (string, bool) {
	if i := ns.IndexByPrefix(p); i >= 0 {
		return ns[i].URL, true
	}

	return "", false
}

// IndexByPrefix searches Namespace by name prefix.
//
// It returns index of the found element or -1, if there is
// no match.
//
// See also: [Namespace.ByPrefix]
func (ns Namespace) IndexByPrefix(p string) int {
	for i, ent := range ns {
		if p == ent.Prefix {
			return i
		}
	}

	return -1
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

// nsEqualURLs reports if two namespace URLs are equal.
// It ignores the difference between http: and https: schemes.
func nsEqualURLs(u1, u2 string) bool {
	// Equal strings are equal URLs
	if u1 == u2 {
		return true
	}

	// Canonicalize schemes: replace "https:" with "http"
	if strings.HasPrefix(u1, "https:") {
		u1 = "http" + u1[5:]
	}

	if strings.HasPrefix(u2, "https:") {
		u2 = "http" + u2[5:]
	}

	return u1 == u2
}
