package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {

	ans := make([]string, 0)

	var helper func(root *TreeNode, path string)

	helper = func(root *TreeNode, path string) {
		pathB := path
		if root == nil {
			return
		}
		pathB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			ans = append(ans, pathB)
			return
		}
		pathB += "->"
		helper(root.Left, pathB)
		helper(root.Right, pathB)

	}

	helper(root, "")

	return ans
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	res := binaryTreePaths(root)
	fmt.Printf("%v\n", res)
}
