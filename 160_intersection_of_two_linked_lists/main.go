package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	l1 := headA
	l2 := headB
	m := make(map[*ListNode]bool)
	for l1 != nil || l2 != nil {
		if l1 != nil {
			if _, ok := m[l1]; ok {
				return l1
			}
			m[l1] = true
			l1 = l1.Next
		}
		if l2 != nil {
			if _, ok := m[l2]; ok {
				return l2
			}
			m[l2] = true
			l2 = l2.Next
		}
	}

	return nil
}

func main() {

}
