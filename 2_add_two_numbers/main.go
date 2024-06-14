package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}

	cur := dummy
	remain := 0
	for l1 != nil || l2 != nil || remain != 0 {
		l1Num := 0
		l2Num := 0
		if l1 != nil {
			l1Num = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Num = l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{
			Val: (l1Num + l2Num + remain) % 10,
		}

		remain = (l1Num + l2Num + remain) / 10

		cur = cur.Next
	}

	return dummy.Next
}

func main() {

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	res := addTwoNumbers(l1, l2)

	fmt.Printf("res: %v\n", res)
}
