// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common options

package cmdcups

import (
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/cups"
	"github.com/alexpevzner/mfp/ipp"
	"github.com/alexpevzner/mfp/transport"
)

// optAttrs describes the --attrs option.
// It specifies a list of requested attributes.
var optAttrs = argv.Option{
	Name:     "--attrs",
	Help:     "Additional attributes",
	HelpArg:  "attr,...",
	Validate: argv.ValidateAny,
	Complete: optAttrsComplete,
}

// optAttrsGet returns --attrs option (list of requested attributes).
func optAttrsGet(inv *argv.Invocation) (attrs []string) {
	params := inv.Values("--attrs")
	for _, param := range params {
		for _, name := range strings.Split(param, ",") {
			if name != "" {
				attrs = append(attrs, name)
			}
		}
	}

	return
}

// optAttrsGet is the completion callback for the --attrs option.
func optAttrsComplete(arg string) (compl []argv.Completion) {
	infos := ((*ipp.PrinterAttributes)(nil)).KnownAttrs()

	attrName := arg
	prefix := ""

	if i := strings.LastIndex(attrName, ","); i >= 0 {
		attrName = arg[i+1:]
		prefix = arg[:i+1]
	}

	for _, info := range infos {
		if strings.HasPrefix(info.Name, attrName) {
			c := argv.Completion{
				String: prefix + info.Name + ",",
				Flags:  argv.CompletionNoSpace,
			}
			compl = append(compl, c)
		}
	}

	return
}

// optID describes the --id option.
// It specifies the printer-id.
var optID = argv.Option{
	Name:     "--id",
	Help:     "Printer ID (1...65535)",
	HelpArg:  "id",
	Validate: argv.ValidateIntRange(0, 1, 65535),
}

// optIDGet returns --id option value.
func optIDGet(inv *argv.Invocation) int {
	id := 0
	if opt, ok := inv.Get("--id"); ok {
		id, _ = strconv.Atoi(opt)
	}
	return id
}

// optLimit describes the --limit option.
// It specifies the maximum number of returned printers
var optLimit = argv.Option{
	Name:     "--limit",
	Help:     "Maximum number of printers",
	HelpArg:  "N",
	Validate: argv.ValidateIntRange(0, 1, math.MaxInt32),
}

// optLimitGet returns --limit option value.
func optLimitGet(inv *argv.Invocation) int {
	lim := 0
	if opt, ok := inv.Get("--limit"); ok {
		lim, _ = strconv.Atoi(opt)
	}
	return lim
}

// optLocation describes the --location option.
// It specified the desired printer location (e.g. "2nd Floor Computer Lab")
var optLocation = argv.Option{
	Name: "--location",
	Help: "" +
		`Printer location ` +
		`(e.g., "2nd Floor Computer Lab")`,
	HelpArg:  "where",
	Validate: argv.ValidateAny,
}

// optLocationGet returns --location option value.
func optLocationGet(inv *argv.Invocation) string {
	opt, _ := inv.Get("--location")
	return opt
}

// optTimeout describes the --timeout=seconds option.
// It specifies operation timeout
var optTimeout = argv.Option{
	Name:     "--timeout",
	Help:     "operation timeout",
	HelpArg:  "seconds",
	Validate: argv.ValidateIntRange(0, 1, math.MaxInt32),
}

// optTimeoutGet returns --timeout option value.
func optTimeoutGet(inv *argv.Invocation) time.Duration {
	if opt, ok := inv.Get(optTimeout.Name); ok {
		v, _ := strconv.Atoi(opt)
		return time.Duration(v) * time.Second
	}
	return cups.DefaultGetDevicesTimeout
}

// optUser describes the --user option.
// It allows to filter printers by the user name
// Only printers accessible to that user will be returned.
var optUser = argv.Option{
	Name:     "--user",
	Help:     "Show only printers accessible to that user",
	HelpArg:  "name",
	Validate: argv.ValidateAny,
}

// optUserGet returns --user option value.
func optUserGet(inv *argv.Invocation) string {
	opt, _ := inv.Get("--user")
	return opt
}

// optCUPSURL returns CUPS URL (-u/--cups option).
// If option is not set, it uses default destination.
func optCUPSURL(inv *argv.Invocation) *url.URL {
	dest := transport.DefaultCupsUNIX

	if addr, ok := inv.Parent().Get("-u"); ok {
		dest = transport.MustParseAddr(addr, "ipp://localhost/")
	}

	return dest
}
