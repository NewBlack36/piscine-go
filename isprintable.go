package piscine

func IsPrintable(s string) bool {
	cmpt := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 32 && s[i] <= 126 {
			cmpt += 1
		}
	}
	if cmpt == len(s) {
		return true
	} else {
		return false
	}
}
