package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil {
		return q == nil
	}
	if q == nil {
		return p == nil
	}

	valid := isSameTree(p.Left, q.Left)
	if !valid || p.Val != q.Val {
		return false
	}
	valid = isSameTree(p.Right, q.Right)
	if !valid || p.Val != q.Val {
		return false
	}

	return true
}

func check(p *TreeNode, q *TreeNode) bool {

	if p == nil {
		return q == nil
	}
	if q == nil {
		return p == nil
	}

	valid := check(p.Left, q.Left)
	if !valid || p.Val != q.Val {
		return false
	}
	valid = check(p.Right, q.Right)
	if !valid || p.Val != q.Val {
		return false
	}

	return true
}

func main() {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	res := isSameTree(p, q)

	fmt.Printf("res: %s \n", res)
}
