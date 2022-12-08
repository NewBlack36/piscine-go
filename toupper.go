package piscine

func ToUpper(s string) string {
	a := []rune(s)
	for i := range a {
		if a[i] >= 'a' && a[i] <= 'z' {
			a[i] = a[i] - 32
		}
	}
	return string(a)
}
