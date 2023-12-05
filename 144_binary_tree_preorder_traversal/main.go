package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {

	ans := make([]int, 0)

	ans = helper(root, ans)

	return ans
}

func helper(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)

	ans = helper(root.Left, ans)
	ans = helper(root.Right, ans)

	return ans
}

func max(num1, num2 int) {
	if num1 > num2 {
		return
	}
}

func main() {

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

	ans := preorderTraversal(root)

	fmt.Printf("ans: %v\n", ans)
}
