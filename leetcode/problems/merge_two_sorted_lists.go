package problems

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil {
			cur.Next = l2
			break
		}

		if l2 == nil {
			cur.Next = l1
			break
		}

		v1 := getVal(l1)
		v2 := getVal(l2)
		if v1 < v2 {
			cur.Next = l1
			cur = l1
			l1 = getNext(l1)
		} else {
			cur.Next = l2
			cur = l2
			l2 = getNext(l2)
		}
	}
	return getNext(res)
}
