// MFP - Miulti-Function Printers and scanners toolkit
// The "model" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Query printer and scanner attributes

package model

import (
	"context"
	"net/url"

	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/transport"
)

// queryESCLScannerCapabilities queries escl.ScannerCapabilities from
// the provided endpoints (assuming they all are aliases of the same
// device).
func queryESCLScannerCapabilities(ctx context.Context,
	endpoints []string) (*escl.ScannerCapabilities, error) {

	var err error

	for _, ep := range endpoints {
		log.Debug(ctx, "escl: trying %q", ep)

		var u *url.URL
		u, err2 := transport.ParseAddr(ep, "")
		if err2 != nil {
			if err == nil {
				err = err2
			}

			log.Debug(ctx, "escl: %q: %s", ep, err2)
			continue
		}

		clnt := escl.NewClient(u, nil)
		caps, _, err2 := clnt.GetScannerCapabilities(ctx)

		if err2 != nil {
			if err == nil {
				err = err2
			}

			log.Debug(ctx, "escl: %q: %s", ep, err2)
			continue
		}

		return caps, nil
	}

	return nil, err
}
