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

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/testutils"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/xmldoc"
)

// TestKyoceraECOSYSM2040dnScannerCapabilities ScannerCapabilities decoding
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
						SupportedResolutions: SupportedResolutions{
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
						SupportedResolutions: SupportedResolutions{
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
						},
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
						SupportedResolutions: SupportedResolutions{
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
						},
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
