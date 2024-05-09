package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {

	var dfs func(root *TreeNode, val int) int

	dfs = func(root *TreeNode, val int) int {
		if root == nil {
			return 0
		}
		val = val<<1 | root.Val
		if root.Left == nil && root.Right == nil {
			return val
		}

		return dfs(root.Left, val) + dfs(root.Right, val)
	}

	res := dfs(root, 0)

	return res

}

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	res := sumRootToLeaf(root)

	fmt.Printf("res: %d\n", res)
}
