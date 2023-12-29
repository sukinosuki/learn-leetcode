package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	m := make(map[int]int)
	countMap := make(map[int]int)
	var helper func(root *TreeNode, level int)
	helper = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		m[level] += root.Val
		countMap[level]++

		helper(root.Left, level+1)
		helper(root.Right, level+1)
	}

	helper(root, 0)
	ans := make([]float64, len(m))
	for k, v := range m {
		count := countMap[k]
		ans[k] = float64(v) / float64(count)

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
