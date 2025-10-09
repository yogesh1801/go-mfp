// MFP - Miulti-Function Printers and scanners toolkit
// IEEE 1284 definitions
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IEEE-1284 tests

package ieee1284

import (
	"reflect"
	"testing"
)

// TestDeviceIDParse tests DeviceIDParse function
func TestDeviceIDParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *DeviceID
	}{
		{
			name:  "standard device ID",
			input: "MFG:Hewlett-Packard;CMD:PJL,POSTSCRIPT;MDL:HP LaserJet 4000;CLS:PRINTER;",
			expected: &DeviceID{
				RawData: "MFG:Hewlett-Packard;CMD:PJL,POSTSCRIPT;MDL:HP LaserJet 4000;CLS:PRINTER;",
				Records: []DeviceIDRecord{
					{Key: "MFG", RawData: "Hewlett-Packard", Values: []string{"Hewlett-Packard"}},
					{Key: "CMD", RawData: "PJL,POSTSCRIPT", Values: []string{"PJL", "POSTSCRIPT"}},
					{Key: "MDL", RawData: "HP LaserJet 4000", Values: []string{"HP LaserJet 4000"}},
					{Key: "CLS", RawData: "PRINTER", Values: []string{"PRINTER"}},
				},
			},
		},
		{
			name:  "with spaces and empty fields",
			input: " MFG: Canon ; CMD: PCL, POSTSCRIPT ; MDL: Canon MF4400 ; ; SERN: ABC123 ; ",
			expected: &DeviceID{
				RawData: " MFG: Canon ; CMD: PCL, POSTSCRIPT ; MDL: Canon MF4400 ; ; SERN: ABC123 ; ",
				Records: []DeviceIDRecord{
					{Key: "MFG", RawData: "Canon", Values: []string{"Canon"}},
					{Key: "CMD", RawData: "PCL, POSTSCRIPT", Values: []string{"PCL", "POSTSCRIPT"}},
					{Key: "MDL", RawData: "Canon MF4400", Values: []string{"Canon MF4400"}},
					{Key: "SERN", RawData: "ABC123", Values: []string{"ABC123"}},
				},
			},
		},
		{
			name:  "record without value",
			input: "MFG:Epson;UNKNOWN_KEY;MDL:XP-4100;",
			expected: &DeviceID{
				RawData: "MFG:Epson;UNKNOWN_KEY;MDL:XP-4100;",
				Records: []DeviceIDRecord{
					{Key: "MFG", RawData: "Epson", Values: []string{"Epson"}},
					{Key: "UNKNOWN_KEY", RawData: "", Values: []string{""}},
					{Key: "MDL", RawData: "XP-4100", Values: []string{"XP-4100"}},
				},
			},
		},
		{
			name:  "empty string",
			input: "",
			expected: &DeviceID{
				RawData: "",
				Records: []DeviceIDRecord{},
			},
		},
		{
			name:  "only semicolons",
			input: ";;;",
			expected: &DeviceID{
				RawData: ";;;",
				Records: []DeviceIDRecord{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DeviceIDParse(tt.input)

			if result.RawData != tt.expected.RawData {
				t.Errorf("RawData mismatch: got %q, want %q", result.RawData, tt.expected.RawData)
			}

			if len(result.Records) != len(tt.expected.Records) {
				t.Errorf("Records count mismatch: got %d, want %d", len(result.Records), len(tt.expected.Records))
			}

			for i, rec := range result.Records {
				if i >= len(tt.expected.Records) {
					break
				}

				expectedRec := tt.expected.Records[i]
				if rec.Key != expectedRec.Key {
					t.Errorf("Record %d Key mismatch: got %q, want %q", i, rec.Key, expectedRec.Key)
				}
				if rec.RawData != expectedRec.RawData {
					t.Errorf("Record %d RawData mismatch: got %q, want %q", i, rec.RawData, expectedRec.RawData)
				}
				if !reflect.DeepEqual(rec.Values, expectedRec.Values) {
					t.Errorf("Record %d Values mismatch: got %v, want %v", i, rec.Values, expectedRec.Values)
				}
			}
		})
	}
}

// TestDeviceID_Get tests the DeviceID.Get method
func TestDeviceID_Get(t *testing.T) {
	devid := DeviceIDParse("MFG:Brother;CMD:PCL;MDL:HL-L2340D;SN:ABCD1234;")

	tests := []struct {
		name     string
		key      string
		expected *DeviceIDRecord
	}{
		{
			name:     "existing key uppercase",
			key:      "MFG",
			expected: &DeviceIDRecord{Key: "MFG", RawData: "Brother", Values: []string{"Brother"}},
		},
		{
			name:     "existing key lowercase",
			key:      "mfg",
			expected: &DeviceIDRecord{Key: "MFG", RawData: "Brother", Values: []string{"Brother"}},
		},
		{
			name:     "existing key mixed case",
			key:      "MfG",
			expected: &DeviceIDRecord{Key: "MFG", RawData: "Brother", Values: []string{"Brother"}},
		},
		{
			name:     "non-existing key",
			key:      "UNKNOWN",
			expected: nil,
		},
		{
			name:     "empty key",
			key:      "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := devid.Get(tt.key)

			if tt.expected == nil {
				if result != nil {
					t.Errorf("Expected nil, got %+v", result)
				}
				return
			}

			if result == nil {
				t.Errorf("Expected record, got nil")
				return
			}

			if result.Key != tt.expected.Key {
				t.Errorf("Key mismatch: got %q, want %q", result.Key, tt.expected.Key)
			}
			if result.RawData != tt.expected.RawData {
				t.Errorf("RawData mismatch: got %q, want %q", result.RawData, tt.expected.RawData)
			}
			if !reflect.DeepEqual(result.Values, tt.expected.Values) {
				t.Errorf("Values mismatch: got %v, want %v", result.Values, tt.expected.Values)
			}
		})
	}
}

// TestDeviceID_Manufacturer tests the DeviceID.Manufacturer method
func TestDeviceID_Manufacturer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"MANUFACTURER key", "MANUFACTURER:Canon;MDL:PIXMA;", "Canon"},
		{"MFG key", "MFG:Canon;MDL:PIXMA;", "Canon"},
		{"MANUFACTURER preferred over MFG", "MANUFACTURER:Canon Inc;MFG:Canon;MDL:PIXMA;", "Canon Inc"},
		{"no manufacturer", "MDL:PIXMA;CMD:PCL;", ""},
		{"empty manufacturer", "MANUFACTURER:;MDL:PIXMA;", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			devid := DeviceIDParse(tt.input)
			result := devid.Manufacturer()
			if result != tt.expected {
				t.Errorf("Manufacturer() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestDeviceID_Model tests the DeviceID.Model method
func TestDeviceID_Model(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"MODEL key", "MFG:Epson;MODEL:XP-4100;", "XP-4100"},
		{"MDL key", "MFG:Epson;MDL:XP-4100;", "XP-4100"},
		{"MODEL preferred over MDL", "MFG:Epson;MODEL:XP-4100 Series;MDL:XP-4100;", "XP-4100 Series"},
		{"no model", "MFG:Epson;CMD:PCL;", ""},
		{"empty model", "MFG:Epson;MODEL:;", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			devid := DeviceIDParse(tt.input)
			result := devid.Model()
			if result != tt.expected {
				t.Errorf("Model() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestDeviceID_CommandSet tests the DeviceID.CommandSet method
func TestDeviceID_CommandSet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "COMMAND SET key",
			input:    "MFG:HP;COMMAND SET:PJL,POSTSCRIPT,PCL;MDL:LaserJet;",
			expected: []string{"PJL", "POSTSCRIPT", "PCL"},
		},
		{
			name:     "CMD key",
			input:    "MFG:HP;CMD:PJL,POSTSCRIPT;MDL:LaserJet;",
			expected: []string{"PJL", "POSTSCRIPT"},
		},
		{
			name:     "COMMAND SET preferred over CMD",
			input:    "MFG:HP;COMMAND SET:PJL,POSTSCRIPT;CMD:PCL;MDL:LaserJet;",
			expected: []string{"PJL", "POSTSCRIPT"},
		},
		{
			name:     "no command set",
			input:    "MFG:HP;MDL:LaserJet;",
			expected: nil,
		},
		{
			name:     "empty command set",
			input:    "MFG:HP;COMMAND SET:;MDL:LaserJet;",
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			devid := DeviceIDParse(tt.input)
			result := devid.CommandSet()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CommandSet() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestDeviceID_SerialNumber tests the DeviceID.SerialNumber method
func TestDeviceID_SerialNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"SERIALNUMBER key", "MFG:Dell;MDL:2355;SERIALNUMBER:12345ABC;", "12345ABC"},
		{"SERN key", "MFG:Dell;MDL:2355;SERN:12345ABC;", "12345ABC"},
		{"SN key", "MFG:Dell;MDL:2355;SN:12345ABC;", "12345ABC"},
		{"SERIALNUMBER preferred over SERN", "MFG:Dell;MDL:2355;SERIALNUMBER:12345ABC;SERN:67890DEF;", "12345ABC"},
		{"SERIALNUMBER preferred over SN", "MFG:Dell;MDL:2355;SERIALNUMBER:12345ABC;SN:67890DEF;", "12345ABC"},
		{"SERN preferred over SN", "MFG:Dell;MDL:2355;SERN:12345ABC;SN:67890DEF;", "12345ABC"},
		{"no serial", "MFG:Dell;MDL:2355;", ""},
		{"empty serial", "MFG:Dell;MDL:2355;SERIALNUMBER:;", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			devid := DeviceIDParse(tt.input)
			result := devid.SerialNumber()
			if result != tt.expected {
				t.Errorf("SerialNumber() = %q, want %q", result, tt.expected)
			}
		})
	}
}
