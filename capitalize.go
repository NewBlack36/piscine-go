package piscine

func aze(m rune) bool {
	if (m >= 'A' && m <= 'Z') || (m >= 'a' && m <= 'z') || (m >= '0' && m <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	a := []rune(s)

	n := true

	for i := 0; i < len(s); i++ {
		if aze(a[i]) == true && n {
			if a[i] >= 'a' && a[i] <= 'z' {
				a[i] = 'A' - 'a' + a[i]
			}
			n = false
		} else if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] = 'a' - 'A' + a[i]
		} else if aze(a[i]) == false {
			n = true
		}
	}
	return string(a)
}
