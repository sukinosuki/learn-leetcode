package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 当前节点和左边最大、右边最小相比
func minDiffInBST(root *TreeNode) int {

	ans := math.MaxInt
	prev := -1
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {

		if root == nil {
			return
		}
		helper(root.Left)
		if prev != -1 && root.Val-prev < ans {
			ans = root.Val - prev
		}
		prev = root.Val
		helper(root.Right)
	}
	helper(root)

	return ans
}

func main() {
	root := &TreeNode{
		Val: 90,
		Left: &TreeNode{
			Val: 69,
			Left: &TreeNode{
				Val: 49,
				Right: &TreeNode{
					Val: 68,
				},
			},
			Right: &TreeNode{
				Val: 88,
			},
		},
	}
	res := minDiffInBST(root)
	fmt.Printf("res: %v\n", res)
}
