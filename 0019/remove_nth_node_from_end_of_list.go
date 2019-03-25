/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    first := dummy
    second := dummy
    for n >= 0 {
        second = second.Next
        n --
    }
    
    for second != nil {
        first = first.Next
        second = second.Next
    }
    first.Next = first.Next.Next
    
    return dummy.Next
}
