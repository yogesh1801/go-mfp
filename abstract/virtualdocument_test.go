// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// VirtualDocument tests

package abstract

import (
	"bytes"
	"io"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// TestDocumentFromBytes tests documents, created by NewVirtualDocument
func TestDocumentFromBytes(t *testing.T) {
	files := [][]byte{
		[]byte("000"),
		[]byte("111"),
		[]byte("222"),
	}

	res := Resolution{200, 200}

	newdoc := func() Document {
		return NewVirtualDocument(res, files...)
	}

	// Test normal usage
	doc := newdoc()

	buf := &bytes.Buffer{}
	cnt := 0
	file, err := doc.Next()
	for err == nil {
		io.Copy(buf, file)
		cnt++

		file, err = doc.Next()
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

	if doc.Resolution() != res {
		t.Errorf("Resolution mismatch: %v != %v",
			doc.Resolution(), res)
	}

	// Test reading from closed document
	doc = newdoc()
	file, _ = doc.Next()
	doc.Close()
	_, err = file.Read(make([]byte, 5))
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

// TestDocumentFromBytesFileFormat tests DocumentFile format
// for documents, created by NewVirtualDocument
func TestDocumentFromBytesFileFormat(t *testing.T) {
	type testData struct {
		data   []byte
		format string
	}

	tests := []testData{
		{
			data:   testutils.Images.BMP100x75,
			format: DocumentFormatBMP,
		},

		{
			data:   testutils.Images.PNG100x75rgb8,
			format: DocumentFormatPNG,
		},

		{
			data:   testutils.Images.PNG100x75gray8,
			format: DocumentFormatPNG,
		},

		{
			data:   testutils.Images.JPEG100x75,
			format: DocumentFormatJPEG,
		},
	}

	// Create a test document
	files := [][]byte{}
	for _, test := range tests {
		files = append(files, test.data)
	}

	res := Resolution{200, 200}
	doc := NewVirtualDocument(res, files...)

	// Verify that formats are properly recognized
	for _, test := range tests {
		file, err := doc.Next()
		if err != nil {
			panic(err) // Should not happen
		}

		format := file.Format()
		if format != test.format {
			t.Errorf("DocumentFile.Format: expected %q, present %q",
				test.format, format)
		}
	}
}
