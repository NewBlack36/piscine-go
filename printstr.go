package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}
