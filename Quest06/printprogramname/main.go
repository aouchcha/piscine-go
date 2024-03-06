package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nameprograme := os.Args
	for _, char := range nameprograme[0][2:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
