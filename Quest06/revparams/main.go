package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args

	for i := len(argument) - 1; i >= 1; i-- {
		for _, char := range argument[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
