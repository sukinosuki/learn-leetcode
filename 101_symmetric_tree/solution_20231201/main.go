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

	queue1, queue2 := []*TreeNode{root.Left}, []*TreeNode{root.Right}

	for len(queue1) > 0 && len(queue2) > 0 {

		node1, node2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		if node1 == nil && node2 == nil {
			continue
		}
		// 等价(因为上面已经判断过node1、node2全部为nil时的情况了): if node == nil || node2 == nil {
		if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		queue1 = append(queue1, node1.Left)
		queue2 = append(queue2, node2.Right)
		queue1 = append(queue1, node1.Right)
		queue2 = append(queue2, node2.Left)
	}

	return len(queue1) == 0 && len(queue2) == 0
}

func main() {

}
