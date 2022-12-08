package piscine

func IsUpper(s string) bool {
	cmpt := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			cmpt += 1
		}
	}
	if cmpt == len(s) {
		return true
	} else {
		return false
	}
}
