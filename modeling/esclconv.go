// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL type conversions

package modeling

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/cpython"
	"github.com/OpenPrinting/go-mfp/proto/escl"
)

// esclDecodeADFOption decodes escl.ADFOption from the Python object
func esclDecodeADFOption(obj *cpython.Object) (escl.ADFOption, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	opt := escl.DecodeADFOption(s)
	if opt == escl.UnknownADFOption {
		return 0, fmt.Errorf("%s: invalid eSCL AdfOption", s)
	}

	return opt, nil
}

// esclDecodeADFState decodes escl.ADFState from the Python object
func esclDecodeADFState(obj *cpython.Object) (escl.ADFState, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	st := escl.DecodeADFState(s)
	if st == escl.UnknownADFState {
		return 0, fmt.Errorf("%s: invalid eSCL AdfState", s)
	}

	return st, nil
}

// esclDecodeBinaryRendering decodes escl.BinaryRendering from the
// Python object
func esclDecodeBinaryRendering(obj *cpython.Object) (
	escl.BinaryRendering, error) {

	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	rnd := escl.DecodeBinaryRendering(s)
	if rnd == escl.UnknownBinaryRendering {
		return 0, fmt.Errorf("%s: invalid eSCL BinaryRendering(", s)
	}

	return rnd, nil
}

// esclDecodeCCDChannel decodes escl.CCDChannel from the Python object
func esclDecodeCCDChannel(obj *cpython.Object) (escl.CCDChannel, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	ccd := escl.DecodeCCDChannel(s)
	if ccd == escl.UnknownCCDChannel {
		return 0, fmt.Errorf("%s: invalid eSCL CcdChannel", s)
	}

	return ccd, nil
}

// esclDecodeColorMode decodes escl.ColorMode from the Python object
func esclDecodeColorMode(obj *cpython.Object) (escl.ColorMode, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	cm := escl.DecodeColorMode(s)
	if cm == escl.UnknownColorMode {
		return 0, fmt.Errorf("%s: invalid eSCL ColorMode(", s)
	}

	return cm, nil
}

// esclDecodeColorSpace decodes escl.ColorSpace from the Python object
func esclDecodeColorSpace(obj *cpython.Object) (escl.ColorSpace, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	sps := escl.DecodeColorSpace(s)
	if sps == escl.UnknownColorSpace {
		return 0, fmt.Errorf("%s: invalid eSCL ColorSpace", s)
	}

	return sps, nil
}

// esclDecodeContentType decodes escl.ContentType from the Python object
func esclDecodeContentType(obj *cpython.Object) (escl.ContentType, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	ct := escl.DecodeContentType(s)
	if ct == escl.UnknownContentType {
		return 0, fmt.Errorf("%s: invalid eSCL ContentType", s)
	}

	return ct, nil
}

// esclDecodeFeedDirection decodes escl.FeedDirection from the Python object
func esclDecodeFeedDirection(obj *cpython.Object) (escl.FeedDirection, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	feed := escl.DecodeFeedDirection(s)
	if feed == escl.UnknownFeedDirection {
		return 0, fmt.Errorf("%s: invalid eSCL FeedDirection", s)
	}

	return feed, nil
}

// esclDecodeImagePosition decodes escl.ImagePosition from the Python object
func esclDecodeImagePosition(obj *cpython.Object) (escl.ImagePosition, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	pos := escl.DecodeImagePosition(s)
	if pos == escl.UnknownImagePosition {
		return 0, fmt.Errorf("%s: invalid eSCL ImagePosition", s)
	}

	return pos, nil
}

// esclDecodeInputSource decodes escl.InputSource from the Python object
func esclDecodeInputSource(obj *cpython.Object) (escl.InputSource, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	src := escl.DecodeInputSource(s)
	if src == escl.UnknownInputSource {
		return 0, fmt.Errorf("%s: invalid eSCL InputSource", s)
	}

	return src, nil
}

// esclDecodeIntent decodes escl.Intent from the Python object
func esclDecodeIntent(obj *cpython.Object) (escl.Intent, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	intent := escl.DecodeIntent(s)
	if intent == escl.UnknownIntent {
		return 0, fmt.Errorf("%s: invalid eSCL Intent", s)
	}

	return intent, nil
}

// esclDecodeJobState decodes escl.JobState from the Python object
func esclDecodeJobState(obj *cpython.Object) (escl.JobState, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	st := escl.DecodeJobState(s)
	if st == escl.UnknownJobState {
		return 0, fmt.Errorf("%s: invalid eSCL JobState", s)
	}

	return st, nil
}

// esclDecodeJobStateReason decodes escl.JobStateReason from the Python object
func esclDecodeJobStateReason(obj *cpython.Object) (escl.JobStateReason, error) {
	s, err := obj.Str()
	if err != nil {
		return "", err
	}

	return escl.JobStateReason(s), nil
}

// esclDecodeVersion decodes escl.Version from the Python object
func esclDecodeVersion(obj *cpython.Object) (escl.Version, error) {
	s, err := obj.Str()
	if err != nil {
		return 0, err
	}

	ver, err := escl.DecodeVersion(s)
	if err != nil {
		return 0, err
	}

	return ver, nil
}

// esclDecodeUnits decodes escl.Units from the Python object
func esclDecodeUnits(obj *cpython.Object) (escl.Units, error) {
	un, err := obj.Str()
	if err != nil {
		return 0, err
	}

	st := escl.DecodeUnits(un)
	if st == escl.UnknownUnits {
		return 0, fmt.Errorf("%s: invalid eSCL Units", un)
	}

	return st, nil
}
