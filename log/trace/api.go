// MFP - Miulti-Function Printers and scanners toolkit
// Protocol tracer
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Public API

package trace

import (
	"io"

	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/transport"
)

// OnRequest records a request message to the trace log and tracer file.
//
// It logs the message at [log.LevelTrace] and writes full details to the
// tracer file using the logger and tracer from the [query.RequestContext].
//
// To intercept the request body, the tracer may wrap it in a custom
// [io.ReadCloser]. The caller MUST replace their original body with the
// one returned by this function to ensure data is properly captured.
//
// An empty body may be represented by nil.
func OnRequest(query *transport.ServerQuery,
	msg Message, body io.ReadCloser) io.ReadCloser {

	ctx := query.RequestContext()
	writer := CtxWriter(ctx)

	log.CtxLogger(ctx).Begin("").
		Trace("%s request received:\n%s",
			msg.Protocol(), query.DumpRequest()).
		Object(log.LevelTrace, 0, msg).
		Commit()

	if writer != nil {
		cc := body
		if body != nil {
			body, cc = transport.TeeReadCloser2(body)
		}

		writer.OnRequest(query, msg, cc)
	}

	return body
}

// OnResponse records a response message to the trace log and tracer file.
//
// It logs the message at [log.LevelTrace] and writes full details to the
// tracer file using the logger and tracer from the [query.RequestContext].
//
// To intercept the response body, the tracer may wrap it in a custom
// [io.ReadCloser]. The caller MUST replace their original body with the
// one returned by this function to ensure data is properly captured.
//
// An empty body may be represented by nil.
func OnResponse(query *transport.ServerQuery,
	msg Message, body io.ReadCloser) io.ReadCloser {

	ctx := query.RequestContext()
	writer := CtxWriter(ctx)

	log.CtxLogger(ctx).Begin("").
		Trace("%s response sent:\n%s",
			msg.Protocol(), query.DumpResponse()).
		Object(log.LevelTrace, 0, msg).
		Commit()

	if writer != nil {
		cc := body
		if body != nil {
			body, cc = transport.TeeReadCloser2(body)
		}

		writer.OnResponse(query, msg, cc)
	}

	return body
}
