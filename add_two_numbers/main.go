package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	root := &ListNode{}
	tail := root
	carry := 0

	for l1 != nil || l2 != nil {
		var num1, num2 int

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		sum := num1 + num2 + carry
		if sum >= 10 {
			tail.Val = sum % 10
			carry = 1
		} else {
			carry = 0
			tail.Val = sum
		}

		if l1 != nil || l2 != nil {
			tail.Next = &ListNode{}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}

	return root
}

func main() {
	list1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	//list1 := &ListNode{
	//	Val: 9,
	//	Next: &ListNode{
	//		Val: 9,
	//		Next: &ListNode{
	//			Val: 9,
	//			Next: &ListNode{
	//				Val: 9,
	//				Next: &ListNode{
	//					Val: 9,
	//					Next: &ListNode{
	//						Val: 9,
	//						Next: &ListNode{
	//							Val:  9,
	//							Next: nil,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//list2 := &ListNode{
	//	Val: 9,
	//	Next: &ListNode{
	//		Val: 9,
	//		Next: &ListNode{
	//			Val: 9,
	//			Next: &ListNode{
	//				Val:  9,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}

	res := addTwoNumbers(list1, list2)

	println("res %s", res)
}
