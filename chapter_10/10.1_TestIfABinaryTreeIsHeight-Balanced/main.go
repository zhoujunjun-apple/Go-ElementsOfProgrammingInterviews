package main

//TreeNode struct represent a tree node
type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

//GetTreeHeight function return the height of root tree
func GetTreeHeight(root *TreeNode) int {
	if root == nil {
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
func IsBalancedTreeBruteForce(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return true
}


func main(){

}
