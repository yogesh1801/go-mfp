// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML iterator

package xmldoc

import "bytes"

// Iter allows iteration of the XML tree as a linear sequence of elements.
//
// Assuming the tree has the following structure:
//
//	Root
//	  |--> Child 1
//	  |      |--> Child 1.1
//	  |      `--> Child 1.2
//	  `--> Child 2
//	         |--> Child 2.1
//	         `--> Child 2.2
//
// The nodes will be returned in the following order:
//
//   - Root
//   - Child 1
//   - Child 1.1
//   - Child 1.2
//   - Child 2
//   - Child 2.1
//   - Child 2.2
type Iter struct {
	stack    []iterStackLevel
	pathname bytes.Buffer
}

func (i *Iter) stackTop() *iterStackLevel {
	return &i.stack[len(i.stack)-1]
}

type iterStackLevel struct {
	elements []Element
	pathlen  int
}

// Iterate begins iteration of the XML Element tree.
//
// The newly created iterator points to the dummy pseudo-element.
// After Iter.Next is called for the very first time, the current
// node becomes root.
//
// So the valid usage pattern is following
//
//	iter := root.Iterate()
//	for iter.Next() {
//	  // Do something
//	}
//
// Important note: the structure of the XML tree MUST NOT be
// modified during iterations. Otherwise, behavior is undefined.
func (root Element) Iterate() *Iter {
	elements := []Element{
		Element{},
		root,
	}

	i := &Iter{
		stack: []iterStackLevel{{elements, 0}},
	}

	return i
}

// Next moves iterator to the next Element in order.
// It returns false at the end.
func (i *Iter) Next() bool {
	for len(i.stack) > 0 {
		level := i.stack[len(i.stack)-1]
		cur := level.elements[0]
		tail := level.elements[1:]

		switch {
		case len(cur.Children) > 0:
			// Enter into the current element
			top := i.stackTop()
			top.elements = tail

			i.stack = append(i.stack,
				iterStackLevel{
					cur.Children, i.pathname.Len(),
				})

			cur = cur.Children[0]
			i.pathname.WriteByte('/')
			i.pathname.WriteString(cur.Name)
			return true

		case len(tail) > 0:
			// Switch to the next element
			top := i.stackTop()
			top.elements = tail

			cur = tail[0]

			i.pathname.Truncate(top.pathlen)
			i.pathname.WriteByte('/')
			i.pathname.WriteString(cur.Name)
			return true

		default:
			// Move up the stack
			depth := len(i.stack) - 1
			for depth > 0 && len(i.stack[depth-1].elements) == 0 {
				depth--
			}
			i.stack = i.stack[:depth]
			if len(i.stack) > 0 {
				i.pathname.Truncate(i.stack[depth-1].pathlen)
			} else {
				i.pathname.Truncate(1)
			}
		}
	}

	return false
}

// Done reports if iteration is done.
func (i *Iter) Done() bool {
	return len(i.stack) == 0
}

// Elem returns the current element.
func (i *Iter) Elem() *Element {
	if !i.Done() {
		return &i.stack[len(i.stack)-1].elements[0]
	}
	return nil
}

// Path returns a full path from root to the current element.
// Path starts with the '/' character and uses '/' as a separator.
func (i *Iter) Path() string {
	return i.pathname.String()
}
