// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD image edges test

package escl

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

// TestSupportedEdgesAddDel tests SupportedEdges.Add and SupportedEdges.Del operations
func TestSupportedEdgesAddDel(t *testing.T) {
	type testOp struct {
		action string
		val    SupportedEdge
	}

	type testData struct {
		seq []testOp
		res SupportedEdges
	}

	tests := []testData{
		{
			seq: nil,
			res: SupportedEdges{},
		},

		{
			seq: []testOp{
				{"add", TopEdge},
			},
			res: SupportedEdges{
				1 << TopEdge,
			},
		},

		{
			seq: []testOp{
				{"add", TopEdge},
				{"add", LeftEdge},
			},
			res: SupportedEdges{
				1<<TopEdge | 1<<LeftEdge,
			},
		},

		{
			seq: []testOp{
				{"add", TopEdge},
				{"add", LeftEdge},
				{"del", TopEdge},
				{"add", BottomEdge},
			},
			res: SupportedEdges{
				1<<LeftEdge | 1<<BottomEdge,
			},
		},
	}

	for _, test := range tests {
		var edges SupportedEdges
		seq := ""

		for _, op := range test.seq {
			seq += fmt.Sprintf("  %s %s\n", op.action, op.val)

			switch op.action {
			case "add":
				edges.Add(op.val)
			case "del":
				edges.Del(op.val)

			default:
				panic(fmt.Errorf("unknown action %q", op.action))
			}
		}

		if edges != test.res {
			t.Errorf("\nfor the sequence:\n%s"+
				"expected: %s\npresent:  %s",
				seq, test.res, edges)
		}
	}
}

// TestMakeSupportedEdges tests MakeSupportedEdges
func TestMakeSupportedEdges(t *testing.T) {
	type testData struct {
		in  []SupportedEdge
		res SupportedEdges
	}

	tests := []testData{
		{[]SupportedEdge{}, SupportedEdges{}},
		{[]SupportedEdge{TopEdge},
			SupportedEdges{
				1 << TopEdge,
			},
		},
		{[]SupportedEdge{TopEdge, BottomEdge, LeftEdge},
			SupportedEdges{
				1<<TopEdge |
					1<<BottomEdge |
					1<<LeftEdge,
			},
		},
	}

	for _, test := range tests {
		edges := MakeSupportedEdges(test.in...)
		if edges != test.res {
			t.Errorf("\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.res, edges)
		}
	}
}

// TestSupportedEdgesString tests SupportedEdges.String
func TestSupportedEdgesString(t *testing.T) {
	type testData struct {
		edges SupportedEdges
		s     string
	}

	tests := []testData{
		{SupportedEdges{}, ""},
		{MakeSupportedEdges(TopEdge), "TopEdge"},
		{MakeSupportedEdges(TopEdge, LeftEdge),
			"TopEdge,LeftEdge"},
		{MakeSupportedEdges(TopEdge, LeftEdge),
			"LeftEdge,TopEdge"},
		{MakeSupportedEdges(TopEdge, BottomEdge, LeftEdge),
			"TopEdge,BottomEdge,LeftEdge"},
	}

	for _, test := range tests {
		s := test.edges.String()

		// Compare resulting strings, ignoring the order
		// of color modes in the output.
		out := strings.Split(s, ",")
		exp := strings.Split(test.s, ",")

		slices.Sort(out)
		slices.Sort(exp)

		if !slices.Equal(out, exp) {
			t.Errorf("%s:\n"+
				"expected: %s\n"+
				"present:  %s",
				test.edges, exp, out)
		}
	}
}
