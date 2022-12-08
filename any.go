package piscine

func Any(f func(string) bool, a []string) bool {
	for _, s := range a {
		if f(s) == true {
			return true
		}
	}
	return false
}
