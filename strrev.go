package piscine

func StrRev(s string) string {
	var rvs string
	for i := len(s) - 1; i >= 0; i-- {
		rvs += string(s[i])
	}
	return rvs
}
