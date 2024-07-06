// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package discovery

// KwAuthInfo contains possible values for AuthInfoRequired parameter
type KwAuthInfo string

// KwAuthInfo values:
const (
	KwAuthInfoCertificate KwAuthInfo = "certificate" // TLS certificate
	KwAuthInFonegotiate   KwAuthInfo = "negotiate"   // Kerberos, RFC4559
	KwAuthInfoNone        KwAuthInfo = "none"        // No authentication
	KwAuthInfoOAuth       KwAuthInfo = "oauth"       //  OAuth 2.0 RFC6749
	KwAuthInfoPasswd      KwAuthInfo = "username,passwor"
)

// KwTLSVersion contains known TLS versions
type KwTLSVersion string

// KwTLSVersion values:
const (
	KwTLSVersionNone KwTLSVersion = "none" // TLS not supported
	KwTLSVersion10   KwTLSVersion = "1.0"  // TLS 1.0
	KwTLSVersion11   KwTLSVersion = "1.1"  // TLS 1.1
	KwTLSVersion12   KwTLSVersion = "1.2"  // TLS 1.2
	KwTLSVersion13   KwTLSVersion = "1.3"  // TLS 1.3
)
