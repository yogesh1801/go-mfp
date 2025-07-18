package wsscan

import (
	"testing"
)

func TestBooleanElement_IsValid(t *testing.T) {
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
		if got := c.input.IsValid(); got != c.expected {
			t.Errorf("IsValid(%q) = %v; want %v", c.input, got, c.expected)
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
		got, err := c.input.Bool()
		if c.expectError {
			if err == nil {
				t.Errorf("Bool(%q) expected error, got nil", c.input)
			}
			continue
		}
		if err != nil {
			t.Errorf("Bool(%q) unexpected error: %v", c.input, err)
			continue
		}
		if got != c.expected {
			t.Errorf("Bool(%q) = %v; want %v", c.input, got, c.expected)
		}
	}
}
