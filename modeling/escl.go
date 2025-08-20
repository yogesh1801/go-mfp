// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL part of Model

package modeling

import (
	"fmt"
	"net/http"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/transport"
)

// esclLoad decodes eSCL part of model. The model file assumed to
// be preloaded into the Model's Python interpreter (model.py).
func (model *Model) esclLoad() error {
	// Load and decode ScannerCapabilities
	obj, err := model.py.Eval("escl.caps")
	if err != nil {
		err = fmt.Errorf("escl.caps: %s", err)
		return err
	}

	if !obj.IsNone() {
		var caps *escl.ScannerCapabilities
		err = model.pyImportStruct(&caps, obj)
		if err != nil {
			err = fmt.Errorf("escl.caps: %s", err)
			return err
		}

		model.esclScanCaps = caps
	}

	// Load eSCL hooks
	const esclOnScanJobsRequestName = "escl_onScanJobsRequest"
	obj, err = model.py.GetGlobal(esclOnScanJobsRequestName)
	if err != nil {
		return err
	}

	if obj != nil && !obj.IsCallable() {
		return fmt.Errorf("%s is not function",
			esclOnScanJobsRequestName)
	}

	model.esclOnScanJobsRequestScriptlet = obj

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

	// Setup hooks
	if model.esclOnScanJobsRequestScriptlet != nil {
		options.Hooks.OnScanJobsRequest = model.esclOnScanJobsRequest
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

// esclOnScanJobsRequest implements the [escl.ServerHooks.OnScanJobsRequest]
// hook for he modeled eSCL scanner.
func (model *Model) esclOnScanJobsRequest(
	query *transport.ServerQuery,
	ss *escl.ScanSettings) *escl.ScanSettings {

	// Export request to Python
	q, err := model.queryToPython(query)
	if err != nil {
		query.Reject(http.StatusServiceUnavailable, err)
		return nil
	}

	rq, err := model.pyExportStruct(ss)
	if err != nil {
		query.Reject(http.StatusServiceUnavailable, err)
		return nil
	}

	// Call the hook
	_, err = model.esclOnScanJobsRequestScriptlet.Call(q, rq)
	if err != nil {
		query.Reject(http.StatusServiceUnavailable, err)
		return nil
	}

	// Convert possibly modified request back to Go
	err = model.queryFromPython(query, q)
	if err != nil {
		query.Reject(http.StatusServiceUnavailable, err)
		return nil
	}

	var ss2 *escl.ScanSettings
	err = model.pyImportStruct(&ss2, rq)
	if err != nil {
		query.Reject(http.StatusServiceUnavailable, err)
		return nil
	}

	return ss2
}
