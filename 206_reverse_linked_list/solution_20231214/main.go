package main

type ListNode struct {
	Val  int
	Next *ListNode
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

}
