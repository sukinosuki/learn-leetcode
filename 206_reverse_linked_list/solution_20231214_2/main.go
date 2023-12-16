package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	stack := []*ListNode{}
	cur := head
	for cur != nil {
		stack = append(stack, cur)
		next := cur.Next
		cur = next
	}

	p := &ListNode{}
	cur = p
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node.Next = nil

		cur.Next = node
		cur = cur.Next
	}

	return p.Next
}

func main() {

}
