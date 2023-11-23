package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	helper(root.Left)

}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		return root.Left.Val + root.Val
	}
	if root.Right != nil {
		return root.Right.Val + root.Val
	}
	return 0
}
func main() {

}
