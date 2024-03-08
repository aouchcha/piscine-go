package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
	fmt.Println(WeAreUnique("hello", "yoall"))
	fmt.Println(WeAreUnique("+154", "154"))
	fmt.Println(WeAreUnique("abc", " "))
	fmt.Println(WeAreUnique("abc", ""))
}

func WeAreUnique(str1, str2 string) int {
	unique := make(map[rune]int)
	str := str1 + str2
	//existance := 0
	result := 0
	if str1 == "" && str2 == "" {
		return -1
	} else {
		for _, char := range str {
			unique[char]++
		}
	}
	for _, char := range str {
		if unique[char] == 1 {
			result++
		}
	}
	return result
}
