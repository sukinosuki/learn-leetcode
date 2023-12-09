package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	var p *ListNode
	for head != nil {
		nums = append(nums, head.Val)
		next := head.Next
		head.Next = p
		p = head
		head = next
	}
	for index := range nums {
		if nums[index] != p.Val {
			return false
		}

		p = p.Next
	}

	return true
}

func main() {

	nodes := &ListNode{Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		}}

	isValid := isPalindrome(nodes)

	fmt.Printf("isValid: %v\n", isValid)
}
