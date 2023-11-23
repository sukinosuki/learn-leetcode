package main

//给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
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

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 以升序数组的中间元素作为根节点 root。
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	// 递归的构建 root 的左子树与右子树。
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	res := helper(nums, 0, len(nums)-1)

	return res
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}

	sortedArrayToBST(nums)
}
