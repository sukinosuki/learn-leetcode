package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	extraNode := &ListNode{}

	cur := head
	for cur != nil {
		if cur.Next == extraNode {
			return true
		}

		next := cur.Next
		cur.Next = extraNode
		cur = next
	}

	return false
}

func main() {

}
