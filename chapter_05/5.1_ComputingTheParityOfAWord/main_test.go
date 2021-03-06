package main

import (
	"testing"
)

func Test_BruteForceParity(t *testing.T) {
	tests := []struct {
		word     int8
		expected int8
	}{
		{
			word:     1,
			expected: 1,
		},
		{
			word:     2,
			expected: 1,
		},
		{
			word:     3,
			expected: 0,
		},
		{
			word:     127,
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

func Test_BruteForceXORParity(t *testing.T) {
	tests := []struct {
		word     int8
		expected int8
	}{
		{
			word:     1,
			expected: 1,
		},
		{
			word:     2,
			expected: 1,
		},
		{
			word:     3,
			expected: 0,
		},
		{
			word:     127,
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

func Test_DropLowestSetParity(t *testing.T) {
	tests := []struct {
		word     int8
		expected int8
	}{
		{
			word:     1,
			expected: 1,
		},
		{
			word:     2,
			expected: 1,
		},
		{
			word:     3,
			expected: 0,
		},
		{
			word:     127,
			expected: 1,
		},
	}

	for _, tt := range tests {
		ret := DropLowestSetParity(tt.word)
		if ret != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, ret)
		}
	}
}

func Test_RightShiftParity(t *testing.T) {
	tests := []struct {
		word     int64
		expected int64
	}{
		{
			word:     1,
			expected: 1,
		},
		{
			word:     2,
			expected: 1,
		},
		{
			word:     3,
			expected: 0,
		},
		{
			word:     127,
			expected: 1,
		},
	}

	for _, tt := range tests {
		ret := RightShiftParity(tt.word)
		if ret != tt.expected {
			t.Errorf("expected=%d, got=%d", tt.expected, ret)
		}
	}
}

func Benchmark_DropLowestSetParity(t *testing.B) {
	var word int8 = 112
	for i := 0; i < t.N; i++ {
		DropLowestSetParity(word)
	}
}

func Benchmark_BruteForceXORParity(t *testing.B) {
	var word int8 = 112
	for i := 0; i < t.N; i++ {
		BruteForceXORParity(word)
	}
}

func Benchmark_BruteForceParity(t *testing.B) {
	var word int8 = 112
	for i := 0; i < t.N; i++ {
		BruteForceParity(word)
	}
}

func Benchmark_RightShiftParity(t *testing.B) {
	var word int64 = 112
	for i := 0; i < t.N; i++ {
		RightShiftParity(word)
	}
}
