package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	if n == len(s) {
		return a[len(s)-1]
	} else if n <= 0 || n > len(s) {
		return 0
	} else {
		return a[n-1]
	}
}
