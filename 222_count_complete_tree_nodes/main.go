package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	count := countNodes(root)

	fmt.Printf("count: %v\n", count)
}
