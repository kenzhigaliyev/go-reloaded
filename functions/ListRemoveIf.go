package functions

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

	// for current != nil {
	// 	if l.Head.Data == data_ref {
	// 		l.Head = l.Head.Next
	// 		current = l.Head
	// 		// current = current.Next
	// 		// previous = l.Head
	// 	} else if current.Next != nil && current.Data == data_ref {
	// 		oldnext := current.Next
	// 		current = oldnext
	// 		// fmt.Println("here")
	// 		current = current.Next
	// 	} else if current.Next == nil && current.Data == data_ref {
	// 		current = nil
	// 	}
	// 	// current = current.Next
	// }

	// if l.Head == nil {
	// 	return
	// } else if l.Head.Data == data_ref {
	// 	new := &NodeI{}
	// 	new.Next = l.Head.Next
	// 	// l = new
	// } else {
	// 	val := l.Head
	// 	for val.Next != nil && val.Next.Data != data_ref {
	// 		val.Next = val.Next.Next
	// 	}

	// }
}
