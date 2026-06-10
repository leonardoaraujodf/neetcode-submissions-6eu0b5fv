/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

    curr := head
	for i := 0; i < n; i++ {
		curr = curr.Next
	}

	node := head
	var prev *ListNode
	for curr != nil {
		prev = node
		node = node.Next
		curr = curr.Next
	}
	// fmt.Println(prev.Val)
	// fmt.Println(node.Val)
	if prev == nil {
		head = head.Next
	} else {
		prev.Next = node.Next
	}

	return head
}
