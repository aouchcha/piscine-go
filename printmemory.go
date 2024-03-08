package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"
	l := len(base)
	for i := 0; i < len(arr); i++ {
		mod := int(arr[i]) % l
		div := int(arr[i]) / l
		fmt.Print(string(rune(base[div])))
		fmt.Print(string(rune(base[mod])))
		if i != 3 && i != 7 && i != 9 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}

	for _, char := range arr {
		if char < 32 || char > 126 {
			fmt.Print(".")
		} else {
			fmt.Print(string(char))
		}
	}
}
