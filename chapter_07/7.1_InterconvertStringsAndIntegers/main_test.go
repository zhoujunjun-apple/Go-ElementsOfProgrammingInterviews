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