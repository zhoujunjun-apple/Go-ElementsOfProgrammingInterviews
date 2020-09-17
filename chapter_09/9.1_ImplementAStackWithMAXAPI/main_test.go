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

func Test_StackMaxPush(t *testing.T) {
	tests := []struct {
		eleSlice []int
		maxSlice []int
	}{
		{
			eleSlice: []int{1, 2, 3, 5, 3, 6, 2, 1, 1},
			maxSlice: []int{1, 2, 3, 5, 5, 6, 6, 6, 6},
		},
		{
			eleSlice: []int{9, 8, 7, 9, 10, 4, 3, 12},
			maxSlice: []int{9, 9, 9, 9, 10, 10, 10, 12},
		},
	}

	for _, tt := range tests {
		sm := StackMax{}
		sm.Pushs(tt.eleSlice)
		checkMaxSlice(t, tt.maxSlice, sm.maxer)
	}
}

func checkMaxSlice(t *testing.T, expMax []int, retMax []int) {
	expLen, retLen := len(expMax), len(retMax)
	if expLen != retLen {
		t.Fatalf("max slice length expected=%d, got=%d", expLen, retLen)
	}
	for idx, expItem := range expMax {
		retItem := retMax[idx]
		if retItem != expItem {
			t.Errorf("[%d] max value unmatch. expected=%d, got=%d", idx, expItem, retItem)
		}
	}
}
