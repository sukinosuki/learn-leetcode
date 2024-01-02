package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) (ans bool) {

	defer func(ans *bool) {
		recover()
		//if err != nil {
		//	*ans = true
		//}
	}(&ans)

	m := make(map[int]int)

	var helper func(root *TreeNode) bool
	helper = func(root *TreeNode) bool {

		if root == nil {
			return false
		}

		_, ok := m[root.Val]
		if ok {
			ans = true
			panic(errors.New("2"))
		}

		m[k-root.Val] = root.Val

		return ok || helper(root.Left) || helper(root.Right)
	}

	helper(root)

	return
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	}

	res := findTarget(root, 9)
	fmt.Printf("res: %v \n", res)
}
