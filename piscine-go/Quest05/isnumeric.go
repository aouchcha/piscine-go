package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}*/

func IsNumeric(s string) bool {
	isnumeric := true
	for i := 0; i < len(s); i++ {
		if s[i] < 48 || s[i] > 57 {
			isnumeric = false
			break
		} else {
			isnumeric = true
		}
	}
	return isnumeric
}
