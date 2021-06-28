package main

import "log"

type ListNode struct {
	Val int
	Next *ListNode
}

func formatList(nums int) *ListNode{
	var tail *ListNode

	data := nums % 10
	nums = nums / 10

	head := &ListNode{Val: data}
	tail = head

	for {
		if nums == 0 {
			break
		}
		data = nums % 10
		nums = nums / 10
		tail.Next = &ListNode{Val: data}
		tail = tail.Next
	}

	return  head
}

func formatInt(nums *ListNode) int {
	multiple := 1
	data := 0
	for nums!=nil {
		data += nums.Val * multiple
		multiple *= 10
		nums = nums.Next
	}
	return data
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	d1 := formatInt(l1)
	d2 := formatInt(l2)
	data := d1+d2
	l := formatList(data)
	return l
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	sum := 0
	n1,n2 := 0,0
	for l1!=nil || l2 != nil {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum = n1+n2+carry
		carry = sum/10
		sum = sum%10
		n1,n2 = 0,0
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}
	return head
}


func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l2.Val = l1.Val + l2.Val
	if l2.Val >= 10 {
		l2.Val = l2.Val % 10
		if l2.Next != nil {
			l2.Next.Val = l2.Next.Val + 1
			if l2.Next.Val == 10 {
				l2.Next = addTwoNumbers3(&ListNode{0,nil},l2.Next)
			}
		} else {
			l2.Next = &ListNode{1,nil}
		}
	}
	l2.Next = addTwoNumbers3(l1.Next,l2.Next)
	return l2
}

func main() {
	l1 := formatList(10000000001)
	l2 := formatList(465)
	l := addTwoNumbers3(l1,l2)
	data := formatInt(l)
	log.Println("result is ",data)
}
