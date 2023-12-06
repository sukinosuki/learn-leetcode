package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if leftHeight == rightHeight {
		return (1 << leftHeight) + countNodes(root.Right)
	} else {
		return (1 << (leftHeight - 1)) + countNodes(root.Left)
	}
}

func getHeight(root *TreeNode) int {

	height := 0
	for root != nil {
		height++
		root = root.Left
	}
	return height
}

//func getNum(node *TreeNode) int {
//	if node == nil {
//		return 0
//	}
//	left := node.Left
//	right := node.Right
//	leftHeight, rightHeight := 0, 0
//	for left != nil {
//		leftHeight++
//		left = left.Left
//	}
//	for right != nil {
//		rightHeight++
//		right = right.Right
//	}
//	if leftHeight == rightHeight {
//		return (2 << leftHeight) - 1
//	}
//	return 0
//}

func main() {

}
