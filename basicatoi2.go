package piscine

func BasicAtoi2(s string) int {
	a := []rune(s)
	n := len(s)
	m := 0
	for i := 0; i < n; i++ {
		if a[i] < '0' || a[i] > '9' {
			return 0
		} else {
			m *= 10
			m += int(a[i]) - '0'
		}
	}
	return m
}
