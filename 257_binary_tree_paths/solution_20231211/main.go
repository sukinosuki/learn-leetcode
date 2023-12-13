package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {

	ans := make([]string, 0)

	var helper func(node *TreeNode, path string)
	helper = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		pathB := path
		pathB += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			ans = append(ans, pathB)
			return
		}

		pathB += "->"
		helper(node.Left, pathB)
		helper(node.Right, pathB)
	}

	helper(root, "")
	return ans
}

func main() {

}
