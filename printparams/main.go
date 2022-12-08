package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	samba := os.Args
	for i := range samba {
		if i != 0 {
			for a := range samba[i] {
				b := []rune(samba[i])
				{
					z01.PrintRune(b[a])
				}
			}
			z01.PrintRune('\n')
		}
	}
}
