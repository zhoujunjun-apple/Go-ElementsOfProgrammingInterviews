package main

import (
	"testing"
)

var tests = []struct{
	root *TreeNode
	expected bool
}{
	{
		root: nil,
		expected: false,
	},
	{
		root: &TreeNode{
			value: 0,
		},
		expected: true,
	},
	{
		root: &TreeNode{
			value: 0,
			left: &TreeNode{
				value: 1,
			},
		},
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
		expected: false,
	},
}

func Test_IsBalancedTreeBruteForce(t *testing.T) {
	for idx, tt := range tests {
		ret := IsBalancedTreeBruteForce(tt.root)
		if ret != tt.expected {
			t.Errorf("[%d] expected=%t, got=%t", idx, tt.expected, ret)
		}
	}
}