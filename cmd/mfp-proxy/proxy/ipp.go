// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Functions for IPP messages and attributes

package proxy

import (
	"net/url"

	"github.com/OpenPrinting/goipp"
	"github.com/alexpevzner/mfp/transport"
)

// ippMsgXlatURLs translates URLs in the IPP message attributes,
// recursively scanning entire message.
//
// Each found URL is translated using the provided `xlat` function.
func ippMsgXlatURLs(msg *goipp.Message,
	xlat func(*url.URL) *url.URL) *goipp.Message {

	// Obtain a deep copy of all message attributes, packed
	// into groups. Roll over all attributes, translating
	// values in place.
	groups := msg.AttrGroups().DeepCopy()
	for i := range groups {
		group := &groups[i]
		for j := range group.Attrs {
			attr := &group.Attrs[j]
			ippAttrXlatURLs(attr, xlat)
		}
	}

	// Rebuild the message
	msg2 := goipp.NewMessageWithGroups(msg.Version, msg.Code,
		msg.RequestID, groups)

	return msg2
}

// ippValXlatURLs translates URLs in the goipp.Attribute, recursively
// scanning nested collections.
//
// Each found URL is translated using the provided `xlat` function.
//
// Translation is performed "in place".
func ippAttrXlatURLs(attr *goipp.Attribute, xlat func(*url.URL) *url.URL) {
	for i := range attr.Values {
		v := &attr.Values[i]
		ippValXlatURLs(&v.V, v.T, xlat)
	}
}

// ippValXlatURLs translates URLs in the goipp.Value, recursively
// scanning nested collections.
//
// Each found URL is translated using the provided `xlat` function.
//
// Translation is performed "in place".
func ippValXlatURLs(v *goipp.Value, t goipp.Tag, xlat func(*url.URL) *url.URL) {
	switch v2 := (*v).(type) {
	case goipp.Collection:
		for i := range v2 {
			attr := &v2[i]
			ippAttrXlatURLs(attr, xlat)
		}

	case goipp.String:
		if t != goipp.TagURI {
			return
		}

		u, err := transport.ParseURL(string(v2))
		if err == nil {
			u2 := xlat(u)
			*v = goipp.String(u2.String())
		}
	}
}
