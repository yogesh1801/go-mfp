// MFP      - Miulti-Function Printers and scanners toolkit
// mainfunc - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common paths

package commands

import (
	"os/user"
	"path/filepath"
	"strings"
)

var (
	// pathHomeDir contains home directory of the calling user
	pathHomeDir string
)

// PathHomeDir returns home directory of the calling user
func PathHomeDir() string {
	return pathHomeDir
}

// PathUserConfDir returns user configuration directory
// for the program
func PathUserConfDir(program string) string {
	return filepath.Join(pathHomeDir, "."+strings.ToLower(program))
}

// init initializes paths
func init() {
	user, err := user.Current()
	if err != nil {
		panic(err.Error())
	}

	pathHomeDir = user.HomeDir
}
