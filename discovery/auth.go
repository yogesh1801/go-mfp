// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Authentication modes

package discovery

import "strings"

// AuthMode defines the type of authentication information, required
// by printer or scanner
type AuthMode int

// AuthMode values:
const (
	AuthNone        AuthMode = 1 << iota // No authentication
	AuthCertificate                      // TLS certificate
	AuthKerberos                         // Kerberos (RFC4559)
	AuthOAuth2                           // OAuth 2.0 (RFC6749)
	AuthPasswd                           // User name+password
	AuthOther                            // Other (unknown) mode
)

// String formats AuthMode as string, for printing and logging
func (auth AuthMode) String() string {
	s := []string{}

	if auth&AuthNone != 0 {
		s = append(s, "none")
	}
	if auth&AuthCertificate != 0 {
		s = append(s, "certificate")
	}
	if auth&AuthKerberos != 0 {
		s = append(s, "Kerberos")
	}
	if auth&AuthOAuth2 != 0 {
		s = append(s, "OAuth2")
	}
	if auth&AuthPasswd != 0 {
		s = append(s, "login+password")
	}
	if auth&AuthOther != 0 {
		s = append(s, "other")
	}

	return strings.Join(s, ",")
}
