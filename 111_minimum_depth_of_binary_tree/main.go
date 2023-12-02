package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	ans := 1
	for len(stack) > 0 {

		size := len(stack)
		for size > 0 {
			node := stack[0]
			stack = stack[1:]

			if node.Left == nil && node.Right == nil {
				return ans
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			size--
		}
		ans++
	}

	return ans
}

func main() {
	//node := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 3,
	//			Left: &TreeNode{
	//				Val: 4,
	//				Left: &TreeNode{
	//					Val: 5,
	//					Left: &TreeNode{
	//						Val: 6,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	node := &TreeNode{
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
	}

	res := minDepth(node)
	fmt.Printf("res %d\n", res)
}
