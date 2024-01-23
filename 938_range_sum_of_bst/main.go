package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {

	sum := 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Val > high {
			helper(root.Left)
		} else if root.Val < low {
			helper(root.Right)
		} else {
			if root.Val <= high && root.Val >= low {
				sum += root.Val
				helper(root.Left)
				helper(root.Right)
			}
		}
	}

	helper(root)

	return sum
}

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 18,
				Left: &TreeNode{
					Val: 19,
				},
			},
		},
	}

	sum := rangeSumBST(root, 7, 15)

	fmt.Printf("sum: %v\n", sum)
}
