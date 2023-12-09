package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}

	constructPaths(root, "")

	return paths
}

func constructPaths(root *TreeNode, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			constructPaths(root.Left, pathSB)
			constructPaths(root.Right, pathSB)
		}
	}
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

	binaryTreePaths(root)
}
