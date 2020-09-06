package main

import (
	"testing"
)

func Test_NativeStrToInt(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{
			str:      "321",
			expected: 321,
		},
		{
			str:      "-124",
			expected: -124,
		},
		{
			str:      "-1",
			expected: -1,
		},
	}

	for _, tt := range tests {
		ret := NativeStrToInt(tt.str)
		if ret != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, ret)
		}
	}
}

func Test_SubstractStrToInt(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{
			str:      "321",
			expected: 321,
		},
		{
			str:      "-124",
			expected: -124,
		},
		{
			str:      "-1",
			expected: -1,
		},
	}

	for _, tt := range tests {
		ret := SubstractStrToInt(tt.str)
		if ret != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, ret)
		}
	}
}

func Test_IntToStr(t *testing.T) {
	tests := []struct {
		integer  int
		expected string
	}{
		{
			expected: "321",
			integer:  321,
		},
		{
			expected: "-124",
			integer:  -124,
		},
		{
			expected: "-1",
			integer:  -1,
		},
		{
			expected: "0",
			integer:  0,
		},
	}

	for _, tt := range tests {
		ret := IntToStr(tt.integer)
		if ret != tt.expected {
			t.Errorf("expected=%s, got=%s", tt.expected, ret)
		}
	}
}
