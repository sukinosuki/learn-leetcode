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

	stack := []*TreeNode{root}
	subStack := []int{targetSum}

	for len(stack) > 0 {
		size := len(stack)
		for size > 0 {
			node := stack[0]
			stack = stack[1:]
			subTarget := subStack[0]
			subStack = subStack[1:]

			if node.Left == nil && node.Right == nil && subTarget == node.Val {
				return true
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
				subStack = append(subStack, subTarget-node.Val)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
				subStack = append(subStack, subTarget-node.Val)
			}
			size--
		}
	}

	return false
}

func main() {

}
