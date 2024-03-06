package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}*/

func IsUpper(s string) bool {
	count := 0
	isupper := false
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 {
			count++
		}
	}
	if count == len(s) {
		isupper = true
	} else {
		isupper = false
	}

	return isupper
}
