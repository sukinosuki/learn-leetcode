package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {

	ans := []int{}
	if root == nil {
		return ans
	}

	// 先进后出
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if node.Left == nil && node.Right == nil {
			stack = stack[:len(stack)-1]
			ans = append(ans, node.Val)
			continue
		}

		// 先添加右节点
		if node.Right != nil {
			stack = append(stack, node.Right)
			node.Right = nil
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil
		}
	}

	return ans
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
