/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var dummy = &ListNode{}
	tail := dummy
	var next *ListNode
	next = head.Next.Next
	dummy.Next = head.Next
	tail = tail.Next
	tail.Next = head
	tail = tail.Next

	tail.Next = nil

	for next != nil {
		head = next
		if head.Next == nil {
			tail.Next = head
			return dummy.Next
		}
		next = head.Next.Next
		tail.Next = head.Next
		tail = tail.Next
		tail.Next = head
		tail = tail.Next
		tail.Next = nil

	}

	return dummy.Next

}
