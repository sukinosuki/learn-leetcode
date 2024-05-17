package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	ans := 0
	for head != nil {
		ans = ans<<1 + head.Val

		head = head.Next
	}

	return ans
}

func main() {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	res := getDecimalValue(head)
	fmt.Printf("res: %v\n", res)
}
