/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// k = 3
// 1 -> 2 -> 3 -> 4 -> 5 -> 6
// 1 <- 2 <- 3
// curr: 4, head: 1
// prevHead.Next = curr
// 


func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead, newTail := reverseUpToK(head, k)
	head = newHead
	for newTail != nil {
		tmp := newTail
		newHead, newTail = reverseUpToK(newTail.Next, k)
		tmp.Next = newHead
	}
	return head
}

// First return vale is the new head
// Second return value is the new tail
func reverseUpToK(head *ListNode, k int) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	if !isAllowedToReverse(head, k) {
		return head, nil
	}

	var prev *ListNode
	curr := head
	count := 0
	for curr != nil && count < k {
		//fmt.Println("curr: ", curr.Val)
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
		count++
	}
	// fmt.Println("curr: ", curr)
	// fmt.Println("prev: ", prev)
	// fmt.Println("head: ", head)
	head.Next = curr
	return prev, head
}

func isAllowedToReverse(head *ListNode, k int) bool {
	if head == nil {
		return false
	}
	curr := head
	count := 0
	for curr != nil && count < k {
		curr = curr.Next
		count++
	}

	return count == k
}



