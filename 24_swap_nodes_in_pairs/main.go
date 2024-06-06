package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	dummy := &ListNode{
		Next: head,
	}
	cur := dummy
	for head != nil && head.Next != nil {

		// node3
		next2 := head.Next.Next
		// node1
		next := head.Next

		// node2指向node1
		next.Next = head
		// node1指向node3
		head.Next = next2

		// cur.next指向node2
		cur.Next = next
		// cur移动到node1, 因为node1.next会指向node3,后面node3和node4交换时，cur.next就指向node4了
		cur = head

		// 移动到node3,下一次循环执行node3和node4的交换
		// 优化: 可以直接使用cur.next cur.next.next来判断循环是否结束
		head = next2
	}

	return dummy.Next
}

func main() {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	swapPairs(head)
}
