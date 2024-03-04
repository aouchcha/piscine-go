package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func PrintStr(str string) {
	slice := []rune(str)
	for _, char := range slice {
		z01.PrintRune(char)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintResult(nbr int) {
	numberdiv := '0'
	numbermod := '0'

	for i := 1; i <= nbr/10; i++ {
		numberdiv++
	}

	z01.PrintRune(numberdiv)

	for j := 1; j <= nbr%10; j++ {
		numbermod++
	}

	z01.PrintRune(numbermod)
}

func main() {
	points := &point{}

	setPoint(points)

	x := points.x
	y := points.y

	PrintStr("x = ")
	PrintResult(x)
	PrintStr(", y = ")
	PrintResult(y)
	PrintStr("\n")
}
