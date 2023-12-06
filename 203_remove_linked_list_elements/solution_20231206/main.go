package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	var helper func(node *ListNode)

	p := &ListNode{}
	cur := p
	helper = func(node *ListNode) {
		if node == nil {
			return
		}
		if node.Val == val {
			helper(node.Next)
		} else {
			cur.Next = &ListNode{
				Val: node.Val,
			}
			cur = cur.Next
			helper(node.Next)
		}
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
