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

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/cups"
	"github.com/alexpevzner/mfp/internal/env"
)

// cmdGetPrinters defines the "get-printers" sub-command.
var cmdGetPrinters = argv.Command{
	Name:    "get-printers",
	Help:    "Get information on configured printers",
	Handler: cmdGetPrintersHandler,
	Options: []argv.Option{
		optAttrs,
		optID,
		optLimit,
		optLocation,
		optUser,
		argv.HelpOption,
	},
}

// cmdGetPrintersHandler is the "get-printers" command handler
func cmdGetPrintersHandler(ctx context.Context, inv *argv.Invocation) error {
	// Prepare arguments
	dest := optCUPSURL(inv)

	sel := &cups.GetPrintersSelection{
		PrinterID:       optIDGet(inv),
		Limit:           optLimitGet(inv),
		PrinterLocation: optLocationGet(inv),
		User:            optUserGet(inv),
	}

	attrList := optAttrsGet(inv)
	attrList = append(attrList, prnAttrsRequested...)

	// Perform the query
	clnt := cups.NewClient(dest, nil)
	printers, err := clnt.CUPSGetPrinters(ctx, sel, attrList)
	if err != nil {
		return err
	}

	// Format output
	pager := env.NewPager()

	pager.Printf("CUPS: %s", dest)
	for _, prn := range printers {
		pager.Printf("")
		prnAttrsFormat(pager, prn)
	}

	return pager.Display()
}
