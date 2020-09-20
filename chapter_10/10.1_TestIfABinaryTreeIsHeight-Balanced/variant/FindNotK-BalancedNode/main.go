package main

//TreeNode is a tree node
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
}

//NodeInfo is the recored information in FindNonKBalancedNode
type NodeInfo struct {
	node *TreeNode
}

//FindNonKBalancedNode funtion returns a node in root tree that
//not k-balanced, but all of its descendants are k-balanced.
func FindNonKBalancedNode(root *TreeNode, k int) NodeInfo {
	if root == nil {
		return NodeInfo{
			node: nil,
		}
	}
	return NodeInfo{}
}

func main() {

}
