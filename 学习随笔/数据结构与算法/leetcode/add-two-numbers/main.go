package main

import "fmt"

func main() {
	a := &ListNode{Val:9}
	b := &ListNode{Val:9}
	//c := &ListNode{Val:9}
	d := &ListNode{Val:9}
	e := &ListNode{Val:9}
	f := &ListNode{Val:9}
	a.Next = b
	b.Next = nil
	//c.Next = nil

	d.Next = e
	e.Next = f
	f.Next = nil
	listListNode(addTwoNumbers(a, d))
}


type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
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
	}
	return
}


func listListNode(l1 *ListNode) {
	for l1 != nil {
		fmt.Print(l1.Val, "  ")
		l1 = l1.Next
	}
	fmt.Println("list")
}