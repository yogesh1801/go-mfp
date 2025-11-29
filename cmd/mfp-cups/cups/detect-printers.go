// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The "get-printers" command.

package cups

import (
	"context"

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/cups"
	"github.com/OpenPrinting/go-mfp/internal/env"
)

// cmdDetectPrinters defines the "detect-printers" sub-command.
var cmdDetectPrinters = argv.Command{
	Name:    "detect-printers",
	Help:    "Search for available devices",
	Handler: cmdDetectPrintersHandler,
	Options: []argv.Option{
		optSchemesExclude,
		optSchemesInclude,
		optLimit,
		optTimeout,
		argv.HelpOption,
	},
}

// cmdDetectPrintersHandler is the "detect-printers" command handler
func cmdDetectPrintersHandler(ctx context.Context, inv *argv.Invocation) error {
	// Prepare arguments
	dest := optCUPSURL(inv)

	sel := &cups.GetDevicesSelection{
		Limit:          optLimitGet(inv),
		Timeout:        optTimeoutGet(inv),
		ExcludeSchemes: optSchemesExcludeGet(inv),
		IncludeSchemes: optSchemesIncludeGet(inv),
	}

	// Perform the query
	clnt := cups.NewClient(dest, nil)
	devices, err := clnt.CUPSGetDevices(ctx, sel, []string{"all"})
	if err != nil {
		return err
	}

	// Format output
	pager := env.NewPager()

	pager.Printf("CUPS: %s", dest)
	for _, dev := range devices {
		pager.Printf("")
		devAttrsFormat(pager, dev)
	}

	return pager.Display()
}
