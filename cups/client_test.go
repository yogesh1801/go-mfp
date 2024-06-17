// MFP - Miulti-Function Printers and scanners toolkit
// CUPS Client and Server
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS client test

package cups

import (
	"testing"

	"github.com/alexpevzner/mfp/transport"
)

func TestCUPS(t *testing.T) {
	c := NewClient(transport.DefaultCupsUNIX, nil)
	rsp, err := c.CUPSGetDefault([]string{"all"})

	if err != nil {
		t.Errorf("%s", err)
		return
	}

	_ = rsp
	//fmt.Printf("%#v", rsp)
}
