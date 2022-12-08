package piscine

func Map(f func(int) bool, a []int) []bool {
	length := 0
	for l := range a {
		length = l + 1
	}
	alpha := make([]bool, length)
	for i := range a {
		alpha[i] = f(a[i])
	}
	return alpha
}
