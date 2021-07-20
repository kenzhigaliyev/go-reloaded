package functions

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	// PrintList(l)
	new := &NodeI{Data: data_ref}
	if l == nil || l.Data >= data_ref {
		new.Next = l
		return new
	} else {
		val := l
		for val.Next != nil && val.Next.Data < data_ref {
			val = val.Next
		}
		new.Next = val.Next
		val.Next = new
	}
	return l
}
