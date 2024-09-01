// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// DNS-SD TXT records

package dnssd

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/alexpevzner/mfp/discovery"
	"github.com/alexpevzner/mfp/uuid"
)

// txtDecodePrinterParameters decodes PrinterParameters out of
// the TXT record
func txtDecodePrinterParameters(txt []string) (discovery.UnitID,
	discovery.PrinterParameters, error) {

	id := discovery.UnitID{}
	params := discovery.PrinterParameters{
		Priority: 50,
	}

	found := map[string]struct{}{}
	missed := map[string]struct{}{
		"txtvers": struct{}{},
		"uuid":    struct{}{},
	}

	for _, t := range txt {
		// Split record into key and value.
		// Ignore keys without values -- CUPS does the same.
		key, value, novalue := txtParse(t)

		if novalue {
			continue
		}

		// Update found/missed sets and check for duplicates
		if _, dup := found[key]; dup {
			continue
		}

		delete(missed, key)

		// Decode the value
		var err error
		switch txToLower(key) {
		case "adminurl":
			params.AdminURL = value
		case "air":
			params.Auth, err = txtAuth(value)
		case "bind":
			params.Bind, err = txtBool(value)
		case "color":
			params.Color, err = txtBool(value)
		case "copies":
			params.Copies, err = txtBool(value)
		case "duplex":
			params.Duplex, err = txtBool(value)
		case "fax":
		case "kind":
			params.Media, err = txtMediaKind(value)
		case "note":
			params.Location = value
		case "papermax":
			params.Paper, err = txtPaperMax(value)
		case "pdl":
			params.PDL, err = txtKeywords(value)
		case "priority":
			params.Priority, err = strconv.Atoi(value)
			if err != nil {
				break
			}
			if params.Priority < 0 || params.Priority >= 100 {
				err = errors.New("out of range (0...99)")
			}
		case "product":
			params.PPD, err = txtPPD(value)
		case "punch":
			// Punch can take values "0", "1", "2", "3",
			// "4" and "U", according to the number of holes
			// the puncher can make. "0" means "no punching"
			if value != "0" && txToLower(value) != "u" {
				params.Punch = true
			}
		case "rp":
			params.Queue = value
		case "sort":
			params.Sort, err = txtBool(value)
		case "staple":
			params.Staple, err = txtBool(value)
		case "txtvers":
			if value != "1" {
				err = fmt.Errorf("unknown version %q", value)
			}
		case "ty":
			id.MakeModel = value
		case "usb_mdl":
		case "usb_mfg":
		case "uuid":
			id.UUID, err = uuid.Parse(value)

		// These parameters are ignored so far, but it may change
		// in the future.
		//
		// TODO: review these parameters.
		case "print_wfds":
		case "qtotal":
		case "scan":
		case "tls":
		case "urf":
		}

		if err != nil {
			err = fmt.Errorf("%s: %w", key, err)
			return discovery.UnitID{},
				discovery.PrinterParameters{}, err
		}
	}

	// Check for missed required keys
	if len(missed) != 0 {
		keys := make([]string, 0, len(missed))
		for key := range missed {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		err := fmt.Errorf("missed: %s", strings.Join(keys, ","))
		return discovery.UnitID{}, discovery.PrinterParameters{}, err
	}

	return id, params, nil
}

// txtAuth decodes an authentication mode
func txtAuth(value string) (discovery.AuthMode, error) {
	switch txToLower(value) {
	case "none":
		return discovery.AuthNone, nil
	case "certificate":
		return discovery.AuthCertificate, nil
	case "negotiate":
		return discovery.AuthNegotiate, nil
	case "oauth":
		return discovery.AuthOAuth, nil
	case "username,passwor":
		return discovery.AuthPasswd, nil
	}

	return discovery.AuthOther, nil
}

// txtBool decodes a boolean value
func txtBool(value string) (bool, error) {
	return txToLower(value) == "t", nil
}

// txtKeywords decodes comma-separated list of keywords
func txtKeywords(value string) ([]string, error) {
	// Split value into comma-separated list
	pdl := strings.Split(value, ",")

	// Drop empty items (just in case)
	o := 0
	for i := range pdl {
		if pdl[i] != "" {
			pdl[o] = pdl[i]
			o++
		}
	}

	pdl = pdl[:o]
	return pdl, nil
}

// txtMediaKind decodes discovery.MediaKind bits
func txtMediaKind(value string) (discovery.MediaKind, error) {
	keywords, _ := txtKeywords(value)

	var media discovery.MediaKind
	for _, kw := range keywords {
		switch txToLower(kw) {
		case "disc":
			media |= discovery.MediaDisk
		case "document":
			media |= discovery.MediaDocument
		case "envelope":
			media |= discovery.MediaEnvelope
		case "label":
			media |= discovery.MediaLabel
		case "large-format":
			media |= discovery.MediaLargeFormat
		case "photo":
			media |= discovery.MediaPhoto
		case "postcard":
			media |= discovery.MediaPostcard
		case "receipt":
			media |= discovery.MediaReceipt
		case "roll":
			media |= discovery.MediaRoll
		default:
			media |= discovery.MediaOther
		}
	}

	return media, nil
}

// txtPPD decodes the max paper size
func txtPaperMax(value string) (discovery.PaperSize, error) {
	switch txToLower(value) {
	case "<legal-A4":
		return discovery.PaperA4Minus, nil
	case "legal-A4":
		return discovery.PaperA4, nil
	case "tabloid-A3":
		return discovery.PaperA3, nil
	case "isoC-A2":
		return discovery.PaperA2, nil
	case ">isoC-A2":
		return discovery.PaperA2Plus, nil
	}

	return discovery.PaperUnknown, nil
}

// txtPPD decodes a PPD file name
func txtPPD(value string) (string, error) {
	ppd, _ := strings.CutPrefix(value, "(")
	ppd, _ = strings.CutSuffix(value, ")")
	return ppd, nil
}

// txtParse parses "key=value" string into key and value parts.
func txtParse(s string) (key, value string, novalue bool) {
	i := strings.IndexByte(s, '=')
	if i >= 0 {
		return s[0:i], s[i+1:], false
	}

	return s, "", true
}

// txToLower converts ASCII string to lowercase.
func txToLower(s string) string {
	buf := []byte(s)

	for i := range buf {
		if c := buf[i]; 'A' <= c && c <= 'Z' {
			buf[i] = c - 'A' + 'a'
		}
	}

	return string(buf)
}
