package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root.Left, root.Right)
}

func check(l *TreeNode, r *TreeNode) bool {

	if l == nil && r == nil {
		return true
	}

	// 等价(因为上面已经判断过node1、node2全部为nil时的情况了): if (l == nil && r != nil) || (l != nil && r == nil) {
	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return check(l.Left, r.Right) && check(l.Right, r.Left)
}

func main() {

}
