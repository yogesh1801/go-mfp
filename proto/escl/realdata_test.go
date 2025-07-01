// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Real-data tests

package escl

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestKyoceraECOSYSM2040dnScannerCapabilities tests ScannerCapabilities decoding
// for Kyocera ECOSYS M2040dn MFP
func TestKyoceraECOSYSM2040dnScannerCapabilities(t *testing.T) {
	// Parse XML
	data := bytes.NewReader(testutils.
		Kyocera.ECOSYS.M2040dn.ESCL.ScannerCapabilities)
	xml, err := xmldoc.Decode(NsMap, data)
	if err != nil {
		panic(err)
	}

	// Decode ScannerCapabilities
	scancaps, err := DecodeScannerCapabilities(xml)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	// Verify ScannerCapabilities
	expected := ScannerCapabilities{
		Version:      MakeVersion(2, 62),
		MakeAndModel: optional.New("Kyocera ECOSYS M2040dn"),
		SerialNumber: optional.New("VCF9192281"),
		UUID: optional.New(
			uuid.Must(uuid.Parse("4509a320-00a0-008f-00b6-002507510eca"))),
		AdminURI: optional.New("https://KM7B6A91.local/airprint"),
		IconURI:  optional.New("https://KM7B6A91.local/printer-icon/machine_128.png"),
		Platen: &Platen{
			PlatenInputCaps: &InputSourceCaps{
				MinWidth:  118,
				MaxWidth:  2551,
				MinHeight: 118,
				MaxHeight: 3508,
				SupportedIntents: []Intent{
					Document,
					TextAndGraphic,
					Photo,
					Preview,
				},
				SettingProfiles: []SettingProfile{
					SettingProfile{
						ColorModes: []ColorMode{
							BlackAndWhite1,
							Grayscale8,
							RGB24,
						},
						DocumentFormats: []string{
							"image/jpeg",
							"application/pdf",
						},
						SupportedResolutions: []SupportedResolutions{{
							DiscreteResolutions: DiscreteResolutions{
								DiscreteResolution{
									XResolution: 200,
									YResolution: 100,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 200,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 400,
								},
								DiscreteResolution{
									XResolution: 300,
									YResolution: 300,
								},
								DiscreteResolution{
									XResolution: 400,
									YResolution: 400,
								},
								DiscreteResolution{
									XResolution: 600,
									YResolution: 600,
								},
							}},
						},
					},
				},
				FeedDirections: []FeedDirection{
					ShortEdgeFeed,
					LongEdgeFeed,
				},
			},
		},
		ADF: &ADF{
			ADFSimplexInputCaps: &InputSourceCaps{
				MinWidth:  591,
				MaxWidth:  2551,
				MinHeight: 591,
				MaxHeight: 4205,
				SupportedIntents: []Intent{
					Document,
					TextAndGraphic,
					Photo,
					Preview,
				},
				SettingProfiles: []SettingProfile{
					SettingProfile{
						ColorModes: []ColorMode{
							BlackAndWhite1,
							Grayscale8,
							RGB24,
						},
						DocumentFormats: []string{
							"image/jpeg",
							"application/pdf",
						},
						SupportedResolutions: []SupportedResolutions{{
							DiscreteResolutions: DiscreteResolutions{
								DiscreteResolution{
									XResolution: 200,
									YResolution: 100,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 200,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 400,
								},
								DiscreteResolution{
									XResolution: 300,
									YResolution: 300,
								},
								DiscreteResolution{
									XResolution: 400,
									YResolution: 400,
								},
								DiscreteResolution{
									XResolution: 600,
									YResolution: 600,
								},
							},
						}},
					},
				},
			},
			ADFDuplexInputCaps: &InputSourceCaps{
				MinWidth:  591,
				MaxWidth:  2551,
				MinHeight: 591,
				MaxHeight: 4205,
				SupportedIntents: []Intent{
					Document,
					TextAndGraphic,
					Photo,
					Preview,
				},
				SettingProfiles: []SettingProfile{
					SettingProfile{
						ColorModes: []ColorMode{
							BlackAndWhite1,
							Grayscale8,
							RGB24,
						},
						DocumentFormats: []string{
							"image/jpeg",
							"application/pdf",
						},
						SupportedResolutions: []SupportedResolutions{{
							DiscreteResolutions: DiscreteResolutions{
								DiscreteResolution{
									XResolution: 200,
									YResolution: 100,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 200,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 400,
								},
								DiscreteResolution{
									XResolution: 300,
									YResolution: 300,
								},
								DiscreteResolution{
									XResolution: 400,
									YResolution: 400,
								},
								DiscreteResolution{
									XResolution: 600,
									YResolution: 600,
								},
							},
						}},
					},
				},
				FeedDirections: []FeedDirection{
					ShortEdgeFeed,
					LongEdgeFeed,
				},
			},
			FeederCapacity: optional.New(75),
			ADFOptions: []ADFOption{
				DetectPaperLoaded,
				SelectSinglePage,
				Duplex,
			},
		},
		CompressionFactorSupport: &Range{
			Min:    1,
			Max:    5,
			Normal: 1,
			Step:   optional.New(1),
		},
		SharpenSupport: &Range{
			Min:    -3,
			Max:    3,
			Normal: 0,
			Step:   optional.New(1),
		},
	}

	if !reflect.DeepEqual(scancaps, expected) {
		t.Errorf("decoded data mismatch:\n"+
			"expected: %#v\n"+
			"present:  %#v\n",
			expected, scancaps)
	}
}

// TestKyoceraECOSYSM2040dnScannerStatus tests ScannerStatus decoding
// for Kyocera ECOSYS M2040dn MFP
func TestKyoceraECOSYSM2040dnScannerStatus(t *testing.T) {
	// Parse XML
	data := bytes.NewReader(testutils.
		Kyocera.ECOSYS.M2040dn.ESCL.ScannerStatus)
	xml, err := xmldoc.Decode(NsMap, data)
	if err != nil {
		panic(err)
	}

	// Decode ScannerStatus
	status, err := DecodeScannerStatus(xml)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	// Verify ScannerStatus
	expected := ScannerStatus{
		Version:  MakeVersion(2, 62),
		State:    ScannerProcessing,
		ADFState: optional.New(ScannerAdfProcessing),
		Jobs: []JobInfo{
			{
				JobURI:          "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32",
				JobUUID:         optional.New("urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				Age:             optional.New(2 * time.Second),
				ImagesCompleted: optional.New(0),
				JobState:        JobProcessing,
				JobStateReasons: []JobStateReason{JobScanningAndTransferring},
			},
			{
				JobURI:          "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d31",
				JobUUID:         optional.New("urn:uuid:4509a320-00a0-008f-00b6-00559a327d31"),
				Age:             optional.New(19 * time.Second),
				ImagesCompleted: optional.New(1),
				JobState:        JobCompleted,
				JobStateReasons: []JobStateReason{JobCompletedSuccessfully},
			},
			{
				JobURI:          "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d30",
				JobUUID:         optional.New("urn:uuid:4509a320-00a0-008f-00b6-00559a327d30"),
				Age:             optional.New(35 * time.Second),
				ImagesCompleted: optional.New(1),
				JobState:        JobCompleted,
				JobStateReasons: []JobStateReason{JobCompletedSuccessfully},
			},
			{
				JobURI:          "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d2f",
				JobUUID:         optional.New("urn:uuid:4509a320-00a0-008f-00b6-00559a327d2f"),
				Age:             optional.New(60 * time.Second),
				ImagesCompleted: optional.New(1),
				JobState:        JobCompleted,
				JobStateReasons: []JobStateReason{JobCompletedSuccessfully},
			},
			{
				JobURI:          "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d07",
				JobUUID:         optional.New("urn:uuid:4509a320-00a0-008f-00b6-00559a327d07"),
				Age:             optional.New(72 * time.Second),
				ImagesCompleted: optional.New(1),
				JobState:        JobCompleted,
				JobStateReasons: []JobStateReason{JobCompletedSuccessfully},
			},
		},
	}

	if !reflect.DeepEqual(status, expected) {
		t.Errorf("decoded data mismatch:\n"+
			"expected: %#v\n"+
			"present:  %#v\n",
			expected, status)
	}
}

// TestHPLaserJetM426fdnScannerCapabilities tests ScannerCapabilities decoding
// for the HP LaserJet M426fdn
func TestHPLaserJetM426fdnScannerCapabilities(t *testing.T) {
	// Parse XML
	data := bytes.NewReader(testutils.
		HP.LaserJet.M426fdn.ESCL.ScannerCapabilities)
	xml, err := xmldoc.Decode(NsMap, data)
	if err != nil {
		panic(err)
	}

	// Decode ScannerCapabilities
	scancaps, err := DecodeScannerCapabilities(xml)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	// Verify ScannerCapabilities
	expected := ScannerCapabilities{
		Version:      MakeVersion(2, 5),
		MakeAndModel: optional.New("HP LaserJet MFP M426fdn"),
		SerialNumber: optional.New("PHBLL6F5GB"),
		Platen: &Platen{
			PlatenInputCaps: &InputSourceCaps{
				MinWidth:       576,
				MaxWidth:       2550,
				MinHeight:      576,
				MaxHeight:      3508,
				MaxScanRegions: optional.New(1),
				SupportedIntents: []Intent{
					Document,
					Photo,
					Preview,
					TextAndGraphic,
				},
				MaxOpticalXResolution: optional.New(1200),
				MaxOpticalYResolution: optional.New(1200),
				SettingProfiles: []SettingProfile{
					SettingProfile{
						ColorModes: []ColorMode{
							RGB24,
							Grayscale8,
						},
						ContentTypes: []ContentType{
							ContentTypePhoto,
							ContentTypeText,
							ContentTypeTextAndPhoto,
						},
						DocumentFormats: []string{
							"image/jpeg",
							"application/pdf",
						},
						DocumentFormatsExt: []string{
							"image/jpeg",
							"application/pdf",
						},
						SupportedResolutions: []SupportedResolutions{{
							DiscreteResolutions: DiscreteResolutions{
								DiscreteResolution{
									XResolution: 75,
									YResolution: 75,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 200,
								},
								DiscreteResolution{
									XResolution: 300,
									YResolution: 300,
								},
								DiscreteResolution{
									XResolution: 600,
									YResolution: 600,
								},
								DiscreteResolution{
									XResolution: 1200,
									YResolution: 1200,
								},
							},
						}},
						ColorSpaces: []ColorSpace{SRGB},
					},
				},
			},
		},
		ADF: &ADF{
			ADFSimplexInputCaps: &InputSourceCaps{
				MinWidth:       576,
				MaxWidth:       2550,
				MinHeight:      576,
				MaxHeight:      4500,
				MaxScanRegions: optional.New(1),
				SupportedIntents: []Intent{
					Document,
					Photo,
					Preview,
					TextAndGraphic,
				},
				MaxOpticalXResolution: optional.New(300),
				MaxOpticalYResolution: optional.New(300),
				SettingProfiles: []SettingProfile{
					SettingProfile{
						ColorModes: []ColorMode{
							RGB24,
							Grayscale8,
						},
						ContentTypes: []ContentType{
							ContentTypePhoto,
							ContentTypeText,
							ContentTypeTextAndPhoto,
						},
						DocumentFormats: []string{
							"image/jpeg",
							"application/pdf",
						},
						DocumentFormatsExt: []string{
							"image/jpeg",
							"application/pdf",
						},
						SupportedResolutions: []SupportedResolutions{{
							DiscreteResolutions: DiscreteResolutions{
								DiscreteResolution{
									XResolution: 75,
									YResolution: 75,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 200,
								},
								DiscreteResolution{
									XResolution: 300,
									YResolution: 300,
								},
							},
						}},
						ColorSpaces: []ColorSpace{SRGB},
					},
				},
			},
			ADFDuplexInputCaps: &InputSourceCaps{
				MinWidth:       576,
				MaxWidth:       2550,
				MinHeight:      576,
				MaxHeight:      4500,
				MaxScanRegions: optional.New(1),
				SupportedIntents: []Intent{
					Document,
					Photo,
					Preview,
					TextAndGraphic,
				},
				MaxOpticalXResolution: optional.New(300),
				MaxOpticalYResolution: optional.New(300),
				SettingProfiles: []SettingProfile{
					SettingProfile{
						ColorModes: []ColorMode{
							RGB24,
							Grayscale8,
						},
						ContentTypes: []ContentType{
							ContentTypePhoto,
							ContentTypeText,
							ContentTypeTextAndPhoto,
						},
						DocumentFormats: []string{
							"image/jpeg",
							"application/pdf",
						},
						DocumentFormatsExt: []string{
							"image/jpeg",
							"application/pdf",
						},
						SupportedResolutions: []SupportedResolutions{{
							DiscreteResolutions: DiscreteResolutions{
								DiscreteResolution{
									XResolution: 75,
									YResolution: 75,
								},
								DiscreteResolution{
									XResolution: 200,
									YResolution: 200,
								},
								DiscreteResolution{
									XResolution: 300,
									YResolution: 300,
								},
							},
						}},
						ColorSpaces: []ColorSpace{SRGB},
					},
				},
			},
			FeederCapacity: optional.New(50),
			ADFOptions: []ADFOption{
				DetectPaperLoaded,
				Duplex,
				SelectSinglePage,
			},
		},
		ContrastSupport: &Range{
			Min:    -127,
			Max:    127,
			Normal: 0,
			Step:   optional.New(1),
		},
	}

	if !reflect.DeepEqual(scancaps, expected) {
		t.Errorf("decoded data mismatch:\n"+
			"expected: %#v\n"+
			"present:  %#v\n",
			expected, scancaps)
	}
}

// TestHPLaserJetM426fdnScannerStatus tests ScannerStatus decoding
// for the HP LaserJet M426fdn
func TestHPLaserJetM426fdnScannerStatus(t *testing.T) {
	// Parse XML
	data := bytes.NewReader(testutils.
		HP.LaserJet.M426fdn.ESCL.ScannerStatus)
	xml, err := xmldoc.Decode(NsMap, data)
	if err != nil {
		panic(err)
	}

	// Decode ScannerStatus
	status, err := DecodeScannerStatus(xml)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	// Verify ScannerStatus
	expected := ScannerStatus{
		Version:  MakeVersion(2, 5),
		State:    ScannerIdle,
		ADFState: optional.New(ScannerAdfEmpty),
		Jobs: []JobInfo{
			{
				JobURI:             "/eSCL/ScanJobs/1005",
				JobUUID:            optional.New("166-1005"),
				Age:                optional.New(2 * time.Second),
				ImagesCompleted:    optional.New(1),
				ImagesToTransfer:   optional.New(0),
				TransferRetryCount: optional.New(0),
				JobState:           JobCompleted,
				JobStateReasons:    []JobStateReason{JobCompletedSuccessfully},
			},
			{
				JobURI:             "/eSCL/ScanJobs/1004",
				JobUUID:            optional.New("166-1004"),
				Age:                optional.New(50 * time.Second),
				ImagesCompleted:    optional.New(1),
				ImagesToTransfer:   optional.New(0),
				TransferRetryCount: optional.New(0),
				JobState:           JobCompleted,
				JobStateReasons:    []JobStateReason{JobCompletedSuccessfully},
			},
		},
	}

	if !reflect.DeepEqual(status, expected) {
		t.Errorf("decoded data mismatch:\n"+
			"expected: %#v\n"+
			"present:  %#v\n",
			expected, status)
	}
}
