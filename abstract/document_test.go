// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Document tests

package abstract

import (
	"bytes"
	"io"
	"testing"
)

// TestDocumentFromBytes tests documents, created by NewDocumentFromBytes
func TestDocumentFromBytes(t *testing.T) {
	files := [][]byte{
		[]byte("000"),
		[]byte("111"),
		[]byte("222"),
	}

	format := "application/data"
	res := Resolution{200, 200}

	newdoc := func() Document {
		return NewDocumentFromBytes(format, res,
			files...)
	}

	// Test normal usage
	doc := newdoc()

	buf := &bytes.Buffer{}
	cnt := 0
	rd, err := doc.Next()
	for err == nil {
		io.Copy(buf, rd)
		cnt++

		rd, err = doc.Next()
	}

	if err != io.EOF {
		t.Errorf("Error mismatch: %s != %s", err, io.EOF)
	}

	if cnt != 3 {
		t.Errorf("Files count mismatch: %d != 3", cnt)
	}

	joined := string(bytes.Join(files, nil))
	if buf.String() != joined {
		t.Errorf("Returned data mismatch: %q != %q", buf, joined)
	}

	if doc.Format() != format {
		t.Errorf("Format mismatch: %q != %q", doc.Format(), format)
	}

	if doc.Resolution() != res {
		t.Errorf("Resolution mismatch: %v != %v",
			doc.Resolution(), res)
	}

	// Test reading from closed document
	doc = newdoc()
	rd, _ = doc.Next()
	doc.Close()
	_, err = rd.Read(make([]byte, 5))
	if err != ErrDocumentClosed {
		t.Errorf("Read from closed document:\n"+
			"error expected: %s\n"+
			"error present:  %s\n",
			ErrDocumentClosed, err)
	}

	// Test Document.Next from closed document
	doc = newdoc()
	doc.Close()
	_, err = doc.Next()
	if err != ErrDocumentClosed {
		t.Errorf("Document.Next from closed document:\n"+
			"error expected: %s\n"+
			"error present:  %s\n",
			ErrDocumentClosed, err)
	}
}
