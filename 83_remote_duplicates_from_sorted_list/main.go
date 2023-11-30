package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next == nil {
			break
		}
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}

	return head
}
func main() {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val:  2,
	//			Next: nil,
	//		},
	//	},
	//}

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	res := deleteDuplicates(head)

	fmt.Printf("res: %s\n", res)
}
