package piscine

func IsAlpha(s string) bool {
	cmpt := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' || s[i] >= '0' && s[i] <= '9' || s[i] == ' ' {
			cmpt += 1
		}
	}
	if cmpt == len(s) {
		return true
	} else {
		return false
	}
}
