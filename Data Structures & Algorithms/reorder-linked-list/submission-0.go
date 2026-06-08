/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // 1 -> 2 -> 3 -> 4
	// prev = nil
	// curr = 1
	//
	// map[*ListNode]*ListNode
	// 1 -> nil
	// 2 -> 1
	// 3 -> 2


	// 1
	// 4
	// 2
	// 3
	// be careful about even and ood lenghts
	// for odd lenghts, the last case the left and right will be equal at some point and
	// just one final insertion needs to be made

	m := map[*ListNode]*ListNode{}
	var prev *ListNode
	curr := head
	for curr != nil {
		m[curr] = prev
		prev = curr
		curr = curr.Next
	}

	// previous now holds the last node
	left, right := head, prev
	for left != right  {
		next := left.Next // 3
		// fmt.Printf("left: %d\n", left.Val)
		// fmt.Printf("right: %d\n", right.Val)
		// fmt.Printf("next: %d\n", next.Val)
		left.Next = right // 3, 2 ->3
		if right == next { // 3 == 3
			break
		}
		right.Next = next // 2, 1->4->2
		left = right.Next // 2
		right = m[right] // 3
		// fmt.Printf("left: %d\n", left.Val)
		// fmt.Printf("right: %d\n", right.Val)
		// fmt.Printf("next: %d\n", next.Val)
	}
	right.Next = nil
}
