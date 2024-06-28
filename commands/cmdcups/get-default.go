// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The "get-default" command.

package cmdcups

import (
	"context"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/cups"
	"github.com/alexpevzner/mfp/env"
)

// cmdGetDefault defines the "get-default" sub-command
var cmdGetDefault = argv.Command{
	Name:    "get-default",
	Help:    "Get default printer",
	Handler: cmdGetDefaultHandler,
	Options: []argv.Option{
		{
			Name:     "--attrs",
			Help:     "Additional attributes",
			HelpArg:  "attr,...",
			Validate: argv.ValidateAny,
			Complete: optAttrsComplete,
		},
		argv.HelpOption,
	},
}

// cmdGetDefaultHandler is the "get-default" command handler
func cmdGetDefaultHandler(ctx context.Context, inv *argv.Invocation) error {
	dest := optCUPSURL(inv)

	attrList := optAttrsGet(inv)
	attrList = append(attrList, prnAttrsRequested...)

	pager := env.NewPager()

	pager.Printf("CUPS: %s", dest)

	clnt := cups.NewClient(dest, nil)
	prn, err := clnt.CUPSGetDefault(ctx, attrList)
	if err != nil {
		return err
	}

	prnAttrsFormat(pager, prn)

	return pager.Display()
}
