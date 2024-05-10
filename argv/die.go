// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Fatal exit

package argv

import (
	"fmt"
	"io"
	"os"
)

var (
	// dieOutput and dieExit variables allows to hook
	// fatal exit on testing, having 100% coverage evem
	// on that case
	dieOutput io.Writer = os.Stderr
	dieExit             = os.Exit
)

// die writes message into the os.Stderr and dies.
func die(err error) {
	fmt.Fprintf(dieOutput, "%s\n", err)
	dieExit(1)
}
