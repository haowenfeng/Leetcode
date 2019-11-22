package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func insert(val int, head *ListNode) *ListNode {
	var node *ListNode = &ListNode{Val: val}
	node.Next = head
	return node
}

func dump(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("cur.Val:%d cur.Next:%v\n", cur.Val, cur.Next)
		cur = cur.Next
	}
}

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5
	dump(head)

	fmt.Print("\n\nafter reverseList...\n")
	newlist := reverseList(head)
	dump(newlist)

	fmt.Print("\n\nafter insert before head...\n")
	newhead := insert(100, newlist)
	dump(newhead)
}
