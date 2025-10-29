// MFP - Miulti-Function Printers and scanners toolkit
// The "virtual" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Virtual MFP simulator

package virtual

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/internal/env"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/modeling"
	"github.com/OpenPrinting/go-mfp/transport"
)

// simulate runs scanner simulator.
//
// If argv is not empty, it specifies the external command that will
// be run under the simulator.
func simulate(ctx context.Context, model *modeling.Model,
	port int, argv []string) error {

	// Create a virtual server
	pathmux := transport.NewPathMux()
	server := transport.NewServer(ctx, nil, pathmux)

	addr := fmt.Sprintf("localhost:%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// Add eSCL handler
	if esclcaps := model.GetESCLScanCaps(); esclcaps != nil {

		s := &abstract.VirtualScanner{
			ScanCaps: esclcaps.ToAbstract(),
			Resolution: abstract.Resolution{
				XResolution: 600,
				YResolution: 600,
			},
			PlatenImage: testutils.Images.PNG5100x7016,
			ADFImages: [][]byte{
				testutils.Images.PNG5100x7016,
				testutils.Images.PNG5100x7016,
				testutils.Images.PNG5100x7016,
			},
		}

		handler := model.NewESCLServer(s)
		pathmux.Add("/eSCL", handler)
	}

	// Add IPP handler
	if handler := model.NewIPPServer(); handler != nil {
		pathmux.Add("/ipp/print", handler)
	}

	// Check that we have added at least something
	if pathmux.Empty() {
		return errors.New("model is emoty")
	}

	// Run external command if specified
	if len(argv) != 0 {
		runner := env.Runner{
			ESCLPort: port,
			ESCLPath: "/eSCL",
			ESCLName: "Virtual MFP Scanner",
		}

		go func() {
			err = runner.Run(ctx, argv[0], argv[1:]...)
			server.Close()
		}()
	}

	// Make sure the program terminates when ctx is canceled.
	go func() {
		<-ctx.Done()
		server.Close()
	}()

	// Serve requests
	log.Info(ctx, "starting virtual MFP at %s", addr)
	server.Serve(ln)

	return err
}
