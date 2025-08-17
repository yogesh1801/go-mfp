// MFP - Miulti-Function Printers and scanners toolkit
// CUPS Client and Server
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS client test

package cups

import (
	"context"
	"testing"
)

func TestCUPS(t *testing.T) {
	c := NewClient(DefaultUNIXURL, nil)
	rsp, err := c.CUPSGetDefault(context.Background(), []string{"all"})

	if err != nil {
		t.Errorf("%s", err)
		return
	}

	_ = rsp
	//fmt.Printf("%#v", rsp)
}
