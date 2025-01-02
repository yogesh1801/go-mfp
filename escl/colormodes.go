// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of scan color modes

package escl

import "strings"

// ColorModes contains a set (bitmask) of [ColorMode]s.
type ColorModes int

// MakeColorModes makes [ColorModes] from the list of [ColorMode]s.
func MakeColorModes(list ...ColorMode) ColorModes {
	var cmodes ColorModes

	for _, cm := range list {
		cmodes.Add(cm)
	}

	return cmodes
}

// String returns a string representation of the [ColorModes],
// for debugging.
func (cmodes ColorModes) String() string {
	modes := [...]ColorMode{BlackAndWhite1, Grayscale8, Grayscale16,
		RGB24, RGB48}
	s := make([]string, 0, len(modes))

	for _, cm := range modes {
		if cmodes.Contains(cm) {
			s = append(s, cm.String())
		}
	}

	return strings.Join(s, ",")
}

// Add adds [ColorMode] to the set.
func (cmodes *ColorModes) Add(cm ColorMode) *ColorModes {
	*cmodes |= 1 << cm
	return cmodes
}

// Del deletes [ColorMode] from the set.
func (cmodes *ColorModes) Del(cm ColorMode) *ColorModes {
	*cmodes &^= 1 << cm
	return cmodes
}

// Contains reports if [ColorMode] exists in the set.
func (cmodes ColorModes) Contains(cm ColorMode) bool {
	return (cmodes & (1 << cm)) != 0
}
