package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	cur2 := head
	index1 := 1
	index2 := 1
	for cur2.Next != nil && cur2.Next.Next != nil {
		cur = cur.Next
		cur2 = cur2.Next.Next
		index1 += 1
		index2 += 2
	}
	// 要删除的node在cur之后
	//if cur2.Next!= nil{
	//	n-=1
	//}
	if n < index1 || (cur2.Next != nil && n <= index1) {
		size := index2 - index1 - n
		if cur2.Next != nil {
			size += 1
		}

		for size > 0 {
			cur = cur.Next
			size--
		}

		cur.Next = cur.Next.Next

		return head
	}

	//要删除的node在cur之后，重新按次数循环
	size := index2
	if cur2.Next != nil {
		size += 1
	}
	size -= n

	dummy := &ListNode{
		Next: head,
	}
	cur = dummy
	for size > 0 {

		cur = cur.Next
		size--
	}
	cur.Next = cur.Next.Next

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
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	removeNthFromEnd(head, 4)

}
