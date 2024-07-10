// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// File Backend

package log

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// backendFile is the Backend that writes log to file.
type backendFile struct {
	mutex   sync.Mutex // Access lock
	path    string     // Path to file
	maxsize int        // Maximum file size before rotation
	backups int        // Maximum number of created backups
	file    *os.File   // Output file
}

// NewFileBackend returns a Backend that writes log to file.
// It also supports log file rotation.
//
// When file size exceeds maxsize, the content of the log file
// file gzip-ed and copied into backup. Already existent backups
// renamed in circular order, according to the following scheme:
//
//	file.log -> file.log.0.gz -> file.log.1.gz -> ... -> file.log.N.gz
//
// The backups parameter defines maximum number of backup files
// (so N is backups-1).
//
// If the previous run of the program has left more backup files,
// that current configuration allows, the excessive files will
// be deleted during rotation.
//
// The file Backend doesn't break multi-line logical record between
// two log files, so the actual maximum file size may exceed the
// configured threshold.
//
// Setting maxsize to 0 disables rotation and setting backups
// to 0 disables creation of the backup files.
//
// Note, file Backend ignores any I/O errors when writing to
// log files, as it has no method to report them.
func NewFileBackend(path string, maxsize, backups int) Backend {
	return &backendFile{
		path:    path,
		maxsize: maxsize,
		backups: backups,
	}
}

// Send implements the [Backend.Send] interface.
func (bk *backendFile) Send(levels []Level, lines [][]byte) {
	// Lock the Backend
	bk.mutex.Lock()
	defer bk.mutex.Unlock()

	// Open log file on demand
	if bk.file == nil {
		os.MkdirAll(filepath.Dir(bk.path), 0755)
		bk.file, _ = os.OpenFile(bk.path,
			os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

		if bk.file == nil {
			return
		}
	}

	// Acquire file lock
	fl, err := FileLockEx(bk.file)
	if err != nil {
		return
	}

	defer fl.Close()

	// Rotate now
	bk.rotate()

	// Format time prefix
	now := time.Now()

	year, month, day := now.Date()
	hour, min, sec := now.Clock()

	prefix := fmt.Sprintf("%2.2d-%2.2d-%4.4d %2.2d:%2.2d:%2.2d: ",
		day, month, year,
		hour, min, sec)

	// Write log lines
	buf := bufAlloc()
	defer bufFree(buf)

	for _, line := range lines {
		buf.WriteString(prefix)
		buf.Write(line)
		buf.WriteByte('\n')
	}

	buf.WriteTo(bk.file)
}

// rotate performs auto-rotation
func (bk *backendFile) rotate() {
	// Do we need to rotate?
	if bk.maxsize <= 0 {
		return // Rotation is disabled
	}

	stat, err := bk.file.Stat()
	if err != nil || stat.Size() <= int64(bk.maxsize) {
		return // Rotation not required
	}

	// Remove all excessive backup files.
	err = bk.backupPrune()
	if err != nil {
		return
	}

	// gzip current log into temporary backup:
	//   file.log->file.log.~.gz
	tmpbackup := bk.backupName("~")
	if bk.backups > 0 {
		err := bk.gzip(tmpbackup)
		if err != nil {
			return
		}
	}

	// Rotate backup files. For bk.backups == 5 it will be:
	//   file.log.4.gz -> file.log.4.gz
	//   file.log.2.gz -> file.log.3.gz
	//   file.log.1.gz -> file.log.2.gz
	//   file.log.0.gz -> file.log.1.gz
	if bk.backups > 1 {
		prev := bk.backupName(strconv.Itoa(bk.backups - 1))
		for i := bk.backups - 1; i > 0; i-- {
			next := bk.backupName(strconv.Itoa(i - 1))
			os.Rename(next, prev)
			prev = next
		}

		// Rename temporary backup the 0th backup
		os.Rename(tmpbackup, prev)
	}

	// And finally, truncate output file
	bk.file.Truncate(0)
}

// backupPrune removes all excessive backup files.
func (bk *backendFile) backupPrune() error {
	// List all files in the directory that look as our backups
	backups, err := filepath.Glob(bk.backupName("*"))
	if err != nil {
		return err
	}

	// Prepare a set of file names that we will keep
	keep := make(map[string]struct{})
	for i := 0; i < bk.backups-1; i++ {
		keep[bk.backupName(strconv.Itoa(i))] = struct{}{}
	}

	// And delete all files not in the set
	for _, name := range backups {
		if _, found := keep[name]; !found {
			os.Remove(name)
		}
	}

	return nil
}

// backupName returns a backup file name
func (bk *backendFile) backupName(index string) string {
	return fmt.Sprintf("%s.%s.gz", bk.path, index)
}

// gzip compresses current log file into the newpath
func (bk *backendFile) gzip(newpath string) error {
	// Open input file
	ifile, err := os.Open(bk.path)
	if err != nil {
		return err
	}

	defer ifile.Close()

	// Open output file
	ofile, err := os.OpenFile(newpath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	// gzip ifile->ofile
	w := gzip.NewWriter(ofile)
	_, err = io.Copy(w, ifile)
	err2 := w.Close()
	err3 := ofile.Close()

	switch {
	case err == nil && err2 != nil:
		err = err2
	case err == nil && err3 != nil:
		err = err3
	}

	// Cleanup and exit
	if err != nil {
		os.Remove(newpath)
	}

	return err

}
