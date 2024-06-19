// MFP - Miulti-Function Printers and scanners toolkit
// Execution environment
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Pager -- scroll text in a terminal

package env

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Parameters for external Pager program:
var (
	DefaultPagerCommand     = "less"
	DefaultPagerEnvironment = []string{"LESS=FRX", "LV=-c"}
)

// Pager scrolls potentially large text in a terminal.
// Implements io.Writer interface
type Pager struct {
	buf bytes.Buffer
}

// NewPager creates a new [Pager]
func NewPager() *Pager {
	return &Pager{}
}

// Write collects text for display. All subsequent writes
// are concatenated and displayed as a whole.
func (p *Pager) Write(text []byte) (n int, err error) {
	return p.buf.Write(text)
}

// Printf writes formatted line into the [Pager], adding newline
// character after end of the line.
//
// It returns the number of bytes written and any write error encountered.
func (p *Pager) Printf(format string, args ...any) (n int, err error) {
	return fmt.Fprintf(&p.buf, format+"\n", args...)
}

// Display shows collected text in terminal.
func (p *Pager) Display() error {
	command := os.Getenv("PAGER")
	environment := []string{}
	if command == "" {
		command = DefaultPagerCommand
		environment = DefaultPagerEnvironment
	}

	// Prepare pager command
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), environment...)

	// Create pager's stdin stream
	pipe, err := cmd.StdinPipe()
	if err != nil {
		goto ERROR
	}

	go func() {
		io.Copy(pipe, &p.buf)
		pipe.Close()
	}()

	// Run the command
	err = cmd.Run()
	if err != nil {
		goto ERROR
	}

	return nil

ERROR:
	return fmt.Errorf("Pager: %w", err)
}
