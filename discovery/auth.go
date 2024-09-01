// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Authentication modes

package discovery

// AuthMode defines the type of authentication information, required
// by printer or scanner
type AuthMode int

// AuthMode values:
const (
	AuthNone        AuthMode = 1 << iota // No authentication
	AuthCertificate                      // TLS certificate
	AuthNegotiate                        // Kerberos (RFC4559)
	AuthOAuth                            // OAuth 2.0 (RFC6749)
	AuthPasswd                           // User name+password
	AuthOther                            // Other (unknown) mode
)
