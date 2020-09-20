package main

//TreeNode represent a binary tree node
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

//SizeComplete struct record complete subtree info
type SizeComplete struct {
	root        *TreeNode //the root node of the largest complete subtree
	height      int       //height of this tree
	isFull      bool      //whether this tree is full
	isComplete  bool      //whether this tree is complete
	subtreeSize int       //the largest complete subtree size under this tree
}

//LargestCompleteSubtree function find the largest complete subtree in root tree
//and record this subtree information on return
func LargestCompleteSubtree(root *TreeNode) SizeComplete {
	if root == nil {
		return SizeComplete{
			isFull:      true,
			isComplete:  true,
			subtreeSize: 0,
			height:      -1,
			root:        nil,
		}
	}

	leftRet := LargestCompleteSubtree(root.left)
	rightRet := LargestCompleteSubtree(root.right)

	defaultTreeInfo := SizeComplete{
		isFull:      false,
		isComplete:  false,
		root:        GetLargerRoot(leftRet, rightRet),
		height:      MaxInt(leftRet.height, rightRet.height),
		subtreeSize: MaxInt(leftRet.subtreeSize, rightRet.subtreeSize),
	}

	if leftRet.isComplete && rightRet.isComplete {
		//left and right subtree are both complete subtree.
		//try to combine them to get a bigger complete tree.
		if leftRet.subtreeSize < rightRet.subtreeSize {
			//right complete subtree is bigger than left complete subtree.
			//cannot combine into a bigger complete tree
			return defaultTreeInfo
		} else if leftRet.subtreeSize == rightRet.subtreeSize {
			//right complete subtree has the same size with the left complete subtree
			//need to recheck if the left subtree is full
			if leftRet.isFull {
				//left complete subtree is a full tree.
				//left and right subtree can combine into a bigger complete tree
				return SizeComplete{
					root:        root,
					isFull:      leftRet.isFull && rightRet.isFull,
					isComplete:  true,
					height:      leftRet.height + 1,
					subtreeSize: leftRet.subtreeSize + rightRet.subtreeSize + 1,
				}
			} else {
				//left complete subtree is not a full tree.
				//cannot combine right subtree into a bigger complete tree
				return defaultTreeInfo
			}
		} else {
			//right complete subtree is less than left complete subtree
			if leftRet.isFull {
				//left complete subtree is a full tree
				//recheck right subtree's size
				if rightRet.subtreeSize >= int(leftRet.subtreeSize/2) {
					return SizeComplete{
						root:        root,
						isFull:      false,
						isComplete:  true,
						height:      leftRet.height + 1,
						subtreeSize: leftRet.subtreeSize + rightRet.subtreeSize + 1,
					}
				} else {
					return defaultTreeInfo
				}
			} else {
				//left complete subtree is not full
				//can combine right subtree into a bigger tree only when right tree is full
				//and right subtree size is right
				wantRightSize := PowerInt(2, leftRet.height) - 1
				if rightRet.isFull && rightRet.subtreeSize == wantRightSize {
					return SizeComplete{
						root:        root,
						isFull:      false,
						isComplete:  true,
						height:      leftRet.height + 1,
						subtreeSize: leftRet.subtreeSize + rightRet.subtreeSize + 1,
					}
				} else {
					return defaultTreeInfo
				}
			}
		}
	} else {
		//one of subtree isn't a complete tree. only need to record size value of the larger one
		return defaultTreeInfo
	}
}

//IsTreeComplete function check if root tree is a complete tree
//It use LargestCompleteSubtree function to find the largest
//complete subtree in root tree, then compare the subtree and root tree
func IsTreeComplete(root *TreeNode) bool {
	ret := LargestCompleteSubtree(root)
	return ret.root == root
}

//MaxInt function return the maximum value between two integer
func MaxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//PowerInt function return base**power
func PowerInt(base, power int) int {
	ret := 1
	if base == 1 || power <= 0 {
		return ret
	}

	for power > 0 {
		ret *= base
		power--
	}
	return ret
}

//GetLargerRoot function return the root node pointer of larger complete subtree
func GetLargerRoot(left, right SizeComplete) *TreeNode {
	if left.subtreeSize > right.subtreeSize {
		return left.root
	}
	return right.root
}

func main() {

}
