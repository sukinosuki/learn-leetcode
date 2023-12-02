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

	//1.左孩子和有孩子都为空的情况，说明到达了叶子节点，直接返回1即可
	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	//这里其中一个节点为空，说明m1和m2有一个必然为0，所以可以返回m1 + m2 + 1;
	if root.Left == nil || root.Right == nil {
		return left + right + 1
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
