package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	samba := os.Args[0]
	for i, as := range samba {
		if i > 1 {
			z01.PrintRune(as)
		}
	}
	z01.PrintRune('\n')
}
