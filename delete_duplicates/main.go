package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//func deleteDuplicates(head *ListNode) *ListNode {
//
//	if head == nil {
//		return nil
//	}
//	list := &ListNode{Val: head.Val}
//	tail := head.Next
//	prehead := list
//
//	for tail != nil {
//		if tail.Val == prehead.Val {
//			prehead.Next = tail.Next
//			//tail = tail.Next
//		} else {
//			prehead.Next = tail
//			prehead = tail
//		}
//		tail = tail.Next
//
//	}
//
//	return list
//}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

func main() {
	//list := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val: 3,
	//				Next: &ListNode{
	//					Val: 3,
	//				},
	//			},
	//		},
	//	},
	//}

	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}

	//list := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//	},
	//}
	res := deleteDuplicates(list)

	fmt.Println("res ", res)
}
