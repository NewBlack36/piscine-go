package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	samba := os.Args
	var a int
	for i := range samba {
		a = i
	}
	for i := a; i > 0; i-- {
		for _, j := range samba[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
