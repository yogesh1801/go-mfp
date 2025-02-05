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
	"strconv"
	"strings"

	"github.com/alexpevzner/mfp/transport"
)

// mapping defines the mapping between local port and destination URL
type mapping struct {
	proto     string   // Proxy protocol
	localPort int      // Local port
	targetURL *url.URL // Destination URL
}

// parseMapping parses mapping from the command-line option
// string of the following form:
//
//	local-port=target-url
func parseMapping(option, param string) (m mapping, err error) {
	// Split parameter into the local-port and target-url
	var local, target string
	if i := strings.IndexByte(param, '='); i >= 0 {
		local = param[:i]
		target = param[i+1:]
	}

	if local == "" || target == "" {
		err = fmt.Errorf("syntax must be \"%s local-port=target-url\"",
			option)
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
