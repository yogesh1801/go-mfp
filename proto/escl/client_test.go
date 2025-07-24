// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL Client test

package escl

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestClient tests the basic functionality of the eSCL Client
func TestClient(t *testing.T) {
	// Create ScannerCapabilities
	xml, err := xmldoc.Decode(
		NsMap,
		bytes.NewReader(testutils.
			Kyocera.ECOSYS.M2040dn.ESCL.ScannerCapabilities))
	assert.NoError(err)

	caps, err := DecodeScannerCapabilities(xml)
	assert.NoError(err)

	// Create loopback transport
	tr, loopback := transport.NewLoopback()

	// Start virtual scanner
	s := &abstract.VirtualScanner{
		ScanCaps: caps.ToAbstract(),
		Resolution: abstract.Resolution{
			XResolution: 600,
			YResolution: 600,
		},
		PlatenImage: testutils.Images.PNG5100x7016,
		ADFImages: [][]byte{
			testutils.Images.PNG5100x7016,
			testutils.Images.PNG5100x7016,
			testutils.Images.PNG5100x7016,
		},
	}

	base := transport.MustParseURL("http://localhost/eSCL")
	options := AbstractServerOptions{
		Version:  caps.Version,
		Scanner:  s,
		BasePath: base.Path,
	}

	handler := NewAbstractServer(context.TODO(), options)
	server := transport.NewServer(nil, handler)

	go server.Serve(loopback)
	defer server.Close()

	// Create a client
	clnt := NewClient(base, tr)

	// Test Client.GetScannerCapabilities
	caps2, _, err := clnt.GetScannerCapabilities(context.TODO())
	if err != nil {
		t.Errorf("Client.GetScannerCapabilities: %s", err)
		return
	}

	capsExpected := FromAbstractScannerCapabilities(caps.Version, s.ScanCaps)
	diff := testutils.Diff(caps2, capsExpected)
	if diff != "" {
		t.Errorf("Client.GetScannerCapabilities:\n%s", diff)
		return
	}

	// Test Client.GetScannerStatus
	status, _, err := clnt.GetScannerStatus(context.TODO())
	if err != nil {
		t.Errorf("Client.GetScannerStatus: %s", err)
		return
	}

	if status.State != ScannerIdle {
		t.Errorf("Client.GetScannerStatus: state mismatch:\n"+
			"expected: %s\n"+
			"present:  %s\n",
			ScannerIdle, status.State)
	}

	// Test Client.Scan
	rq := ScanSettings{
		Version:     caps.Version,
		InputSource: optional.New(InputFeeder),
		XResolution: optional.New(s.Resolution.XResolution),
		YResolution: optional.New(s.Resolution.YResolution),
	}

	job, _, err := clnt.Scan(context.TODO(), rq)
	if err != nil {
		t.Errorf("Client.Scan: %s", err)
		return
	}

	// Fetch scanned images
	images := 0
	for err == nil {
		var doc io.ReadCloser
		doc, _, err = clnt.NextDocument(context.TODO(), job)
		if doc != nil {
			images++
			defer doc.Close()
		}

		if err != nil && err != io.EOF {
			t.Errorf("Client.NextDocument: %s", err)
			return
		}
	}

	if images != len(s.ADFImages) {
		t.Errorf("Client.NextDocument:\n"+
			"images expected: %d\n"+
			"images present: %d\n",
			len(s.ADFImages), images)
	}

	// Test Client.Cancel
	_, err = clnt.Cancel(context.TODO(), job)
	if err != nil && err != io.EOF {
		t.Errorf("Client.Scan: %s", err)
		return
	}
}
