// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// (*Command) Main() test

package argv

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"
)

// TestMain is a test for (*Command) Main()
func TestMain(t *testing.T) {
	buf := &bytes.Buffer{}
	cmd := Command{
		Name: "test",
		Parameters: []Parameter{
			{Name: "greeting..."},
		},
		Handler: func(ctx context.Context, inv *Invocation) error {
			buf.WriteString(
				strings.Join(inv.Values("greeting"), ", "))
			return nil
		},
	}

	run := func(args []string) {
		saveArgs := os.Args
		saveDieOutput := dieOutput
		saveDieExit := dieExit
		saveHelpOutput := HelpOutput

		os.Args = args
		dieOutput = buf
		dieExit = func(int) {}
		HelpOutput = buf

		buf.Reset()
		cmd.Main(nil)

		os.Args = saveArgs
		dieOutput = saveDieOutput
		dieExit = saveDieExit
		HelpOutput = saveHelpOutput
	}

	run([]string{"test", "hello", "world"})

	expected := "hello, world"
	received := buf.String()

	if expected != received {
		t.Errorf("test 1:\n"+
			"expected: `%s`\n"+
			"present:  `%s`\n",
			expected, received)
	}

	run([]string{"test"})

	expected = HelpString(&cmd)
	received = buf.String()

	if expected != received {
		t.Errorf("test 2:\n"+
			"expected: `%s`\n"+
			"present:  `%s`\n",
			expected, received)
	}
}
