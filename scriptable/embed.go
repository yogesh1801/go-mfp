// MFP - Miulti-Function Printers and scanners toolkit
// Print and scam servers with added scriptability.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Embedded files

package scriptable

import _ "embed" // For go:embed to work

//go:embed init.py
var embedPyInit string
