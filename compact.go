package piscine

func Compact(ptr *[]string) int {
	ter := 0
	for _, s := range *ptr {
		if s != "" {
			ter++
		}
	}
	cmpt := 0
	as := make([]string, ter)
	for _, s := range *ptr {
		if s != "" {
			as[cmpt] = s
			cmpt++
		}
	}
	*ptr = as
	return ter
}
