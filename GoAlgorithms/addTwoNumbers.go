package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
//	var tail *ListNode
//	carry := 0
//	for l1 != nil || l2 != nil {
//		n1 := 0
//		n2 := 0
//		if l1 != nil {
//			n1 = l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			n2 = l2.Val
//			l2 = l2.Next
//		}
//		sum := n1 + n2 + carry
//		sum, carry = sum%10, sum/10
//		if head == nil {
//			head = &ListNode{Val: sum}
//			tail = head
//		} else {
//			tail.Next = &ListNode{Val: sum}
//			tail = tail.Next
//		}
//	}
//	if carry > 0 {
//		tail.Next = &ListNode{Val: carry}
//	}
//	return
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	for l1 != nil && l2 != nil {
		v1 := l1.Val
		v2 := l2.Val
		l1 = l1.Next
		l2 = l2.Next
		sum = v1 + v2
	}
	return &ListNode{sum, nil}
}

func main() {
	//方法一：
	//设a1=1 a2=2 a3=4
	//a1 + a2 + a3 = 2^3 -1
	//a1^2 + a2^2 + a3^2 = 21 = (4^3-1)/3

	//方法二：
	//a1 + a2 +.....+ an = 2^n -1 = a1(1-q^n)/(1-q)
	//(a1 + a2 +.....+ an)^2 = 4^n + 1 -2^(n+1)
	//a1^2 + a2^2 + a3^2 +an^2 = 构成新等比数列 q=q^2，首项为a1^2
	//结果=a1^2(1-q^2n)/(1-q^2)
	//

	l1 := new(ListNode)
	b := new(ListNode)
	c := new(ListNode)
	l1.Val = 1
	l1.Next = b
	b.Val = 2
	b.Next = c
	c.Val = 3
	c.Next = nil
	fmt.Println("``````````获取l1````````")
	//for l1 != nil{
	//	fmt.Println(l1.Val)
	//	l1 = l1.Next
	//}

	fmt.Println("``````````获取l2``````````")
	l2 := new(ListNode)
	e := new(ListNode)
	f := new(ListNode)
	l2.Val = 4
	l2.Next = e
	e.Val = 5
	e.Next = f
	f.Val = 6
	f.Next = nil
	//for l2 != nil{
	//	fmt.Println(l2.Val)
	//	l2 = l2.Next
	//}

	l3 := addTwoNumbers(l1, l2)
	for i := 0; i < 3; i++ {
		fmt.Println(l3.Val)
		l3.Next = l2
	}

}
