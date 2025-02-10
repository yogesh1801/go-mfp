// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
//

package proxy

import (
	"bytes"
	"fmt"

	"github.com/OpenPrinting/goipp"
)

// changeSetMessage represents set of changes, applied to
// the [goipp.Message] during translation
type changeSetMessage struct {
	Groups []changeSetGroup // Changes per group
}

// changeSetGroup represents set of changes, applied to the
// particular group of attributes during message translation
type changeSetGroup struct {
	Tag    goipp.Tag        // Group tag
	Values []changeSetValue // Changes in values
}

// changeSetValue represents per-value changes
type changeSetValue struct {
	Path     string      // Path to the value from the Message root
	Old, New goipp.Value // Old and new values
}

// Empty reports if changeSetMessage is empty (contains no changes)
func (chg changeSetMessage) Empty() bool {
	return len(chg.Groups) == 0
}

// MarshalLog returns string representation of changeSetMessage for logging.
// It implements [log.Marshaler] interface.
func (chg changeSetMessage) MarshalLog() []byte {
	var buf bytes.Buffer

	for _, g := range chg.Groups {
		fmt.Fprintf(&buf, "GROUP %s:\n", g.Tag)
		for _, v := range g.Values {
			fmt.Fprintf(&buf, "    ATTR %s:\n", v.Path)
			fmt.Fprintf(&buf, "        Old: %s\n", v.Old)
			fmt.Fprintf(&buf, "        New: %s\n", v.New)
		}
	}

	return buf.Bytes()
}
