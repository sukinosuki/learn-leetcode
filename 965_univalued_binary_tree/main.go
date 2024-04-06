package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {

	value := root.Val

	var helper func(tree *TreeNode) bool
	helper = func(tree *TreeNode) bool {
		if tree == nil {
			return true
		}

		if tree.Val != value {
			return false
		}

		return helper(tree.Left) && helper(tree.Right)
	}

	return helper(root)
}

func main() {

}
