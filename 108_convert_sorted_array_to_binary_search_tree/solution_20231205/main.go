package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	nodes := helper(nums, 0, len(nums)-1)
	return nodes
}

func helper(nums []int, l, r int) *TreeNode {

	if l > r {
		return nil
	}
	mid := (l + r) / 2
	node := &TreeNode{
		Val: nums[mid],
	}
	left := helper(nums, l, mid-1)
	node.Left = left
	right := helper(nums, mid+1, r)
	node.Right = right

	return node
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	nodes := sortedArrayToBST(nums)

	fmt.Printf("nodes: %v\n", nodes)
}
