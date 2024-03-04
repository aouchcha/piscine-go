package piscine

/*import (
	"github.com/01-edu/z01"
)*/

/*func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}*/

func NRune(s string, n int) rune {
	// str := []rune(s)
	if n <= 0 || n > len(s) {
		return 0
	}
	/*for i := 0; i < len(s); i++ {
		if i == n-1 {
			break
		}
	}*/
	return []rune(s)[n-1]
}
