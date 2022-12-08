package piscine

func TrimAtoi(s string) int {
	var a []int
	r := 0
	min := -1
	p := 0
	i := 0
	as := 0
	for _, rune := range s {
		if rune == '-' {
			min = i
		}
		if fonc(rune) {
			if p == 0 {
				p = i
			}
			a = append(a, int(rune-'0'))
		}
		i++
	}

	for count := range a {
		as = count + 1
	}

	for i := 0; i < as; i++ {
		r = r*10 + a[i]
	}

	if min < p && min != -1 {
		r = r * -1
	}
	return r
}

func fonc(x rune) bool {
	for a := '0'; a <= '9'; a++ {
		if x == a {
			return true
		}
	}
	return false
}
