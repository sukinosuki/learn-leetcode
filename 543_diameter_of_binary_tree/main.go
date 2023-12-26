package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter := deep(root.Left) + deep(root.Right)

	return diameter
}

func deep(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(deep(root.Left), deep(root.Right)) + 1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
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
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	deep := deep(root.Left)
	fmt.Printf("deep: %v\n", deep)
}
