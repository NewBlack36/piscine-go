package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	iterator := l.Head
	for iterator != nil {
		if comp(iterator.Data, ref) {
			return &iterator.Data
		}

		iterator = iterator.Next
	}
	return nil
}