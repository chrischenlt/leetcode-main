package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{}
	p := dummy
	for true {

		flag := true
		min := math.MaxInt
		index := 0
		for i := 0; i < len(lists); i++ {
			l := lists[i]
			if l != nil {
				flag = false
			} else {
				continue
			}
			if l.Val < min {
				min = l.Val
				index = i
			}
		}
		if flag {
			break
		}
		p.Next = lists[index]
		p = p.Next
		lists[index] = lists[index].Next
	}

	return dummy.Next
}
