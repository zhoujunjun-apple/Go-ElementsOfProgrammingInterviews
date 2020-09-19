package main

import (
	"math"
)

//TreeNode struct represent a tree node
type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

//GetTreeHeight function return the height of root tree. O(N) time and O(1) space
func GetTreeHeight(root *TreeNode) int {
	if root == nil {
		//A empty node has height -1, so a tree with a single node has height 0 
		return -1
	}
	leftHeight := GetTreeHeight(root.left)
	rightHeight := GetTreeHeight(root.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}else{
		return rightHeight + 1
	}
}

//IsBalancedTreeBruteForce function use brute-force check if the root tree is a balanced tree
//Under Average-Case, we assume that left subtree and right subtree have approximately O(N/2) nodes
//time complexity: 
//F(N) = F(GetTreeHeight(left)) + F(GetTreeHeight(right)) + 2*F(N/2)
//= (N-1) + 2*F(N/2)
//= 2*N - (1+2) + (2^2) * F(N/(2^2))
//= 3*N - (1+2+4) + (2^3) * F(N/(2^3))
//= k*N - (1+2+4+...+2^(k-1)) + (2^k) * F(N/(2^k))
//= N*(lgN) - (N) + N
//= N*(lgN)
func IsBalancedTreeBruteForce(root *TreeNode) bool {
	if root == nil {
		//A empty tree is a balanced tree
		return true
	}

	leftHeight := GetTreeHeight(root.left)
	rightHeight := GetTreeHeight(root.right)
	if math.Abs(float64(leftHeight - rightHeight)) > 1.0 {
		//if the difference of left subtree height and right 
		//subtree height is large than 1, this tree is not balanced
		return false
	}

	//if this tree is balanced, it also need to check  if its subtree is both balanced
	return IsBalancedTreeBruteForce(root.left) && IsBalancedTreeBruteForce(root.right)
}


func main(){

}
