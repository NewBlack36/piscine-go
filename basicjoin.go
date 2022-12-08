package piscine

func BasicJoin(elems []string) string {
	var r string
	for _, i := range elems {
		r += i
	}
	return r
}
