package piscine

func LastRune(s string) rune {
	a := []rune(s)
	return a[len(a)-1]
}
