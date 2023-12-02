package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	valid1 := hasPathSum(root.Left, targetSum-root.Val)
	valid2 := hasPathSum(root.Right, targetSum-root.Val)

	return valid1 || valid2
}

func main() {

}
