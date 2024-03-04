package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}*/

func IsAlpha(s string) bool {
	isalpha := true
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 || s[i] >= 97 && s[i] <= 122 || s[i] >= 48 && s[i] <= 57 {
			count++
		}
	}
	if count == len(s) {
		isalpha = true
	} else {
		isalpha = false
	}
	return isalpha
}
