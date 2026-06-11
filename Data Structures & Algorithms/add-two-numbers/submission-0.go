/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    curr1, curr2 := l1, l2
	var resList, curr3 *ListNode
	remaining := 0
	for curr1 != nil && curr2 != nil {
		val := curr1.Val + curr2.Val + remaining
		
		if curr3 == nil {
			curr3 = &ListNode{}
			resList = curr3
		} else {
			curr3.Next = &ListNode{}
			curr3 = curr3.Next
		}
		
		if val < 10 {
			curr3.Val = val
			remaining = 0
		} else {
			curr3.Val = val - 10
			remaining = 1
		}
		
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	var curr *ListNode
	if curr1 != nil {
		curr = curr1
	} else {
		curr = curr2
	}

	for curr != nil {
		val := curr.Val + remaining
		curr3.Next = &ListNode{}
		curr3 = curr3.Next
		if val < 10 {
			curr3.Val = val
			remaining = 0
		} else {
			curr3.Val = val - 10
			remaining = 1
		}
		curr = curr.Next
	}

	if remaining > 0 {
		curr3.Next = &ListNode{Val: remaining}
	}

	return resList
}
