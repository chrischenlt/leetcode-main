package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	// 1. 遍历，用哈希表存储，直至能在哈希表中找到或者为空
	// TODO: 2.快慢指针
	m := make(map[*ListNode]struct{})

	for true {
		if head == nil {
			return nil
		}
		_, exist := m[head]
		if exist {
			return head
		} else {
			m[head] = struct{}{}
		}
		head = head.Next
	}
	return nil
}
