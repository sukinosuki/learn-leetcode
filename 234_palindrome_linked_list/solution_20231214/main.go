package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	slow = reverseList(slow)

	for slow != nil {
		if slow.Val != head.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {

	var p *ListNode

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = p
		p = cur

		cur = next
	}

	return p
}

func main() {

	nodes := &ListNode{Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		}}

	isValid := isPalindrome(nodes)

	fmt.Printf("isValid: %v\n", isValid)
}
