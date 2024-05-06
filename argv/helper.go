// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Help page generation

package argv

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Constants (formatting parameters)
//
//                                     |<>|<-- hlpMinColumnSpace
//
//                     Options are:
//                       -c, --compress    compress output
//   hlpOffOptionName -->
//   hlpOffOptionHelp -------------------->
const (
	hlpOffOptionName  = 2
	hlpOffOptionHelp  = 20
	hlpMinColumnSpace = 2
)

// Precomputed strings
var (
	// Prefix before option name
	hlpPfxOptionName = strings.Repeat(" ", hlpOffOptionName)

	// Prefix before option usage text
	hlpPfxOptionHelp = strings.Repeat(" ", hlpOffOptionHelp)
)

// helper builds help
type helper struct {
	cmd *Command  // Target command
	out io.Writer // Output goes here
	err error     // Sticky I/O error
}

// Help generates a help page and writes it into output io.Writer.
// It panics, if cmd.Verify() returns an error.
//
// The returned error, if any, is the I/O error from the destination
// io.Writer.
func Help(cmd *Command, out io.Writer) error {
	hlp := newHelper(cmd, out)
	hlp.generate()
	return hlp.err
}

// HelpString generates a help page and returns it as a single string.
// It panics, if cmd.Verify() returns an error.
func HelpString(cmd *Command) string {
	buf := &bytes.Buffer{}
	Help(cmd, buf)
	return buf.String()
}

// newHelper creates a new helper.
//
// It panics, if cmd.Verify() returns an error.
func newHelper(cmd *Command, out io.Writer) *helper {
	err := cmd.Verify()
	if err != nil {
		panic(err)
	}

	return &helper{
		cmd: cmd,
		out: out,
	}
}

// generate generates a help page
func (hlp *helper) generate() {
	hlp.describeUsageLine()
	hlp.describeOptions()
	hlp.describeCommandLong()
}

// describeUsageLine describes usage in a single line
func (hlp *helper) describeUsageLine() {
	cmd := hlp.cmd

	hlp.printf("usage: %s", cmd.Name)

	if cmd.hasOptions() {
		hlp.printf(" [options]")
	}

	for i := range cmd.Parameters {
		param := &cmd.Parameters[i]
		hlp.printf(" %s", param.Name)
	}

	if cmd.hasSubCommands() {
		hlp.printf(" command [arguments]")
	}

	hlp.nl()
}

// describeOptions describes command options
func (hlp *helper) describeOptions() {
	cmd := hlp.cmd

	if !cmd.hasOptions() {
		return
	}

	hlp.nl()
	hlp.puts("Options are:\n")

	for i := range cmd.Options {
		opt := &cmd.Options[i]
		names := hlpPfxOptionName + hlp.optionNames(opt)
		hlp.puts(names)

		usage := strings.Split(opt.Help, "\n")
		if len(usage) > 0 {
			if len(names)+hlpMinColumnSpace <= hlpOffOptionHelp {
				hlp.space(hlpOffOptionHelp - len(names))
				hlp.puts(usage[0] + "\n")
				usage = usage[1:]
			} else {
				hlp.nl()
			}

			for _, line := range usage {
				hlp.puts(hlpPfxOptionHelp + line + "\n")
			}
		}
	}
}

// describeCommandLong writes a long command description
func (hlp *helper) describeCommandLong() {
	cmd := hlp.cmd

	if cmd.Description != "" {
		hlp.nl()
		hlp.puts(cmd.Description)
		hlp.nl()
	}

}

// optionNames merges together Option names
func (hlp *helper) optionNames(opt *Option) string {
	names := opt.Name
	for _, name := range opt.Aliases {
		names += ", " + name
	}
	return names
}

// putc writes a character into the help page
func (hlp *helper) putc(c byte) {
	if hlp.err == nil {
		_, hlp.err = hlp.out.Write([]byte{c})
	}
}

// putc writes a string into the help page
func (hlp *helper) puts(s string) {
	if hlp.err == nil {
		_, hlp.err = hlp.out.Write([]byte(s))
	}
}

// space putc n space characters
func (hlp *helper) space(n int) {
	for n > 0 {
		hlp.putc(' ')
		n--
	}
}

// space nl putc NL character
func (hlp *helper) nl() {
	hlp.putc('\n')
}

// printf writes formatted string into the help page
func (hlp *helper) printf(format string, args ...interface{}) {
	if hlp.err == nil {
		_, hlp.err = fmt.Fprintf(hlp.out, format, args...)
	}
}
