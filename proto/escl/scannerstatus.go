// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner status

package escl

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScannerStatus is the scanner response, that represents the current
// scanner status.
//
// eSCL Technical Specification, 9.
//
// GET /{root}/ScannerStatus
type ScannerStatus struct {
	Version  Version                // eSCL protocol version
	State    ScannerState           // Overall scanner state
	ADFState optional.Val[ADFState] // ADF state
	Jobs     []JobInfo              // State of particular jobs
}

// PushJobInfo pushes [JobInfo] into the beginning of the
// [ScannerStatus.Jobs] slice.
//
// If max > 0, resulting Jobs slice size will not exceed the
// maximum, and exceeding elements will be thrown away.
func (status *ScannerStatus) PushJobInfo(info JobInfo, max int) {
	if max > 0 && len(status.Jobs) >= max {
		status.Jobs = status.Jobs[0 : max-1]
	}

	status.Jobs = append(status.Jobs, JobInfo{})
	copy(status.Jobs[1:], status.Jobs)
	status.Jobs[0] = info
}

// DecodeScannerStatus decodes [ScannerStatus] from the XML tree.
func DecodeScannerStatus(root xmldoc.Element) (
	ret *ScannerStatus, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	var status ScannerStatus

	// Lookup relevant XML elements
	ver := xmldoc.Lookup{Name: NsPWG + ":Version", Required: true}
	state := xmldoc.Lookup{Name: NsPWG + ":State", Required: true}
	adfState := xmldoc.Lookup{Name: NsScan + ":AdfState"}
	jobs := xmldoc.Lookup{Name: NsScan + ":Jobs"}

	missed := root.Lookup(&ver, &state, &adfState, &jobs)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	status.Version, err = decodeVersion(ver.Elem)
	if err != nil {
		return
	}

	status.State, err = decodeScannerState(state.Elem)
	if err != nil {
		return
	}

	if adfState.Found {
		var s ADFState
		s, err = decodeADFState(adfState.Elem)
		if err != nil {
			return
		}
		status.ADFState = optional.New(s)
	}

	if jobs.Found {
		for _, elem := range jobs.Elem.Children {
			if elem.Name == NsScan+":JobInfo" {
				var info JobInfo
				info, err = decodeJobInfo(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(jobs.Elem, err)
					return
				}

				status.Jobs = append(status.Jobs, info)
			}
		}
	}

	ret = &status
	return
}

// ToXML generates XML tree for the [ScannerStatus].
func (status *ScannerStatus) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsScan + ":ScannerStatus",
		Children: []xmldoc.Element{
			status.Version.toXML(NsPWG + ":Version"),
			status.State.toXML(NsPWG + ":State"),
		},
	}

	if status.ADFState != nil {
		elm.Children = append(elm.Children,
			(*status.ADFState).toXML(NsScan+":AdfState"))
	}

	if status.Jobs != nil {
		chld := xmldoc.Element{Name: NsScan + ":Jobs"}
		for _, job := range status.Jobs {
			chld2 := job.toXML(NsScan + ":JobInfo")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	return elm
}
