package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var array []int
		a := 0

		b := 0
		var min int
		for n != 0 {
			a = n % 10
			n /= 10
			array = append(array, a)
		}

		for count := range array {
			b = count + 1
		}
		for i := 0; i < b-1; i++ {
			for j := 0; j < b-i-1; j++ {
				if array[j] > array[j+1] {
					min = array[j]
					array[j] = array[j+1]
					array[j+1] = min
				}
			}
		}
		for i := 0; i < b; i++ {
			z01.PrintRune(rune(array[i] + 48))
		}
	}
}
