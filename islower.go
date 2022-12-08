package piscine

func IsLower(s string) bool {
	cmpt := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			cmpt += 1
		}
	}
	if cmpt == len(s) {
		return true
	} else {
		return false
	}
}
