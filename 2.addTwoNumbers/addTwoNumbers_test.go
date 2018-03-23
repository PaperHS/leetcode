package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_Add(t *testing.T) {

	n1 := rand.Intn(100000000000)
	n2 := rand.Intn(100000000000)
	n1 = 81
	n2 = 0
	l1 := int2ListNode(n1)
	l2 := int2ListNode(n2)
	fmt.Printf("n1=%d,n2=%d \n,n1+n2= %d\n", n1, n2, n1+n2)
	l3 := addTwoNumbers(l1, l2)
	fmt.Printf("result=%d \n", ListNode2int(l3))
}

func Test_180(t *testing.T) {
	l1 := int2ListNode(1)
	l3 := int2ListNode(999999)
	ret := addTwoNumbers(l1, l3)
	sum := ListNode2int(ret)
	fmt.Printf("sum:%d\n", sum)
	if sum != 1000000 {
		t.Errorf("sum = %d", sum)
	}
}

func ListNode2int(l *ListNode) int {
	var l2 *ListNode = l
	num := l2.Val
	tims := 10
	for l2.Next != nil {
		l2 = l2.Next
		num += l2.Val * tims
		tims *= 10
	}
	return num
}

func int2ListNode(n int) *ListNode {
	f := getListNode(n)
	r, num := f()
	for {
		_, num = f()
		if num == 0 {
			break
		}
	}
	return r
}

func getListNode(num int) func() (*ListNode, int) {
	l := new(ListNode)
	return func() (*ListNode, int) {
		ret := new(ListNode)
		ret.Val = num % 10
		//		fmt.Println("val", ret.Val)
		l.Next = ret
		num = num / 10
		//		fmt.Println("num:", num)
		l = ret
		return ret, num
	}
}
