package main

import (
	"fmt"
)

func main() {
	x := Atoi("8333325")
	if IsPrime(x) {
		fmt.Println(int(x - '0'))
		return
	} else {
		FPrime(x)
	}
}

func FPrime(n int) {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if IsPrime(i) {
				fmt.Print(i)
				if n != i {
					fmt.Print("*")
				}
				FPrime(n / i)
				return
			}
		}
	}
}

func Atoi(s string) int {
	result := 0
	for _, num := range s {
		if num >= '0' && num <= '9' {
			result = result*10 + int(num-'0')
		}
	}
	return result
}

func IsPrime(nbr int) bool {
	if nbr == 0 || nbr == 1 {
		return false
	} else if nbr == 2 {
		return true
	} else {
		for i := 3; i < nbr; i++ {
			if nbr%i == 0 {
				return false
			} else {
				return true
			}
		}
	}
	return true
}
