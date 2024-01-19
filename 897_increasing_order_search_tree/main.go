package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {

	ans := &TreeNode{
		Right: nil,
	}
	curr := ans

	var helper func(root *TreeNode)

	helper = func(root *TreeNode) {

		if root == nil {
			return
		}

		helper(root.Left)
		curr.Right = root
		root.Left = nil
		curr = curr.Right

		helper(root.Right)
	}

	helper(root)

	return ans.Right
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
	}
	res := increasingBST(root)
	fmt.Printf("res: %v\n", res)
}
