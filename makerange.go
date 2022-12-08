package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	} else {
		code := make([]int, max-min)
		for i := 0; i < (max - min); i++ {
			code[i] = min + i
		}
		return code
	}
}
