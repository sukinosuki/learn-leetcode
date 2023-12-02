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

	res := helper(root)
	return res + 1
}

func helper(root *TreeNode) int {

	if root == nil {
		return -1
	}

	left := helper(root.Left)
	right := helper(root.Right)

	if left == -1 && right == -1 {
		return 0
	}
	if left == -1 {
		return right + 1
	}
	if right == -1 {
		return left + 1
	}

	if left > right {
		return right + 1
	} else {
		return left + 1
	}
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	//node := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 4,
	//		},
	//		Right: &TreeNode{
	//			Val: 5,
	//		},
	//	},
	//}

	res := minDepth(node)
	fmt.Printf("res %d\n", res)
}
