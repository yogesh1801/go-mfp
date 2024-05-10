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
	"os"
)

// die function writes message into the os.Stderr and dies.
// Having it a pointer to function allows to hook it in testing.
var die = func(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
