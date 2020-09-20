package main

//TreeNode is a tree node
type TreeNode struct {
	name  string
	left  *TreeNode
	right *TreeNode
}

//NodeInfo is the recored information in FindNonKBalancedNode
type NodeInfo struct {
	node *TreeNode
	size int
}

//FindNonKBalancedNode funtion returns a node in root tree that
//not k-balanced, but all of its descendants are k-balanced.
//O(N) time and O(height of tree) space
func FindNonKBalancedNode(root *TreeNode, k int) NodeInfo {
	if root == nil {
		return NodeInfo{}
	}
	leftRet := FindNonKBalancedNode(root.left, k)
	if leftRet.node != nil {
		return leftRet
	}

	rightRet := FindNonKBalancedNode(root.right, k)
	if rightRet.node != nil {
		return rightRet
	}

	newSize := leftRet.size + rightRet.size + 1
	if AbsInt(leftRet.size, rightRet.size) > k {
		return NodeInfo{
			node: root,
			size: newSize,
		}
	} else {
		return NodeInfo{
			node: nil,
			size: newSize,
		}
	}

}

func AbsInt(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func main() {

}
