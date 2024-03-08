package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoi(12345))  // "12345"
	fmt.Println(Itoi(0))      // "0"
	fmt.Println(Itoi(-12345)) // "-12345"
}

func StrRev(s string) string {
	sli := []rune(s)
	for i := 0; i < len(sli)/2; i++ {
		sli[i], sli[len(sli)-i-1] = sli[len(sli)-i-1], sli[i]
	}
	return string(sli)
}

var word string

func Itoi(nbr int) string {

	if nbr == 0 {
		return "0"
	} else if nbr < 0 {
		word = word + "-"
		nbr = -nbr
	} else {
		if nbr%10 != 0 {
			word = word + string(rune(nbr%10+'0'))
			Itoi(nbr / 10)
		}
	}

	return StrRev(word)
}
