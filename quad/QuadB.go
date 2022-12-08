package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
	} else {
		for a := 1; a <= y; a++ {
			for b := 1; b <= x; b++ {
				if a == 1 && b == 1 {
					z01.PrintRune('/')
				} else if a == 1 && b == x {
					z01.PrintRune(92)
				} else if a == y && b == 1 {
					z01.PrintRune(92)
				} else if a == y && b == x {
					z01.PrintRune('/')
				} else if b == 1 && (a != y || a != 1) {
					z01.PrintRune('*')
				} else if b == x && (a != 1 || a != y) {
					z01.PrintRune('*')
				} else if a == 1 && (b != 1 || b != x) {
					z01.PrintRune('*')
				} else if a == y && (b != 1 || b != x) {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
