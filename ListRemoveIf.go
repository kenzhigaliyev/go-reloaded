package student

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head
	l.Head = nil
	if current == nil {
		return
	}
	for current.Next != nil {
		if current.Data != data_ref {
			ListPushBackNodeL(l, current.Data)
		}
		current = current.Next
	}
	return
}
