package piscine

func Atoi(s string) int {
	return conv(may(s))
}

func may(s string) (sot int, abs []rune) {
	a := []rune(s)
	a = trimspace1(a)
	if s == "" {
		return
	}
	switch a[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sot, abs = 1, a
	case '+':
		sot, abs = 1, a[1:]
	case '-':
		sot, abs = -1, a[1:]
	default:
		return
	}
	for i, b := range abs {
		if b < '0' || '9' < b {
			abs = abs[:i]
			break
		}
	}
	return
}

func conv(sot int, absStr []rune) int {
	abs := 0
	for _, b := range absStr {
		abs = abs*10 + int(b-'0')
	}
	return sot * abs
}

func trimspace1(a []rune) []rune {
	for i := 0; i <= len(a)-1; i++ {
		if a[i] == '-' && i != 0 || a[i] == '+' && i != 0 {
			i = '0'
			return []rune{'0'}
		}
		if a[i] == ' ' || a[i] == '.' || a[i] >= 'a' {
			return []rune{'0'}
		}
	}
	return a
}
