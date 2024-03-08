package main

import "fmt"

func PrintMemory(arr [10]byte) {
	hix := "0123456789abcdef"
	for i := 0; i < len(arr); i++ {
		d := int(arr[i]) / 16
		m := int(arr[i]) % 16
		fmt.Print(string(rune(hix[d])))
		fmt.Print(string(rune(hix[m])))
		if i != 3 && i != 7 && i != 9 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
	for _, v := range arr {
		if v >= 32 && v <= 126 {
			fmt.Print(string(v))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
