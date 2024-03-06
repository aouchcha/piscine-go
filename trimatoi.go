package main

import (
	"fmt"
)

func main() {
	fmt.Println(TrimAtoi("12345"))        //123245
	fmt.Println(TrimAtoi("str123ing45"))  //12345
	fmt.Println(TrimAtoi("012 345"))      //12345
	fmt.Println(TrimAtoi("Hello World!")) //0
	fmt.Println(TrimAtoi("sd+x1fa2W3s4")) //1234
	fmt.Println(TrimAtoi("sd-x1fa2W3s4")) //-1234
	fmt.Println(TrimAtoi("sdx1-fa2W3s4")) // 1234
	fmt.Println(TrimAtoi("sdx1+fa2W3s4")) //1234
}

func TrimAtoi(s string) int {
	word := ""
	result := 0
	signe := 1
	for _, char := range s {
		if char == '-' && word == "" {
			signe = -1
			continue
		}
		if char >= '0' && char <= '9' {
			word = word + string(char)
		}
	}
	for _, val := range word {
		result = result*10 + int(val-'0')
	}

	return result * signe
}
