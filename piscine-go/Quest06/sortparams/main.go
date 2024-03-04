package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args

	var table []string

	for i := 1; i < len(argument); i++ {
		table = append(table, argument[i])
	}

	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1] {
				swap := table[j]
				table[j] = table[j+1]
				table[j+1] = swap
			}
		}
	}

	for i := 0; i < len(table); i++ {
		for _, char := range table[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
