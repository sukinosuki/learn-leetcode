package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	//if head == nil {
	//	return nil
	//}
	//
	//node := reverseList(head.Next)
	//if node == nil {
	//	return head
	//}
	//
	//node.Next = head
	//head.Next = nil
	//
	//return head
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
