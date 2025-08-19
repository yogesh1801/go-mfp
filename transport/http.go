// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Helper functions for HTTP

package transport

import (
	"net/http"
	"strings"
)

// HTTPRemoveHopByHopHeaders removes HTTP hop-by-hop headers,
// per [RFC 7230, section 6.1].
//
// [RFC 7230, section 6.1]: https://www.rfc-editor.org/rfc/rfc7230.html#section-6.1
func HTTPRemoveHopByHopHeaders(hdr http.Header) {
	// Per RFC 7230, section 6.1:
	//
	// Hence, the Connection header field provides a declarative way of
	// distinguishing header fields that are only intended for the immediate
	// recipient ("hop-by-hop") from those fields that are intended for all
	// recipients on the chain ("end-to-end"), enabling the message to be
	// self-descriptive and allowing future connection-specific extensions
	// to be deployed without fear that they will be blindly forwarded by
	// older intermediaries.
	if c := hdr.Get("Connection"); c != "" {
		for _, f := range strings.Split(c, ",") {
			if f = strings.TrimSpace(f); f != "" {
				hdr.Del(f)
			}
		}
	}

	// These headers are always considered hop-by-hop.
	for _, c := range []string{"Connection", "Keep-Alive",
		"Proxy-Authenticate", "Proxy-Connection",
		"Proxy-Authorization", "Te", "Trailer", "Transfer-Encoding"} {
		hdr.Del(c)
	}
}

// HTTPCopyHeaders copies HTTP headers from src to dst.
//
// Headers that already present in the `dst` but not present
// in the `src` are not touched.
func HTTPCopyHeaders(dst, src http.Header) {
	for k, v := range src {
		if strings.ToLower(k) != "content-length" {
			dst[k] = v
		}
	}
}

// HTTPPurgeHeaders removes all headers from the [http.Header].
func HTTPPurgeHeaders(hdr http.Header) {
	for key := range hdr {
		delete(hdr, key)
	}
}
