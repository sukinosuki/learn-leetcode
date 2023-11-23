package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}
	res := &ListNode{}
	curr := res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curr.Val = list2.Val
			curr.Next = list2.Next
			return res
		}
		if list2 == nil {
			curr.Val = list1.Val
			curr.Next = list1.Next
			return res
		}

		if list1.Val <= list2.Val {
			curr.Val = list1.Val
			list1 = list1.Next
		} else {
			curr.Val = list2.Val
			list2 = list2.Next
		}

		curr.Next = &ListNode{}
		curr = curr.Next
	}

	return res
}

func main() {
	//list1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 4,
	//		},
	//	},
	//}
	//list2 := &ListNode{Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val: 4,
	//		},
	//	}}

	list1 := &ListNode{Val: 2}
	list2 := &ListNode{Val: 1}

	res := mergeTwoLists(list1, list2)

	fmt.Println("res ", res)
}
