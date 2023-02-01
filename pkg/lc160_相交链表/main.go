package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	//tagA, tagB := 0, 0
	//for a != nil && b != nil {
	//	if a == b {
	//		return a
	//	}
	//	a = a.Next
	//	b = b.Next
	//	if a == nil{
	//		a = headB
	//		tagA++
	//	}
	//	if b == nil{
	//		b = headA
	//		tagB++
	//	}
	//}

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}
