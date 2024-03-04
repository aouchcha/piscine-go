package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}*/

func IsPrintable(s string) bool {
	// str := []rune(s)
	isprintable := false
	for _, char := range s {
		if char < 32 || char > 126 {
			isprintable = false
		} else {
			isprintable = true
		}
	}
	return isprintable
}
