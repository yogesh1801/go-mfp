// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Action -- contains parsed Command's arguments.

package argv

// Action defines action to be taken when Command is
// applied to the command line.
type Action struct {
	options map[string][]string
}

// Getopt returns value of option or parameter as a single string.
//
// For multi-value options and repeated parameters values
// are concatenated into the single string using CSV encoding.
func (act *Action) Getopt(name string) (val string, found bool) {
	return "", false
}

// GetoptSlice returns value of option or parameter as a slice of string.
// If option is not found, it returns nil
func (act *Action) GetoptSlice(name string) (val []string) {
	return nil
}
