package wsscan

import (
	"testing"
)

// Test for BooleanElement
func TestBooleanElement_Validate(t *testing.T) {
	cases := []struct {
		input    BooleanElement
		expected bool
	}{
		{"1", true},
		{"0", true},
		{"true", true},
		{"false", true},
		{"TRUE", true},
		{"FALSE", true},
		{" yes ", false},
		{"no", false},
		{"2", false},
		{"", false},
	}
	for _, c := range cases {
		err := c.input.Validate()
		if c.expected {
			if err != nil {
				t.Errorf("Validate(%q) = %v; want nil", c.input, err)
			}
		} else {
			if err == nil {
				t.Errorf("Validate(%q) = nil; want error", c.input)
			}
		}
	}
}

func TestBooleanElement_Bool(t *testing.T) {
	cases := []struct {
		input       BooleanElement
		expected    bool
		expectError bool
	}{
		{"1", true, false},
		{"0", false, false},
		{"true", true, false},
		{"false", false, false},
		{"TRUE", true, false},
		{"FALSE", false, false},
		{" yes ", false, true},
		{"no", false, true},
		{"2", false, true},
		{"", false, true},
	}
	for _, c := range cases {
		got := c.input.Bool()
		if c.expectError {
			if err := c.input.Validate(); err == nil {
				t.Errorf("Bool(%q) expected error, got nil", c.input)
			}
			continue
		}
		if got != c.expected {
			t.Errorf("Bool(%q) = %v; want %v", c.input, got, c.expected)
		}
	}
}
