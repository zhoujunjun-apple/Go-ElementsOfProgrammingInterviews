package main

import (
	"testing"
)

var tests = []struct{
	root *TreeNode
	height int
	expected bool
}{
	{
		root: nil,
		expected: false,
		height: -1,
	},
	{
		root: &TreeNode{
			value: 0,
		},
		height: 0,
		expected: true,
	},
	{
		root: &TreeNode{
			value: 0,
			left: &TreeNode{
				value: 1,
			},
		},
		height: 1,
		expected: true,
	},
	{
		root: &TreeNode{
			value: 0,
			left: &TreeNode{
				value: 1,
				left: &TreeNode{
					value: 2,
				},
			},
		},
		height: 2,
		expected: false,
	},
	{
		root: &TreeNode{
			value: 0,
			left: &TreeNode{
				value: 1,
				left: &TreeNode{
					value:2,
				},
			},
			right: &TreeNode{
				value: 3,
			},
		},
		height: 2,
		expected: true,
	},
	{
		root: &TreeNode{
			value: 0,
			left: &TreeNode{
				value: 1,
			},
			right: &TreeNode{
				value: 2,
				left: &TreeNode{
					value: 3,
					left: &TreeNode{
						value: 4,
					},
				},
			},
		},
		height: 3,
		expected: false,
	},
	{
		root: &TreeNode{
			value: 0,
			left: &TreeNode{
				value: 1,
				left: &TreeNode{
					value: 6,
					right: &TreeNode{
						value: 7,
						right: &TreeNode{
							value: 8,
						},
					},
				},
			},
			right: &TreeNode{
				value: 2,
				left: &TreeNode{
					value: 3,
				},
				right: &TreeNode{
					value: 4,
					left: &TreeNode{
						value: 5,
					},
				},
			},
		},
		height: 4,
		expected: false,
	},
}

func Test_GetTreeHeight(t *testing.T) {
	for idx, tt := range tests {
		ret := GetTreeHeight(tt.root)
		if ret != tt.height {
			t.Errorf("[%d] height expected=%d, got=%d", idx, tt.height, ret)
		}
	}
}

func Test_IsBalancedTreeBruteForce(t *testing.T) {
	for idx, tt := range tests {
		ret := IsBalancedTreeBruteForce(tt.root)
		if ret != tt.expected {
			t.Errorf("[%d] expected=%t, got=%t", idx, tt.expected, ret)
		}
	}
}