// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL Proxy

package escl

import (
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/transport"
	"github.com/OpenPrinting/go-mfp/util/missed"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Proxy is the forwarding eSCL proxy.
//
// It implements the http.Handler interface for the eSCL requests,
// forwards eSCL requests, represented as the http.Request  to the
// destination and responses in the reverse direction and rewrites
// the eSCL request and response bodies to properly translate URLs,
// embedded into the protocol messages.
type Proxy struct {
	localPath string             // Path portion of the local URL
	remoteURL *url.URL           // Remote URLs
	clnt      *Client            // eSCL client part of proxy
	urlxlat   *transport.URLXlat // URL translator
	//sniffer   Sniffer           // Sniffer callbacks
	hooks  ServerHooks   // eSCL server hooks
	seqnum atomic.Uint64 // Sequence number, for sniffer
}

// NewProxy creates the new [Proxy].
//
// The `clnt` is the client side of the proxy. If nil is passed,
// the client will be created automatically.
func NewProxy(localPath string, remoteURL *url.URL) *Proxy {
	localPath = transport.CleanURLPath(localPath + "/")

	localURL, err := url.Parse("http://localhost")
	assert.NoError(err)
	localURL.Path = localPath

	proxy := &Proxy{
		localPath: localPath,
		remoteURL: remoteURL,
		clnt:      NewClient(remoteURL, nil),
		urlxlat:   transport.NewURLXlat(localURL, remoteURL),
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
		proxy.reject(query, details, err)
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
		proxy.reject(query, details, err)
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
		proxy.reject(query, details, err)
		return
	}

	// Translate joburi
	joburi = proxy.reverseJobURI(query, joburi)

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

	println("getJobURINextDocument: joburi:", joburi)

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

	// Translate joburi
	joburi = proxy.forwardJobURI(query, joburi)

	// Forward request
	ctx := query.RequestContext()
	body, details, err := proxy.clnt.NextDocument(ctx, joburi)
	if err != nil {
		proxy.reject(query, details, err)
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

	// Translate joburi
	joburi = proxy.forwardJobURI(query, joburi)

	// Forward request
	ctx := query.RequestContext()
	details, err := proxy.clnt.Cancel(ctx, joburi)
	if err != nil {
		proxy.reject(query, details, err)
		return
	}

	// Send the response
	query.WriteHeader(details.StatusCode)
}

// reject rejects the query due to the error, returned by proxy.clnt.
//
// Note, details may be nil.
func (proxy *Proxy) reject(query *transport.ServerQuery,
	details *HTTPDetails, err error) {

	status := http.StatusServiceUnavailable
	if details != nil {
		status = details.StatusCode
	}

	query.Reject(status, err)
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
func (proxy *Proxy) forwardJobURI(query *transport.ServerQuery,
	joburi string) string {

	translated := proxy.urlxlat.ForwardPath(joburi)

	ctx := query.RequestContext()
	log.Begin(ctx).
		Debug("eSCL: JobUri translated:").
		Debug("  - %s", joburi).
		Debug("  + %s", translated).
		Commit()

	return translated
}

// reverseJobURI translates the JobUri in the remote->local direction.
func (proxy *Proxy) reverseJobURI(query *transport.ServerQuery,
	joburi string) string {

	translated := proxy.urlxlat.ReversePath(joburi)

	ctx := query.RequestContext()
	log.Begin(ctx).
		Debug("eSCL: JobUri translated:").
		Debug("  - %s", joburi).
		Debug("  + %s", translated).
		Commit()

	return translated
}
