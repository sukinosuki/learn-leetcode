package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {

	depth := findDepth(root, x, 0)
	fmt.Printf("depth: %d", depth)

	return false
}

func findDepth(tree *TreeNode, value int, deep int) int {
	if tree == nil {
		return 0
	}
	if tree.Val == value {
		return 1
	}

	leftDepth := findDepth(tree.Left, value, deep+1)
	rightDepth := findDepth(tree.Right, value, deep+1)

	if leftDepth == 0 {
		return rightDepth + 1
	}

	return leftDepth + 1
}

func main() {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	isCousins(root, 3, 4)
}
