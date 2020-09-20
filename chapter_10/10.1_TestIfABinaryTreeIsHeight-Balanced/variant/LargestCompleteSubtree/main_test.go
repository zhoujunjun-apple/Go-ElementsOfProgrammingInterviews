package main

import (
	"testing"
)

//add new test case at the end of tests slice
var tests = []struct {
	root           *TreeNode //the root node of test tree
	expSize        int       //the expected size of largest subtree
	isRootComplete bool      //is the test tree a complete tree?
}{
	{ //0, index value of tests slice
		root:           nil,
		expSize:        0,
		isRootComplete: true,
	},
	{ //1
		root:           &TreeNode{},
		expSize:        1,
		isRootComplete: true,
	},
	{ //2
		root: &TreeNode{
			left: &TreeNode{},
		},
		expSize:        2,
		isRootComplete: true,
	},
	{ //3
		root: &TreeNode{
			right: &TreeNode{},
		},
		expSize:        1,
		isRootComplete: false,
	},
	{ //4
		root: &TreeNode{
			left:  &TreeNode{},
			right: &TreeNode{},
		},
		expSize:        3,
		isRootComplete: true,
	},
	{ //5
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{},
			},
		},
		expSize:        2,
		isRootComplete: false,
	},
	{ //6
		root: &TreeNode{
			left: &TreeNode{
				right: &TreeNode{},
			},
		},
		expSize:        1,
		isRootComplete: false,
	},
	{ //7
		root: &TreeNode{
			right: &TreeNode{
				right: &TreeNode{},
			},
		},
		expSize:        1,
		isRootComplete: false,
	},
	{ //8
		root: &TreeNode{
			right: &TreeNode{
				left: &TreeNode{},
			},
		},
		expSize:        2,
		isRootComplete: false,
	},
	{ //9
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{},
			},
			right: &TreeNode{},
		},
		expSize:        4,
		isRootComplete: true,
	},
	{ //10
		root: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{
					left: &TreeNode{},
				},
			},
		},
		expSize:        2,
		isRootComplete: false,
	},
	{ //11
		root: &TreeNode{
			right: &TreeNode{
				right: &TreeNode{
					right: &TreeNode{},
				},
			},
		},
		expSize:        1,
		isRootComplete: false,
	},
	{ //12
		root: &TreeNode{
			left: &TreeNode{
				right: &TreeNode{},
			},
			right: &TreeNode{},
		},
		expSize:        1,
		isRootComplete: false,
	},
	{ //13
		root: &TreeNode{
			left: &TreeNode{},
			right: &TreeNode{
				left: &TreeNode{},
			},
		},
		expSize:        2,
		isRootComplete: false,
	},
	{ //14
		root: &TreeNode{
			left: &TreeNode{},
			right: &TreeNode{
				right: &TreeNode{},
			},
		},
		expSize:        1,
		isRootComplete: false,
	},
	{ //15
		root: &TreeNode{
			left: &TreeNode{
				left:  &TreeNode{},
				right: &TreeNode{},
			},
		},
		expSize:        3,
		isRootComplete: false,
	},
	{ //16
		root: &TreeNode{
			right: &TreeNode{
				left:  &TreeNode{},
				right: &TreeNode{},
			},
		},
		expSize:        3,
		isRootComplete: false,
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
		expSize:        6,
		isRootComplete: false,
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
		expSize:        5,
		isRootComplete: false,
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
		expSize:        2,
		isRootComplete: false,
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
		expSize:        6,
		isRootComplete: true,
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
		expSize:        7,
		isRootComplete: true,
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
		expSize:        3,
		isRootComplete: false,
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

func Test_IsTreeComplete(t *testing.T) {
	for idx, tt := range tests {
		ret := IsTreeComplete(tt.root)
		if ret != tt.isRootComplete {
			t.Errorf("[%d] expected=%t, got=%t", idx, tt.isRootComplete, ret)
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
