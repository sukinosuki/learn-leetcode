package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {

		if node == nil {
			return
		}
		fmt.Printf("LL: %d\n", node.Val)

		inorder(node.Left)
		res = append(res, node.Val)
		fmt.Printf("APPEND: %d\n", node.Val)
		inorder(node.Right)
		fmt.Printf("RR: %d\n", node.Val)
	}

	inorder(root)
	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			//Right: &TreeNode{
			//	Val: 5,
			//},
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 7,
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	res := inorderTraversal(root)

	fmt.Printf("res %v\n", res)
}
