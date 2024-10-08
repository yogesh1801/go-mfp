// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML element

package xml

import (
	"os"
	"slices"
)

// Element represents a single decoded XML Element
type Element struct {
	Name     string    // Name of this element (ns:name)
	Text     string    // Element body
	Attrs    []Attr    // Element attributes
	Children []Element // All children
}

// Attr represents an XML element attribute
type Attr struct {
	Name  string // Attribute name
	Value string // Attribute value
}

// Expand replaces ${var} or $var in the XML [Element.Text] and
// [Attr.Value] of the entire XML tree based on the mapping
// function and returns rewritten XML tree.
//
// It is uses [os.Expand] function for strings substitution.
func (root Element) Expand(mapping func(string) string) Element {
	root.Text = os.Expand(root.Text, mapping)
	root.Attrs = slices.Clone(root.Attrs)
	root.Children = slices.Clone(root.Children)

	for i := range root.Attrs {
		attr := &root.Attrs[i]
		attr.Value = os.Expand(attr.Value, mapping)
	}

	for i := range root.Children {
		root.Children[i] = root.Children[i].Expand(mapping)
	}

	return root
}

// AttrByName searches element attributes by name.
//
// It returns (attr, true) if element was found or (Attr{}, false) if not.
func (root Element) AttrByName(name string) (Attr, bool) {
	for _, attr := range root.Attrs {
		if attr.Name == name {
			return attr, true
		}
	}

	return Attr{}, false
}

// AttrsMap returns map of the Element's attributes, indexed by name.
//
// If there are multiple attributes with the same name,
// only the first is returned.
func (root Element) AttrsMap() map[string]Attr {
	attrs := make(map[string]Attr)

	for _, attr := range root.Attrs {
		if _, found := attrs[attr.Name]; !found {
			attrs[attr.Name] = attr
		}
	}

	return attrs
}

// ChildByName searches child element by name.
//
// It returns (child, true) if element was found or (Element{}, false) if not.
func (root Element) ChildByName(name string) (Element, bool) {
	for _, elm := range root.Children {
		if elm.Name == name {
			return elm, true
		}
	}

	return Element{}, false
}

// ChildrenMap returns map of the Element's children, indexed by name.
//
// If there are multiple children with the same name,
// only the first is returned.
func (root Element) ChildrenMap() map[string]Element {
	children := make(map[string]Element)

	for _, elm := range root.Children {
		if _, found := children[elm.Name]; !found {
			children[elm.Name] = elm
		}
	}

	return children
}

// Lookup searches for children elements, multiple elements a time.
//
//	Use pattern:
//
//	  element1 = Lookup(Name: "Element1"}
//	  element2 = Lookup(Name: "Element2"}
//	  element3 = Lookup(Name: "Element3"}
//
//	  root.Lookup(&element1,&element2,&element3)
//	  if element1.Found{
//	      . . .
//	  }
//	  if element2.Found{
//	      . . .
//	  }
//	  if element2.Found{
//	      . . .
//	  }
//
// If root contains multiple children with the same name, the
// first child always returned.
//
// It returns the first not found Lookup element with the
// [Lookup.Required] flag set, or nil of all required elements
// are found.
//
// Note, if there are missed required elements, it is not
// guaranteed that all lookup requests will be processed.
func (root Element) Lookup(lookups ...*Lookup) *Lookup {
	if len(lookups)*len(root.Children) <= 16 {
		// If we have a small amount of Lookup/Children combination,
		// just go straightforward, it will be less expensive, that
		// doing via maps of children
		for _, l := range lookups {
			l.Elem, l.Found = root.ChildByName(l.Name)
			if !l.Found && l.Required {
				return l
			}
		}
	} else {
		// Otherwise, obtain map of children and then fill the answer.
		children := root.ChildrenMap()
		for _, l := range lookups {
			l.Elem, l.Found = children[l.Name]
			if !l.Found && l.Required {
				return l
			}
		}
	}

	return nil
}

// LookupAttrs is like [Element.Lookup], but for attributes, not for
// children elements.
func (root Element) LookupAttrs(lookups ...*LookupAttr) *LookupAttr {
	if len(lookups)*len(root.Attrs) <= 16 {
		// If we have a small amount of Lookup/Attrs combination,
		// just go straightforward, it will be less expensive, that
		// doing via maps of attributes
		for _, l := range lookups {
			l.Attr, l.Found = root.AttrByName(l.Name)
			if !l.Found && l.Required {
				return l
			}
		}
	} else {
		// Otherwise, obtain map of attributes and then fill the answer.
		attrs := root.AttrsMap()
		for _, l := range lookups {
			l.Attr, l.Found = attrs[l.Name]
			if !l.Found && l.Required {
				return l
			}
		}
	}

	return nil
}
