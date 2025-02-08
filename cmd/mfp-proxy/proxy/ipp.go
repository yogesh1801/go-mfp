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

// msgXlat translates URLs embedded in the IPP message attributes
// when message is being forwarded via proxy.
type msgXlat struct {
	urlxlat *transport.URLXlat // URL translator
}

// newMsgXlat creates a new msgXlat
func newMsgXlat(urlxlat *transport.URLXlat) *msgXlat {
	return &msgXlat{
		urlxlat: urlxlat,
	}
}

// Forward translates message in the forward (client->server)
// direction.
func (mx *msgXlat) Forward(msg *goipp.Message) *goipp.Message {
	return mx.translateMsg(msg, mx.urlxlat.Forward)
}

// Forward translates message in the reverse (server->client)
// direction.
func (mx *msgXlat) Reverse(msg *goipp.Message) *goipp.Message {
	return mx.translateMsg(msg, mx.urlxlat.Reverse)
}

// translateMsg performs the actual goipp.Message translation.
//
// Each found URL is translated using the provided `xlat` function.
func (mx *msgXlat) translateMsg(msg *goipp.Message,
	xlat func(*url.URL) *url.URL) *goipp.Message {

	// Obtain a deep copy of all message attributes, packed
	// into groups. Roll over all attributes, translating
	// values in place.
	groups := msg.AttrGroups().DeepCopy()
	for i := range groups {
		group := &groups[i]
		for j := range group.Attrs {
			attr := &group.Attrs[j]
			mx.translateAttr(attr, xlat)
		}
	}

	// Rebuild the message
	msg2 := goipp.NewMessageWithGroups(msg.Version, msg.Code,
		msg.RequestID, groups)

	return msg2
}

// translateAttr translates URLs found in the goipp.Attribute, recursively
// scanning nested collections.
//
// Each found URL is translated using the provided `xlat` function.
//
// Translation is performed "in place".
func (mx *msgXlat) translateAttr(attr *goipp.Attribute,
	xlat func(*url.URL) *url.URL) {

	for i := range attr.Values {
		v := &attr.Values[i]
		mx.translateVal(&v.V, v.T, xlat)
	}
}

// translateVal translates URLs in the goipp.Value, recursively
// scanning nested collections.
//
// Each found URL is translated using the provided `xlat` function.
//
// Translation is performed "in place".
func (mx *msgXlat) translateVal(v *goipp.Value, t goipp.Tag, xlat func(*url.URL) *url.URL) {
	switch v2 := (*v).(type) {
	case goipp.Collection:
		for i := range v2 {
			attr := &v2[i]
			mx.translateAttr(attr, xlat)
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
