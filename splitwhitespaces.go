package piscine

func SplitWhiteSpaces(s string) []string {
	a := []rune(s)
	cmpt := 0
	lettre := false
	for _, val := range a {
		if val == '\n' || val == '\t' || val == ' ' {
			if lettre {
				cmpt++
				lettre = false
			}
		} else {
			lettre = true
		}
	}
	if lettre {
		cmpt++
	}
	alpha := make([]string, cmpt)
	i := 0
	start := 0
	lettre = false
	for ind, val := range a {
		if val == '\n' || val == '\t' || val == ' ' {
			if lettre {
				alpha[i] = string(a[start:ind])
				i++
				lettre = false
			}
			start = ind + 1
		} else {
			lettre = true
		}
	}
	if lettre {
		alpha[i] = string(a[start:])
	}
	return alpha
}
