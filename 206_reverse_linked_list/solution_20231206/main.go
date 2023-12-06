package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var p *ListNode

	for head != nil {
		next := head.Next
		head.Next = p

		p = head
		head = next
	}

	return p
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
					},
				},
			},
		},
	}

	res := reverseList(head)

	fmt.Printf("res: %v\n", res)
}
