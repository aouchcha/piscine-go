package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}*/

func IsLower(s string) bool {
	counter := 0
	islower := false
	for i := 0; i < len(s); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			counter++
		}
	}
	if counter == len(s) {
		islower = true
	} else {
		islower = false
	}
	return islower
}
