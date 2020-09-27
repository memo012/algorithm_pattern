package 链表中倒数第k个节点

type ListNode struct {
	Val int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	node := new(ListNode)
	node.Next = head
	tail := node
	for i := 0; i < k; i++ {
		tail = tail.Next
	}
	for tail != nil {
		node = node.Next
		tail = tail.Next
	}
	return node
}
