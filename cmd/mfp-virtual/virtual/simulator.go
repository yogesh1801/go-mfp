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
	"fmt"
	"net"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/internal/env"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/uuid"
)

func scannerCapabilities() *abstract.ScannerCapabilities {
	colorModes := generic.MakeBitset(
		abstract.ColorModeBinary,
		abstract.ColorModeMono,
		abstract.ColorModeColor,
	)

	depths := generic.MakeBitset(
		abstract.ColorDepth8,
	)

	renderings := generic.MakeBitset(
		abstract.BinaryRenderingHalftone,
		abstract.BinaryRenderingThreshold,
	)

	intents := generic.MakeBitset(
		abstract.IntentDocument,
		abstract.IntentTextAndGraphic,
		abstract.IntentPhoto,
		abstract.IntentPreview,
	)

	resolutions := []abstract.Resolution{
		{XResolution: 75, YResolution: 75},
		{XResolution: 150, YResolution: 150},
		{XResolution: 300, YResolution: 300},
		{XResolution: 600, YResolution: 600},
	}

	profile := abstract.SettingsProfile{
		ColorModes:       colorModes,
		Depths:           depths,
		BinaryRenderings: renderings,
		Resolutions:      resolutions,
	}

	inputcaps := &abstract.InputCapabilities{
		MinWidth:   0,
		MaxWidth:   abstract.A4Width,
		MinHeight:  0,
		MaxHeight:  abstract.A4Height,
		MaxXOffset: abstract.A4Width / 2,
		MaxYOffset: abstract.A4Height / 2,
		Intents:    intents,
		Profiles:   []abstract.SettingsProfile{profile},
	}

	caps := &abstract.ScannerCapabilities{
		UUID: uuid.Must(uuid.Parse(
			"169e8d94-9a17-4f14-ae81-52b9176ee9be")),
		MakeAndModel:     "OpenPrinting eSCL scanner",
		SerialNumber:     "OP-0000223321",
		Manufacturer:     "OpenPrinting",
		DocumentFormats:  []string{"image/jpeg", "application/pdf"},
		ADFCapacity:      50,
		CompressionRange: abstract.Range{Min: 2, Normal: 5, Max: 10},
		BrightnessRange:  abstract.Range{Min: -100, Normal: 0, Max: 100},
		ContrastRange:    abstract.Range{Min: -100, Normal: 0, Max: 100},
		Platen:           inputcaps,
		ADFSimplex:       inputcaps,
		ADFDuplex:        inputcaps,
	}
	return caps
}

// simulate runs scanner simulator.
//
// If argv is not empty, it specifies the external command that will
// be run under the simulator.
func simulate(ctx context.Context, port int, argv []string) error {
	s := &abstract.VirtualScanner{
		ScanCaps: scannerCapabilities(),
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

	// Create a virtual server
	options := escl.AbstractServerOptions{
		Scanner:  s,
		BasePath: "/eSCL",
	}

	handler := escl.NewAbstractServer(ctx, options)
	server := transport.NewServer(nil, handler)

	addr := fmt.Sprintf("localhost:%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	// Run external command if specified
	if len(argv) != 0 {
		runner := env.Runner{
			ESCLPort: port,
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
