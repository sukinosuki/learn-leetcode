package main

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q != nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}
	rightFlag := isSameTree(p.Right, q.Right)
	if rightFlag == false {
		return false
	}
	leftFlag := isSameTree(p.Left, q.Left)

	return leftFlag
}

func main() {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	isSameTree(p, q)
}
