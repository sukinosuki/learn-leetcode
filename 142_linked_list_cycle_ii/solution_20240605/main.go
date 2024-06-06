package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	flag := false
	for fast != nil && fast.Next != nil {
		if fast.Next == fast.Next.Next {
			return fast
		}

		if flag {
			fast = fast.Next
		} else {
			fast = fast.Next.Next
		}
		slow = slow.Next

		if flag && fast == slow {
			return fast
		}
		if fast == slow {
			slow = head
			flag = true
		}
	}

	return nil
}

func main() {
	node3 := &ListNode{
		Val: 3,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node0 := &ListNode{
		Val: 0,
	}
	node4 := &ListNode{
		Val: 4,
	}

	node3.Next = node2
	node2.Next = node0
	node0.Next = node4

	node4.Next = node2
	head := node3

	res := detectCycle(head)

	fmt.Printf("res: %v\n", res)
}
