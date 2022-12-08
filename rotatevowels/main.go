package main

import (
	"os"

	"github.com/01-edu/z01"
)

func check(t rune) bool {
	if t == 'a' || t == 'A' || t == 'e' || t == 'E' || t == 'o' || t == 'O' || t == 'u' || t == 'U' || t == 'i' || t == 'I' {
		return true
	}
	return false
}

func main() {
	samba := os.Args[1:]
	s := []rune{}
	v := ""
	len := 0
	q := true
	for _, samba := range samba {
		for _, j := range samba {
			if check(j) {
				s = append(s, j)
				len++
			}
		}
		if q {
			v = samba
			q = false
			continue
		}
		v = v + " " + samba
	}
	cur := 0
	for _, c := range v {
		if check(c) {
			z01.PrintRune(s[len-cur-1])
			cur++
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
