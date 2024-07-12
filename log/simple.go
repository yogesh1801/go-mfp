// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Simple logging API

package log

import (
	"context"
	"encoding"
)

// Trace writes a Trace-level message to the [Logger] associated
// with the Context.
//
// If Logger is not available, [DefaultLogger] will be used.
// The [context.Context] parameter may be safely passed as nil.
func Trace(ctx context.Context, format string, v ...any) {
	CtxLogger(ctx).Trace(format, v...)
}

// Debug writes a Debug-level message to the [Logger] associated
// with the Context.
//
// If Logger is not available, [DefaultLogger] will be used.
// The [context.Context] parameter may be safely passed as nil.
func Debug(ctx context.Context, format string, v ...any) {
	CtxLogger(ctx).Debug(format, v...)
}

// Info writes a Info-level message to the [Logger] associated
// with the Context.
//
// If Logger is not available, [DefaultLogger] will be used.
// The [context.Context] parameter may be safely passed as nil.
func Info(ctx context.Context, format string, v ...any) {
	CtxLogger(ctx).Info(format, v...)
}

// Warning writes a Warning-level message to the [Logger] associated
// with the Context.
//
// If Logger is not available, [DefaultLogger] will be used.
// The [context.Context] parameter may be safely passed as nil.
func Warning(ctx context.Context, format string, v ...any) {
	CtxLogger(ctx).Warning(format, v...)
}

// Error writes a Error-level message to the [Logger] associated
// with the Context.
//
// If Logger is not available, [DefaultLogger] will be used.
// The [context.Context] parameter may be safely passed as nil.
func Error(ctx context.Context, format string, v ...any) {
	CtxLogger(ctx).Error(format, v...)
}

// Fatal writes a Fatal-level message to the [Logger] associated
// with the Context.
//
// It calls os.Exit(1) and never returns.
//
// If Logger is not available, [DefaultLogger] will be used.
// The [context.Context] parameter may be safely passed as nil.
func Fatal(ctx context.Context, format string, v ...any) {
	CtxLogger(ctx).Fatal(format, v...)
}

// Object writes any object that implements [encoding.TextMarshaler]
// interface to the Logger associated with the Context.
//
// If Logger is not available, [DefaultLogger] will be used.
// The [context.Context] parameter may be safely passed as nil.
//
// If [encoding.TextMarshaler.MarshalText] returns an error, it
// will be written to log with the [Error] log level, regardless
// of the level specified by the first parameter.
func Object(ctx context.Context,
	level Level, obj encoding.TextMarshaler) context.Context {
	CtxLogger(ctx).Object(level, obj)
	return ctx
}
