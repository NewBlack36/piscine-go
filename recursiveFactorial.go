package piscine

import (
	"math"
)

func RecursiveFactorial(x int) int {
	if 0 > x {
		return 0
	} else if 0 == x || 1 == x {
		return 1
	} else {
		res := 1
		for i := 1; i <= x; i++ {
			res *= i
			if math.MaxInt32 < res {
				res = 0
				break
			}
		}
		return res
	}
}
