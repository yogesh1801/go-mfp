// MFP - Miulti-Function Printers and scanners toolkit
// Execution environment
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package env

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/OpenPrinting/go-mfp/transport"
)

// Runner runs external program in the environment, where interactions
// with CUPS and the eSCL scanner are redirected to the specified ports
// and can be intercepted.
//
// It does its work by setting the following environment variables:
//   - CUPS_SERVER=localhost:port
//   - SANE_AIRSCAN_DEVICE=escl:Scanner Name:http://localhost:port/eSCL
//
// In the context of the program being executed, these variables are
// interpreted by the libcups.so and sane-airscan, respectively.
//
// This is used by the mfp-proxy and mfp-virtual commands.
type Runner struct {
	CUPSPort int    // CUPS server port, 0 if none
	ESCLPort int    // eSCL server port, 0 if none
	ESCLPath string // Path part of the eSCL URL
	ESCLName string // eSCL scanner name (will be visible as SANE name)
}

// Run executes the command and waits for its completion.
// Command execution can be terminated by canceling the
// supplied [context.Context].
func (r *Runner) Run(ctx context.Context,
	command string, args ...string) error {

	// Prepare the command
	cmd := exec.CommandContext(ctx, command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	const envCUPS = "CUPS_SERVER="
	const envESCL = "SANE_AIRSCAN_DEVICE="

	// Copy system environment
	for _, env := range os.Environ() {
		switch {
		case r.CUPSPort != 0 && strings.HasPrefix(env, envCUPS):
		case r.ESCLPort != 0 && strings.HasPrefix(env, envESCL):

		default:
			cmd.Env = append(cmd.Env, env)
		}
	}

	// Set CUPS_SERVER and SANE_AIRSCAN_DEVICE
	if r.CUPSPort != 0 {
		env := fmt.Sprintf(envCUPS+"localhost:%d", r.CUPSPort)
		cmd.Env = append(cmd.Env, env)
	}

	if r.ESCLPort != 0 {
		name := r.ESCLName
		if name == "" {
			name = "eSCL scanner"
		}

		env := fmt.Sprintf(
			envESCL+"escl:%s:http://localhost:%d%s",
			name, r.ESCLPort, transport.CleanURLPath(r.ESCLPath))
		cmd.Env = append(cmd.Env, env)
	}

	// Run the command
	err := cmd.Run()
	return err
}
