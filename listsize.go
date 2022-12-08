package piscine

func ListSize(l *List) int {
	let := l.Head
	i := 0
	for let != nil {
		i++
		let = let.Next
	}
	return i
}
