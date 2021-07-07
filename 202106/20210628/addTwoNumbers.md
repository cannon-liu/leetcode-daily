### 两数之和

#### 解题思路一：进位方法
采用n1+n2+carry % 10是该位的值，较短链表的值

##### 复杂度分析
时间复杂度: O(max(m,n)),两个链表的最大值
空间复杂度: O(1)。注意返回值不计入空间复杂度

#### 解题思路二：数字和链表的相互转换
这边测试没有问题，但是网页上总是报错

#### 解题思路三：递归法

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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
				l2.Next = addTwoNumbers(&ListNode{0,nil},l2.Next)
			}
		} else {
			l2.Next = &ListNode{1,nil}
		}
	}
	l2.Next = addTwoNumbers(l1.Next,l2.Next)
	return l2
}
```

复杂度分析：


```g```