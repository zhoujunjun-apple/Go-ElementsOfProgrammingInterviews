package main

import (
	"testing"
)

//add new test case at the end of tests slice
var tests = []struct {
	root    *TreeNode
	expSize int
}{
	{ //0, index value of tests slice
		root:    nil,
		expSize: 0,
	},
	{ //1
		root:    &TreeNode{},
		expSize: 1,
	},
	{ //2
		root: &TreeNode{
			left: &TreeNode{},
		},
		expSize: 2,
	},
	{ //3
		root: &TreeNode{
			right: &TreeNode{},
		},
		expSize: 1,
	},
	{ //4
		root: &TreeNode{
			left:  &TreeNode{},
			right: &TreeNode{},
		},
		expSize: 3,
	},
	{ //5
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{},
			},
		},
		expSize: 2,
	},
	{ //6
		root: &TreeNode{
			left: &TreeNode{
				right: &TreeNode{},
			},
		},
		expSize: 1,
	},
	{ //7
		root: &TreeNode{
			right: &TreeNode{
				right: &TreeNode{},
			},
		},
		expSize: 1,
	},
	{ //8
		root: &TreeNode{
			right: &TreeNode{
				left: &TreeNode{},
			},
		},
		expSize: 2,
	},
	{ //9
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{},
			},
			right: &TreeNode{},
		},
		expSize: 4,
	},
	{ //10
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{
					left: &TreeNode{},
				},
			},
		},
		expSize: 2,
	},
	{ //11
		root: &TreeNode{
			right: &TreeNode{
				right: &TreeNode{
					right: &TreeNode{},
				},
			},
		},
		expSize: 1,
	},
	{ //12
		root: &TreeNode{
			left: &TreeNode{
				right: &TreeNode{},
			},
			right: &TreeNode{},
		},
		expSize: 1,
	},
	{ //13
		root: &TreeNode{
			left: &TreeNode{},
			right: &TreeNode{
				left: &TreeNode{},
			},
		},
		expSize: 2,
	},
	{ //14
		root: &TreeNode{
			left: &TreeNode{},
			right: &TreeNode{
				right: &TreeNode{},
			},
		},
		expSize: 1,
	},
	{ //15
		root: &TreeNode{
			left: &TreeNode{
				left:  &TreeNode{},
				right: &TreeNode{},
			},
		},
		expSize: 3,
	},
	{ //16
		root: &TreeNode{
			right: &TreeNode{
				left:  &TreeNode{},
				right: &TreeNode{},
			},
		},
		expSize: 3,
	},
	{ //17
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{
					right: &TreeNode{
						left: &TreeNode{
							left:  &TreeNode{},
							right: &TreeNode{},
						},
						right: &TreeNode{
							left: &TreeNode{},
						},
					},
				},
			},
			right: &TreeNode{},
		},
		expSize: 6,
	},
	{ //18
		root: &TreeNode{
			left: &TreeNode{
				right: &TreeNode{
					left: &TreeNode{
						left: &TreeNode{
							left:  &TreeNode{},
							right: &TreeNode{},
						},
						right: &TreeNode{},
					},
				},
			},
		},
		expSize: 5,
	},
	{ //19
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{},
			},
			right: &TreeNode{
				left: &TreeNode{},
			},
		},
		expSize: 2,
	},
	{ //20
		root: &TreeNode{
			left: &TreeNode{
				left:  &TreeNode{},
				right: &TreeNode{},
			},
			right: &TreeNode{
				left: &TreeNode{},
			},
		},
		expSize: 6,
	},
	{ //21
		root: &TreeNode{
			left: &TreeNode{
				left:  &TreeNode{},
				right: &TreeNode{},
			},
			right: &TreeNode{
				left:  &TreeNode{},
				right: &TreeNode{},
			},
		},
		expSize: 7,
	},
	{ //22
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{
					left:  &TreeNode{},
					right: &TreeNode{},
				},
			},
			right: &TreeNode{},
		},
		expSize: 3,
	},
}

func Test_LargestCompleteSubtree(t *testing.T) {
	for idx, tt := range tests {
		ret := LargestCompleteSubtree(tt.root)
		if ret.subtreeSize != tt.expSize {
			t.Errorf("[%d] largest complete subtree expected=%d, got=%d.",
				idx, tt.expSize, ret.subtreeSize)
		}
	}
}

func Test_PowerInt(t *testing.T) {
	tests := []struct {
		base   int
		power  int
		expRet int
	}{
		{base: 0, power: 1, expRet: 0},
		{base: 1, power: 2, expRet: 1},
		{base: 2, power: 0, expRet: 1},
		{base: 2, power: 1, expRet: 2},
		{base: 2, power: 5, expRet: 32},
	}

	for idx, tt := range tests {
		gotRet := PowerInt(tt.base, tt.power)
		if gotRet != tt.expRet {
			t.Errorf("[%d] expected=%d, got=%d", idx, tt.expRet, gotRet)
		}
	}
}
