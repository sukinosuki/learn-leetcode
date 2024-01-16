package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	l1 := head
	l2 := head

	for l2 != nil && l2.Next != nil {

		l1 = l1.Next
		l2 = l2.Next.Next
	}

	return l1
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

	res := middleNode(head)
	fmt.Printf("res: %v\n", res.Val)
}
