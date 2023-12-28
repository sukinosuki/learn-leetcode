package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	//maxDepth := max(max(depth(root), depth(root.Left)-1), depth(root.Right)-1)

	maxDepth := max(max(depth(root.Left)+depth(root.Right), diameterOfBinaryTree(root.Left)), diameterOfBinaryTree(root.Right))
	return maxDepth
}

func depth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func main() {

}
