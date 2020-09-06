package main

import (
	"testing"
)

func Test_ThreeBucket(t *testing.T) {
	tests := []struct {
		slice []int //array to be arranged
		pidx  int   //pivot index
		npidx int   //expected new pivot index
	}{
		{
			slice: []int{3, 5, 6, 1, 2, 7, 9, 0},
			pidx:  1,
			npidx: 4,
		},
		{
			slice: []int{2,4,1,6,3,9,6,5,0,7},
			pidx: 2,
			npidx: 1,
		},
		{
			slice: []int{4,3,2,1,0},
			pidx: 4,
			npidx: 0,
		},
	}

	for _, tt := range tests {
		retIdx := ThreeBucket(tt.slice, tt.pidx)
		checkResult(t, tt.slice, tt.npidx, retIdx)
	}
}

func checkResult(t *testing.T, tts []int, ttn, ridx int) {
	if ttn != ridx {
		t.Errorf("new pivot index: expected=%d, got=%d", ttn, ridx)
	}
	pivot := tts[ttn]
	for i, item := range tts {
		if i < ttn {
			if item >= pivot {
				t.Errorf("items before index %d should less than %d, but got item %d at index %d", ttn, pivot, item, i)
			}
		} else {
			if item < pivot {
				t.Errorf("items after index %d should larger than %d, but got item %d at index %d", ttn, pivot, item, i)
			}
		}
	}
}
