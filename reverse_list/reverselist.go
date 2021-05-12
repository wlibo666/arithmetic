package reverse_list

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func printList(head *ListNode) {
	tmp := head

	for tmp != nil {
		fmt.Printf("node value:%d\n", tmp.Value)
		tmp = tmp.Next
	}
}

func Reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var tmp, pre *ListNode

	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
