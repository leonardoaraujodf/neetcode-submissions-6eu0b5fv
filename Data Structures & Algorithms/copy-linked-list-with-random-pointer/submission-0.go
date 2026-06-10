/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
		return nil
	}
	// nodesMap maps a Node from the provided linked list to nodes created in the new
	// linked list. It is a 1 to 1 mapping, like the head from list 1 maps to head of
	// list 2.
	nodesMap := make(map[*Node]*Node)
	// curr iterates over the first list
	curr := head
	for curr != nil {
		// tmp iterates over the second list
		curr2, ok := nodesMap[curr]
		if !ok {
			curr2 = &Node{Val: curr.Val,}
			nodesMap[curr] = curr2
		}
		
		if curr.Next != nil {
			next2, ok := nodesMap[curr.Next]
			if !ok {
				next2 = &Node{Val: curr.Next.Val,}
				nodesMap[curr.Next] = next2
			}
			curr2.Next = next2
		}

		if curr.Random != nil {
			random2, ok := nodesMap[curr.Random]
			if !ok {
				random2 = &Node{Val: curr.Random.Val,}
				nodesMap[curr.Random] = random2
			}
			curr2.Random = random2
		}

		curr = curr.Next
	}
	return nodesMap[head]
}
