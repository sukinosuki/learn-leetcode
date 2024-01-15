package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	var stack = []int{}
	var helper func(root1 *TreeNode)

	helper = func(root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			stack = append(stack, root.Val)
			return
		}

		if root.Left != nil {
			helper(root.Left)
		}

		if root.Right != nil {
			helper(root.Right)
		}
	}

	helper(root1)

	var helper2 func(root *TreeNode) bool

	helper2 = func(root *TreeNode) bool {
		if root.Left == nil && root.Right == nil {
			if len(stack) == 0 {
				return false
			}
			if stack[0] != root.Val {
				return false
			}

			stack = stack[1:]
			return true
		}
		left := true
		right := true

		if root.Left != nil {
			left = helper2(root.Left)
		}
		if root.Right != nil {
			right = helper2(root.Right)
		}
		return left && right
	}

	res := helper2(root2)

	return res && len(stack) == 0
}

func main() {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	root1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	root2 = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
	}

	res := leafSimilar(root1, root2)
	fmt.Printf("res: %v\n", res)
}
