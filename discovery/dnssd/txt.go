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
	"strconv"
	"strings"

	"github.com/alexpevzner/mfp/discovery"
	"github.com/alexpevzner/mfp/uuid"
)

// txtPrinter represents a decoded TXT record for printer
type txtPrinter struct {
	uuid    uuid.UUID                    // Device UUID
	svcType string                       // Service type
	meta    discovery.Metadata           // Unit metadata
	params  *discovery.PrinterParameters // Printer parameters
}

// txtScanner represents a decoded TXT record for scanner
type txtScanner struct {
	uuid    uuid.UUID                    // Device UUID
	svcType string                       // Service type
	uriPath string                       // Path part of URI
	meta    discovery.Metadata           // Unit metadata
	params  *discovery.ScannerParameters // Printer parameters
}

// txtPrinter decodes record for printer
func decodeTxtPrinter(svcType, svcInstance string,
	txt []string) (txtPrinter, error) {

	// Set defaults
	p := txtPrinter{
		// The default UUID, in a very unlikely case UUID is missed
		// in the TXT record
		uuid: uuid.MD5(uuid.NilUUID, svcInstance),

		// Save service type
		svcType: svcType,

		// Default parameters
		params: &discovery.PrinterParameters{
			// 50 is the reasonable default
			Priority: 50,
		},
	}

	// Some defaults depend on service type
	switch svcType {
	case svcTypeIPP, svcTypeIPPS:
		p.params.Queue = "ipp/print"
	case svcTypeLPD:
		p.params.Queue = "auto"
	}

	// Parse key by key
	found := map[string]struct{}{}

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

		// Decode the value
		var err error
		switch txToLower(key) {
		case "adminurl":
			p.params.AdminURL = value
		case "air":
			p.params.Auth, err = txtAuth(value)
		case "bind":
			p.params.Bind, err = txtBool(value)
		case "color":
			p.params.Color, err = txtBool(value)
		case "copies":
			p.params.Copies, err = txtBool(value)
		case "duplex":
			p.params.Duplex, err = txtBool(value)
		case "fax":
		case "kind":
			p.params.Media, err = txtMediaKind(value)
		case "note":
			p.params.Location = value
		case "papermax":
			p.params.Paper, err = txtPaperMax(value)
		case "pdl":
			p.params.PDL, err = txtKeywords(value)
		case "priority":
			p.params.Priority, err = strconv.Atoi(value)
			if err != nil {
				break
			}
			if p.params.Priority < 0 || p.params.Priority >= 100 {
				err = errors.New("out of range (0...99)")
			}
		case "product":
			p.params.PSProduct = value
		case "punch":
			// Punch can take values "0", "1", "2", "3",
			// "4" and "U", according to the number of holes
			// the puncher can make. "0" means "no punching"
			if value != "0" && txToLower(value) != "u" {
				p.params.Punch = true
			}
		case "rp":
			p.params.Queue = value
		case "sort":
			p.params.Sort, err = txtBool(value)
		case "staple":
			p.params.Staple, err = txtBool(value)
		case "txtvers":
			if value != "1" {
				err = fmt.Errorf("unknown version %q", value)
			}
		case "ty":
			p.meta.MakeModel = value
		case "usb_mdl":
			p.meta.Model = value
		case "usb_mfg":
			p.meta.Manufacturer = value
		case "uuid":
			p.uuid, err = uuid.Parse(value)

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

		// Check for error
		if err != nil {
			err = fmt.Errorf("%s: %w", key, err)
			return txtPrinter{}, err
		}
	}

	return p, nil
}

// txtPrinter decodes record for printer
func decodeTxtScanner(svcType, svcInstance string,
	txt []string) (txtScanner, error) {
	s := txtScanner{
		// The default UUID, in a very unlikely case UUID is missed
		// in the TXT record
		uuid: uuid.MD5(uuid.NilUUID, svcInstance),

		// Save service type
		svcType: svcType,

		// The default path to the eSCL endpoint
		uriPath: "eSCL",

		// Default parameters
		params: &discovery.ScannerParameters{},
	}

	found := map[string]struct{}{}

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

		// Decode the value
		var err error
		switch txToLower(key) {
		case "adminurl":
			s.params.AdminURL = value
		case "cs":
			s.params.Colors, err = txtColors(value)
		case "duplex":
			s.params.Duplex, err = txtBool(value)
		case "is":
			s.params.Sources, err = txtSources(value)
		case "note":
			s.params.Location = value
		case "pdl":
			s.params.PDL, err = txtKeywords(value)
		case "representation":
			s.params.IconURL = value
		case "rs":
			// Strip leading and trailing '/'.
			// sane-airscan does the same.
			rs := value
			for strings.HasPrefix(rs, "/") {
				rs = rs[1:]
			}
			for strings.HasSuffix(rs, "/") {
				rs = rs[:len(rs)-1]
			}

			s.uriPath = value
		case "ty":
			s.meta.MakeModel = value
		case "uuid":
			s.uuid, err = uuid.Parse(value)
		}

		// Check for error
		if err != nil {
			err = fmt.Errorf("%s: %w", key, err)
			return txtScanner{}, err
		}
	}

	return s, nil
}

// txtAuth decodes an authentication mode
func txtAuth(value string) (discovery.AuthMode, error) {
	switch txToLower(value) {
	case "none":
		return discovery.AuthNone, nil
	case "certificate":
		return discovery.AuthCertificate, nil
	case "negotiate":
		return discovery.AuthKerberos, nil
	case "oauth":
		return discovery.AuthOAuth2, nil
	case "username,passwor":
		return discovery.AuthPasswd, nil
	}

	return discovery.AuthOther, nil
}

// txtBool decodes a boolean value
func txtBool(value string) (bool, error) {
	return txToLower(value) == "t", nil
}

// txtColors decodes discovery.ColorMode bits
func txtColors(value string) (discovery.ColorMode, error) {
	keywords, _ := txtKeywords(value)

	var colors discovery.ColorMode
	for _, kw := range keywords {
		switch txToLower(kw) {
		case "color":
			colors |= discovery.ColorRGB
		case "grayscale":
			colors |= discovery.ColorGrayscale
		case "binary":
			colors |= discovery.ColorBW
		default:
			colors |= discovery.ColorOther
		}
	}

	return colors, nil
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

// txtSources decodes discovery.ScanSource bits
func txtSources(value string) (discovery.ScanSource, error) {
	keywords, _ := txtKeywords(value)

	var sources discovery.ScanSource
	for _, kw := range keywords {
		switch txToLower(kw) {
		case "adf":
			sources |= discovery.ScanADF
		case "platen":
			sources |= discovery.ScanPlaten
		default:
			sources |= discovery.ScanOther
		}
	}

	return sources, nil
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
