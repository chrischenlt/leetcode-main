package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var last *ListNode

	for head.Next != nil {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}

	head.Next = last

	return head
}
