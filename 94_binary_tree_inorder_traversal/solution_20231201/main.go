package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历: 左-根-右
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 节点入栈
			stack = append(stack, root)
			// 遍历当前节点所有的左树
			root = root.Left
		}

		// 当前节点和该节点下的所有左树入栈后，开始出栈
		root = stack[len(stack)-1]
		// 更新栈
		stack = stack[:len(stack)-1]

		// 追加当前节点值
		res = append(res, root.Val)

		// 开始遍历出栈节点的右树
		root = root.Right
	}

	return res
}

func main() {

}
