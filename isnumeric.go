package piscine

func IsNumeric(s string) bool {
	cmpt := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			cmpt += 1
		}
	}
	if cmpt == len(s) {
		return true
	} else {
		return false
	}
}
