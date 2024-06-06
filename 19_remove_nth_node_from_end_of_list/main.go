package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var helper func(node *ListNode) (*ListNode, int)
	helper = func(node *ListNode) (*ListNode, int) {
		if node == nil {
			return nil, 0
		}
		_node, num := helper(node.Next)

		if num == n {
			node.Next = _node.Next
		}

		return node, num + 1
	}

	dummy := &ListNode{
		Next: head,
	}
	helper(dummy)

	return dummy.Next
}

func main() {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val: 5,
	//				},
	//			},
	//		},
	//	},
	//}
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//	},
	//}
	head := &ListNode{
		Val: 1,
	}
	removeNthFromEnd(head, 1)
}
