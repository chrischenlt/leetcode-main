package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入：head = [1,2,3,4]
//输出：[2,1,4,3]

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(head.Next.Next)
	next.Next = head
	return next
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	res := dummy
	for dummy.Next != nil && dummy.Next.Next != nil {
		tag, next := dummy.Next, dummy.Next.Next.Next
		dummy.Next = dummy.Next.Next
		dummy.Next.Next = tag
		tag.Next = next
		dummy = dummy.Next.Next
	}

	return res.Next
}
