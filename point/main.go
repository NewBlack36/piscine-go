package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func PrintNbr(n int) {
	a := '0'
	b := '0'
	for i := 0; i < n/10; i++ {
		a++
	}
	z01.PrintRune(a)
	for j := 0; j < n%10; j++ {
		b++
	}
	z01.PrintRune(b)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	out1 := "x = "
	for _, a := range out1 {
		z01.PrintRune(a)
	}
	PrintNbr(points.x)

	out2 := ", y = "
	for _, b := range out2 {
		z01.PrintRune(b)
	}
	PrintNbr(points.y)

	z01.PrintRune('\n')
}
