/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    curr1 := list1
    curr2 := list2
    var currMerged, headMerged *ListNode
    for curr1 != nil && curr2 != nil {
        var minor *ListNode 
        if curr1.Val < curr2.Val {
            minor = curr1
            curr1 = curr1.Next
        } else {
            minor = curr2
            curr2 = curr2.Next
        }

        if currMerged == nil {
            currMerged = minor
            headMerged = minor
        } else {
            currMerged.Next = minor
            currMerged = currMerged.Next
        }
    }

    if curr1 != nil {
        if currMerged == nil {
            currMerged = curr1
            headMerged = curr1
        } else {
            currMerged.Next = curr1
        }
    }

    if curr2 != nil {
        if currMerged == nil {
            currMerged = curr2
            headMerged = curr2
        } else {
            currMerged.Next = curr2
        }
    }

    return headMerged
}
