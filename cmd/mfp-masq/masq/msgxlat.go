// MFP - Miulti-Function Printers and scanners toolkit
// The "masq" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// goipp.Message translation for forwarding via proxy

package masq

import (
	"fmt"
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
func (mx *msgXlat) Forward(
	msg *goipp.Message) (*goipp.Message, changeSetMessage) {

	return mx.translateMsg(msg, mx.urlxlat.Forward)
}

// Forward translates message in the reverse (server->client)
// direction.
func (mx *msgXlat) Reverse(
	msg *goipp.Message) (*goipp.Message, changeSetMessage) {

	return mx.translateMsg(msg, mx.urlxlat.Reverse)
}

// translateMsg performs the actual goipp.Message translation.
//
// It returns the translated goipp.Message and a set of applied
// changes.
//
// Each found URL is translated using the provided `xlat` function.
func (mx *msgXlat) translateMsg(msg *goipp.Message,
	xlat func(*url.URL) *url.URL) (*goipp.Message, changeSetMessage) {

	chgmsg := changeSetMessage{}

	// Obtain a deep copy of all message attributes, packed
	// into groups. Roll over all attributes, translating
	// values in place.
	groups := msg.AttrGroups().DeepCopy()
	for i := range groups {
		group := &groups[i]
		chggrp := changeSetGroup{Tag: group.Tag}

		for j := range group.Attrs {
			attr := &group.Attrs[j]
			chg := mx.translateAttr(attr, xlat)
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
// Each found URL is translated using the provided `xlat` function.
//
// Translation is performed "in place".
func (mx *msgXlat) translateAttr(attr *goipp.Attribute,
	xlat func(*url.URL) *url.URL) []changeSetValue {

	chg := []changeSetValue{}

	for i := range attr.Values {
		v := &attr.Values[i]
		morechg := mx.translateVal(&v.V, v.T, xlat)

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
// Each found URL is translated using the provided `xlat` function.
//
// Translation is performed "in place".
func (mx *msgXlat) translateVal(v *goipp.Value, t goipp.Tag,
	xlat func(*url.URL) *url.URL) []changeSetValue {

	switch oldval := (*v).(type) {
	case goipp.Collection:
		chg := []changeSetValue{}

		for i := range oldval {
			attr := &oldval[i]
			morechg := mx.translateAttr(attr, xlat)
			chg = append(chg, morechg...)
		}

		return chg

	case goipp.String:
		if t != goipp.TagURI {
			return nil
		}

		u, err := transport.ParseURL(string(oldval))
		if err == nil {
			u2 := xlat(u)
			newval := goipp.String(u2.String())

			if oldval != newval {
				*v = goipp.String(u2.String())

				chg := []changeSetValue{
					{Old: oldval, New: newval},
				}

				return chg
			}
		}
	}

	return nil
}
