package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	c := len(a)
	d := len(b)
	for i := 0; i <= c-d; i++ {
		if toFind == s[i:i+d] {
			return i
		}
	}
	return -1
}
