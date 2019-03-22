/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1.Next == nil && l2.Next == nil{
        if l1.Val+l2.Val >= 10 {
            return &ListNode{(l1.Val+l2.Val)%10, &ListNode{1,nil}}
        }
        return &ListNode{(l1.Val+l2.Val)%10,nil}
    }
    
    if l1.Next == nil {
        l1.Next = &ListNode{0,nil}
    }
    
    if l2.Next == nil {
        l2.Next = &ListNode{0,nil}
    }
    
    if l1.Val + l2.Val >= 10{
        l1.Next.Val += 1
    }
    
    return &ListNode{(l1.Val + l2.Val)%10,addTwoNumbers(l1.Next,l2.Next)}
}
