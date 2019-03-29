/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var next = head

	dummy := &ListNode{}
	tail := dummy
	for next != nil {
		sl := make([]*ListNode, 0)
		for i := 0; i < k; i++ {
			if head != nil {
				sl = append(sl, head)
				head = head.Next
			} else {
				break
			}
		}
		if len(sl) == k {
			next = sl[k-1].Next
			for i := k - 1; i > -1; i-- {
				tail.Next = sl[i]
				tail = tail.Next
			}
			tail.Next = nil
		} else {
			tail.Next = next
			return dummy.Next
		}
	}
	return dummy.Next
}
