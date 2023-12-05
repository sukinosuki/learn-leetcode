package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	low := head
	quick := head.Next

	for {
		if low == nil || quick == nil || quick.Next == nil {
			return false
		}

		if low == quick || low == quick.Next {
			return true
		}

		low = low.Next
		quick = quick.Next.Next
	}
}

func main() {

}
