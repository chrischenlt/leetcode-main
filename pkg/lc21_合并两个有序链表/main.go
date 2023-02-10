package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			res.Next = list1
			list1 = list1.Next
		} else {
			res.Next = list2
			list2 = list2.Next
		}
		res = res.Next
	}

	if list1 != nil {
		res.Next = list1
	}

	if list2 != nil {
		res.Next = list2
	}

	return head.Next
}
