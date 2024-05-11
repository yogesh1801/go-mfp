// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// String functions

package argv

// strCommonPrefix returns a common prefix of two strings
func strCommonPrefix(s1, s2 string) string {
	maxlen := len(s1)
	if len(s2) < maxlen {
		maxlen = len(s2)
	}

	pfxlen := 0
	for pfxlen < maxlen && s1[pfxlen] == s2[pfxlen] {
		pfxlen++
	}

	return s1[:pfxlen]
}

// strCommonPrefixSlice returns common prefix of the slice of strings
func strCommonPrefixSlice(slice []string) (pfx string) {
	if len(slice) > 0 {
		// Note, if we choose the "minimal" and "maximal" candidates,
		// in lexicographical order, the longest common prefix of all
		// candidates will be the longest common prefix of these two
		// strings

		min := slice[0]
		max := slice[0]

		for _, s := range slice[1:] {
			if s < min {
				min = s
			}

			if s > max {
				max = s
			}
		}

		pfx = strCommonPrefix(min, max)
	}

	return
}
