// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Filesystem paths completer

package argv

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// fscompleter performs pathname completion on a top of fs.FS
type fscompleter struct {
	fsys  fs.FS                  // Underlying file system
	getwd func() (string, error) // Get working directory
}

// newFscompleter creates new fscompleter
func newFscompleter(fsys fs.FS, getwd func() (string, error)) *fscompleter {
	if getwd == nil {
		getwd = func() (string, error) {
			return string(filepath.Separator), nil
		}
	}

	return &fscompleter{
		fsys:  fsys,
		getwd: getwd,
	}
}

// complete performs filesystem paths completion and returns
// completion candidates.
func (fscompl *fscompleter) complete(arg string) (compl []Completion) {
	// Split argument into directory and relative path
	dir, file := fscompl.splitPath(arg)
	_ = file

	// Read the directory
	entries, err := fscompl.readDir(dir)
	if err != nil {
		return nil
	}

	// Match file name against the directory entries
	var lasIsDir bool

	for _, ent := range entries {
		name := ent.Name()
		lasIsDir = ent.IsDir()

		if strings.HasPrefix(name, file) {
			candidate := fscompl.mergePath(dir, name)
			compl = append(compl, Completion{candidate, 0})
		}
	}

	if len(compl) == 1 && lasIsDir {
		compl[0].String += string(filepath.Separator)
		compl[0].Flags = CompletionNoSpace
	}

	return compl
}

func (fscompl *fscompleter) readDir(dir string) ([]fs.DirEntry, error) {
	// Obtain absolute path
	absdir, err := fscompl.absPath(dir)
	if err != nil {
		return nil, err
	}

	// Read the directory
	entries, err := fs.ReadDir(fscompl.fsys, absdir)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// splitPath splits the path into directory and file
func (fscompl fscompleter) splitPath(path string) (dir, file string) {
	i := len(path) - 1
	for i >= 0 && !os.IsPathSeparator(path[i]) {
		i--
	}

	switch {
	case i < 0:
		return "", path
	case i == 0:
		return string(filepath.Separator), path[1:]
	case i == len(path)-1:
		return path[0:i], ""
	default:
		return path[0:i], path[i+1:]
	}
}

// mergePath joins directory prefix with file in that directory.
// Unlike filepath.Join, it doesn't call filepath.Clean() on result.
func (fscompl fscompleter) mergePath(dir, file string) string {
	if dir != "" && !os.IsPathSeparator(dir[len(dir)-1]) {
		dir += string(filepath.Separator)
	}
	return dir + file
}

// absPath makes an absolute path
func (fscompl *fscompleter) absPath(path string) (string, error) {
	cwd, err := fscompl.getwd()
	if err != nil {
		return "", err
	}

	if !filepath.IsAbs(path) {
		path = cwd + string(filepath.Separator) + path
	}

	abspath := filepath.Clean(path)

	// Adjust abspath. fs.FS requires it to be without starting '/'
	if abspath == "" ||
		(len(abspath) == 1 && os.IsPathSeparator(abspath[0])) {
		abspath = "."
	} else {
		abspath = abspath[1:]
	}

	return abspath, nil
}
