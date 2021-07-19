package functions

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	result := &NodeI{}

	for val := n1; n1.Next != nil; n1 = n1.Next {
		result = SortListInsert(result, val.Data)
	}
	for val := n2; n2.Next != nil; n2 = n2.Next {
		result = SortListInsert(result, val.Data)
	}
	return result
}
