package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	let := l.Head
	for let != nil {
		if cond(let) == true {
			f(let)
		}
		let = let.Next
	}
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
