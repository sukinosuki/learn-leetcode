package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	ans := make([]int, 0)
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		ans = append(ans, node.Val)
		helper(node.Right)
	}

	helper(root)
	return ans
}

func main() {

}
