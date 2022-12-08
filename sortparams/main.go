package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	samba := os.Args
	tr := 0
	for s := range samba {
		tr = s + 1
	}
	for a := 1; a < tr; a++ {
		for b := a + 1; b < tr; b++ {
			if samba[a] > samba[b] {
				samba[a], samba[b] = samba[b], samba[a]
			}
		}
	}
	for b := 1; b <= tr-1; b++ {
		for _, a := range samba[b] {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
