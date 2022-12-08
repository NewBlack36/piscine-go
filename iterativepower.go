package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	res := 1

	for power >= 1 {
		res *= nb
		power--
	}
	return res
}
