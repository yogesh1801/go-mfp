// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// CreateScanJobResponse: the Scan Service's response to a scan job request

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// CreateScanJobResponse contains the WSD Scan Service's response to a
// client's scan request. All four child elements are required.
type CreateScanJobResponse struct {
	DocumentFinalParameters DocumentParameters
	ImageInformation        ImageInformation
	JobId                   int
	JobToken                string
}

// toXML generates XML tree for the [CreateScanJobResponse].
func (r CreateScanJobResponse) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			r.DocumentFinalParameters.toXML(NsWSCN +
				":DocumentFinalParameters"),
			r.ImageInformation.toXML(NsWSCN + ":ImageInformation"),
			{Name: NsWSCN + ":JobId", Text: strconv.Itoa(r.JobId)},
			{Name: NsWSCN + ":JobToken", Text: r.JobToken},
		},
	}
}

// decodeCreateScanJobResponse decodes [CreateScanJobResponse] from the XML
// tree.
func decodeCreateScanJobResponse(root xmldoc.Element) (
	CreateScanJobResponse, error,
) {
	var r CreateScanJobResponse

	documentFinalParameters := xmldoc.Lookup{
		Name:     NsWSCN + ":DocumentFinalParameters",
		Required: true,
	}
	imageInformation := xmldoc.Lookup{
		Name:     NsWSCN + ":ImageInformation",
		Required: true,
	}
	jobId := xmldoc.Lookup{
		Name:     NsWSCN + ":JobId",
		Required: true,
	}
	jobToken := xmldoc.Lookup{
		Name:     NsWSCN + ":JobToken",
		Required: true,
	}

	if missed := root.Lookup(
		&documentFinalParameters, &imageInformation,
		&jobId, &jobToken,
	); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	var err error

	if r.DocumentFinalParameters, err = decodeDocumentParameters(
		documentFinalParameters.Elem); err != nil {
		return r, fmt.Errorf("DocumentFinalParameters: %w", err)
	}

	if r.ImageInformation, err = decodeImageInformation(
		imageInformation.Elem); err != nil {
		return r, fmt.Errorf("ImageInformation: %w", err)
	}

	if r.JobId, err = decodeNonNegativeInt(jobId.Elem); err != nil {
		return r, fmt.Errorf("JobId: %w", err)
	}
	if r.JobId < 1 {
		return r, fmt.Errorf("JobId: must be at least 1, got %d", r.JobId)
	}

	r.JobToken = jobToken.Elem.Text

	return r, nil
}
