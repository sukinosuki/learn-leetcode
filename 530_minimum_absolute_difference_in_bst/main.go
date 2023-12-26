package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	if root == nil {
		return 0
	}
	arr := []int{}
	var helper func(node *TreeNode)

	helper = func(node *TreeNode) {

		if node == nil {
			return
		}
		arr = append(arr, node.Val)
		helper(node.Left)
		helper(node.Right)
	}
	helper(root)

	if len(arr) < 2 {
		return 0
	}
	sort.Ints(arr)
	minSub := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		minSub = min(minSub, arr[i]-arr[i-1])
	}
	return minSub
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}

	return n1
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	sub := getMinimumDifference(root)

	fmt.Printf("sub: %v\n", sub)
}
