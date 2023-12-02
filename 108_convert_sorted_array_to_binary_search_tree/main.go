package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	return helper(nums, 0, len(nums)-1)
}

// 1-> 0:4
// 2-> 0:1
// 3-> 0:-1

func helper(nums []int, left, right int) *TreeNode {

	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	nodes := sortedArrayToBST(nums)

	fmt.Printf("nodes: %v\n", nodes)
}
