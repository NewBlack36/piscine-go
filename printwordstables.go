package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			z01.PrintRune(v2)
		}
		z01.PrintRune('\n')
	}
}
