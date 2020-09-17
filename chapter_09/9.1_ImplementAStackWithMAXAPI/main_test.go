package main

import (
	"testing"
)

func Test_Max(t *testing.T) {
	tests := []struct {
		es          []int
		expectedMax int
	}{
		{
			es:          []int{1, 4, 2, 3, 12, 6, 7, 8},
			expectedMax: 12,
		},
		{
			es:          []int{3, 5, 9, 3, 6, 1},
			expectedMax: 9,
		},
	}

	for _, tt := range tests {
		stk := Stack{}
		stk.Pushs(tt.es)

		max, _ := stk.Max()
		if max != tt.expectedMax {
			t.Errorf("max expected=%d, got=%d", tt.expectedMax, max)
		}
	}
}
