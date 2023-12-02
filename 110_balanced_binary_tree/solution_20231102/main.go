package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	depthSub, valid := helper(root, true)

	fmt.Printf("depthSub: %v, sub: %d.\n", depthSub, valid)
	return valid
}

func helper(root *TreeNode, valid bool) (int, bool) {
	if root == nil {
		return 0, valid
	}
	if !valid {
		return 0, false
	}

	leftDepth, valid1 := helper(root.Left, valid)
	rightDepth, valid2 := helper(root.Right, valid)
	if !valid1 || !valid2 {
		return 0, false
	}

	if leftDepth > rightDepth {
		return leftDepth + 1, leftDepth-rightDepth <= 1
	} else {
		return rightDepth + 1, rightDepth-leftDepth <= 1
	}
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
				},
				//Right: &TreeNode{
				//	Val: 4,
				//},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	valid := isBalanced(root)

	fmt.Printf("valid: %v\n", valid)
}
