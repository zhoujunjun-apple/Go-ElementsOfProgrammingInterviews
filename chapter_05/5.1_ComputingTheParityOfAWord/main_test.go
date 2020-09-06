package main

import (
	"testing"
)

func Test_BruteForceParity(t *testing.T){
	tests := []struct{
		word int8
		expected int8
	}{
		{
			word: 1,
			expected: 1,
		},
		{
			word: 2,
			expected: 1,
		},
		{
			word: 3,
			expected: 0,
		},
		{
			word: 127,
			expected: 1,
		},
	}

	for _, tt := range tests {
		ret := BruteForceParity(tt.word)
		if ret != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, ret)
		}
	}
}

func Test_BruteForceXORParity(t *testing.T){
	tests := []struct{
		word int8
		expected int8
	}{
		{
			word: 1,
			expected: 1,
		},
		{
			word: 2,
			expected: 1,
		},
		{
			word: 3,
			expected: 0,
		},
		{
			word: 127,
			expected: 1,
		},
	}

	for _, tt := range tests {
		ret := BruteForceXORParity(tt.word)
		if ret != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, ret)
		}
	}
}