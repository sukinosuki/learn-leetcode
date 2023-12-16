package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {

		for root.Left != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, node.Val)
		if node.Right != nil {
			root = node.Right
			stack = append(stack, root)
		}
	}

	return ans
}

//[3 4 5 2 1 7 8]

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
