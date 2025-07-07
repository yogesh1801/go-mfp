// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Functions for paths

package transport

import (
	"path"
	"strings"
)

// CleanURLPath returns the canonical form for p:
//   - if p is not started with '/', '/' is prepended
//   - . and .. elements are removed
//   - the trailing '/' is preserved
func CleanURLPath(p string) string {
	if p == "" {
		return "/"
	}

	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}

	trailing := strings.HasSuffix(p, "/")
	p = path.Clean(p)

	if trailing && !strings.HasSuffix(p, "/") {
		p += "/"
	}

	return p
}
