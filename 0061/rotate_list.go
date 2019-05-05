/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0{
        return head
    }
    h := head
    var l []*ListNode
    for h != nil {
        l = append(l,h)
        h = h.Next
    }
    total := len(l)
    l[total-1].Next = head
    index := total - 1 - (k-1) % total
    preIndex := total - 1 - (k) % total
    
    h = l[index]
    l[preIndex].Next = nil
    
    
    return h
    
    
}
