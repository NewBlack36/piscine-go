package piscine

func Unmatch(a []int) int {
	var ter int
	for _, v := range a {
		ter = 0
		for _, z := range a {
			if v == z {
				ter++
			}
		}
		if ter%2 != 0 {
			return v
		}
	}
	return -1
}
