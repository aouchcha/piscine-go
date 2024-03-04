package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		n = -n
		z01.PrintRune('-')
		z01.PrintRune('9')
		n = 223372036854775808
	}
	if n < 0 {
		n = -n
		z01.PrintRune('-')
	}
	if n <= 9 {
		z01.PrintRune(rune('0' + n))
	} else if n > 9 {
		if n != 0 {
			PrintNbr(n / 10)
			z01.PrintRune(rune('0' + n%10))
		}
	}
}
