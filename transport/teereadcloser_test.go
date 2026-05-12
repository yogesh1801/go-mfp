// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// TeeReadCloser tests

package transport

import (
	"bytes"
	"io"
	"sync"
	"testing"
)

// testTeeReadCloserReader is the io.ReadCloser for testing of
// TeeReadCloser
type testTeeReadCloserReader struct {
	io.Reader
	closeCount int
}

func (m *testTeeReadCloserReader) Close() error {
	m.closeCount++
	return nil
}

// testTestTeeReadCloserWriter is the io.WriteCloser for testing of
// TeeReadCloser
type testTestTeeReadCloserWriter struct {
	bytes.Buffer
	closeCount int
}

func (m *testTestTeeReadCloserWriter) Close() error {
	m.closeCount++
	return nil
}

// TestTeeReadCloser performs testing of TeeReadCloser
func TestTeeReadCloser(t *testing.T) {
	inputData := []byte("hello world")
	reader := &testTeeReadCloserReader{Reader: bytes.NewReader(inputData)}
	writer := &testTestTeeReadCloserWriter{}

	trc := TeeReadCloser(reader, writer)

	// 1. Read all data
	buf := make([]byte, len(inputData))
	n, err := trc.Read(buf)
	if err != nil && err != io.EOF {
		t.Fatalf("unexpected read error: %v", err)
	}
	if n != len(inputData) {
		t.Errorf("expected %d bytes, got %d", len(inputData), n)
	}

	// 2. Read again to trigger EOF and automatic writer close
	_, err = trc.Read(make([]byte, 1))
	if err != io.EOF {
		t.Errorf("expected EOF, got %v", err)
	}

	if writer.closeCount != 1 {
		t.Errorf("writer should be closed once after EOF, got %d", writer.closeCount)
	}

	// 3. Explicitly call Close
	err = trc.Close()
	if err != nil {
		t.Errorf("unexpected close error: %v", err)
	}

	if reader.closeCount != 1 {
		t.Errorf("reader should be closed exactly once, got %d", reader.closeCount)
	}
	if writer.closeCount != 1 {
		t.Errorf("writer should still be closed only once, got %d", writer.closeCount)
	}

	// 4. Verify data integrity
	if writer.String() != string(inputData) {
		t.Errorf("expected data %q, got %q", string(inputData), writer.String())
	}
}

// TestTeeReadCloser2 performs testing of TeeReadCloser2
func TestTeeReadCloser2(t *testing.T) {
	data := []byte("synchronous pipe test")
	source := &testTeeReadCloserReader{Reader: bytes.NewReader(data)}

	r1, r2 := TeeReadCloser2(source)

	var wg sync.WaitGroup
	wg.Add(1)

	var r2Data []byte
	var r2Err error

	// Read from r2 in a separate goroutine to prevent deadlocking r1
	go func() {
		defer wg.Done()
		defer r2.Close()
		r2Data, r2Err = io.ReadAll(r2)
	}()

	// Read from r1
	r1Data, r1Err := io.ReadAll(r1)
	r1.Close()

	wg.Wait()

	// Validate r1
	if r1Err != nil {
		t.Errorf("r1 read error: %v", r1Err)
	}
	if !bytes.Equal(r1Data, data) {
		t.Errorf("r1 expected %q, got %q", string(data), string(r1Data))
	}

	// Validate r2
	if r2Err != nil {
		t.Errorf("r2 read error: %v", r2Err)
	}
	if !bytes.Equal(r2Data, data) {
		t.Errorf("r2 expected %q, got %q", string(data), string(r2Data))
	}

	// Validate source closure
	if source.closeCount != 1 {
		t.Errorf("source reader should be closed exactly once, got %d", source.closeCount)
	}
}
