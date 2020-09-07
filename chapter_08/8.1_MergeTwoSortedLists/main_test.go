package main

import (
	"fmt"
	"testing"
)

func Test_NativeMerge(t *testing.T) {
	tests := []struct {
		s1          []int
		s2          []int
		expectedLen int
	}{
		{
			s1:          []int{5, 4, 3, 2, 1},
			s2:          []int{9, 6},
			expectedLen: 7,
		},
		{
			s1:          []int{8, 4, 2, 1},
			s2:          []int{},
			expectedLen: 4,
		},
		{
			s1:          []int{9, 7, 5, 3, 1},
			s2:          []int{8, 6, 4, 2, 0},
			expectedLen: 10,
		},
	}

	for _, tt := range tests {
		left := SingleList{}
		right := SingleList{}

		anums, _ := left.Add(tt.s1)
		fmt.Printf("left: %v\n", left.GetValues())
		if anums != len(tt.s1) {
			t.Fatalf("left length expected=%d, got=%d", len(tt.s1), anums)
		}

		anums, _ = right.Add(tt.s2)
		fmt.Printf("right:%v\n", right.GetValues())
		if anums != len(tt.s2) {
			t.Fatalf("right length expected=%d, got=%d", len(tt.s2), anums)
		}

		ret := NativeMerge(&left, &right)
		fmt.Printf("ret: %v\n", ret.GetValues())
		checkResult(t, ret, tt.expectedLen)
	}
}

func checkResult(t *testing.T, ret *SingleList, expLen int) {
	if expLen != ret.Length {
		t.Fatalf("single list length expected=%d, got=%d", expLen, ret.Length)
	}

	header := ret.Header
	for header != nil {
		front := header.Value
		header = header.Next

		if header != nil {
			back := header.Value
			if front > back {
				t.Errorf("front node value should not large than back node. got: %d > %d. %v", front, back, ret.GetValues())
			}
		} else {
			break
		}
	}
}
