package piscine

func BasicAtoi(s string) int {
	cp := 0
	for _, i := range s {
		cp = cp*10 + int(i) - '0'
	}
	return cp
}
