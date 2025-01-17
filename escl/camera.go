// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Camera capabilities

package escl

import "github.com/alexpevzner/mfp/optional"

// Camera contains scanner capabilities for the Camera source.
type Camera struct {
	CameraInputCaps optional.Val[InputSourceCaps] // Camera capabilities
}
