package main

import (
	"fmt"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// cone 1
			if i == 0 && j == 0 {
				fmt.Print("o")
			} else /* 9nt 2 */ if i == 0 && j == x-1 {
				fmt.Print("o")
			} else /* 9nt 3*/ if i == y-1 && j == 0 {
				fmt.Print("o")
			} else /*9nt 4 */ if i == y-1 && j == x-1 {
				fmt.Print("o")
			} else if i == 0 && j != 0 {
				fmt.Print("-")
			} else if i == y-1 && j != 0 && j < x-1 {
				fmt.Print("-")
			} else if i > 0 && j == 0 || i > 0 && j == x-1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// var x int = 5
	// var y int = 3
	QuadA(5, 3)
	QuadA(1, 1)
	QuadA(1, 5)
	QuadA(10, 3)
}
