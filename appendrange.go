package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	} else {
		var code []int
		for i := min; i < max; i++ {
			code = append(code, i)
		}
		return code
	}
}
