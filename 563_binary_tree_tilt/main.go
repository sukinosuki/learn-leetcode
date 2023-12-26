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

func findTilt(root *TreeNode) int {

	sum := 0

	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {

		if root == nil {
			return 0
		}
		leftSum := helper(root.Left)
		rightSum := helper(root.Right)

		sum += int(math.Abs(float64(leftSum - rightSum)))

		return leftSum + rightSum + root.Val
	}

	helper(root)

	return sum
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	tilt := findTilt(root)

	fmt.Printf("tilt: %v\n", tilt)
}
