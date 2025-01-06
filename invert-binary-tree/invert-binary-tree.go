package invertbinarytree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(input []int) *TreeNode {
	if len(input) < 3 {
		return nil
	}
	root := &TreeNode{
		Val: input[0],
	}
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	panic("implement me.")
}
