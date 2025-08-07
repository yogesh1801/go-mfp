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

// Keys for context.WithValue and context.Value, used by logger.
var (
	// contextKeyLogger specified a Logger, associated with the Context.
	contextKeyLogger = contextKey{"log-logger"}

	// contextKeyLogger specified a prefix, associated with the Context.
	contextKeyPrefix = contextKey{"log-prefix"}
)

// Wrap Context key and values into unexported structures, so Context
// with Logger cannot be constructed outside of this package.
type (
	contextKey         struct{ name string }
	contextValueLogger struct{ logger *Logger }
	contextValuePrefix struct{ prefix string }
)

// NewContext returns new [context.Context] with the associated [Logger].
func NewContext(parent context.Context, dest *Logger) context.Context {
	return context.WithValue(parent,
		contextKeyLogger, contextValueLogger{dest})
}

// WithPrefix returns a new [context.Context] with the associated prefix.
func WithPrefix(parent context.Context, prefix string) context.Context {
	return context.WithValue(parent,
		contextKeyPrefix, contextValuePrefix{prefix})
}

// CtxLogger returns a [Logger] associated with the [context.Context].
// If no Logger is available, [FatalLogger] will be returned.
//
// Note, [context.Context] parameter may be safely passed as nil.
func CtxLogger(ctx context.Context) *Logger {
	if ctx != nil {
		v := ctx.Value(contextKeyLogger)
		if v != nil {
			ctxv, ok := v.(contextValueLogger)
			if ok {
				return ctxv.logger
			}
		}
	}

	return FatalLogger
}

// CtxPrefix returns a log prefix associated with the [context.Context].
// If no Logger is available, an empty string ("") will be returned.
//
// Note, [context.Context] parameter may be safely passed as nil.
func CtxPrefix(ctx context.Context) string {
	if ctx != nil {
		v := ctx.Value(contextKeyPrefix)
		if v != nil {
			ctxv, ok := v.(contextValuePrefix)
			if ok {
				return ctxv.prefix
			}
		}
	}

	return ""
}
