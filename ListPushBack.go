package student

type NodeI struct {
	Data int
	Next *NodeI
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBackNodeI(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func ListPushBackNodeL(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		return
	}

	val := l.Head

	for val.Next != nil {
		val = val.Next
	}
	val.Next = &NodeL{Data: data}
}
