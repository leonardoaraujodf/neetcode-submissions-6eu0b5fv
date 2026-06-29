/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    k := len(lists)
	if k == 0 {
		return nil
	} else if k == 1 {
		return lists[0]
	}

	res := lists[0]
	for i := 1; i < k; i++ {
		res = merge(res, lists[i])
	}
	return res
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	curr1, curr2 := list1, list2
	var head *ListNode
	var resCurr *ListNode
	for curr1 != nil && curr2 != nil {
		var tmp *ListNode
		if curr1.Val < curr2.Val {
			tmp = curr1
			curr1 = curr1.Next
		} else {
			tmp = curr2
			curr2 = curr2.Next
		}
		if head == nil {
			head = tmp
			resCurr = head
		} else {
			resCurr.Next = tmp
			resCurr = resCurr.Next
		}
	}
	if curr1 != nil {
		resCurr.Next = curr1
	} else if curr2 != nil {
		resCurr.Next = curr2
	}

	return head
}
