package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {

	if root == nil || root.Left == nil || root.Right == nil {
		return -1
	}

	secondMin := -1

	var helper func(node *TreeNode)

	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val != root.Val {
			if secondMin == -1 || node.Val < secondMin {
				secondMin = node.Val
			}
		}

		helper(node.Left)
		helper(node.Right)
	}

	helper(root)

	return secondMin
}

func main() {
	//root := &TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val: 2,
	//	},
	//	Right: &TreeNode{
	//		Val: 5,
	//		Left: &TreeNode{
	//			Val: 5,
	//		},
	//		Right: &TreeNode{
	//			Val: 7,
	//		},
	//	},
	//}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	res := findSecondMinimumValue(root)
	fmt.Printf("res: %v\n", res)
}
