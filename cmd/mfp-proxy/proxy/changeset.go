// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
//

package proxy

import "github.com/OpenPrinting/goipp"

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
