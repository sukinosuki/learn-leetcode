package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		helper(node.Left)
		helper(node.Right)
		ans = append(ans, node.Val)
	}

	helper(root)

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
