/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	m := map[*ListNode]*ListNode{}
	var prev *ListNode
	curr := head
	for curr != nil {
		m[curr] = prev
		prev = curr
		curr = curr.Next
	}

	left, right := head, prev
	for left != right  {
		next := left.Next 
		left.Next = right
		if right == next {
			break
		}
		right.Next = next 
		left = right.Next
		right = m[right]
	}
	right.Next = nil
}
