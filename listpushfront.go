package piscine

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		l.Head, l.Tail = &NodeL{Data: data}, l.Head
	} else {
		n := &NodeL{Data: data}
		n.Next, l.Head = l.Head, n
	}
}
