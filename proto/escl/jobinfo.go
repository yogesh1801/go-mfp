// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// JobInfo -- the particular job state.

package escl

import (
	"strconv"
	"time"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// JobInfo reports the state of a particular scan job.
//
// eSCL Technical Specification, 9.1.
//
// Note, JobUUID, despite its name, often used by firmwares as an arbitrary
// string, uniquely identifying the job, but lacking the correct UUID
// syntax. So JobUUID is defined as string, not uuid.UUID.
type JobInfo struct {
	JobURI             string                      // Base URL for the job
	JobUUID            optional.Val[string]        // Unique, persistent
	Age                optional.Val[time.Duration] // Time since last update
	ImagesCompleted    optional.Val[int]           // Images now completed
	ImagesToTransfer   optional.Val[int]           // Images yet to transfer
	TransferRetryCount optional.Val[int]           // Load retries for now
	JobState           JobState                    // Job state
	JobStateReasons    []JobStateReason            // Reason of the state
}

// decodeJobInfo decodes [JobInfo] from the XML tree.
func decodeJobInfo(root xmldoc.Element) (info JobInfo, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	jobURI := xmldoc.Lookup{Name: NsPWG + ":JobUri", Required: true}
	jobUUID := xmldoc.Lookup{Name: NsPWG + ":JobUuid"}
	age := xmldoc.Lookup{Name: NsScan + ":Age"}
	compl := xmldoc.Lookup{Name: NsPWG + ":ImagesCompleted"}
	xfer := xmldoc.Lookup{Name: NsPWG + ":ImagesToTransfer"}
	retry := xmldoc.Lookup{Name: NsScan + ":TransferRetryCount"}
	state := xmldoc.Lookup{Name: NsPWG + ":JobState", Required: true}
	reasons := xmldoc.Lookup{Name: NsPWG + ":JobStateReasons"}

	missed := root.Lookup(&jobURI, &jobUUID, &age, &compl, &xfer,
		&retry, &state, &reasons)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	info.JobURI = jobURI.Elem.Text

	if jobUUID.Found {
		info.JobUUID = optional.New(jobUUID.Elem.Text)
	}

	if age.Found {
		var v int
		v, err = decodeNonNegativeInt(age.Elem)
		if err != nil {
			return
		}
		info.Age = optional.New(time.Duration(v) * time.Second)
	}

	if compl.Found {
		var v int
		v, err = decodeNonNegativeInt(compl.Elem)
		if err != nil {
			return
		}
		info.ImagesCompleted = optional.New(v)
	}

	if xfer.Found {
		var v int
		v, err = decodeNonNegativeInt(xfer.Elem)
		if err != nil {
			return
		}
		info.ImagesToTransfer = optional.New(v)
	}

	if retry.Found {
		var v int
		v, err = decodeNonNegativeInt(retry.Elem)
		if err != nil {
			return
		}
		info.TransferRetryCount = optional.New(v)
	}

	info.JobState, err = decodeJobState(state.Elem)
	if err != nil {
		return
	}

	if reasons.Found {
		for _, elem := range reasons.Elem.Children {
			if elem.Name == NsPWG+":JobStateReason" {
				var reason JobStateReason
				reason, err = decodeJobStateReason(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						reasons.Elem, err)
					return
				}

				info.JobStateReasons = append(
					info.JobStateReasons, reason)
			}
		}
	}

	return
}

// toXML generates XML tree for the [JobInfo].
func (info JobInfo) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{Name: NsPWG + ":JobUri", Text: info.JobURI},
		},
	}

	if info.JobUUID != nil {
		chld := xmldoc.Element{
			Name: NsPWG + ":JobUuid",
			Text: *info.JobUUID,
		}
		elm.Children = append(elm.Children, chld)
	}

	if info.Age != nil {
		age := *info.Age
		age += time.Second / 2
		age /= time.Second

		chld := xmldoc.Element{
			Name: NsScan + ":Age",
			Text: strconv.FormatUint(uint64(age), 10),
		}
		elm.Children = append(elm.Children, chld)
	}

	if info.ImagesCompleted != nil {
		chld := xmldoc.Element{
			Name: NsPWG + ":ImagesCompleted",
			Text: strconv.Itoa(*info.ImagesCompleted),
		}
		elm.Children = append(elm.Children, chld)
	}

	if info.ImagesToTransfer != nil {
		chld := xmldoc.Element{
			Name: NsPWG + ":ImagesToTransfer",
			Text: strconv.Itoa(*info.ImagesToTransfer),
		}
		elm.Children = append(elm.Children, chld)
	}

	if info.TransferRetryCount != nil {
		chld := xmldoc.Element{
			Name: NsScan + ":TransferRetryCount",
			Text: strconv.Itoa(*info.TransferRetryCount),
		}
		elm.Children = append(elm.Children, chld)
	}

	elm.Children = append(elm.Children,
		info.JobState.toXML(NsPWG+":JobState"))

	if info.JobStateReasons != nil {
		chld := xmldoc.Element{Name: NsPWG + ":JobStateReasons"}
		for _, reason := range info.JobStateReasons {
			chld2 := reason.toXML(NsPWG + ":JobStateReason")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	return elm
}
