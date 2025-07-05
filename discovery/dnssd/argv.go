// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Argv helpers for DNS-SD

package dnssd

import (
	"context"
	"strings"

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/discovery"
	"github.com/OpenPrinting/go-mfp/util/generic"
)

// ArgvCompleter is the [argv.Completer] that performs DNS-SD based
// auto-completion.
func ArgvCompleter(prefix string) []argv.Completion {
	// Prepare the Client
	ctx := context.Background()
	clnt := discovery.NewClient(ctx)
	defer clnt.Close()

	backend, err := NewBackend(ctx, "", 0)
	if err != nil {
		return nil
	}

	defer backend.Close()
	clnt.AddBackend(backend)

	// Perform device discovery and gather results
	devices, err := clnt.GetDevices(ctx, discovery.ModeNormal)
	if err != nil {
		return nil
	}

	found := generic.NewSet[string]()
	for _, dev := range devices {
		if strings.HasPrefix(dev.DNSSDName, prefix) {
			found.Add(dev.DNSSDName)
		}
	}

	// Export results
	compl := make([]argv.Completion, 0, found.Count())
	found.ForEach(func(name string) {
		compl = append(compl, argv.Completion{String: name})
	})

	return compl
}
