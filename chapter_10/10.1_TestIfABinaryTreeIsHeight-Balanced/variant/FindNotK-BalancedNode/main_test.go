package main

import "testing"

var tests = []struct {
	root    *TreeNode //the root node of test tree
	kValue  int       //the wanted k-banlanced degree
	expNode *TreeNode //the expected node pointer
}{
	{
		root:    nil,
		kValue:  1,
		expNode: nil,
	},
}

func Test_FindNonKBalancedNode(t *testing.T) {
	for idx, tt := range tests {
		ret := FindNonKBalancedNode(tt.root, tt.kValue)
		if ret.node != tt.expNode {
			t.Errorf("[%d] expected=%v, got=%v", idx, tt.expNode, ret.node)
		}
	}
}
