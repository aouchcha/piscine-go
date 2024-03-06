package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}*/

func BasicAtoi2(s string) int {
	var result int
	str := []rune(s)
	for _, char := range str {
		if char < 48 || char > 57 {
			result = 0
			break
		} else {
			result = result*10 + int(char-'0')
		}
	}
	return result
}
