// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL part of Model

package modeling

import (
	"net/http"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/transport"
)

// esclWrite writes eSCL part of model into the [formatter].
func (model *Model) esclWrite(f *formatter) error {
	if model.esclScanCaps != nil {
		obj, err := model.pyExportStruct(model.esclScanCaps)
		if err != nil {
			return err
		}

		f.Printf("# eSCL scanner parameters:\n")
		f.Printf("escl.caps = ")

		return f.Format(obj)
	}

	return nil
}

// NewESCLServer creates a virtual eSCL server on a top of
// the existent abstract.Scanner implementation.
//
// It will return nil, if model doesn't have the eSCL scanner capabilities.
func (model *Model) NewESCLServer(
	scanner abstract.Scanner) *escl.AbstractServer {

	// Obtain scanner capabilities
	caps := model.GetESCLScanCaps()
	if caps == nil {
		return nil
	}

	// Setup options
	options := escl.AbstractServerOptions{
		Version:  caps.Version,
		Scanner:  scanner,
		BasePath: "/eSCL",
		Hooks: escl.ServerHooks{
			OnScannerCapabilitiesResponse: model.esclOnScannerCapabilitiesResponse,
		},
	}

	// Create the eSCL server
	return escl.NewAbstractServer(options)
}

// esclOnScannerCapabilitiesResponse implements the
// [escl.ServerHooks.OnScannerCapabilitiesResponse] hook
// for the modeled eSCL scanner.
func (model *Model) esclOnScannerCapabilitiesResponse(
	query *transport.ServerQuery,
	caps *escl.ScannerCapabilities) *escl.ScannerCapabilities {

	caps2 := model.GetESCLScanCaps()
	if caps2 == nil {
		query.Reject(http.StatusServiceUnavailable, nil)
	}

	return caps2
}
