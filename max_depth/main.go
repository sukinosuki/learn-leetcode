package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	func maxDepth(root *TreeNode) int {
//		if root == nil {
//			return 0
//		}
//
//		depth := findDepth(root, 0)
//
//		return depth
//	}
//
//	func findDepth(root *TreeNode, l int) int {
//		if root == nil {
//			return l
//		}
//
//		leftDepth := findDepth(root.Left, l+1)
//		rightDepth := findDepth(root.Right, l+1)
//
//		if leftDepth > rightDepth {
//			return leftDepth
//		}
//
//		return rightDepth
//	}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return 1 + leftDepth
	}
	return 1 + rightDepth
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	depth := maxDepth(root)
	fmt.Println("depth ", depth)
}
