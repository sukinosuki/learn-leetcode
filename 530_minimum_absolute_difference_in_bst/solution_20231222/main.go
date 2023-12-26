package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	var helper func(node *TreeNode)
	var prev *TreeNode
	minSub := math.MaxInt
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		helper(node.Left)
		if prev != nil && node.Val-prev.Val < minSub {
			minSub = node.Val - prev.Val
		}
		prev = node
		helper(node.Right)
	}

	helper(root)

	return minSub
}

func main() {
	//root := &TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 1,
	//		},
	//		Right: &TreeNode{
	//			Val: 3,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 6,
	//	},
	//}
	root := &TreeNode{
		Val: 236,
		Left: &TreeNode{
			Val: 104,
			Right: &TreeNode{
				Val: 227,
			},
		},
		Right: &TreeNode{
			Val: 701,
			Right: &TreeNode{
				Val: 911,
			},
		},
	}

	sub := getMinimumDifference(root)

	fmt.Printf("sub: %v\n", sub)
}
