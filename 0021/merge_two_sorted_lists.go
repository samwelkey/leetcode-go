/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}

	if l2 == nil && l1 != nil {
		return l1
	}

	if l1 == nil && l2 == nil {
		return l1
	}

	var dummy = &ListNode{}
	var tail *ListNode

	if l1.Val <= l2.Val {
		dummy.Next = l1
		tail = dummy.Next
		l1 = l1.Next
	} else {
		dummy.Next = l2
		tail = dummy.Next
		l2 = l2.Next
	}

	for {

		if l1 == nil && l2 != nil {
			tail.Next = l2
			return dummy.Next
		}

		if l2 == nil && l1 != nil {
			tail.Next = l1
			return dummy.Next
		}

		if l1.Val <= l2.Val {
			tail.Next = l1
			tail = tail.Next
			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = tail.Next
			l2 = l2.Next
		}
	}
}
