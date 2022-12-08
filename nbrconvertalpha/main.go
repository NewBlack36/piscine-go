package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	n := os.Args[1:]
	var r []int
	k := 96
	for _, w := range n {
		if w == "--upper" {
			k = 64
			continue
		}
		num := 0
		for _, j := range w {
			num = num*10 + int(rune(j)-'0')
		}
		r = append(r, num)
	}
	for i := 0; i < len(r); i++ {
		if len(r) == 0 {
			break
		} else if r[i] > 26 {
			z01.PrintRune(rune(' '))
			continue
		}
		z01.PrintRune(rune(r[i]) + rune(k))
	}
	if len(r) > 0 {
		z01.PrintRune('\n')
	}
}
