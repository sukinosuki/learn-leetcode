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

// 当前节点和左边最大、右边最小相比
func minDiffInBST(root *TreeNode) int {

	if root == nil || root.Left == nil && root.Right == nil {
		return math.MaxInt
	}

	_min := math.MaxInt
	if root.Left != nil {
		_min = min(_min, int(math.Abs(float64(root.Val-findMax(root.Left)))))
	}
	if root.Right != nil {
		_min = min(_min, int(math.Abs(float64(root.Val-findMin(root.Right)))))
	}

	leftMin := min(minDiffInBST(root.Left), _min)
	rightMin := min(minDiffInBST(root.Right), _min)

	return min(leftMin, rightMin)
}

func findMin(root *TreeNode) int {
	if root == nil {
		return math.MaxInt
	}

	left := findMin(root.Left)
	right := findMin(root.Right)

	return min(min(left, right), root.Val)
}

func findMax(root *TreeNode) int {
	if root == nil {
		return -1
	}

	left := findMax(root.Left)
	right := findMax(root.Right)

	return max(max(left, right), root.Val)
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func main() {
	//root := &TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val: 1,
	//	},
	//	Right: &TreeNode{
	//		Val: 48,
	//		Left: &TreeNode{
	//			Val: 12,
	//		},
	//		Right: &TreeNode{
	//			Val: 51,
	//		},
	//	},
	//}

	root := &TreeNode{
		Val: 90,
		Left: &TreeNode{
			Val: 69,
			Left: &TreeNode{
				Val: 49,
				Right: &TreeNode{
					Val: 68,
				},
			},
			Right: &TreeNode{
				Val: 88,
			},
		},
	}
	res := minDiffInBST(root)
	fmt.Printf("res: %v\n", res)
}
