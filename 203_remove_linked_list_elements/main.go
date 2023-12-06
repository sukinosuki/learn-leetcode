package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	if head != nil && head.Val == val {
		return head.Next
	}

	return head
}

func main() {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
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
		},
	}

	res := removeElements(head, 6)

	fmt.Printf("res: %v\n", res)
}
