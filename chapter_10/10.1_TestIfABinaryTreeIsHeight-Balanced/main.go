package main

import (
	"math"
)

//TreeNode struct represent a tree node
type TreeNode struct {
	value int
	left  *TreeNode
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
	} else {
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
//Space complexity: O(N) because each recursive function call
//allocate O(1) space onto stack under the Worst-Case
func IsBalancedTreeBruteForce(root *TreeNode) bool {
	if root == nil {
		//A empty tree is a balanced tree
		return true
	}

	leftHeight := GetTreeHeight(root.left)
	rightHeight := GetTreeHeight(root.right)
	if math.Abs(float64(leftHeight-rightHeight)) > 1.0 {
		//if the difference of left subtree height and right
		//subtree height is large than 1, this tree is not balanced
		return false
	}

	//if this tree is balanced, it also need to check  if its subtree is both balanced
	return IsBalancedTreeBruteForce(root.left) && IsBalancedTreeBruteForce(root.right)
}

//HeightBalanced struct represent a tree's height and balance
type HeightBalanced struct {
	height   int
	balanced bool
}

//IsBalancedTreeCombine function improve the time complexity of IsBalancedTreeBruteForce
//function by combing the height-calculation recursion and the balance-calculation recursion
//since they both use postorder recursion. But the returned height is not the real height of the
//tree when the tree is not balanced
func IsBalancedTreeCombine(root *TreeNode) HeightBalanced {
	if root == nil {
		return HeightBalanced{
			height:   -1,
			balanced: true,
		}
	}

	leftRet := IsBalancedTreeCombine(root.left)
	if !leftRet.balanced {
		return HeightBalanced{
			//left subtree of root tree is not balanced. no need to check right subtree.
			//so we don't know the height of right subtree. Reset root tree height to zero
			height:   0,
			balanced: false, //cause all the function call return
		}
	}

	rightRet := IsBalancedTreeCombine(root.right)
	if !rightRet.balanced {
		return HeightBalanced{
			height:   0,
			balanced: false, //cause all the function call return
		}
	}

	//left and right subtree are both balanced, check if root tree is balanced
	//and update the height of root tree
	ret := HeightBalanced{
		height:   maxAddOne(leftRet.height, rightRet.height),
		balanced: math.Abs(float64(leftRet.height-rightRet.height)) <= 1.0,
	}
	return ret
}

func maxAddOne(i, j int) int {
	if i > j {
		return i + 1
	}
	return j + 1
}

func main() {

}
