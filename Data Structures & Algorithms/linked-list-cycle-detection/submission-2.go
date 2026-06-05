/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    var fast *ListNode
    slow := head
    if slow.Next != nil {
        fast = slow.Next.Next
    }
    for slow != nil && fast != nil {
        fmt.Printf("slow: %p, fast: %p\n", slow, fast)
        if slow == fast {
            return true
        }

        if slow.Next != nil {
            slow = slow.Next
        } else {
            return false
        }

        if fast.Next != nil {
            fast = fast.Next.Next
        } else {
            return false
        }
    }

    return false;
}
