// MFP - Miulti-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scan job status

package wsscan

import (
	"fmt"
	"strconv"
	"time"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// JobStatus represents the <wscn:JobStatus> element,
// containing all information about the status of the current scan job.
type JobStatus struct {
	JobCompletedTime optional.Val[time.Time]
	JobCreatedTime   optional.Val[time.Time]
	JobID            int
	JobState         JobState
	JobStateReasons  []JobStateReason
	ScansCompleted   int
}

// toXML generates XML tree for the [JobStatus].
func (js JobStatus) toXML(name string) xmldoc.Element {
	children := []xmldoc.Element{}

	if js.JobCompletedTime != nil {
		completedTime := optional.Get(js.JobCompletedTime)
		children = append(children, xmldoc.Element{
			Name: NsWSCN + ":JobCompletedTime",
			Text: completedTime.Format(time.RFC3339),
		})
	}

	if js.JobCreatedTime != nil {
		createdTime := optional.Get(js.JobCreatedTime)
		children = append(children, xmldoc.Element{
			Name: NsWSCN + ":JobCreatedTime",
			Text: createdTime.Format(time.RFC3339),
		})
	}

	children = append(children, xmldoc.Element{
		Name: NsWSCN + ":JobId",
		Text: strconv.Itoa(js.JobID),
	})

	children = append(children, js.JobState.toXML(NsWSCN+":JobState"))

	if len(js.JobStateReasons) > 0 {
		jsrChildren := make([]xmldoc.Element, len(js.JobStateReasons))
		for i, v := range js.JobStateReasons {
			jsrChildren[i] = v.toXML(NsWSCN + ":JobStateReason")
		}
		children = append(children, xmldoc.Element{
			Name:     NsWSCN + ":JobStateReasons",
			Children: jsrChildren,
		})
	}

	children = append(children, xmldoc.Element{
		Name: NsWSCN + ":ScansCompleted",
		Text: strconv.Itoa(js.ScansCompleted),
	})

	return xmldoc.Element{
		Name:     name,
		Children: children,
	}
}

// decodeJobStatus decodes [JobStatus] from the XML tree.
func decodeJobStatus(root xmldoc.Element) (js JobStatus, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	jobID := xmldoc.Lookup{
		Name:     NsWSCN + ":JobId",
		Required: true,
	}
	jobState := xmldoc.Lookup{
		Name:     NsWSCN + ":JobState",
		Required: true,
	}
	scansCompleted := xmldoc.Lookup{
		Name:     NsWSCN + ":ScansCompleted",
		Required: true,
	}

	jobCompletedTime := xmldoc.Lookup{
		Name:     NsWSCN + ":JobCompletedTime",
		Required: false,
	}
	jobCreatedTime := xmldoc.Lookup{
		Name:     NsWSCN + ":JobCreatedTime",
		Required: false,
	}
	jobStateReasons := xmldoc.Lookup{
		Name:     NsWSCN + ":JobStateReasons",
		Required: false,
	}

	missed := root.Lookup(
		&jobID,
		&jobState,
		&scansCompleted,
		&jobCompletedTime,
		&jobCreatedTime,
		&jobStateReasons,
	)
	if missed != nil && missed.Required {
		return js, xmldoc.XMLErrMissed(missed.Name)
	}

	if js.JobID, err = decodeNonNegativeInt(jobID.Elem); err != nil {
		return js, fmt.Errorf("jobId: %w", err)
	}
	if js.JobState, err = decodeJobState(jobState.Elem); err != nil {
		return js, fmt.Errorf("jobState: %w", err)
	}
	if js.ScansCompleted, err = decodeNonNegativeInt(scansCompleted.Elem); err != nil {
		return js, fmt.Errorf("scansCompleted: %w", err)
	}

	if jobCompletedTime.Found {
		var completedTime time.Time
		if completedTime, err = decodeTime(jobCompletedTime.Elem); err != nil {
			return js, fmt.Errorf("jobCompletedTime: %w", err)
		}
		js.JobCompletedTime = optional.New(completedTime)
	}

	if jobCreatedTime.Found {
		var createdTime time.Time
		if createdTime, err = decodeTime(jobCreatedTime.Elem); err != nil {
			return js, fmt.Errorf("jobCreatedTime: %w", err)
		}
		js.JobCreatedTime = optional.New(createdTime)
	}

	if jobStateReasons.Elem.Children != nil {
		for _, child := range jobStateReasons.Elem.Children {
			if child.Name == NsWSCN+":JobStateReason" {
				val, err := decodeJobStateReason(child)
				if err != nil {
					return js, fmt.Errorf("JobStateReasons: "+
						"invalid JobStateReason: %w", err)
				}
				js.JobStateReasons = append(js.JobStateReasons, val)
			}
		}
	}

	return js, nil
}
