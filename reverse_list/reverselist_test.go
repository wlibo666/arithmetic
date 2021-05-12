package reverse_list

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	node1 := &ListNode{Value: 1}
	node2 := &ListNode{Value: 2}
	node3 := &ListNode{Value: 3}
	node4 := &ListNode{Value: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = nil

	fmt.Printf("before reverse:\n")
	printList(node1)

	fmt.Printf("\nend reverse:\n")
	reverse := Reverse(node1)
	printList(reverse)
}
