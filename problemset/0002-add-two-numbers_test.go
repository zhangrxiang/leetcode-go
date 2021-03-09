package problemset

import (
	"log"
	"testing"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，
// 并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//示例：
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}
	l1.Next.Next.Next = &ListNode{5, nil}

	l2 := &ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next.Next = &ListNode{4, nil}
	l2.Next.Next.Next = &ListNode{4, nil}
	log.Println(l1.Val, l1.Next.Val, l1.Next.Next.Val, l1.Next.Next.Next.Val)
	log.Println(l2.Val, l2.Next.Val, l2.Next.Next.Val, l2.Next.Next.Next.Val)
	l3 := addTwoNumbers(l1, l2)
	log.Println(l3.Val, l3.Next.Val, l3.Next.Next.Val, l3.Next.Next.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	h := list
	sum, carry := 0, 0
	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		sum += carry
		carry = sum / 10
		h.Next = &ListNode{
			Val: sum % 10,
		}
		h = h.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		h.Next = &ListNode{1, nil}
	}
	return list.Next
}
