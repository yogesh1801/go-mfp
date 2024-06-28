// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The "get-printers" command.

package cmdcups

import (
	"context"
	"math"
	"strconv"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/cups"
	"github.com/alexpevzner/mfp/env"
)

// cmdGetPrinters defines the "get-printers" sub-command.
var cmdGetPrinters = argv.Command{
	Name:    "get-printers",
	Help:    "Get information on configured printers",
	Handler: cmdGetPrintersHandler,
	Options: []argv.Option{
		{
			Name:     "--attrs",
			Help:     "Additional attributes",
			HelpArg:  "attr,...",
			Validate: argv.ValidateAny,
			Complete: optAttrsComplete,
		},
		{
			Name:     "--id",
			Help:     "Printer ID (1...65535)",
			HelpArg:  "id",
			Validate: argv.ValidateIntRange(0, 1, 65535),
		},
		{
			Name:     "--limit",
			Help:     "Maximum number of printers",
			HelpArg:  "N",
			Validate: argv.ValidateIntRange(0, 1, math.MaxInt32),
		},
		{
			Name: "--location",
			Help: "" +
				`Printer location ` +
				`(e.g., "2nd Floor Computer Lab")`,
			HelpArg:  "where",
			Validate: argv.ValidateAny,
		},
		{
			Name:     "--user",
			Help:     "Show only printers accessible to that user",
			HelpArg:  "name",
			Validate: argv.ValidateAny,
		},
		argv.HelpOption,
	},
}

// cmdGetPrintersHandler is the "get-printers" command handler
func cmdGetPrintersHandler(ctx context.Context, inv *argv.Invocation) error {
	// Prepare arguments
	dest := optCUPSURL(inv)

	sel := &cups.GetPrintersSelection{}

	if opt, ok := inv.Get("--id"); ok {
		sel.PrinterID, _ = strconv.Atoi(opt)
	}

	if opt, ok := inv.Get("--limit"); ok {
		sel.Limit, _ = strconv.Atoi(opt)
	}

	if opt, ok := inv.Get("--location"); ok {
		sel.PrinterLocation = opt
	}

	attrList := optAttrsGet(inv)
	attrList = append(attrList, prnAttrsRequested...)

	sel.User, _ = inv.Get("--user")

	// Perform the query
	pager := env.NewPager()

	pager.Printf("CUPS: %s", dest)

	clnt := cups.NewClient(dest, nil)
	printers, err := clnt.CUPSGetPrinters(ctx, sel, attrList)
	if err != nil {
		return err
	}

	for _, prn := range printers {
		pager.Printf("")
		prnAttrsFormat(pager, prn)
	}

	return pager.Display()
}
