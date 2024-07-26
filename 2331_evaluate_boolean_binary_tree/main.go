package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {

	var helper func(node *TreeNode) bool

	helper = func(node *TreeNode) bool {
		if node.Left == nil {

			return node.Val == 1
		}

		bool1 := helper(node.Left)
		bool2 := helper(node.Right)
		if node.Val == 2 {
			return bool1 || bool2
		}

		return bool1 && bool2
	}

	return helper(root)
}

func main() {

	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	res := evaluateTree(root)

	fmt.Printf("res: %v\n", res)
}
