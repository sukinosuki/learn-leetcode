package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}
	root.Val = targetSum - root.Val

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		size := len(stack)
		for size > 0 {

			node := stack[0]
			stack = stack[1:]
			if node.Left == nil && node.Right == nil && node.Val == 0 {
				return true
			}

			if node.Left != nil {
				node.Left.Val = node.Val - node.Left.Val
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val - node.Right.Val
				stack = append(stack, node.Right)
			}
			size--
		}
	}
	return false
}

func main() {

}
