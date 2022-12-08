package piscine

func Split(s, sep string) []string {
	a := 0
	for i := range sep {
		a = i + 1
	}
	b := 0
	for i := range s {
		b = i + 1
	}
	for i := 0; i < b-a; i++ {
		if s[i:i+a] == sep {
			s = s[:i] + " " + s[i+a:]
			b -= a
		}
	}
	return SplitWhiteSpaces(s)
}
