// MFP - Miulti-Function Printers and scanners toolkit
// IPv6 zone suffixes handling
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Functions for IPv6 zone suffixes, the UNIX version.

package zone

import (
	"strconv"
	"unsafe"
)

// #include <stdlib.h>
// #include <net/if.h>
import "C"

// Name returns IPv6 zone name (which is the same as the
// network interface name) by interface index.
func Name(ifindex int) string {
	if ifindex == 0 {
		return ""
	}

	var buf [C.IF_NAMESIZE]C.char

	// Try if_indextoname
	s := C.if_indextoname(C.uint(ifindex), &buf[0])
	if s != nil {
		return C.GoString(s)
	}

	// Fallback to numerical name. Go stdlib does the same.
	return strconv.Itoa(int(ifindex))
}

// Index returns IPv6 zone index by zone name.
func Index(name string) int {
	if name == "" {
		return 0
	}

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	i := C.if_nametoindex(cname)
	if i > 0 {
		return int(i)
	}

	// Last resort. This is ugly, but Go stdlib does the same :(
	n, _ := strconv.Atoi(name)
	return n
}
