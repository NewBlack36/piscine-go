package piscine

func Max(a []int) int {
	ter := a[0]
	for i := 0; i < len(a); i++ {
		if ter < a[i] {
			ter = a[i]
		}
	}
	return ter
}
