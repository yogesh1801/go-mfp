// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP Proxy

package ipp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/log/trace"
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/goipp"
)

// Proxy is the forwarding IPP proxy.
//
// It implements the http.Handler interface for the IPP requests,
// forwards IPP requests, represented as the http.Request  to the
// destination and responses in the reverse direction and rewrites
// the IPP request and response bodies to properly translate URLs,
// embedded into the protocol messages.
type Proxy struct {
	localPath string            // Path portion of the local URL
	remoteURL *url.URL          // Remote URLs
	clnt      *transport.Client // HTTP client part of proxy
}

// proxyMsgXlat performs URL translation in the IPP requests
// and responses.
type proxyMsgXlat struct {
	urlxlat *transport.URLXlat
}

// proxyMsgChanges contains changes applied to the message by the
// proxyMsgXlat.Forward or proxyMsgXlat.Reverse functions, for logging.
type proxyMsgChanges struct {
	local, remote *url.URL                 // Local and remote URLs
	Groups        []proxyMsgChangesByGroup // Changes per group
}

// proxyMsgChangesByGroup per-group changes
type proxyMsgChangesByGroup struct {
	Tag    goipp.Tag                // Group tag
	Values []proxyMsgChangesByValue // Changed values
}

// proxyMsgChangesByValue represents per-value changes
type proxyMsgChangesByValue struct {
	Path     string      // Path to the value from the Message root
	Old, New goipp.Value // Old and new values
}

// NewProxy creates the new [Proxy].
//
// The `clnt` is the client side of the proxy. If nil is passed,
// the client will be created automatically.
func NewProxy(localPath string, remoteURL *url.URL) *Proxy {
	proxy := &Proxy{
		localPath: localPath,
		remoteURL: remoteURL,
		clnt:      transport.NewClient(nil),
	}
	return proxy
}

// ServeHTTP handles incoming HTTP requests.
// It implements [http.Handler] interface.
func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	// Setup things
	query := transport.NewServerQuery(w, rq)
	ctx := query.RequestContext()

	// Validate the request
	switch query.RequestMethod() {
	case "POST":
	case "GET":
		query.Reject(http.StatusNotFound, nil)
		return
	default:
		query.Reject(http.StatusBadRequest, nil)
		return
	}

	if query.RequestContentType() != "application/ipp" {
		query.Reject(http.StatusBadRequest, nil)
		return
	}

	// Create goipp.Message translator
	xlat, err := proxy.newMsgXlat(query)
	if err != nil {
		log.Debug(ctx, "%s", err)
		query.Reject(http.StatusBadGateway, err)
		return
	}

	// Prepare outgoing request
	out, err := proxy.doRequest(query, xlat)
	if err != nil {
		err = fmt.Errorf("IPP error: %w", err)
		log.Debug(ctx, "%s", err)
		query.Reject(http.StatusBadGateway, err)
		return
	}

	// Execute outgoing request
	log.Debug(ctx, "IPP: forward request to: %s", out.URL)

	rsp, err := proxy.clnt.Do(out)
	if err != nil {
		log.Debug(ctx, "IPP: %s", err)
		query.Reject(http.StatusBadGateway, err)
		return
	}

	// Close response body when done.
	//
	// Note, rsp.Body may change during translation, hence the closure,
	// not the direct rsp.Body.Close() call.
	defer func() { rsp.Body.Close() }()

	// For non-IPP response we are just HTTP proxy
	ct := strings.ToLower(rsp.Header.Get("Content-Type"))
	if ct != "application/ipp" {
		transport.HTTPRemoveHopByHopHeaders(rsp.Header)
		transport.HTTPCopyHeaders(query.ResponseHeader(), rsp.Header)
		query.WriteHeader(rsp.StatusCode)
		io.Copy(query, rsp.Body)
		return
	}

	// Translate IPP response
	err = proxy.doResponse(query, rsp, xlat)
	if err != nil {
		log.Debug(ctx, "IPP: %s", err)
		query.Reject(http.StatusBadGateway, err)
	}
}

// doRequest performs (client->server) part of the IPP request handling
//
// It returns modified request ready to be send to the server or error.
func (proxy *Proxy) doRequest(query *transport.ServerQuery,
	xlat *proxyMsgXlat) (*http.Request, error) {

	// Fetch IPP Request message
	//
	// Note, for IPP requests with data attachment (e.g., Send-Document)
	// the data attachment immediately follows the IPP part.
	body := query.RequestBody()

	var msg goipp.Message
	var consumed transport.DiscardCounter

	ops := goipp.DecoderOptions{EnableWorkarounds: true}
	err := msg.DecodeEx(io.TeeReader(body, &consumed), ops)
	if err != nil {
		return nil, err
	}

	// Translate IPP message
	msg2, chg := xlat.Forward(&msg)

	// Notify tracer
	//
	// Note, we must close the wrapped body to notify tracer
	// that data transfer is done. From another hand, we should
	// not close the original body here; it is controlled by
	// the Proxy.ServeHTTP. So we protect it from being closed
	// by wrapping into io.NopCloser.
	body = trace.OnRequest(query, goippRequest{&msg}, io.NopCloser(body))
	defer body.Close()

	// Log the message changes
	if !chg.isEmpty() {
		ctx := query.RequestContext()
		log.Debug(ctx, "IPP: translated attributes:")
		log.Object(ctx, log.LevelDebug, 4, chg)
	}

	// Replace IPP part of the request body with the translated message
	msg2bytes, _ := msg2.EncodeBytes()
	body = io.NopCloser(io.MultiReader(bytes.NewReader(msg2bytes), body))

	// Setup outgoing request
	out := proxy.outreq(query, xlat, body)
	out.ContentLength = query.RequestContentLength()
	if out.ContentLength >= 0 {
		out.ContentLength += int64(len(msg2bytes))
		out.ContentLength -= consumed.Count
	}

	return out, nil
}

// doResponse performs (client->server) part of the IPP request handling
func (proxy *Proxy) doResponse(query *transport.ServerQuery,
	rsp *http.Response, xlat *proxyMsgXlat) error {

	// Fetch IPP response message
	body := rsp.Body

	var msg goipp.Message
	var consumed transport.DiscardCounter

	ops := goipp.DecoderOptions{EnableWorkarounds: true}
	err := msg.DecodeEx(io.TeeReader(body, &consumed), ops)
	if err != nil {
		return err
	}

	// Translate IPP response
	msg2, chg := xlat.Reverse(&msg)
	if !chg.isEmpty() {
		ctx := rsp.Request.Context()
		log.Debug(ctx, "IPP: translated attributes:")
		log.Object(ctx, log.LevelDebug, 4, chg)
	}

	// Replace http.Response body
	msg2bytes, _ := msg2.EncodeBytes()
	body = io.NopCloser(io.MultiReader(bytes.NewReader(msg2bytes), body))

	// Adjust rsp.ContentLength
	if rsp.ContentLength >= 0 {
		rsp.ContentLength += int64(len(msg2bytes))
		rsp.ContentLength -= consumed.Count
	}

	// Copy response headers and status to the client
	transport.HTTPRemoveHopByHopHeaders(rsp.Header)
	transport.HTTPCopyHeaders(query.ResponseHeader(), rsp.Header)

	if rsp.ContentLength >= 0 {
		rsp.Header.Set("Content-Length",
			strconv.FormatInt(rsp.ContentLength, 10))
	}

	query.WriteHeader(rsp.StatusCode)

	// Notify tracer
	body = trace.OnResponse(query, goippResponse{&msg}, body)
	rsp.Body = body

	// Forward response body
	io.Copy(query, rsp.Body)
	rsp.Body.Close()

	return nil
}

// outreq creates an outgoing HTTP request based on request
// received by the server side of proxy.
func (proxy *Proxy) outreq(query *transport.ServerQuery,
	xlat *proxyMsgXlat, body io.ReadCloser) *http.Request {

	target := xlat.urlxlat.Forward(query.RequestFullURL())

	// Create request
	out, _ := transport.NewRequest(
		query.RequestContext(),
		query.RequestMethod(),
		target,
		body)

	out.Header = query.RequestHeader().Clone()
	transport.HTTPRemoveHopByHopHeaders(out.Header)

	return out
}

// newMsgXlat returns the new translateMsg for the query.
func (proxy *Proxy) newMsgXlat(query *transport.ServerQuery) (
	*proxyMsgXlat, error) {

	// Guess Proxy's local (server) URL out of request.
	s := query.RequestScheme() + "://" + query.RequestHost()
	local, err := transport.ParseURL(s)
	if err != nil {
		err = fmt.Errorf("%q: can't parse local URL", s)
		return nil, err
	}

	local.Path = proxy.localPath

	// Fill the proxyMsgXlat structure
	xlat := &proxyMsgXlat{
		urlxlat: transport.NewURLXlat(local, proxy.remoteURL),
	}

	return xlat, nil
}

// Forward translates message in the forward (client->server)
// direction.
func (xlat *proxyMsgXlat) Forward(
	msg *goipp.Message) (*goipp.Message, proxyMsgChanges) {

	return xlat.translateMsg(msg, xlat.urlxlat.Forward)
}

// Forward translates message in the reverse (server->client)
// direction.
func (xlat *proxyMsgXlat) Reverse(
	msg *goipp.Message) (*goipp.Message, proxyMsgChanges) {

	return xlat.translateMsg(msg, xlat.urlxlat.Reverse)
}

// translateMsg performs the actual goipp.Message translation.
//
// It returns the translated goipp.Message and a set of applied
// changes.
//
// Each found URL is translated using the provided `callback` function.
func (xlat *proxyMsgXlat) translateMsg(msg *goipp.Message,
	callback func(*url.URL) *url.URL) (*goipp.Message, proxyMsgChanges) {

	chgmsg := proxyMsgChanges{
		local:  xlat.urlxlat.Local(),
		remote: xlat.urlxlat.Remote(),
	}

	// Obtain a deep copy of all message attributes, packed
	// into groups. Roll over all attributes, translating
	// values in place.
	groups := msg.AttrGroups().DeepCopy()
	for i := range groups {
		group := &groups[i]
		chggrp := proxyMsgChangesByGroup{Tag: group.Tag}

		for j := range group.Attrs {
			attr := &group.Attrs[j]
			chg := xlat.translateAttr(attr, callback)
			chggrp.Values = append(chggrp.Values, chg...)
		}

		if len(chggrp.Values) > 0 {
			chgmsg.Groups = append(chgmsg.Groups, chggrp)
		}
	}

	// Rebuild the message
	msg2 := goipp.NewMessageWithGroups(msg.Version, msg.Code,
		msg.RequestID, groups)

	return msg2, chgmsg
}

// translateAttr translates URLs found in the goipp.Attribute, recursively
// scanning nested collections.
//
// Each found URL is translated using the provided `callback` function.
//
// Translation is performed "in place".
func (xlat *proxyMsgXlat) translateAttr(attr *goipp.Attribute,
	callback func(*url.URL) *url.URL) []proxyMsgChangesByValue {

	chg := []proxyMsgChangesByValue{}

	for i := range attr.Values {
		v := &attr.Values[i]
		morechg := xlat.translateVal(&v.V, v.T, callback)

		for _, c := range morechg {
			path := attr.Name
			if len(attr.Values) > 1 {
				path += fmt.Sprintf("[%d]", i)
			}

			if c.Path != "" && len(attr.Values) == 0 {
				path += "."
			}

			c.Path = path + c.Path

			chg = append(chg, c)
		}
	}

	return chg
}

// translateVal translates URLs in the goipp.Value, recursively
// scanning nested collections.
//
// Each found URL is translated using the provided `callback` function.
//
// Translation is performed "in place".
func (xlat *proxyMsgXlat) translateVal(v *goipp.Value, t goipp.Tag,
	callback func(*url.URL) *url.URL) []proxyMsgChangesByValue {

	switch oldval := (*v).(type) {
	case goipp.Collection:
		chg := []proxyMsgChangesByValue{}

		for i := range oldval {
			attr := &oldval[i]
			morechg := xlat.translateAttr(attr, callback)
			chg = append(chg, morechg...)
		}

		return chg

	case goipp.String:
		if t != goipp.TagURI {
			return nil
		}

		u, err := transport.ParseURL(string(oldval))
		if err == nil {
			u2 := callback(u)
			newval := goipp.String(u2.String())

			if oldval != newval {
				*v = goipp.String(u2.String())

				chg := []proxyMsgChangesByValue{
					{Old: oldval, New: newval},
				}

				return chg
			}
		}
	}

	return nil
}

// isEmpty reports if proxyMsgChanges contains no changes.
func (chg proxyMsgChanges) isEmpty() bool {
	return len(chg.Groups) == 0
}

// MarshalLog returns string representation of proxyMsgChanges for logging.
// It implements [log.Marshaler] interface.
func (chg proxyMsgChanges) MarshalLog() []byte {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Local URL:  %s\n", chg.local)
	fmt.Fprintf(&buf, "Remote URL: %s\n", chg.remote)
	fmt.Fprintf(&buf, "\n")

	for _, g := range chg.Groups {
		fmt.Fprintf(&buf, "GROUP %s:\n", g.Tag)
		for _, v := range g.Values {
			fmt.Fprintf(&buf, "    ATTR %s:\n", v.Path)
			fmt.Fprintf(&buf, "        - %s\n", v.Old)
			fmt.Fprintf(&buf, "        + %s\n", v.New)
		}
	}

	return buf.Bytes()
}
