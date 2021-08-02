package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	result := &NodeI{}
	result = nil
	for val := n1; val != nil; val = val.Next {
		result = SortListInsert(result, val.Data)
	}
	for val := n2; val != nil; val = val.Next {
		result = SortListInsert(result, val.Data)
	}
	return result
}
