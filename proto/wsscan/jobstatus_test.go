// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scan job status tests

package wsscan

import (
	"testing"
	"time"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestJobStatus_toXML(t *testing.T) {
	createdTime, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00Z")
	completedTime, _ := time.Parse(time.RFC3339, "2024-01-01T13:00:00Z")

	js := JobStatus{
		JobCompletedTime: optional.New(completedTime),
		JobCreatedTime:   optional.New(createdTime),
		JobID:            123,
		JobState:         JobStateCompleted,
		JobStateReasons:  []JobStateReason{None},
		ScansCompleted:   5,
	}

	expected := xmldoc.Element{
		Name: "wscn:JobStatus",
		Children: []xmldoc.Element{
			{
				Name: "wscn:JobCompletedTime",
				Text: "2024-01-01T13:00:00Z",
			},
			{
				Name: "wscn:JobCreatedTime",
				Text: "2024-01-01T12:00:00Z",
			},
			{
				Name: "wscn:JobId",
				Text: "123",
			},
			{
				Name: "wscn:JobState",
				Text: "Completed",
			},
			{
				Name: "wscn:JobStateReasons",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobStateReason",
						Text: "None",
					},
				},
			},
			{
				Name: "wscn:ScansCompleted",
				Text: "5",
			},
		},
	}

	got := js.toXML("wscn:JobStatus")

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

func TestJobStatus_toXML_WithoutCompletedTime(t *testing.T) {
	createdTime, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00Z")

	js := JobStatus{
		JobCreatedTime: optional.New(createdTime),
		JobID:          123,
		JobState:       JobStatePending,
		ScansCompleted: 0,
	}

	got := js.toXML("wscn:JobStatus")

	// JobCompletedTime should not be present
	for _, child := range got.Children {
		if child.Name == "wscn:JobCompletedTime" {
			t.Error("toXML() should not include JobCompletedTime when zero")
		}
	}
}

func Test_decodeJobStatus(t *testing.T) {
	createdTime, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00Z")
	completedTime, _ := time.Parse(time.RFC3339, "2024-01-01T13:00:00Z")

	tests := []struct {
		name     string
		element  xmldoc.Element
		expected JobStatus
		wantErr  bool
	}{
		{
			name: "Valid JobStatus with all fields",
			element: xmldoc.Element{
				Name: "wscn:JobStatus",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobCompletedTime",
						Text: "2024-01-01T13:00:00Z",
					},
					{
						Name: "wscn:JobCreatedTime",
						Text: "2024-01-01T12:00:00Z",
					},
					{
						Name: "wscn:JobId",
						Text: "123",
					},
					{
						Name: "wscn:JobState",
						Text: "Completed",
					},
					{
						Name: "wscn:JobStateReasons",
						Children: []xmldoc.Element{
							{
								Name: "wscn:JobStateReason",
								Text: "None",
							},
						},
					},
					{
						Name: "wscn:ScansCompleted",
						Text: "5",
					},
				},
			},
			expected: JobStatus{
				JobCompletedTime: optional.New(completedTime),
				JobCreatedTime:   optional.New(createdTime),
				JobID:            123,
				JobState:         JobStateCompleted,
				JobStateReasons:  []JobStateReason{None},
				ScansCompleted:   5,
			},
			wantErr: false,
		},
		{
			name: "Valid JobStatus without JobCompletedTime",
			element: xmldoc.Element{
				Name: "wscn:JobStatus",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobCreatedTime",
						Text: "2024-01-01T12:00:00Z",
					},
					{
						Name: "wscn:JobId",
						Text: "456",
					},
					{
						Name: "wscn:JobState",
						Text: "Pending",
					},
					{
						Name: "wscn:ScansCompleted",
						Text: "0",
					},
				},
			},
			expected: JobStatus{
				JobCreatedTime: optional.New(createdTime),
				JobID:          456,
				JobState:       JobStatePending,
				ScansCompleted: 0,
			},
			wantErr: false,
		},
		{
			name: "Missing JobId",
			element: xmldoc.Element{
				Name: "wscn:JobStatus",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobCreatedTime",
						Text: "2024-01-01T12:00:00Z",
					},
					{
						Name: "wscn:JobState",
						Text: "Pending",
					},
					{
						Name: "wscn:ScansCompleted",
						Text: "0",
					},
				},
			},
			expected: JobStatus{},
			wantErr:  true,
		},
		{
			name: "Missing JobState",
			element: xmldoc.Element{
				Name: "wscn:JobStatus",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobCreatedTime",
						Text: "2024-01-01T12:00:00Z",
					},
					{
						Name: "wscn:JobId",
						Text: "123",
					},
					{
						Name: "wscn:ScansCompleted",
						Text: "0",
					},
				},
			},
			expected: JobStatus{},
			wantErr:  true,
		},
		{
			name: "Missing ScansCompleted",
			element: xmldoc.Element{
				Name: "wscn:JobStatus",
				Children: []xmldoc.Element{
					{
						Name: "wscn:JobCreatedTime",
						Text: "2024-01-01T12:00:00Z",
					},
					{
						Name: "wscn:JobId",
						Text: "123",
					},
					{
						Name: "wscn:JobState",
						Text: "Pending",
					},
				},
			},
			expected: JobStatus{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeJobStatus(tt.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeJobStatus() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Check JobCreatedTime
				if (got.JobCreatedTime != nil) != (tt.expected.JobCreatedTime != nil) {
					t.Errorf("decodeJobStatus().JobCreatedTime presence mismatch")
				} else if got.JobCreatedTime != nil {
					if optional.Get(got.JobCreatedTime) != optional.Get(tt.expected.JobCreatedTime) {
						t.Errorf("decodeJobStatus().JobCreatedTime = %v, want %v",
							optional.Get(got.JobCreatedTime), optional.Get(tt.expected.JobCreatedTime))
					}
				}
				// Check JobCompletedTime
				if (got.JobCompletedTime != nil) != (tt.expected.JobCompletedTime != nil) {
					t.Errorf("decodeJobStatus().JobCompletedTime presence mismatch")
				} else if got.JobCompletedTime != nil {
					if optional.Get(got.JobCompletedTime) != optional.Get(tt.expected.JobCompletedTime) {
						t.Errorf("decodeJobStatus().JobCompletedTime = %v, want %v",
							optional.Get(got.JobCompletedTime), optional.Get(tt.expected.JobCompletedTime))
					}
				}
				if got.JobID != tt.expected.JobID {
					t.Errorf("decodeJobStatus().JobId = %v, want %v",
						got.JobID, tt.expected.JobID)
				}
				if got.JobState != tt.expected.JobState {
					t.Errorf("decodeJobStatus().JobState = %v, want %v",
						got.JobState, tt.expected.JobState)
				}
				if got.ScansCompleted != tt.expected.ScansCompleted {
					t.Errorf("decodeJobStatus().ScansCompleted = %v, want %v",
						got.ScansCompleted, tt.expected.ScansCompleted)
				}
				if len(got.JobStateReasons) != len(tt.expected.JobStateReasons) {
					t.Errorf("decodeJobStatus().JobStateReasons length = %v, want %v",
						len(got.JobStateReasons), len(tt.expected.JobStateReasons))
				}
			}
		})
	}
}
