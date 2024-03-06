package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}*/

func BasicAtoi(s string) int {
	result := 0
	for _, r := range s {
		result = result*10 + int(r-48)
	}
	return result
}
