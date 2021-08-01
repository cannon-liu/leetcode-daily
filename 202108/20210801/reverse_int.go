package main

import "fmt"

func main() {
	x := reverse(1534236469)
	fmt.Println(x)

	x = reverse(-123)
	fmt.Println(x)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(x int) *ListNode {
	var tail *ListNode
	mod := x % 10
	x = x / 10
	head := &ListNode{
		Val: mod,
	}
	tail = head
	for {
		if x <= 0 {
			break
		}
		mod = x % 10
		x = x / 10
		tail.Next = &ListNode{
			Val: mod,
		}
		tail = tail.Next
	}
	return head
}

func reverse(x int) int {
	var pg int = 1
	var data int = 0
	if x < 0 {
		pg = -1
	} else {
		pg = 1
	}
	x = x * pg
	node := makeList(x)
	for {
		if node == nil {
			break
		}
		data = data*10 + node.Val
		node = node.Next
	}
	return data * pg
}
