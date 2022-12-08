package piscine

func ListReverse(l *List) {
	ter := l.Head
	ex := l.Head
	ex = nil
	for ter != nil {
		next := ter.Next
		ter.Next = ex
		ex = ter
		ter = next
	}
	s := l.Head
	l.Head = l.Tail
	l.Tail = s
}
