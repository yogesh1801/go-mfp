// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL Proxy

package escl

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/missed"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
	"github.com/OpenPrinting/goipp"
)

// Proxy is the forwarding eSCL proxy.
//
// It implements the http.Handler interface for the eSCL requests,
// forwards eSCL requests, represented as the http.Request  to the
// destination and responses in the reverse direction and rewrites
// the eSCL request and response bodies to properly translate URLs,
// embedded into the protocol messages.
type Proxy struct {
	localPath string   // Path portion of the local URL
	remoteURL *url.URL // Remote URLs
	clnt      *Client  // eSCL client part of proxy
	//sniffer   Sniffer           // Sniffer callbacks
	hooks  ServerHooks   // eSCL server hooks
	seqnum atomic.Uint64 // Sequence number, for sniffer
}

// proxyMsgXlat performs URL translation in the eSCL requests
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
	localPath = transport.CleanURLPath(localPath + "/")

	proxy := &Proxy{
		localPath: localPath,
		remoteURL: remoteURL,
		clnt:      NewClient(remoteURL, nil),
	}
	return proxy
}

// Sniff installs the sniffer callback.
//
// Don't use this function when proxy is already active (i.e., concurrently
// with the [Proxy.ServeHTTP], it can cause race conditions.
//func (proxy *Proxy) Sniff(sniffer Sniffer) {
//	proxy.sniffer = sniffer
//}

// ServeHTTP handles incoming HTTP requests.
// It implements [http.Handler] interface.
func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	// Create a transport.ServerQuery
	query := transport.NewServerQuery(w, rq)
	defer query.Finish()

	// Call the OnHTTPRequest hook
	if proxy.hooks.OnHTTPRequest != nil {
		proxy.hooks.OnHTTPRequest(query)
		if query.IsStatusSet() {
			return
		}
	}

	// Dispatch the request
	if !strings.HasPrefix(query.RequestURL().Path, proxy.localPath) {
		query.Reject(http.StatusNotFound, nil)
		return
	}

	path := query.RequestURL().Path
	subpath, _ := missed.StringsCutPrefix(path, proxy.localPath)
	method := query.RequestMethod()

	// Dispatch the request
	var action func(*transport.ServerQuery)

	const NextDocument = "/NextDocument"
	const ScanImageInfo = "/ScanImageInfo"

	switch {
	// Handle {root}-relative requests
	case method == "GET" && subpath == "ScannerCapabilities":
		action = proxy.getScannerCapabilities

	case method == "GET" && subpath == "ScannerStatus":
		action = proxy.getScannerStatus

	case method == "POST" && subpath == "ScanJobs":
		action = proxy.postScanJobs

	// Handle {JobUri}-relative requests
	case method == "GET" && strings.HasSuffix(path, NextDocument):
		joburi := path[:len(path)-len(NextDocument)]
		action = func(*transport.ServerQuery) {
			proxy.getJobURINextDocument(query, joburi)
		}

	case method == "GET" && strings.HasSuffix(path, ScanImageInfo):
		joburi := path[:len(path)-len(ScanImageInfo)]
		action = func(*transport.ServerQuery) {
			proxy.getJobURIScanImageInfo(query, joburi)
		}

	case method == "DELETE":
		action = func(*transport.ServerQuery) {
			proxy.deleteJobURI(query, path)
		}
	}

	if action != nil {
		action(query)
	} else {
		query.Reject(http.StatusNotFound, nil)
	}

}

// getScannerCapabilities handles GET /{root}/ScannerCapabilities request
func (proxy *Proxy) getScannerCapabilities(query *transport.ServerQuery) {
	// Call OnScannerCapabilitiesRequest hook
	if proxy.hooks.OnScannerCapabilitiesRequest != nil {
		proxy.hooks.OnScannerCapabilitiesRequest(query)
		if query.IsStatusSet() {
			return
		}
	}

	// Forward request
	ctx := query.RequestContext()
	caps, details, err := proxy.clnt.GetScannerCapabilities(ctx)

	if err != nil {
		query.Reject(details.StatusCode, err)
		return
	}

	// Call OnScannerCapabilitiesResponse hook
	if proxy.hooks.OnScannerCapabilitiesResponse != nil {
		caps2 := proxy.hooks.OnScannerCapabilitiesResponse(
			query, caps)
		if query.IsStatusSet() {
			return
		}

		if caps2 != nil {
			caps = caps2
		}
	}

	// Generate and send XML response
	proxy.sendXML(query, HookScannerCapabilities, caps)
}

// getScannerStatus handles GET /{root}/ScannerStatus request
func (proxy *Proxy) getScannerStatus(query *transport.ServerQuery) {
	// Call OnScannerStatusRequest hook
	if proxy.hooks.OnScannerStatusRequest != nil {
		proxy.hooks.OnScannerStatusRequest(query)
		if query.IsStatusSet() {
			return
		}
	}

	// Forward request
	ctx := query.RequestContext()
	status, details, err := proxy.clnt.GetScannerStatus(ctx)

	if err != nil {
		query.Reject(details.StatusCode, err)
		return
	}

	// Call OnScannerStatusResponse hook
	if proxy.hooks.OnScannerStatusResponse != nil {
		status2 := proxy.hooks.OnScannerStatusResponse(
			query, status)
		if query.IsStatusSet() {
			return
		}

		if status2 != nil {
			status = status2
		}
	}

	// Generate and send XML response
	proxy.sendXML(query, HookScannerStatus, status)
}

// postScanJobs handles POST /{root}/ScanJobs
func (proxy *Proxy) postScanJobs(query *transport.ServerQuery) {
	// Fetch the XML request body
	xml, err := xmldoc.Decode(NsMap, query.RequestBody())
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

	// Call OnXMLRequest hook
	if proxy.hooks.OnXMLRequest != nil {
		xml2 := proxy.hooks.OnXMLRequest(query, HookScanJobs, xml)
		if query.IsStatusSet() {
			return
		}

		if !xml2.IsZero() {
			xml = xml2
		}
	}

	// Decode ScanSettings request
	ss, err := DecodeScanSettings(xml)
	if err != nil {
		query.Reject(http.StatusBadRequest, err)
		return
	}

	// Call OnScanJobsRequest hook
	if proxy.hooks.OnScanJobsRequest != nil {
		ss2 := proxy.hooks.OnScanJobsRequest(query, ss)
		if query.IsStatusSet() {
			return
		}

		if ss2 != nil {
			ss = ss2
		}
	}

	// Forward request
	ctx := query.RequestContext()
	joburi, details, err := proxy.clnt.Scan(ctx, *ss)

	if err != nil {
		query.Reject(details.StatusCode, err)
		return
	}

	joburi = proxy.reverseJobURI(joburi)

	// Call OnScanJobsResponse hook
	if proxy.hooks.OnScanJobsResponse != nil {
		joburi2 := proxy.hooks.OnScanJobsResponse(query, ss)
		if query.IsStatusSet() {
			return
		}

		if joburi2 != "" {
			joburi = joburi2
		}
	}

	// Complete the request
	query.Created(joburi)
}

// getJobURINextDocument handles GET /{JobUri}/NextDocument
func (proxy *Proxy) getJobURINextDocument(
	query *transport.ServerQuery, joburi string) {

	// Call OnNextDocumentRequest hook
	if proxy.hooks.OnNextDocumentRequest != nil {
		joburi2 := proxy.hooks.OnNextDocumentRequest(
			query, joburi)
		if query.IsStatusSet() {
			return
		}

		if joburi2 != "" {
			joburi = joburi2
		}
	}

	// Forward request
	ctx := query.RequestContext()
	joburi = proxy.forwardJobURI(joburi)

	body, details, err := proxy.clnt.NextDocument(ctx, joburi)
	if err != nil {
		query.Reject(details.StatusCode, err)
		return
	}

	defer body.Close()

	// Call OnNextDocumentResponse hook
	if proxy.hooks.OnNextDocumentResponse != nil {
		body2 := proxy.hooks.OnNextDocumentResponse(
			query, body)
		if query.IsStatusSet() {
			return
		}

		if body2 != nil {
			body = body2
		}
	}

	// Send the response
	query.SendData(http.StatusOK, details.ContentType, body)
}

// getJobURIScanImageInfo handles GET /{JobUri}/ScanImageInfo
func (proxy *Proxy) getJobURIScanImageInfo(
	query *transport.ServerQuery, joburi string) {
	query.Reject(http.StatusNotImplemented, nil)
}

// deleteJobURI handles DELETE /{JobUri}
func (proxy *Proxy) deleteJobURI(
	query *transport.ServerQuery, joburi string) {

	// Call OnDeleteRequest hook
	if proxy.hooks.OnDeleteRequest != nil {
		joburi2 := proxy.hooks.OnDeleteRequest(query, joburi)
		if query.IsStatusSet() {
			return
		}

		if joburi2 != "" {
			joburi = joburi2
		}
	}

	// Forward request
	ctx := query.RequestContext()
	joburi = proxy.forwardJobURI(joburi)

	details, err := proxy.clnt.Cancel(ctx, joburi)
	if err != nil {
		query.Reject(details.StatusCode, err)
		return
	}

	// Send the response
	query.WriteHeader(details.StatusCode)
}

// sendXML generates and sends the XML response to the query.
func (proxy *Proxy) sendXML(query *transport.ServerQuery,
	action HookAction, rsp interface{ ToXML() xmldoc.Element }) {

	xml := rsp.ToXML()
	if proxy.hooks.OnXMLResponse != nil {
		xml2 := proxy.hooks.OnXMLResponse(query, action, xml)

		if query.IsStatusSet() {
			return
		}

		if !xml2.IsZero() {
			xml = xml2
		}
	}

	query.SendXML(http.StatusOK, NsMap, xml)
}

// forwardJobURI translates the JobUri in the local->remote direction.
func (proxy *Proxy) forwardJobURI(joburl string) string {
	return joburl
}

// reverseJobURI translates the JobUri in the remote->local direction.
func (proxy *Proxy) reverseJobURI(joburl string) string {
	return joburl
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
