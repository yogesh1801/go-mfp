// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// URLs handling tests

package wsdd

import "testing"

// TestURLParse tests urlParse function
func TestURLParse(t *testing.T) {
	type testData struct {
		s  string // Input string
		ok bool   // Must parse
	}

	tests := []testData{
		{"http://www.example.com", true},
		{"https://www.example.com", true},
		{"http://127.0.0.1", true},
		{"http://127.0.0.1/", true},
		{"example.com", false},
		{"http:///example.com", false},
		{"ftp://example.com", false},
		{"ipp://example.com", false},
		{"%%%", false},
	}

	for _, test := range tests {
		u := urlParse(test.s)
		ok := u != nil
		if ok != test.ok {
			msg := "MUST"
			if !test.ok {
				msg = "MUST NOT"
			}
			t.Errorf("%q: %s parse", test.s, msg)
		}
	}
}

// TestURLIsLiteral tests urlIsLiteral function
func TestURLIsLiteral(t *testing.T) {
	type testData struct {
		s   string // Input string
		lit bool   // Must be literal
	}

	tests := []testData{
		{"http://www.example.com", false},
		{"http://127.0.0.1", true},
		{"http://[::1]", true},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]", true},
	}

	for _, test := range tests {
		lit := urlIsLiteral(urlParse(test.s))
		if lit != test.lit {
			msg := "MUST"
			if !test.lit {
				msg = "MUST NOT"
			}
			t.Errorf("%q: %s be literal", test.s, msg)
		}
	}
}

// TestURLIsIP4 tests urlIsIP4 function
func TestURLIsIP4(t *testing.T) {
	type testData struct {
		s   string // Input string
		lit bool   // Must be literal
	}

	tests := []testData{
		{"http://www.example.com", false},
		{"http://127.0.0.1", true},
		{"http://[::1]", false},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]", false},
	}

	for _, test := range tests {
		lit := urlIsIP4(urlParse(test.s))
		if lit != test.lit {
			msg := "MUST"
			if !test.lit {
				msg = "MUST NOT"
			}
			t.Errorf("%q: %s be IP4", test.s, msg)
		}
	}
}

// TestURLIsIP6 tests urlIsIP6 function
func TestURLIsIP6(t *testing.T) {
	type testData struct {
		s   string // Input string
		lit bool   // Must be literal
	}

	tests := []testData{
		{"http://www.example.com", false},
		{"http://127.0.0.1", false},
		{"http://[::1]", true},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]", true},
	}

	for _, test := range tests {
		lit := urlIsIP6(urlParse(test.s))
		if lit != test.lit {
			msg := "MUST"
			if !test.lit {
				msg = "MUST NOT"
			}
			t.Errorf("%q: %s be IP6", test.s, msg)
		}
	}
}

// TestURLWithZone tests urlWithZone function
func TestURLWithZone(t *testing.T) {
	type testData struct {
		in   string // Input URL string
		zone string // The zone
		out  string // Expected output URL string
	}

	tests := []testData{
		{"http://www.example.com", "eth0",
			"http://www.example.com/"},
		{"http://127.0.0.1", "eth0",
			"http://127.0.0.1/"},
		{"http://127.0.0.1:8080", "eth0",
			"http://127.0.0.1:8080/"},
		{"http://[::1]", "eth0",
			"http://[::1]/"},
		{"http://[::1]:8080", "eth0",
			"http://[::1]:8080/"},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]", "eth0",
			"http://[fe80::1cc0:3e8c:119f:c2e1%25eth0]/"},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]:8080", "eth0",
			"http://[fe80::1cc0:3e8c:119f:c2e1%25eth0]:8080/"},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]", "",
			"http://[fe80::1cc0:3e8c:119f:c2e1]/"},
	}

	for _, test := range tests {
		out := urlWithZone(urlParse(test.in), test.zone).String()
		if out != test.out {
			t.Errorf("%q with zone %q:\n"+
				"expected: %q\n"+
				"present:  %q\n",
				test.in, test.zone, test.out, out)
		}
	}
}

// TestURLWithHostname tests urlWithHostname function
func TestURLWithHostname(t *testing.T) {
	type testData struct {
		in   string // Input URL string
		host string // The hostname
		out  string // Expected output URL string
	}

	tests := []testData{
		{"http://www.example.com", "127.0.0.1",
			"http://127.0.0.1/"},
		{"http://www.example.com", "::1",
			"http://[::1]/"},
		{"http://www.example.com:8080", "127.0.0.1",
			"http://127.0.0.1:8080/"},
		{"http://www.example.com:8080", "::1",
			"http://[::1]:8080/"},
		{"http://www.example.com", "fe80::1cc0:3e8c:119f:c2e1%ens18",
			"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]/"},
		{"http://www.example.comi:8080", "fe80::1cc0:3e8c:119f:c2e1%ens18",
			"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]:8080/"},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25ens18]/", "127.0.0.1",
			"http://127.0.0.1/"},
	}

	for _, test := range tests {
		out := urlWithHostname(urlParse(test.in), test.host).String()
		if out != test.out {
			t.Errorf("%q with hostname %q:\n"+
				"expected: %q\n"+
				"present:  %q\n",
				test.in, test.host, test.out, out)
		}
	}
}

// TestURLZone tests urlWithZone function
func TestURLZone(t *testing.T) {
	type testData struct {
		s    string // Input string
		zone string // Expected zone
	}

	tests := []testData{
		{"http://www.example.com", ""},
		{"http://127.0.0.1", ""},
		{"http://127.0.0.1:8080", ""},
		{"http://[::1%25eth0]", ""},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25eth0]", "eth0"},
		{"http://[fe80::1cc0:3e8c:119f:c2e1%25eth0]:8080", "eth0"},
	}

	for _, test := range tests {
		zone := urlZone(urlParse(test.s))
		if zone != test.zone {
			t.Errorf("%q: zone expected %q, present %q",
				test.s, test.zone, zone)
		}
	}
}
