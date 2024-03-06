package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}*/

func Compare(a, b string) int {
	// str1 := []rune(a)
	// str2 := []rune(b)
	result_of_comp := 3
	if a == b {
		result_of_comp = 0
	} else if a < b {
		result_of_comp = -1
	} else {
		result_of_comp = 1
	}
	return result_of_comp
}
