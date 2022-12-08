package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	i := 0
	cpt := 0
	j := l

	for j != nil {
		i++
		j = j.Next
	}
	if pos <= i {
		for l != nil {
			if cpt == pos {
				return l
			}
			cpt++
			l = l.Next
		}
	}
	return nil
}
