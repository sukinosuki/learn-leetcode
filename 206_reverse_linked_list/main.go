package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	p := &ListNode{}
	cur := p
	var helper func(node *ListNode)

	helper = func(node *ListNode) {

		if node == nil {
			return
		}
		helper(node.Next)

		cur.Next = node
		cur = cur.Next
		node.Next = nil
	}

	helper(head)

	return p.Next
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
