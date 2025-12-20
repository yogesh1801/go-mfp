// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scan job description

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// JobDescription represents the <wscn:JobDescription> element,
// containing basic creation information for the currently identified job.
type JobDescription struct {
	JobInformation         optional.Val[string]
	JobName                string
	JobOriginatingUserName string
}

// toXML generates XML tree for the [JobDescription].
func (jd JobDescription) toXML(name string) xmldoc.Element {
	children := []xmldoc.Element{}

	if jd.JobInformation != nil {
		children = append(children, xmldoc.Element{
			Name: NsWSCN + ":JobInformation",
			Text: optional.Get(jd.JobInformation),
		})
	}

	children = append(children, xmldoc.Element{
		Name: NsWSCN + ":JobName",
		Text: jd.JobName,
	})

	children = append(children, xmldoc.Element{
		Name: NsWSCN + ":JobOriginatingUserName",
		Text: jd.JobOriginatingUserName,
	})

	return xmldoc.Element{
		Name:     name,
		Children: children,
	}
}

// decodeJobDescription decodes [JobDescription] from the XML tree.
func decodeJobDescription(root xmldoc.Element) (jd JobDescription, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	jobName := xmldoc.Lookup{
		Name:     NsWSCN + ":JobName",
		Required: true,
	}
	jobOriginatingUserName := xmldoc.Lookup{
		Name:     NsWSCN + ":JobOriginatingUserName",
		Required: true,
	}

	jobInformation := xmldoc.Lookup{
		Name:     NsWSCN + ":JobInformation",
		Required: false,
	}

	missed := root.Lookup(
		&jobName,
		&jobOriginatingUserName,
		&jobInformation,
	)
	if missed != nil && missed.Required {
		return jd, xmldoc.XMLErrMissed(missed.Name)
	}

	jd.JobName = jobName.Elem.Text
	jd.JobOriginatingUserName = jobOriginatingUserName.Elem.Text

	if jobInformation.Found {
		jd.JobInformation = optional.New(jobInformation.Elem.Text)
	}

	return jd, nil
}
