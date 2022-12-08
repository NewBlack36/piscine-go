package piscine

func CountIf(f func(string) bool, tab []string) int {
	cmpt := 0
	for _, s := range tab {
		if f(s) == true {
			cmpt++
		}
	}
	return cmpt
}
