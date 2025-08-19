// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Embedded files

package modeling

import _ "embed" // For go:embed to work

//go:embed init.py
var embedPyInit string

//go:embed query.py
var embedPyQuery string

//go:embed escl.py
var embedPyEscl string
