package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		//nodes := stack[:]
		//stack = stack[0:0]
		size := len(stack)

		for size > 0 {
			node := stack[0]
			stack = stack[1:]

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			size--
		}
		max++
	}

	return max
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
	res := maxDepth(root)

	fmt.Printf("res: %d\n", res)
}
