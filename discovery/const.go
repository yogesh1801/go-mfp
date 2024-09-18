// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Configurable constants

package discovery

import "time"

// Discovery parameters:
const (
	// Warm-up after the cold start.
	WarmUpTime = 5 * time.Second

	// Warm-up time after refresh.
	RefreshTime = 5 * time.Second

	// Stabilization time after discovery of new data.
	StabilizationTime = 1 * time.Second
)
