package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(Atoi("12345"))         // 12345
	fmt.Println(Atoi("0000000012345")) // 12345
	fmt.Println(Atoi("012 345"))       // 0
	fmt.Println(Atoi("Hello World!"))  // 0
	fmt.Println(Atoi("+1234"))         // 1234
	fmt.Println(Atoi("-1234"))         //-1234
	fmt.Println(Atoi("++1234"))        // 0
	fmt.Println(Atoi("--1234"))        // 0
}*/

func Atoi(s string) int {
	str := []rune(s)
	result := 0
	reguilateur := 1
	for i, char := range str {
		if char == '+' && i == 0 {
			continue
		} else if char == '-' && i == 0 {
			reguilateur = -1
			continue
		} else if char < '0' || char > '9' {
			result = 0
			break
		}
		result = result*10 + int(char-'0')

	}
	return result * reguilateur
}
