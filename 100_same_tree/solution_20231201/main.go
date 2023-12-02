package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true

	}
	if p == nil || q == nil {
		return false
	}

	queue1, queue2 := []*TreeNode{p}, []*TreeNode{}
	for len(queue1) > 0 && len(queue2) > 0 {
		node1, node2 := queue1[0], queue2[0]

		queue1, queue2 = queue1[1:], queue2[1:]
		// 比较两个节点的值，如果两个节点的值不相同则两个二叉树一定不同；
		if node1.Val != node2.Val {
			return false
		}

		// 如果两个节点的值相同，则判断两个节点的子节点是否为空，如果只有一个节点的左子节点为空，或者只有一个节点的右子节点为空，则两个二叉树的结构不同，因此两个二叉树一定不同；
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
			return false
		}
		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}
		if left1 != nil {
			queue1 = append(queue1, left1)
		}
		if right1 != nil {
			queue1 = append(queue1, right1)
		}
		if left2 != nil {
			queue2 = append(queue2, left2)
		}
		if right2 != nil {
			queue2 = append(queue2, right2)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0

}

func main() {

}
