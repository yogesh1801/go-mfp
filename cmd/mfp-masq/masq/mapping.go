// MFP - Miulti-Function Printers and scanners toolkit
// The "masq" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package masq

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/OpenPrinting/go-mfp/transport"
)

// mapping defines the mapping between local port and destination URL
type mapping struct {
	param     string   // original parameter
	proto     proto    // Proxy protocol
	localPort int      // Local port
	targetURL *url.URL // Destination URL
}

// parseMapping parses mapping from the command-line option
// string of the following form:
//
//	local-port=target-url
func parseMapping(proto proto, param string) (m mapping, err error) {
	// Save param and proto
	m.param = param
	m.proto = proto

	// Split parameter into the local-port and target-url
	var local, target string
	if i := strings.IndexByte(param, '='); i >= 0 {
		local = param[:i]
		target = param[i+1:]
	}

	if local == "" || target == "" {
		err = fmt.Errorf("parameter must be \"local-port=target-url\"")
		return
	}

	// Parse local-port
	m.localPort, err = strconv.Atoi(local)
	if err != nil || m.localPort < 1 || m.localPort > 65535 {
		err = fmt.Errorf("%q: invalid port", local)
		return
	}

	m.targetURL, err = transport.ParseAddr(target, "")
	if err != nil {
		err = fmt.Errorf("%q: %s", target, err)
		return
	}

	return
}

// mustParseMapping parses mapping like parseMapping and panics
// in a case of errors
func mustParseMapping(proto proto, param string) mapping {
	m, err := parseMapping(proto, param)
	if err != nil {
		err = fmt.Errorf("%s: %s", param, err)
		panic(err)
	}
	return m
}
