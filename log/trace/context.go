// MFP - Miulti-Function Printers and scanners toolkit
// Protocol tracer
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// context.Context with protocol trace

package trace

import (
	"context"
)

var (
	contextKeyTracer = contextKey{"tracer"}
)

// Wrap Context key and values into unexported structures, so Context
// with Tracer cannot be constructed outside of this package.
type (
	contextKey         struct{ name string }
	contextValueTracer struct{ tracer *Writer }
)

// NewContext returns new [context.Context] with the associated tracer.
func NewContext(parent context.Context, tracer *Writer) context.Context {
	return context.WithValue(parent,
		contextKeyTracer, contextValueTracer{tracer})
}

// CtxWriter returns a trace [Writer] associated with the [context.Context].
// If there is no associated Writer, nil will be returned.
func CtxWriter(ctx context.Context) *Writer {
	if ctx != nil {
		switch v := ctx.Value(contextKeyTracer).(type) {
		case contextValueTracer:
			return v.tracer
		}
	}

	return nil
}

// Enabled reports if trace is enabled on the context.
func Enabled(ctx context.Context) bool {
	return CtxWriter(ctx) != nil
}
