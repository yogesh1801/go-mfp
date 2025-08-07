// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package proxy

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/OpenPrinting/go-mfp/transport"
)

// mapping defines the mapping between local port and destination URL
type mapping struct {
	param     string   // original parameter
	proto     proto    // Proxy protocol
	localPath string   // Local path
	targetURL *url.URL // Destination URL
}

// validateMapping mapping validates mapping, defined as the
// command-line option string.
//
// It can be used as argv.Option.Validate callback.
func validateMapping(param string) error {
	_, err := parseMapping(protoIPP, param)
	return err
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
		err = fmt.Errorf("parameter must be \"path=url\"")
		return
	}

	// Parse local path
	m.localPath = local

	// Parse target URL
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
