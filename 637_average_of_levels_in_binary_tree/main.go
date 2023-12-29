package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var ans []float64
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		nodes := stack[:]
		stack = nil

		var sum int
		for _, node := range nodes {
			sum += node.Val
			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}

		ans = append(ans, float64(sum)/float64(len(nodes)))
	}

	return ans
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	avg := averageOfLevels(root)
	fmt.Printf("avg: %v\n", avg)
}
