// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP Query context

package transport

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ServerQuery maintains the HTTP [Server] query processing,
// allowing per-request centralized logging and hooking.
//
// It's interface rules is stronger that of the [http.ResponseWriter]:
//   - [ServerQuery.WriteHeader] must be explicitly called before
//     the body can be written.
//   - [ServerQuery.WriteHeader] cannot be called twice.
//
// It keeps the reference to the original [http.Request] and wraps
// the corresponding [http.ResponseWriter], passed to the
// [http.Handler.ServeHTTP].
type ServerQuery struct {
	log    *log.Record         // Log record for the query
	rq     *http.Request       // Incoming request
	w      http.ResponseWriter // Underlying http.ResponseWriter
	status atomic.Int32        // HTTP status, 0 if not known yet
}

// NewServerQuery returns the new [ServerQuery].
func NewServerQuery(w http.ResponseWriter, rq *http.Request) *ServerQuery {
	ctx := rq.Context()
	query := &ServerQuery{
		log: log.Begin(ctx),
		rq:  rq,
		w:   w,
	}

	return query
}

// Request returns the underlying Request.
// Caller should not modify the request obtained this way.
func (query *ServerQuery) Request() *http.Request {
	return query.rq
}

// RequestContext returns the Request context.
func (query *ServerQuery) RequestContext() context.Context {
	return query.rq.Context()
}

// RequestURL returns the Request URL.
//
// The returned URL is taken unmodified from the underlying [http.Request],
// so in most cases it will not contain scheme and post parts.
func (query *ServerQuery) RequestURL() *url.URL {
	return query.rq.URL
}

// RequestFullURL returns the Request full URL, with the reconstructed
// Scheme and Host parts,
func (query *ServerQuery) RequestFullURL() *url.URL {
	u := URLClone(query.rq.URL)
	u.Scheme = query.RequestScheme()
	u.Host = query.RequestHost()

	return u
}

// RequestMethod returns the Request Method.
func (query *ServerQuery) RequestMethod() string {
	return query.rq.Method
}

// RequestScheme returns the Request Scheme.
func (query *ServerQuery) RequestScheme() string {
	if query.rq.TLS != nil {
		return "https"
	}
	return "http"
}

// RequestHost returns the Request Host.
func (query *ServerQuery) RequestHost() string {
	return query.rq.Host
}

// RequestContentType returns the Request content type,
// normalized to the lowercase.
func (query *ServerQuery) RequestContentType() string {
	return strings.ToLower(query.rq.Header.Get("Content-Type"))
}

// RequestContentLength returns the Request content Length.
func (query *ServerQuery) RequestContentLength() int64 {
	return query.rq.ContentLength
}

// RequestHeader returns http.Header of the request
func (query *ServerQuery) RequestHeader() http.Header {
	return query.rq.Header
}

// Finish must be called when query processing is finished
func (query *ServerQuery) Finish() {
	query.log.Commit()
}

// RequestBody returns body of the http.Request
func (query *ServerQuery) RequestBody() io.ReadCloser {
	return query.rq.Body
}

// ResponseHeader returns http.Header of the response
func (query *ServerQuery) ResponseHeader() http.Header {
	return query.w.Header()
}

// ResponseStatus returns the HTTP response status of the query.
// It returns 0, if status is not yet set.
func (query *ServerQuery) ResponseStatus() int {
	return int(query.status.Load())
}

// IsStatusSet returns true if the HTTP response status is set.
func (query *ServerQuery) IsStatusSet() bool {
	return query.ResponseStatus() != 0
}

// assertStatusSet panics if HTTP response status is not set.
func (query *ServerQuery) assertStatusSet() {
	assert.MustMsg(query.IsStatusSet(),
		"ServerQuery.WriteHeader: HTTP status is not set")
}

// assertStatusNotSet panics if HTTP response status is set.
func (query *ServerQuery) assertStatusNotSet() {
	assert.MustMsg(!query.IsStatusSet(),
		"ServerQuery.WriteHeader: HTTP status already set")
}

// Write writes response body bytes.
func (query *ServerQuery) Write(data []byte) (int, error) {
	query.assertStatusSet()
	return query.w.Write(data)
}

// WriteHeader writes HTTP response header.
func (query *ServerQuery) WriteHeader(status int) {
	assert.MustMsg(status != 0,
		"ServerQuery.WriteHeader: invalid HTTP status %d", status)

	if query.status.CompareAndSwap(0, int32(status)) {
		query.w.WriteHeader(status)
		query.log.Debug("HTTP-SRVR %s %s -- %d %s",
			query.rq.Method, query.rq.URL,
			status, http.StatusText(status))
		query.log.Flush()
		return
	}

	query.assertStatusNotSet()
}

// NoCache set response headers to disable client-side response cacheing.
func (query *ServerQuery) NoCache() {
	hdr := query.ResponseHeader()
	hdr.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	hdr.Set("Pragma", "no-cache")
	hdr.Set("Expires", "0")
}

// Reject completes request with a error.
//
// The error parameter is sent as the text body of the HTTP
// response. If it is nil, the reasonable default will be provided.
func (query *ServerQuery) Reject(status int, err error) {
	query.ResponseHeader().Set("Content-Type", "text/plain; charset=utf-8")
	query.NoCache()
	query.WriteHeader(status)

	if err == nil {
		err = errors.New(http.StatusText(status))
	}

	s := fmt.Sprintf("%3.3d %s\n", status, err)
	query.Write([]byte(s))
	query.Write([]byte("\n"))
}

// Created completes request with the http.StatusCreated
// status and Location: URL
func (query *ServerQuery) Created(location string) {
	scheme := "http"
	if query.rq.TLS != nil {
		scheme = "https"
	}

	location = scheme + "://" + query.rq.Host + location

	query.ResponseHeader().Set("Location", location)
	query.WriteHeader(http.StatusCreated)
}

// SendXML sends the XML response.
func (query *ServerQuery) SendXML(
	status int, ns xmldoc.Namespace, xml xmldoc.Element) {

	query.ResponseHeader().Set("Content-Type", "text/xml")
	query.WriteHeader(status)
	xml.EncodeIndent(query, ns, "  ")
}

// SendData sends the binary data, represented as [io.Reader].
func (query *ServerQuery) SendData(
	status int, contentType string, data io.Reader) {

	query.ResponseHeader().Set("Content-Type", contentType)
	query.WriteHeader(status)
	io.Copy(query, data)
}
