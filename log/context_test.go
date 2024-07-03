// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// context.Context with logging test

package log

import (
	"context"
	"testing"
)

// TestContext tests context.Context with logging
func TestContext(t *testing.T) {
	dest := &Logger{}

	// Create a valid Context, test that expected logger returned
	ctx := NewContext(context.Background(), dest)
	dest2 := CtxLogger(ctx)

	if dest2 != dest {
		t.Errorf("CtxLogger: wrong Logger returned")
	}

	// Use CtxLogger() with alien Context (context not created
	// with the NewContext() function)
	ctx = context.Background()
	dest2 = CtxLogger(ctx)

	if dest2 != DefaultLogger {
		t.Errorf("CtxLogger: alien context must return DefaultLogger")
	}

	// Use CtxLogger() with Context that has invalid value associated
	// with the valid key
	ctx = context.WithValue(context.Background(), ContextKey, 5)
	dest2 = CtxLogger(ctx)

	if dest2 != DefaultLogger {
		t.Errorf("CtxLogger: incompatible context must return DefaultLogger")
	}
}
