// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package cmdcups

import (
	"fmt"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/cups"
	"github.com/alexpevzner/mfp/transport"
)

// cmdGetDefault defines "get-default" command
var cmdGetDefault = argv.Command{
	Name:    "get-default",
	Help:    "Get default printer",
	Handler: cmdGetDefaultHandler,
}

func cmdGetDefaultHandler(*argv.Invocation) error {
	clnt := cups.NewClient(transport.DefaultCupsUNIX, nil)
	prn, err := clnt.CUPSGetDefault(nil)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", prn.PrinterName)
	fmt.Printf("URL:  %s\n", prn.PrinterURISupported)

	return nil
}
