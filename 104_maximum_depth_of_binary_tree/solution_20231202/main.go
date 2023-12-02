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

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
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
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			//Right: &TreeNode{
			//	Val: 3,
			//	Right: &TreeNode{
			//		Val: 4,
			//		Right: &TreeNode{
			//			Val: 5,
			//			Right: &TreeNode{
			//				Val: 6,
			//				Right: &TreeNode{
			//					Val: 7,
			//					Right: &TreeNode{
			//						Val: 8,
			//					},
			//				},
			//			},
			//		},
			//	},
			//},
		},
	}

	res := maxDepth(root)

	fmt.Printf("res: %d\n", res)
}
