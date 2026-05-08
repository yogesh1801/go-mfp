// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// trace.Writter support

package escl

import "github.com/OpenPrinting/go-mfp/proto/trace"

// esclDELETE implements trace.Message interface
// for the eSCL GET DELETE request and response messages
type esclDELETE struct{}

var _ = trace.Message(esclDELETE{})

func (esclDELETE) Ext() string          { return "" }
func (esclDELETE) Name() string         { return "DELETE" }
func (esclDELETE) MarshalLog() []byte   { return nil }
func (esclDELETE) MarshalTrace() []byte { return nil }

// esclNextDocument implements trace.Message interface
// for the eSCL GET NextDocument request and response messages
type esclNextDocument struct{}

var _ = trace.Message(esclNextDocument{})

func (esclNextDocument) Ext() string          { return "" }
func (esclNextDocument) Name() string         { return "NextDocument" }
func (esclNextDocument) MarshalLog() []byte   { return nil }
func (esclNextDocument) MarshalTrace() []byte { return nil }

// esclScanImageInfo implements trace.Message interface
// for the eSCL GET ScanImageInfo request and response messages
type esclScanImageInfo struct{}

var _ = trace.Message(esclScanImageInfo{})

func (esclScanImageInfo) Ext() string          { return "" }
func (esclScanImageInfo) Name() string         { return "ScanImageInfo" }
func (esclScanImageInfo) MarshalLog() []byte   { return nil }
func (esclScanImageInfo) MarshalTrace() []byte { return nil }

// esclScanJobs implements trace.Message interface
// for the eSCL GET ScanJobs request and response messages
type esclScanJobs struct{}

var _ = trace.Message(esclScanJobs{})

func (esclScanJobs) Ext() string          { return "" }
func (esclScanJobs) Name() string         { return "ScanJobs" }
func (esclScanJobs) MarshalLog() []byte   { return nil }
func (esclScanJobs) MarshalTrace() []byte { return nil }

// esclScannerCapabilities implements trace.Message interface
// for the eSCL GET ScannerCapabilities request and response messages
type esclScannerCapabilities struct{}

var _ = trace.Message(esclScannerCapabilities{})

func (esclScannerCapabilities) Ext() string          { return "" }
func (esclScannerCapabilities) Name() string         { return "ScannerCapabilities" }
func (esclScannerCapabilities) MarshalLog() []byte   { return nil }
func (esclScannerCapabilities) MarshalTrace() []byte { return nil }

// esclScannerStatus implements trace.Message interface
// for the eSCL GET ScannerStatus request and response messages
type esclScannerStatus struct{}

var _ = trace.Message(esclScannerStatus{})

func (esclScannerStatus) Ext() string          { return "" }
func (esclScannerStatus) Name() string         { return "ScannerStatus" }
func (esclScannerStatus) MarshalLog() []byte   { return nil }
func (esclScannerStatus) MarshalTrace() []byte { return nil }
