/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * @{link https://leetcode-cn.com/problems/add-two-numbers/description/}
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret *ListNode = new(ListNode)
	x(l1, l2, ret, 0)
	return ret.Next
}

func x(l1 *ListNode, l2 *ListNode, l3 *ListNode, add int) (*ListNode,
	*ListNode, *ListNode, int) {
	var rrr *ListNode = new(ListNode)
	rrr.Val += add
	if l1 != nil {
		rrr.Val += l1.Val
	}
	if l2 != nil {
		rrr.Val += l2.Val
	}
	biger10 := 0
	if rrr.Val/10 > 0 {
		rrr.Val = rrr.Val % 10
		biger10 = 1
	}
	if l1 == nil && l2 == nil && add == 0 {
		l3.Next = nil
		return nil, nil, l3, 0
	} else {
		l3.Next = rrr
		return x(ifHasNext(l1), ifHasNext(l2), rrr, biger10)
	}

}
func ifHasNext(l1 *ListNode) *ListNode {
	if l1 == nil {
		return nil
	} else {
		return l1.Next
	}
}
