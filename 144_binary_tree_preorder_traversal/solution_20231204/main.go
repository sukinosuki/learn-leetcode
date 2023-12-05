package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	rightStack := []*TreeNode{root}

	ans := []int{}
	for len(rightStack) > 0 {
		node := rightStack[len(rightStack)-1]
		rightStack = rightStack[:len(rightStack)-1]

		ans = append(ans, node.Val)

		if node.Right != nil {
			rightStack = append(rightStack, node.Right)
		}

		cur := node
		for cur.Left != nil {
			cur = cur.Left
			if cur.Right != nil {
				rightStack = append(rightStack, cur.Right)
			}
			ans = append(ans, cur.Val)
		}
	}

	return ans
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
