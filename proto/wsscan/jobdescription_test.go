// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scan job description tests

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestJobDescription_toXML(t *testing.T) {
	jd := JobDescription{
		JobInformation:         optional.New("Scan job for accounting"),
		JobName:                "Invoice Scan",
		JobOriginatingUserName: "john.doe",
	}

	expected := xmldoc.Element{
		Name: "wscn:JobDescription",
		Children: []xmldoc.Element{
			{
				Name: "wscn:JobInformation",
				Text: "Scan job for accounting",
			},
			{
				Name: "wscn:JobName",
				Text: "Invoice Scan",
			},
			{
				Name: "wscn:JobOriginatingUserName",
				Text: "john.doe",
			},
		},
	}

	got := jd.toXML("wscn:JobDescription")

	if got.Name != expected.Name {
		t.Errorf("toXML().Name = %v, want %v", got.Name, expected.Name)
	}

	if len(got.Children) != len(expected.Children) {
		t.Errorf("toXML().Children length = %v, want %v", len(got.Children),
			len(expected.Children))
		return
	}

	for i, child := range got.Children {
		if child.Name != expected.Children[i].Name {
			t.Errorf("toXML().Children[%d].Name = %v, want %v",
				i, child.Name, expected.Children[i].Name)
		}
		if child.Text != expected.Children[i].Text {
			t.Errorf("toXML().Children[%d].Text = %v, want %v",
				i, child.Text, expected.Children[i].Text)
		}
	}
}

func TestJobDescription_toXML_WithoutJobInformation(t *testing.T) {
	jd := JobDescription{
		JobName:                "Invoice Scan",
		JobOriginatingUserName: "john.doe",
	}

	got := jd.toXML("wscn:JobDescription")

	// JobInformation should not be present
	for _, child := range got.Children {
		if child.Name == "wscn:JobInformation" {
			t.Error("toXML() should not include JobInformation when nil")
		}
	}

	// Should have JobName and JobOriginatingUserName
	foundJobName := false
	foundJobOriginatingUserName := false
	for _, child := range got.Children {
		if child.Name == "wscn:JobName" {
			foundJobName = true
			if child.Text != "Invoice Scan" {
				t.Errorf("toXML().JobName = %v, want %v", child.Text, "Invoice Scan")
			}
		}
		if child.Name == "wscn:JobOriginatingUserName" {
			foundJobOriginatingUserName = true
			if child.Text != "john.doe" {
				t.Errorf("toXML().JobOriginatingUserName = %v, want %v",
					child.Text, "john.doe")
			}
		}
	}

	if !foundJobName {
		t.Error("toXML() should include JobName")
	}
	if !foundJobOriginatingUserName {
		t.Error("toXML() should include JobOriginatingUserName")
	}
}

func Test_decodeJobDescription(t *testing.T) {
	tests := []struct {
		name     string
		element  xmldoc.Element
		expected JobDescription
		wantErr  bool
	}{
		{
			name: "Valid JobDescription with all fields",
			element: xmldoc.Element{
				Name: "wscn:JobDescription",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobInformation",
						Text: "Scan job for accounting",
					},
					{
						Name: "wscn:JobName",
						Text: "Invoice Scan",
					},
					{
						Name: "wscn:JobOriginatingUserName",
						Text: "john.doe",
					},
				},
			},
			expected: JobDescription{
				JobInformation:         optional.New("Scan job for accounting"),
				JobName:                "Invoice Scan",
				JobOriginatingUserName: "john.doe",
			},
			wantErr: false,
		},
		{
			name: "Valid JobDescription without JobInformation",
			element: xmldoc.Element{
				Name: "wscn:JobDescription",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobName",
						Text: "Document Scan",
					},
					{
						Name: "wscn:JobOriginatingUserName",
						Text: "jane.smith",
					},
				},
			},
			expected: JobDescription{
				JobName:                "Document Scan",
				JobOriginatingUserName: "jane.smith",
			},
			wantErr: false,
		},
		{
			name: "Missing JobName",
			element: xmldoc.Element{
				Name: "wscn:JobDescription",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobOriginatingUserName",
						Text: "john.doe",
					},
				},
			},
			expected: JobDescription{},
			wantErr:  true,
		},
		{
			name: "Missing JobOriginatingUserName",
			element: xmldoc.Element{
				Name: "wscn:JobDescription",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobName",
						Text: "Invoice Scan",
					},
				},
			},
			expected: JobDescription{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeJobDescription(tt.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeJobDescription() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.JobName != tt.expected.JobName {
					t.Errorf("decodeJobDescription().JobName = %v, want %v",
						got.JobName, tt.expected.JobName)
				}
				if got.JobOriginatingUserName != tt.expected.JobOriginatingUserName {
					t.Errorf("decodeJobDescription().JobOriginatingUserName = %v, want %v",
						got.JobOriginatingUserName, tt.expected.JobOriginatingUserName)
				}
				// Check JobInformation
				if (got.JobInformation != nil) != (tt.expected.JobInformation != nil) {
					t.Errorf("decodeJobDescription().JobInformation presence mismatch")
				} else if got.JobInformation != nil {
					if optional.Get(got.JobInformation) != optional.Get(tt.expected.JobInformation) {
						t.Errorf("decodeJobDescription().JobInformation = %v, want %v",
							optional.Get(got.JobInformation), optional.Get(tt.expected.JobInformation))
					}
				}
			}
		})
	}
}
