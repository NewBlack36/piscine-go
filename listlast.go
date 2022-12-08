package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	let := l.Head
	for let.Next != nil {
		let = let.Next
	}
	return let.Data
}
