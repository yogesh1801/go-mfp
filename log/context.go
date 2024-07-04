// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// context.Context with logging

package log

import (
	"context"
)

// ContextKey is a context key. It is used to specify [Logger]
// associated with the [context.Context].
var ContextKey = contextKey{"log-dest"}

// Wrap Context key and value into unexported structures, so Context
// with Logger cannot be constructed outside of this package.
type (
	contextKey   struct{ name string }
	contextValue struct{ *Logger }
)

// NewContext returns new [context.Context] with the associated [Logger].
func NewContext(parent context.Context, dest *Logger) context.Context {
	return context.WithValue(parent, ContextKey, contextValue{dest})
}

// CtxLogger returns a [Logger] associated with the [context.Context].
// If no Logger is available, [DiscardLogger] will be returned.
//
// Note, [context.Context] parameter may be safely passed as nil.
func CtxLogger(ctx context.Context) *Logger {
	if ctx != nil {
		v := ctx.Value(ContextKey)
		if v != nil {
			ctxv, ok := v.(contextValue)
			if ok {
				return ctxv.Logger
			}
		}
	}

	return DiscardLogger
}
