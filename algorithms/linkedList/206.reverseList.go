package linkedList

import (
	"fmt"

	"Leetcode/algorithms/kit"
)

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		fmt.Printf("pre:%v cur:%v\n", kit.List2Ints(pre), kit.List2Ints(cur))
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	fmt.Printf("pre:%v cur:%v\n", kit.List2Ints(pre), kit.List2Ints(cur))
	return pre
}

// 执行过程：
// pre:[nil] cur:[1 2 3 nil]
// pre:[1 nil] cur:[2 3 nil]
// pre:[2 1 nil] cur:[3 nil]
// pre:[3 2 1 nil] cur:[nil]
