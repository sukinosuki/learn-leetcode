package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {

	stack := []*TreeNode{}
	res := []int{}
	var prev *TreeNode

	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 取最后一个
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}

	return res
}

func main() {
	//  [3 4 2 6 5 1]
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

	ans := postorderTraversal(root)

	fmt.Printf("ans: %v\n", ans)
}
