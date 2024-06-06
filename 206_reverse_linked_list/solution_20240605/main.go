package main

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

}
