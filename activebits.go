package piscine

func ActiveBits(n int) int {
	result := 0
	for ; n > 1; n = n / 2 {
		result += n % 2
	}
	result += n
	return result
}
