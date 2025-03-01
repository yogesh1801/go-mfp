// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package escl

import (
	"strconv"
	"time"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/xmldoc"
)

// JobInfo reports the state of a particular scan job.
//
// eSCL Technical Specification, 9.1.
type JobInfo struct {
	JobURI           string                      // Unique Job URI, that identifies the job
	JobUUID          optional.Val[uuid.UUID]     // Unique, persistent
	Age              optional.Val[time.Duration] // Time since last update
	ImagesCompleted  optional.Val[int]           // Images completed so far
	ImagesToTransfer optional.Val[int]           // Images to transfer
	JobState         JobState                    // Job state
	JobStateReasons  []JobStateReason            // Reason of the job state
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
	state := xmldoc.Lookup{Name: NsPWG + ":JobState", Required: true}
	reasons := xmldoc.Lookup{Name: NsPWG + ":JobStateReasons"}

	missed := root.Lookup(&jobURI, &jobUUID, &age, &compl, &xfer,
		&state, &reasons)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	info.JobURI = jobURI.Elem.Text

	if jobUUID.Found {
		var uu uuid.UUID
		uu, err = uuid.Parse(jobUUID.Elem.Text)
		if err != nil {
			return
		}
		info.JobUUID = optional.New(uu)
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
			Text: (*info.JobUUID).URN(),
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

	elm.Children = append(elm.Children,
		info.JobState.toXML(NsPWG+":JobState"))

	if info.JobStateReasons != nil {
		chld := xmldoc.Element{Name: NsPWG + ":JobStatereasons"}
		for _, reason := range info.JobStateReasons {
			chld2 := reason.toXML(NsPWG + ":JobStateReason")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	return elm
}
