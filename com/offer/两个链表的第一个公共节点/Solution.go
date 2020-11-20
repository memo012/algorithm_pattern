package 两个链表的第一个公共节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	hdA, hdB := headA, headB
	count := 0
	for hdA != hdB {
		hdA = hdA.Next
		hdB = hdB.Next
		if hdA == nil {
			hdA = headB
			count++
		}
		if hdB == nil {
			hdB = headA
			count++
		}
		if count > 2 {
			return nil
		}
	}
	return hdA
}
